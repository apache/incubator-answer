package cli

import (
	"github.com/answerdev/answer/internal/base/data"
	"github.com/answerdev/answer/pkg/dir"
)

func CheckConfigFile(configPath string) bool {
	return dir.CheckFileExist(configPath)
}

func CheckUploadDir() bool {
	return dir.CheckDirExist(UploadFilePath)
}

func CheckDB(dataConf *data.Database) bool {
	db, err := data.NewDB(false, dataConf)
	if err != nil {
		return false
	}
	if err = db.Ping(); err != nil {
		return false
	}
	return true
}
