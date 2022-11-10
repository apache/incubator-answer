package cli

import (
	"fmt"

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
		fmt.Printf("connection database failed: %s\n", err)
		return false
	}
	if err = db.Ping(); err != nil {
		fmt.Printf("connection ping database failed: %s\n", err)
		return false
	}
	if !mustInstalled {
		return true
	}

	exist, err := db.IsTableExist(&entity.Version{})
	if err != nil {
		fmt.Printf("check table exist failed: %s\n", err)
		return false
	}
	if !exist {
		fmt.Printf("check table not exist\n")
		return false
	}
	return true
}
