package cli

import (
	"github.com/segmentfault/answer/internal/base/data"
	"github.com/segmentfault/answer/pkg/dir"
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
