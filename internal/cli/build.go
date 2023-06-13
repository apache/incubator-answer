package cli

import (
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/Masterminds/semver/v3"
	"github.com/answerdev/answer/pkg/dir"
	"github.com/answerdev/answer/pkg/writer"
	"github.com/answerdev/answer/ui"
	"github.com/segmentfault/pacman/log"
	"gopkg.in/yaml.v3"
)

const (
	mainGoTpl = `package main

import (
	answercmd "github.com/answerdev/answer/cmd"

  // remote plugins
	{{- range .remote_plugins}}
	_ "{{.}}"
	{{- end}}

  // local plugins
	{{- range .local_plugins}}
	_ "answer/{{.}}"
	{{- end}}
)

func main() {
	answercmd.Main()
}
`
	goModTpl = `module answer

go 1.19
`
)

type answerBuilder struct {
	buildingMaterial *buildingMaterial
	BuildError       error
}

type buildingMaterial struct {
	answerModuleReplacement string
	plugins                 []*pluginInfo
	outputPath              string
	tmpDir                  string
	originalAnswerInfo      OriginalAnswerInfo
}

type OriginalAnswerInfo struct {
	Version  string
	Revision string
	Time     string
}

type pluginInfo struct {
	// Name of the plugin e.g. github.com/answerdev/github-connector
	Name string
	// Path to the plugin. If path exist, read plugin from local filesystem
	Path string
	// Version of the plugin
	Version string
}

func newAnswerBuilder(outputPath string, plugins []string, originalAnswerInfo OriginalAnswerInfo) *answerBuilder {
	material := &buildingMaterial{originalAnswerInfo: originalAnswerInfo}
	parentDir, _ := filepath.Abs(".")
	material.tmpDir, _ = os.MkdirTemp(parentDir, "answer_build")
	if len(outputPath) == 0 {
		outputPath = filepath.Join(parentDir, "new_answer")
	}
	material.outputPath = outputPath
	material.plugins = formatPlugins(plugins)
	material.answerModuleReplacement = os.Getenv("ANSWER_MODULE")
	return &answerBuilder{
		buildingMaterial: material,
	}
}

func (a *answerBuilder) DoTask(task func(b *buildingMaterial) error) {
	if a.BuildError != nil {
		return
	}
	a.BuildError = task(a.buildingMaterial)
}

// BuildNewAnswer builds a new answer with specified plugins
func BuildNewAnswer(outputPath string, plugins []string, originalAnswerInfo OriginalAnswerInfo) (err error) {
	builder := newAnswerBuilder(outputPath, plugins, originalAnswerInfo)
	builder.DoTask(createMainGoFile)
	builder.DoTask(downloadGoModFile)
	builder.DoTask(copyUIFiles)
	builder.DoTask(overwriteIndexTs)
	builder.DoTask(buildUI)
	builder.DoTask(mergeI18nFiles)
	builder.DoTask(buildBinary)
	builder.DoTask(cleanByproduct)
	return builder.BuildError
}

func formatPlugins(plugins []string) (formatted []*pluginInfo) {
	for _, plugin := range plugins {
		plugin = strings.TrimSpace(plugin)
		// plugin description like this 'github.com/answerdev/github-connector@latest=/local/path'
		info := &pluginInfo{}
		plugin, info.Path, _ = strings.Cut(plugin, "=")
		info.Name, info.Version, _ = strings.Cut(plugin, "@")
		formatted = append(formatted, info)
	}
	return formatted
}

// createMainGoFile creates main.go file in tmp dir that content is mainGoTpl
func createMainGoFile(b *buildingMaterial) (err error) {
	fmt.Printf("[build] tmp dir: %s\n", b.tmpDir)
	err = dir.CreateDirIfNotExist(b.tmpDir)
	if err != nil {
		return err
	}

	var (
		remotePlugins []string
	)
	for _, p := range b.plugins {
		remotePlugins = append(remotePlugins, versionedModulePath(p.Name, p.Version))
	}

	mainGoFile := &bytes.Buffer{}
	tmpl, err := template.New("main").Parse(mainGoTpl)
	if err != nil {
		return err
	}
	err = tmpl.Execute(mainGoFile, map[string]any{
		"remote_plugins": remotePlugins,
	})
	if err != nil {
		return err
	}

	err = writer.WriteFile(filepath.Join(b.tmpDir, "main.go"), mainGoFile.String())
	if err != nil {
		return err
	}

	err = writer.WriteFile(filepath.Join(b.tmpDir, "go.mod"), goModTpl)
	if err != nil {
		return err
	}

	for _, p := range b.plugins {
		if len(p.Path) == 0 {
			continue
		}
		replacement := fmt.Sprintf("%s@v%s=%s", p.Name, p.Version, p.Path)
		err = b.newExecCmd("go", "mod", "edit", "-replace", replacement).Run()
		if err != nil {
			return err
		}
	}
	return
}

