package cli

import (
	"bytes"
	"embed"
	"fmt"
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
	builder.DoTask(mergeI18nFiles)
	builder.DoTask(replaceNecessaryFile)
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

func createMainGoFile(b *buildingMaterial) (err error) {
	fmt.Printf("[build] tmp dir: %s\n", b.tmpDir)
	err = dir.CreateDirIfNotExist(b.tmpDir)
	if err != nil {
		return err
	}

	var (
		//localPlugins  []string
		remotePlugins []string
	)
	for _, p := range b.plugins {
		//if len(p.Path) == 0 {
		remotePlugins = append(remotePlugins, versionedModulePath(p.Name, p.Version))
		//} else {
		//localPluginDir := filepath.Base(p.Path)
		//localPlugins = append(localPlugins, localPluginDir)
		//if err := cp.Copy(p.Path, filepath.Join(b.tmpDir, localPluginDir)); err != nil {
		//	return err
		//}
		//}
	}

	mainGoFile := &bytes.Buffer{}
	tmpl, err := template.New("main").Parse(mainGoTpl)
	if err != nil {
		return err
	}
	err = tmpl.Execute(mainGoFile, map[string]any{
		"remote_plugins": remotePlugins,
		//"local_plugins":  localPlugins,
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

func downloadGoModFile(b *buildingMaterial) (err error) {
	// If user specify a module replacement, use it. Otherwise, use the latest version.
	if len(b.answerModuleReplacement) > 0 {
		replacement := fmt.Sprintf("%s=%s", "github.com/answerdev/answer@latest", b.answerModuleReplacement)
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

func replaceNecessaryFile(b *buildingMaterial) (err error) {
	fmt.Printf("try to replace ui build directory\n")
	uiBuildDir := filepath.Join(b.tmpDir, "vendor/github.com/answerdev/answer/ui")
	err = copyDirEntries(ui.Build, ".", uiBuildDir)
	return err
}

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

func copyDirEntries(sourceFs embed.FS, sourceDir string, targetDir string) (err error) {
	entries, err := ui.Build.ReadDir(sourceDir)
	if err != nil {
		return err
	}

	err = dir.CreateDirIfNotExist(targetDir)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		if entry.IsDir() {
			err = copyDirEntries(sourceFs, filepath.Join(sourceDir, entry.Name()), filepath.Join(targetDir, entry.Name()))
			if err != nil {
				return err
			}
			continue
		}
		file, err := sourceFs.ReadFile(filepath.Join(sourceDir, entry.Name()))
		if err != nil {
			return err
		}
		filename := filepath.Join(targetDir, entry.Name())
		err = os.WriteFile(filename, file, 0666)
		if err != nil {
			return err
		}
	}
	return nil
}

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
