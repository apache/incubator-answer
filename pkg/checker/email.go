package checker

import "strings"

func EmailInAllowEmailDomain(email string, allowEmailDomains []string) bool {
	if len(allowEmailDomains) == 0 {
		return true
	}

	for _, domain := range allowEmailDomains {
		if strings.HasSuffix(email, domain) {
			return true
		}
	}

	return false
}
