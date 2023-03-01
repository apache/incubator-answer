package uid

const salt = int64(12345678)

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
func NumToString(id int64) string {
	id = id + salt
	var code []rune
	for id > 0 {
		idx := id % int64(len(AlphanumericSet))
		code = append(code, AlphanumericSet[idx])
		id = id / int64(len(AlphanumericSet))
	}
	return string(code)
}

// StringToNum string to num
func StringToNum(code string) int64 {
	var id int64
	runes := []rune(code)

	for i := len(runes) - 1; i >= 0; i-- {
		ru := runes[i]
		idx := AlphanumericIndex[ru]
		id = id*int64(len(AlphanumericSet)) + int64(idx)
	}
	id = id - salt
	return id
}
