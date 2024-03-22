package checker

import (
	"net/url"
	"strings"
)

func IsURL(str string) bool {
	s := strings.ToLower(str)

	if len(s) == 0 {
		return false
	}

	u, err := url.Parse(s)
	if err != nil || u.Scheme == "" {
		return false
	}

	if u.Host == "" && u.Fragment == "" && u.Opaque == "" {
		return false
	}
	return u.Scheme == "http" || u.Scheme == "https"
}
