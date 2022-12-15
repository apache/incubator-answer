package converter

import (
	"fmt"
	"github.com/segmentfault/pacman/log"
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

// InterfaceToString converts data to string
// It will be used in template render
func InterfaceToString(data interface{}) string {
	switch t := data.(type) {
	case int:
		i := data.(int)
		return strconv.Itoa(i)
	case int8:
		i := data.(int8)
		return strconv.Itoa(int(i))
	case int16:
		i := data.(int16)
		return strconv.Itoa(int(i))
	case int32:
		i := data.(int32)
		return string(i)
	case int64:
		i := data.(int64)
		return strconv.FormatInt(i, 10)
	case string:
		return data.(string)
	default:
		log.Warn("can't convert type:", t)
	}
	return ""
}
