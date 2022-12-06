package converter

import (
	"fmt"
	"strconv"
)

func StringToInt64(str string) int64 {
	num, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0
	}
	return num
}

func StringToInt(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return num
}

func IntToString(data int64) string {
	return fmt.Sprintf("%d", data)
}

func InterfaceToString(data interface{}) string {
	return fmt.Sprintf("%d", data)
}
