package checker

import (
	"fmt"
	"regexp"
)

const (
	levelD = iota
	LevelC
	LevelB
	LevelA
	LevelS
)

// CheckPassword
// minLength: Specifies the minimum length of a password
// maxLength：Specifies the maximum length of a password
// minLevel：Specifies the minimum strength level required for passwords
// pwd：Text passwords
func CheckPassword(minLength, maxLength, minLevel int, pwd string) error {
	// First check whether the password length is within the range
	if len(pwd) < minLength {
		return fmt.Errorf("BAD PASSWORD: The password is shorter than %d characters", minLength)
	}
	if len(pwd) > maxLength {
		return fmt.Errorf("BAD PASSWORD: The password is logner than %d characters", maxLength)
	}

	// The password strength level is initialized to D.
	// The regular is used to verify the password strength.
	// If the matching is successful, the password strength increases by 1
	level := levelD
	patternList := []string{`[0-9]+`, `[a-z]+`, `[A-Z]+`, `[~!@#$%^&*?_-]+`}
	for _, pattern := range patternList {
		match, _ := regexp.MatchString(pattern, pwd)
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
