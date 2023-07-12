package controller

import (
	"github.com/answerdev/answer/internal/base/handler"
	"github.com/answerdev/answer/internal/base/middleware"
	"github.com/answerdev/answer/internal/schema"
	"github.com/answerdev/answer/internal/service/rank"
	"github.com/gin-gonic/gin"
)

// RankController rank controller
type RankController struct {
	rankService *rank.RankService
}

// NewRankController new controller
func NewRankController(
	rankService *rank.RankService) *RankController {
	return &RankController{rankService: rankService}
}

// GetRankPersonalWithPage user personal rank list
// @Summary user personal rank list
// @Description user personal rank list
// @Tags Rank
// @Produce json
// @Param page query int false "page"
// @Param page_size query int false "page size"
// @Param username query string false "username"
// @Success 200 {object} handler.RespBody{data=pager.PageModel{list=[]schema.GetRankPersonalPageResp}}
// @Router /answer/api/v1/personal/rank/page [get]
func (cc *RankController) GetRankPersonalWithPage(ctx *gin.Context) {
	req := &schema.GetRankPersonalWithPageReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}

	req.UserID = middleware.GetLoginUserIDFromContext(ctx)

	resp, err := cc.rankService.GetRankPersonalPage(ctx, req)
	handler.HandleResponse(ctx, err, resp)
}
