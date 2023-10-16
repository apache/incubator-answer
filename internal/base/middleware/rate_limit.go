package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/answerdev/answer/internal/base/handler"
	"github.com/answerdev/answer/internal/base/reason"
	"github.com/answerdev/answer/internal/repo/limit"
	"github.com/answerdev/answer/pkg/encryption"
	"github.com/gin-gonic/gin"
	"github.com/segmentfault/pacman/errors"
	"github.com/segmentfault/pacman/log"
)

type RateLimitMiddleware struct {
	limitRepo *limit.LimitRepo
}

// NewRateLimitMiddleware new rate limit middleware
func NewRateLimitMiddleware(limitRepo *limit.LimitRepo) *RateLimitMiddleware {
	return &RateLimitMiddleware{
		limitRepo: limitRepo,
	}
}

// DuplicateRequestRejection detects and rejects duplicate requests
// It only works for the requests that post content. Such as add question, add answer, comment etc.
func (rm *RateLimitMiddleware) DuplicateRequestRejection(ctx *gin.Context, req any) bool {
	userID := GetLoginUserIDFromContext(ctx)
	fullPath := ctx.FullPath()
	reqJson, _ := json.Marshal(req)
	key := encryption.MD5(fmt.Sprintf("%s:%s:%s", userID, fullPath, string(reqJson)))
	reject, err := rm.limitRepo.CheckAndRecord(ctx, key)
	if err != nil {
		log.Errorf("check and record rate limit error: %s", err.Error())
		return false
	}
	if !reject {
		return false
	}
	log.Debugf("duplicate request: [%s] %s", fullPath, string(reqJson))
	handler.HandleResponse(ctx, errors.BadRequest(reason.DuplicateRequestError), nil)
	return true
}
