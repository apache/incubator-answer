package dir

import "os"

func CreateDirIfNotExist(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

func CheckDirExist(path string) bool {
	f, err := os.Stat(path)
	return err == nil && f.IsDir()
}

func CheckFileExist(path string) bool {
	f, err := os.Stat(path)
	return err == nil && !f.IsDir()
}
