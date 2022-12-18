package day

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestFormat(t *testing.T) {
	sec := time.Now().Unix()
	tz := "Asia/Shanghai"
	actual := Format(sec, "YYYY-MM-DD HH:mm:ss", tz)
	_, _ = time.LoadLocation(tz)
	expected := time.Unix(sec, 0).Format("2006-01-02 15:04:05")
	assert.Equal(t, expected, actual)
}
