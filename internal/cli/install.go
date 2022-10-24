package cli

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"github.com/segmentfault/answer/configs"
	"github.com/segmentfault/answer/i18n"
	"github.com/segmentfault/answer/pkg/dir"
)

const (
	DefaultConfigFileName = "config.yaml"
)

var (
	ConfigFilePath = "/conf/"
	UploadFilePath = "/upfiles/"
	I18nPath       = "/i18n/"
)

// InstallAllInitialEnvironment install all initial environment
func InstallAllInitialEnvironment(dataDirPath string) {
	ConfigFilePath = filepath.Join(dataDirPath, ConfigFilePath)
	UploadFilePath = filepath.Join(dataDirPath, UploadFilePath)
	I18nPath = filepath.Join(dataDirPath, I18nPath)

	installConfigFile()
	installUploadDir()
	installI18nBundle()
	fmt.Println("install all initial environment done")
	return
}

func installConfigFile() {
	fmt.Println("[config-file] try to install...")
	defaultConfigFile := filepath.Join(ConfigFilePath, DefaultConfigFileName)

	// if config file already exists do nothing.
	if CheckConfigFile(defaultConfigFile) {
		fmt.Printf("[config-file] %s already exists\n", defaultConfigFile)
		return
	}

	if err := dir.CreateDirIfNotExist(ConfigFilePath); err != nil {
		fmt.Printf("[config-file] create directory fail %s\n", err.Error())
		return
	}
	fmt.Printf("[config-file] create directory success, config file is %s\n", defaultConfigFile)

	if err := writerFile(defaultConfigFile, string(configs.Config)); err != nil {
		fmt.Printf("[config-file] install fail %s\n", err.Error())
		return
	}
	fmt.Printf("[config-file] install success\n")
}

func installUploadDir() {
	fmt.Println("[upload-dir] try to install...")
	if err := dir.CreateDirIfNotExist(UploadFilePath); err != nil {
		fmt.Printf("[upload-dir] install fail %s\n", err.Error())
	} else {
		fmt.Printf("[upload-dir] install success, upload directory is %s\n", UploadFilePath)
	}
}

func installI18nBundle() {
	fmt.Println("[i18n] try to install i18n bundle...")
	if err := dir.CreateDirIfNotExist(I18nPath); err != nil {
		fmt.Println(err.Error())
		return
	}

	i18nList, err := i18n.I18n.ReadDir(".")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("[i18n] find i18n bundle %d\n", len(i18nList))
	for _, item := range i18nList {
		path := filepath.Join(I18nPath, item.Name())
		content, err := i18n.I18n.ReadFile(item.Name())
		if err != nil {
			continue
		}
		fmt.Printf("[i18n] install %s bundle...\n", item.Name())
		err = writerFile(path, string(content))
		if err != nil {
			fmt.Printf("[i18n] install %s bundle fail: %s\n", item.Name(), err.Error())
		} else {
			fmt.Printf("[i18n] install %s bundle success\n", item.Name())
		}
	}
}

func writerFile(filePath, content string) error {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer func() {
		_ = file.Close()
	}()
	writer := bufio.NewWriter(file)
	if _, err := writer.WriteString(content); err != nil {
		return err
	}
	if err := writer.Flush(); err != nil {
		return err
	}
	return nil
}
