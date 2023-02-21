package checker

import "regexp"

var (
	usernameReg = regexp.MustCompile(`^[a-z0-9._-]{4,30}$`)
)

func IsInvalidUsername(username string) bool {
	return !usernameReg.MatchString(username)
}
