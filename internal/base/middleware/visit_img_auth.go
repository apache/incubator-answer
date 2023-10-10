package middleware

import (
	"github.com/answerdev/answer/internal/base/constant"
	"github.com/gin-gonic/gin"
	"net/http"
)

// VisitAuth when user visit the site image, check visit token. This only for private mode.
func (am *AuthUserMiddleware) VisitAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		siteLogin, err := am.siteInfoCommonService.GetSiteLogin(ctx)
		if err != nil {
			return
		}
		if !siteLogin.LoginRequired {
			ctx.Next()
			return
		}

		visitToken, err := ctx.Cookie(constant.UserVisitCookiesCacheKey)
		if err != nil || len(visitToken) == 0 {
			ctx.Abort()
			ctx.Redirect(http.StatusFound, "/403")
			return
		}

		if !am.authService.CheckUserVisitToken(ctx, visitToken) {
			ctx.Abort()
			ctx.Redirect(http.StatusFound, "/403")
			return
		}
	}
}
