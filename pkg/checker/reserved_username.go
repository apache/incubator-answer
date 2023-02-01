package checker

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sync"

	"github.com/answerdev/answer/configs"
	"github.com/answerdev/answer/internal/cli"
	"github.com/answerdev/answer/pkg/dir"
)

var (
	reservedUsernameMapping = make(map[string]bool)
	reservedUsernameInit    sync.Once
)

func initReservedUsername() {
	reservedUsernamesJsonFilePath := filepath.Join(cli.ConfigFileDir, cli.DefaultReservedUsernamesConfigFileName)
	if dir.CheckFileExist(reservedUsernamesJsonFilePath) {
		// if reserved username file exists, read it and replace configuration
		reservedUsernamesJsonFile, err := os.ReadFile(reservedUsernamesJsonFilePath)
		if err == nil {
			configs.ReservedUsernames = reservedUsernamesJsonFile
		}
	}
	var usernames []string
	_ = json.Unmarshal(configs.ReservedUsernames, &usernames)
	for _, username := range usernames {
		reservedUsernameMapping[username] = true
	}
}

// IsReservedUsername checks whether the username is reserved
func IsReservedUsername(username string) bool {
	reservedUsernameInit.Do(initReservedUsername)
	return reservedUsernameMapping[username]
}
