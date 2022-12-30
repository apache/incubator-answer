package checker

// IsNotZeroString check s is not empty string and is not "0"
func IsNotZeroString(s string) bool {
	return len(s) > 0 && s != "0"
}
