package uid

import (
	"fmt"
	"testing"
)

func Test_getInviteCodeById(t *testing.T) {
	nums := []int64{0, 1, 10, 100, 1000, 10000, 100000, 10010000000001316, 10030000000001316}
	for _, num := range nums {
		code := NumToString(num)
		denum := StringToNum(code)
		fmt.Println(num, code, denum)
	}
}
