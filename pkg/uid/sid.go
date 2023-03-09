package uid

import (
	"strconv"
)

const salt = int64(0)

var ShortIDSwitch = false

var AlphanumericSet = []rune{
	'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
	'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
	'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
}

var AlphanumericIndex map[rune]int

func init() {
	AlphanumericIndex = make(map[rune]int, len(AlphanumericSet))
	for i, ru := range AlphanumericSet {
		AlphanumericIndex[ru] = i
	}
}

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
	id = id + salt
	// fmt.Println("[EN1]", typeCode, id)
	var code []rune
	var tcode []rune
	for id > 0 {
		idx := id % int64(len(AlphanumericSet))
		code = append(code, AlphanumericSet[idx])
		id = id / int64(len(AlphanumericSet))
	}
	for typeCode > 0 {
		idx := typeCode % int64(len(AlphanumericSet))
		tcode = append(tcode, AlphanumericSet[idx])
		typeCode = typeCode / int64(len(AlphanumericSet))
	}
	// fmt.Println("[EN2]", string(tcode), string(code))
	return string(tcode) + string(code)
}

// StringToNum string to num
func ShortIDToNum(code string) int64 {
	if len(code) < 2 {
		return 0
	}
	scodeType := code[0:1]
	code = code[1:int32(len(code))]
	// fmt.Println("[DE1]", scodeType, code)
	var id, codeType int64
	runes := []rune(code)
	codeRunes := []rune(scodeType)

	for i := len(runes) - 1; i >= 0; i-- {
		ru := runes[i]
		idx := AlphanumericIndex[ru]
		id = id*int64(len(AlphanumericSet)) + int64(idx)
	}
	for i := len(codeRunes) - 1; i >= 0; i-- {
		ru := codeRunes[i]
		idx := AlphanumericIndex[ru]
		codeType = codeType*int64(len(AlphanumericSet)) + int64(idx)
	}
	id = id - salt
	// fmt.Println("[DE2]", codeType, id)

	return 10000000000000000 + codeType*10000000000000 + id
}

func EnShortID(id string) string {
	if ShortIDSwitch {
		num, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			return ""
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
