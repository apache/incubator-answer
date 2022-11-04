package writer

import (
	"bufio"
	"os"
)

// ReplaceFile remove old file and write new file
func ReplaceFile(filePath, content string) error {
	_ = os.Remove(filePath)
	return WriteFile(filePath, content)
}

// WriteFile write file to path
func WriteFile(filePath, content string) error {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0o666)
	if err != nil {
		return err
	}
	defer func() {
		_ = file.Close()
	}()
	writer := bufio.NewWriter(file)
	if _, err := writer.WriteString(content); err != nil {
		return err
	}
	if err := writer.Flush(); err != nil {
		return err
	}
	return nil
}
