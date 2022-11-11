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

// CheckDBConnection check database whether the connection is normal
func CheckDBConnection(dataConf *data.Database) bool {
	db, err := data.NewDB(false, dataConf)
	if err != nil {
		fmt.Printf("connection database failed: %s\n", err)
		return false
	}
	if err = db.Ping(); err != nil {
		fmt.Printf("connection ping database failed: %s\n", err)
		return false
	}

	return true
}

// CheckDBTableExist check database whether the table is already exists
func CheckDBTableExist(dataConf *data.Database) bool {
	db, err := data.NewDB(false, dataConf)
	if err != nil {
		fmt.Printf("connection database failed: %s\n", err)
		return false
	}
	if err = db.Ping(); err != nil {
		fmt.Printf("connection ping database failed: %s\n", err)
		return false
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
