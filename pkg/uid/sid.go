package uid

import (
	"strconv"

	"github.com/segmentfault/pacman/utils"
)

const salt = int64(100)

var ShortIDSwitch = false

// NumToString num to string
func NumToShortID(id int64) string {
	sid := strconv.FormatInt(id, 10)
	if len(sid) < 17 {
		return ""
	}
	sTypeCode := sid[1:4]
	sid = sid[4:int32(len(sid))]
	id, err := strconv.ParseInt(sid, 10, 64)
	if err != nil {
		return ""
	}
	typeCode, err := strconv.ParseInt(sTypeCode, 10, 64)
	if err != nil {
		return ""
	}
	code := utils.EnShortID(id, salt)
	tcode := utils.EnShortID(typeCode, salt)
	return string(tcode) + string(code)
}

// StringToNum string to num
func ShortIDToNum(code string) int64 {
	if len(code) < 2 {
		return 0
	}
	scodeType := code[0:2]
	code = code[2:int32(len(code))]

	id := utils.DeShortID(code, salt)
	codeType := utils.DeShortID(scodeType, salt)
	return 10000000000000000 + codeType*10000000000000 + id
}

func EnShortID(id string) string {
	if ShortIDSwitch {
		num, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			return id
		}
		return NumToShortID(num)
	}
	return id
}

func DeShortID(sid string) string {
	num, err := strconv.ParseInt(sid, 10, 64)
	if err != nil {
		return strconv.FormatInt(ShortIDToNum(sid), 10)
	}
	if num < 10000000000000000 {
		return strconv.FormatInt(ShortIDToNum(sid), 10)
	}
	return sid
}

func IsShortID(id string) bool {
	num, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return true
	}
	if num < 10000000000000000 {
		return true
	}
	return false
}
