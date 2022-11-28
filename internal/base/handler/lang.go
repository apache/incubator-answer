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
	switch i18n.Language(acceptLanguage) {
	case i18n.LanguageChinese:
		return i18n.LanguageChinese
	case i18n.LanguageEnglish:
		return i18n.LanguageEnglish
	default:
		return i18n.DefaultLang
	}
}

// GetLangByCtx get language from header
func GetLangByCtx(ctx context.Context) i18n.Language {
	acceptLanguage, ok := ctx.Value(constant.AcceptLanguageFlag).(i18n.Language)
	if ok {
		return acceptLanguage
	}
	return i18n.DefaultLang
}
