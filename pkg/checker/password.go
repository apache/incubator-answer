package checker

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	levelD = iota
	LevelC
	LevelB
	LevelA
	LevelS
)

const (
	PasswordCannotContainSpaces = "error.password.space_invalid"
)

// CheckPassword checks the password strength
func CheckPassword(password string) error {
	if strings.Contains(password, " ") {
		return fmt.Errorf(PasswordCannotContainSpaces)
	}

	// TODO Currently there is no requirement for password strength
	minLevel := 0

	// The password strength level is initialized to D.
	// The regular is used to verify the password strength.
	// If the matching is successful, the password strength increases by 1
	level := levelD
	patternList := []string{`[0-9]+`, `[a-z]+`, `[A-Z]+`, `[~!@#$%^&*?_-]+`}
	for _, pattern := range patternList {
		match, _ := regexp.MatchString(pattern, password)
		if match {
			level++
		}
	}

	// If the final password strength falls below the required minimum strength, return with an error
	if level < minLevel {
		return fmt.Errorf("the password does not satisfy the current policy requirements")
	}
	return nil
}