// downloadGoModFile run go mod commands to download dependencies
func downloadGoModFile(b *buildingMaterial) (err error) {
	// If user specify a module replacement, use it. Otherwise, use the latest version.
	if len(b.answerModuleReplacement) > 0 {
		replacement := fmt.Sprintf("%s=%s", "github.com/answerdev/answer", b.answerModuleReplacement)
		err = b.newExecCmd("go", "mod", "edit", "-replace", replacement).Run()
		if err != nil {
			return err
		}
	}

	err = b.newExecCmd("go", "mod", "tidy").Run()
	if err != nil {
		return err
	}

	err = b.newExecCmd("go", "mod", "vendor").Run()
	if err != nil {
		return err
	}
	return
}

// copyUIFiles copy ui files from answer module to tmp dir
func copyUIFiles(b *buildingMaterial) (err error) {
	goListCmd := b.newExecCmd("go", "list", "-mod=mod", "-m", "-f", "{{.Dir}}", "github.com/answerdev/answer")
	buf := new(bytes.Buffer)
	goListCmd.Stdout = buf
	if err = goListCmd.Run(); err != nil {
		return fmt.Errorf("failed to run go list: %w", err)
	}

	goModUIDir := filepath.Join(strings.TrimSpace(buf.String()), "ui")
	localUIBuildDir := filepath.Join(b.tmpDir, "vendor/github.com/answerdev/answer/ui/")
	if err = copyDirEntries(os.DirFS(goModUIDir), ".", localUIBuildDir); err != nil {
		return fmt.Errorf("failed to copy ui files: %w", err)
	}
	return nil
}

// overwriteIndexTs overwrites index.ts file in ui/src/plugins/ dir
func overwriteIndexTs(b *buildingMaterial) (err error) {
	localUIPluginDir := filepath.Join(b.tmpDir, "vendor/github.com/answerdev/answer/ui/src/plugins/")

	folders, err := getFolders(localUIPluginDir)
	if err != nil {
		return fmt.Errorf("failed to get folders: %w", err)
	}

	content := generateIndexTsContent(folders)
	err = os.WriteFile(filepath.Join(localUIPluginDir, "index.ts"), []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("failed to write index.ts: %w", err)
	}
	return nil
}

func getFolders(dir string) ([]string, error) {
	var folders []string
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		if file.IsDir() && file.Name() != "builtin" {
			folders = append(folders, file.Name())
		}
	}
	return folders, nil
}

func generateIndexTsContent(folders []string) string {
	builder := &strings.Builder{}
	builder.WriteString("export default null;\n\n")
	for _, folder := range folders {
		builder.WriteString(fmt.Sprintf("export { default as %s } from './%s';\n", folder, folder))
	}
	return builder.String()
}

// buildUI run pnpm install and pnpm build commands to build ui
func buildUI(b *buildingMaterial) (err error) {
	localUIBuildDir := filepath.Join(b.tmpDir, "vendor/github.com/answerdev/answer/ui")

	pnpmInstallCmd := b.newExecCmd("pnpm", "install")
	pnpmInstallCmd.Dir = localUIBuildDir
	if err = pnpmInstallCmd.Run(); err != nil {
		return err
	}

	pnpmBuildCmd := b.newExecCmd("pnpm", "build")
	pnpmBuildCmd.Dir = localUIBuildDir
	if err = pnpmBuildCmd.Run(); err != nil {
		return err
	}
	return nil
}

func replaceNecessaryFile(b *buildingMaterial) (err error) {
	fmt.Printf("try to replace ui build directory\n")
	uiBuildDir := filepath.Join(b.tmpDir, "vendor/github.com/answerdev/answer/ui")
	err = copyDirEntries(ui.Build, ".", uiBuildDir)
	return err
}

