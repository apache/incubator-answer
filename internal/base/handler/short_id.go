package handler

import (
	"context"

	"github.com/answerdev/answer/internal/base/constant"
)

// GetEnableShortID get language from header
func GetEnableShortID(ctx context.Context) bool {
	flag, ok := ctx.Value(constant.ShortIDFlag).(bool)
	if ok {
		return flag
	}
	return false
}
