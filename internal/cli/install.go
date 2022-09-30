package cli

import (
	"bufio"
	"fmt"
	"os"

	"github.com/segmentfault/answer/configs"
	"github.com/segmentfault/answer/i18n"
	"github.com/segmentfault/answer/pkg/dir"
)

var SuccessMsg = `
answer initialized successfully.
`

var HasBeenInitializedMsg = `
Has been initialized.
`

func InitConfig() {
	exist, err := PathExists("data/config.yaml")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}
	if exist {
		fmt.Println(HasBeenInitializedMsg)
		os.Exit(0)
	}

	_, err = dir.CreatePathIsNotExist("data")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}
	WriterFile("data/config.yaml", string(configs.Config))
	_, err = dir.CreatePathIsNotExist("data/i18n")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}
	_, err = dir.CreatePathIsNotExist("data/upfiles")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}
	i18nList, err := i18n.I18n.ReadDir(".")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}
	for _, item := range i18nList {
		path := fmt.Sprintf("data/i18n/%s", item.Name())
		content, err := i18n.I18n.ReadFile(item.Name())
		if err != nil {
			continue
		}
		WriterFile(path, string(content))
	}
	fmt.Println(SuccessMsg)
	os.Exit(0)
}

func WriterFile(filePath, content string) error {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	if err != nil {
		return err
	}
	write := bufio.NewWriter(file)
	write.WriteString(content)
	write.Flush()
	return nil
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
