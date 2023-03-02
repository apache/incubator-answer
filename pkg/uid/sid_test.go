package uid

import (
	"fmt"
	"testing"
)

func Test_ShortID(t *testing.T) {
	nums := []int64{0, 1, 10, 100, 1000, 10000, 100000, 10010000000001316, 10030000000001316, 999999999999999999, 1999999999999999999}
	for _, num := range nums {
		code := NumToShortID(num)
		denum := ShortIDToNum(code)
		fmt.Println(num, code, denum)
	}
}

func Test_EnDeShortID(t *testing.T) {
	nums := []string{"0", "1", "10", "100", "1000", "10000", "100000", "1234567", "10010000000001316", "10030000000001316", "99999999999999999", "999999999999999999", "1999999999999999999"}
	for _, num := range nums {
		code := EnShortID(num)
		denum := DeShortID(code)
		fmt.Println(num, code, denum)
	}
}
