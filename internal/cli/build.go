package cli

import (
	"bytes"
	"embed"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/answerdev/answer/pkg/dir"
	"github.com/answerdev/answer/pkg/writer"
	"github.com/answerdev/answer/ui"
	cp "github.com/otiai10/copy"
)

type answerBuilder struct {
	buildingMaterial *buildingMaterial
	BuildError       error
}

type buildingMaterial struct {
	plugins            []*pluginInfo
	outputPath         string
	tmpDir             string
	originalAnswerInfo OriginalAnswerInfo
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
	parentDir, _ := filepath.Abs(".")
	tmpDir, _ := os.MkdirTemp(parentDir, "answer_build")
	if len(outputPath) == 0 {
		outputPath = filepath.Join(parentDir, "new_answer")
	}
	material := &buildingMaterial{
		plugins:            formatPlugins(plugins),
		outputPath:         outputPath,
		tmpDir:             tmpDir,
		originalAnswerInfo: originalAnswerInfo,
	}
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
		localPlugins  []string
		remotePlugins []string
	)
	for _, p := range b.plugins {
		if len(p.Path) == 0 {
			remotePlugins = append(remotePlugins, p.Name)
		} else {
			localPluginDir := filepath.Base(p.Path)
			localPlugins = append(localPlugins, localPluginDir)
			if err := cp.Copy(p.Path, filepath.Join(b.tmpDir, localPluginDir)); err != nil {
				return err
			}
		}
	}

	mainGoFile := &bytes.Buffer{}
	tmpl, err := template.New("main").Parse(mainGoTpl)
	if err != nil {
		return err
	}
	err = tmpl.Execute(mainGoFile, map[string]any{
		"remote_plugins": remotePlugins,
		"local_plugins":  localPlugins,
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
	return
}

func downloadGoModFile(b *buildingMaterial) (err error) {
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
	ldflags := fmt.Sprintf(`-ldflags="-X answercmd.Version=%s -X answercmd.Revision=%s -X answercmd.Time=%s`,
		versionInfo.Version, versionInfo.Revision, versionInfo.Time)
	err = b.newExecCmd("go", "build", ldflags, "-o", b.outputPath, ".").Run()
	if err != nil {
		return err
	}
	return
}

func cleanByproduct(b *buildingMaterial) (err error) {
	//return os.RemoveAll(b.tmpDir)
	return nil
}

func (b *buildingMaterial) newExecCmd(command string, args ...string) *exec.Cmd {
	cmd := exec.Command(command, args...)
	cmd.Dir = b.tmpDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd
}

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

replace github.com/answerdev/answer latest => replace github.com/answerdev/answer feature-plugin
`
)
