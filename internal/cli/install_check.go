package cli

import (
	"github.com/segmentfault/answer/internal/base/data"
	"github.com/segmentfault/answer/pkg/dir"
)

func CheckConfigFile() bool {
	return dir.CheckPathExist(defaultConfigFilePath)
}

func CheckUploadDir() bool {
	return dir.CheckPathExist(defaultConfigFilePath)
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
