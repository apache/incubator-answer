package middleware

import (
	"github.com/answerdev/answer/internal/base/constant"
	"github.com/answerdev/answer/internal/base/handler"
	"github.com/gin-gonic/gin"
)

// ExtractAndSetAcceptLanguage extract accept language from header and set to context
func ExtractAndSetAcceptLanguage(ctx *gin.Context) {
	lang := handler.GetLang(ctx)
	ctx.Set(constant.AcceptLanguageFlag, lang)
}