// mergeI18nFiles merge i18n files
func mergeI18nFiles(b *buildingMaterial) (err error) {
	fmt.Printf("try to merge i18n files\n")

	type YamlPluginContent struct {
		Plugin map[string]any `yaml:"plugin"`
	}

	pluginAllTranslations := make(map[string]*YamlPluginContent)
	for _, plugin := range b.plugins {
		i18nDir := filepath.Join(b.tmpDir, fmt.Sprintf("vendor/%s/i18n", plugin.Name))
		fmt.Println("i18n dir: ", i18nDir)
		if !dir.CheckDirExist(i18nDir) {
			continue
		}

		entries, err := os.ReadDir(i18nDir)
		if err != nil {
			return err
		}

		for _, file := range entries {
			// ignore directory
			if file.IsDir() {
				continue
			}
			// ignore non-YAML file
			if filepath.Ext(file.Name()) != ".yaml" {
				continue
			}
			buf, err := os.ReadFile(filepath.Join(i18nDir, file.Name()))
			if err != nil {
				log.Debugf("read translation file failed: %s %s", file.Name(), err)
				continue
			}

			translation := &YamlPluginContent{}
			if err = yaml.Unmarshal(buf, translation); err != nil {
				log.Debugf("unmarshal translation file failed: %s %s", file.Name(), err)
				continue
			}

			if pluginAllTranslations[file.Name()] == nil {
				pluginAllTranslations[file.Name()] = &YamlPluginContent{Plugin: make(map[string]any)}
			}
			for k, v := range translation.Plugin {
				pluginAllTranslations[file.Name()].Plugin[k] = v
			}
		}
	}

	originalI18nDir := filepath.Join(b.tmpDir, "vendor/github.com/answerdev/answer/i18n")
	entries, err := os.ReadDir(originalI18nDir)
	if err != nil {
		return err
	}

	for _, file := range entries {
		// ignore directory
		if file.IsDir() {
			continue
		}
		// ignore non-YAML file
		filename := file.Name()
		if filepath.Ext(filename) != ".yaml" && filename != "i18n.yaml" {
			continue
		}

		// if plugin don't have this translation file, ignore it
		if pluginAllTranslations[filename] == nil {
			continue
		}

		out, _ := yaml.Marshal(pluginAllTranslations[filename])

		buf, err := os.OpenFile(filepath.Join(originalI18nDir, filename), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Debugf("read translation file failed: %s %s", filename, err)
			continue
		}

		_, _ = buf.WriteString("\n")
		_, _ = buf.Write(out)
		_ = buf.Close()
	}
	return err
}

func copyDirEntries(sourceFs fs.FS, sourceDir string, targetDir string) (err error) {
	err = dir.CreateDirIfNotExist(targetDir)
	if err != nil {
		return err
	}
	err = fs.WalkDir(sourceFs, sourceDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Convert the path to use forward slashes, important because we use embedded FS which always uses forward slashes
		path = filepath.ToSlash(path)

		// Construct the absolute path for the source file/directory
		srcPath := filepath.Join(sourceDir, path)

		// Construct the absolute path for the destination file/directory
		dstPath := filepath.Join(targetDir, path)

		if d.IsDir() {
			// Create the directory in the destination
			err := os.MkdirAll(dstPath, os.ModePerm)
			if err != nil {
				return fmt.Errorf("failed to create directory %s: %w", dstPath, err)
			}
		} else {
			// Open the source file
			srcFile, err := sourceFs.Open(srcPath)
			if err != nil {
				return fmt.Errorf("failed to open source file %s: %w", srcPath, err)
			}
			defer srcFile.Close()

			// Create the destination file
			dstFile, err := os.Create(dstPath)
			if err != nil {
				return fmt.Errorf("failed to create destination file %s: %w", dstPath, err)
			}
			defer dstFile.Close()

			// Copy the file contents
			_, err = io.Copy(dstFile, srcFile)
			if err != nil {
				return fmt.Errorf("failed to copy file contents from %s to %s: %w", srcPath, dstPath, err)
			}
		}

		return nil
	})

	return err
}

// buildBinary build binary file
func buildBinary(b *buildingMaterial) (err error) {
	versionInfo := b.originalAnswerInfo
	cmdPkg := "github.com/answerdev/answer/cmd"
	ldflags := fmt.Sprintf("-X %s.Version=%s -X %s.Revision=%s -X %s.Time=%s",
		cmdPkg, versionInfo.Version, cmdPkg, versionInfo.Revision, cmdPkg, versionInfo.Time)
	err = b.newExecCmd("go", "build",
		"-ldflags", ldflags, "-o", b.outputPath, ".").Run()
	if err != nil {
		return err
	}
	return
}

// cleanByproduct delete tmp dir
func cleanByproduct(b *buildingMaterial) (err error) {
	return os.RemoveAll(b.tmpDir)
}

func (b *buildingMaterial) newExecCmd(command string, args ...string) *exec.Cmd {
	cmd := exec.Command(command, args...)
	fmt.Println(cmd.Args)
	cmd.Dir = b.tmpDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd
}

func versionedModulePath(modulePath, moduleVersion string) string {
	if moduleVersion == "" {
		return modulePath
	}
	ver, err := semver.StrictNewVersion(strings.TrimPrefix(moduleVersion, "v"))
	if err != nil {
		return modulePath
	}
	major := ver.Major()
	if major > 1 {
		modulePath += fmt.Sprintf("/v%d", major)
	}
	return path.Clean(modulePath)
}
