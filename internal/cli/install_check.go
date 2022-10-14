package cli

import (
	"github.com/segmentfault/answer/internal/base/data"
	"github.com/segmentfault/answer/pkg/dir"
)

func CheckConfigFile(configPath string) bool {
	return dir.CheckPathExist(configPath)
}

func CheckUploadDir() bool {
	return dir.CheckPathExist(UploadFilePath)
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
