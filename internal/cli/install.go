package cli

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"path/filepath"

	"github.com/segmentfault/answer/assets"
	"github.com/segmentfault/answer/configs"
	"github.com/segmentfault/answer/i18n"
	"github.com/segmentfault/answer/internal/base/data"
	"github.com/segmentfault/answer/internal/entity"
	"github.com/segmentfault/answer/pkg/dir"
)

const (
	defaultConfigFilePath = "/data/conf/"
	defaultUploadFilePath = "/data/upfiles/"
	defaultI18nPath       = "/data/i18n/"
)

// InstallAllInitialEnvironment install all initial environment
func InstallAllInitialEnvironment() {
	installConfigFile()
	installUploadDir()
	installI18nBundle()
	fmt.Println("install all initial environment done")
	return
}

func installConfigFile() {
	fmt.Println("[config-file] try to install...")
	defaultConfigFile := filepath.Join(defaultConfigFilePath, "config.yaml")

	// if config file already exists do nothing.
	if CheckConfigFile(defaultConfigFile) {
		fmt.Println("[config-file] already exists")
		return
	}

	if _, err := dir.CreatePathIsNotExist(defaultConfigFilePath); err != nil {
		fmt.Printf("[config-file] create directory fail %s\n", err.Error())
		return
	}
	fmt.Printf("[config-file] create directory success\n")

	if err := WriterFile(defaultConfigFile, string(configs.Config)); err != nil {
		fmt.Printf("[config-file] install fail %s\n", err.Error())
		return
	}
	fmt.Printf("[config-file] install success\n")
}

func installUploadDir() {
	fmt.Println("[upload-dir] try to install...")
	if _, err := dir.CreatePathIsNotExist(defaultUploadFilePath); err != nil {
		fmt.Printf("[upload-dir] install fail %s\n", err.Error())
	} else {
		fmt.Printf("[upload-dir] install success\n")
	}
}

func installI18nBundle() {
	fmt.Println("[i18n] try to install i18n bundle...")
	if _, err := dir.CreatePathIsNotExist(defaultI18nPath); err != nil {
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
		path := filepath.Join(defaultI18nPath, item.Name())
		content, err := i18n.I18n.ReadFile(item.Name())
		if err != nil {
			continue
		}
		fmt.Printf("[i18n] install %s bundle...\n", item.Name())
		err = WriterFile(path, string(content))
		if err != nil {
			fmt.Printf("[i18n] install %s bundle fail: %s\n", item.Name(), err.Error())
		} else {
			fmt.Printf("[i18n] install %s bundle success\n", item.Name())
		}
	}
}

func WriterFile(filePath, content string) error {
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

// InitDB init db
func InitDB(dataConf *data.Database) (err error) {
	fmt.Println("[database] try to initialize database")
	db, err := data.NewDB(false, dataConf)
	if err != nil {
		return err
	}
	// check db connection
	if err = db.Ping(); err != nil {
		return err
	}
	fmt.Println("[database] connect success")

	exist, err := db.IsTableExist(&entity.User{})
	if err != nil {
		return err
	}
	if exist {
		fmt.Println("[database] already exists")
		return nil
	}

	// create table if not exist
	s := &bytes.Buffer{}
	s.Write(assets.AnswerSql)
	_, err = db.Import(s)
	if err != nil {
		return err
	}
	fmt.Println("[database] execute sql successfully")
	return nil
}
