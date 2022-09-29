package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/segmentfault/answer/internal/base/constant"
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
