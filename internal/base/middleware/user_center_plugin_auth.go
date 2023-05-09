package middleware

import (
	"github.com/answerdev/answer/internal/base/handler"
	"github.com/answerdev/answer/internal/base/reason"
	"github.com/answerdev/answer/plugin"
	"github.com/gin-gonic/gin"
	"github.com/segmentfault/pacman/errors"
)

// BanAPIWhenUserCenterEnabled ban api when user center enabled
func BanAPIWhenUserCenterEnabled(ctx *gin.Context) {
	if plugin.UserCenterEnabled() {
		handler.HandleResponse(ctx, errors.Forbidden(reason.ForbiddenError), nil)
		ctx.Abort()
		return
	}
	ctx.Next()
}
