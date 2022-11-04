package cli

import (
	"github.com/answerdev/answer/internal/base/data"
	"github.com/answerdev/answer/internal/entity"
	"github.com/answerdev/answer/pkg/dir"
)

func CheckConfigFile(configPath string) bool {
	return dir.CheckFileExist(configPath)
}

func CheckUploadDir() bool {
	return dir.CheckDirExist(UploadFilePath)
}

// CheckDB check database whether the connection is normal
// if mustInstalled is true, will check table if already exists
func CheckDB(dataConf *data.Database, mustInstalled bool) bool {
	db, err := data.NewDB(false, dataConf)
	if err != nil {
		return false
	}
	if err = db.Ping(); err != nil {
		return false
	}
	if !mustInstalled {
		return true
	}

	exist, err := db.IsTableExist(&entity.Version{})
	if err != nil {
		return false
	}
	if !exist {
		return false
	}
	return true
}
