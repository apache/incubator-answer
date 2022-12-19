package handler

import (
	"context"

	"github.com/answerdev/answer/internal/base/constant"
	"github.com/gin-gonic/gin"
	"github.com/segmentfault/pacman/i18n"
)

// GetLang get language from header
func GetLang(ctx *gin.Context) i18n.Language {
	acceptLanguage := ctx.GetHeader(constant.AcceptLanguageFlag)
	if len(acceptLanguage) == 0 {
		return i18n.DefaultLanguage
	}
	return i18n.Language(acceptLanguage)
}

// GetLangByCtx get language from header
func GetLangByCtx(ctx context.Context) i18n.Language {
	acceptLanguage, ok := ctx.Value(constant.AcceptLanguageFlag).(i18n.Language)
	if ok {
		return acceptLanguage
	}
	return i18n.DefaultLanguage
}
