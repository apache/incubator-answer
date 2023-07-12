package checker

import (
	"sync"

	"github.com/answerdev/answer/configs"
	"github.com/segmentfault/pacman/log"
	"gopkg.in/yaml.v3"
)

type PathIgnore struct {
	Users     []string `yaml:"users"`
	Questions []string `yaml:"questions"`
}

var (
	ignorePathInit sync.Once
	pathIgnore     = &PathIgnore{}
)

func initPathIgnore() {
	if err := yaml.Unmarshal(configs.PathIgnore, pathIgnore); err != nil {
		log.Error(err)
	}
}

// IsUsersIgnorePath checks whether the username is in ignore path
func IsUsersIgnorePath(username string) bool {
	ignorePathInit.Do(initPathIgnore)
	for _, u := range pathIgnore.Users {
		if u == username {
			return true
		}
	}
	return false
}

// IsQuestionsIgnorePath checks whether the questionID is in ignore path
func IsQuestionsIgnorePath(questionID string) bool {
	ignorePathInit.Do(initPathIgnore)
	for _, u := range pathIgnore.Questions {
		if u == questionID {
			return true
		}
	}
	return false
}
