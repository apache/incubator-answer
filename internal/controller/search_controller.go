package controller

import (
	"github.com/answerdev/answer/internal/base/handler"
	"github.com/answerdev/answer/internal/base/middleware"
	"github.com/answerdev/answer/internal/base/reason"
	"github.com/answerdev/answer/internal/base/translator"
	"github.com/answerdev/answer/internal/base/validator"
	"github.com/answerdev/answer/internal/entity"
	"github.com/answerdev/answer/internal/schema"
	"github.com/answerdev/answer/internal/service"
	"github.com/answerdev/answer/internal/service/action"
	"github.com/gin-gonic/gin"
	"github.com/segmentfault/pacman/errors"
)

// SearchController tag controller
type SearchController struct {
	searchService *service.SearchService
	actionService *action.CaptchaService
}

// NewSearchController new controller
func NewSearchController(
	searchService *service.SearchService,
	actionService *action.CaptchaService,
) *SearchController {
	return &SearchController{
		searchService: searchService,
		actionService: actionService,
	}
}

// Search godoc
// @Summary search object
// @Description search object
// @Tags Search
// @Produce json
// @Security ApiKeyAuth
// @Param q query string true "query string"
// @Param order query string true "order" Enums(newest,active,score,relevance)
// @Success 200 {object} handler.RespBody{data=schema.SearchListResp}
// @Router /answer/api/v1/search [get]
func (sc *SearchController) Search(ctx *gin.Context) {
	dto := schema.SearchDTO{}

	if handler.BindAndCheck(ctx, &dto) {
		return
	}
	dto.UserID = middleware.GetLoginUserIDFromContext(ctx)
	unit := ctx.ClientIP()
	if dto.UserID != "" {
		unit = dto.UserID
	}
	captchaPass := sc.actionService.ActionRecordVerifyCaptcha(ctx, entity.CaptchaActionSearch, unit, dto.CaptchaID, dto.CaptchaCode)
	if !captchaPass {
		errFields := append([]*validator.FormErrorField{}, &validator.FormErrorField{
			ErrorField: "captcha_code",
			ErrorMsg:   translator.Tr(handler.GetLang(ctx), reason.CaptchaVerificationFailed),
		})
		handler.HandleResponse(ctx, errors.BadRequest(reason.CaptchaVerificationFailed), errFields)
		return
	}

	resp, total, extra, err := sc.searchService.Search(ctx, &dto)
	sc.actionService.ActionRecordAdd(ctx, entity.CaptchaActionSearch, unit)
	handler.HandleResponse(ctx, err, schema.SearchListResp{
		Total:      total,
		SearchResp: resp,
		Extra:      extra,
	})
}
