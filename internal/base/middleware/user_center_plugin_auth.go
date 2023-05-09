package middleware

import (
	"github.com/answerdev/answer/internal/base/handler"
	"github.com/answerdev/answer/internal/base/reason"
	"github.com/answerdev/answer/plugin"
	"github.com/gin-gonic/gin"
	"github.com/segmentfault/pacman/errors"
)

// BanAPIForUserCenter ban api for user center
func BanAPIForUserCenter(ctx *gin.Context) {
	uc, ok := plugin.GetUserCenter()
	if !ok {
		return
	}
	if !uc.Description().EnabledOriginalUserSystem {
		handler.HandleResponse(ctx, errors.Forbidden(reason.ForbiddenError), nil)
		ctx.Abort()
		return
	}
	ctx.Next()
}
