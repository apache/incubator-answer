package controller

import (
	"github.com/answerdev/answer/internal/base/handler"
	"github.com/answerdev/answer/internal/base/middleware"
	"github.com/answerdev/answer/internal/schema"
	"github.com/answerdev/answer/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

// VoteController activity controller
type VoteController struct {
	VoteService *service.VoteService
}

// NewVoteController new controller
func NewVoteController(voteService *service.VoteService) *VoteController {
	return &VoteController{VoteService: voteService}
}

// VoteUp godoc
// @Summary vote up
// @Description add vote
// @Tags Activity
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param data body schema.VoteReq true "vote"
// @Success 200 {object} handler.RespBody{data=schema.VoteResp}
// @Router /answer/api/v1/vote/up [post]
func (vc *VoteController) VoteUp(ctx *gin.Context) {
	req := &schema.VoteReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}
	dto := &schema.VoteDTO{}
	_ = copier.Copy(dto, req)
	dto.UserID = middleware.GetLoginUserIDFromContext(ctx)
	resp, err := vc.VoteService.VoteUp(ctx, dto)
	if err != nil {
		handler.HandleResponse(ctx, err, schema.ErrTypeToast)
	} else {
		handler.HandleResponse(ctx, err, resp)
	}
}

// VoteDown godoc
// @Summary vote down
// @Description add vote
// @Tags Activity
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param data body schema.VoteReq true "vote"
// @Success 200 {object} handler.RespBody{data=schema.VoteResp}
// @Router /answer/api/v1/vote/down [post]
func (vc *VoteController) VoteDown(ctx *gin.Context) {
	req := &schema.VoteReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}
	dto := &schema.VoteDTO{}
	_ = copier.Copy(dto, req)

	dto.UserID = middleware.GetLoginUserIDFromContext(ctx)
	resp, err := vc.VoteService.VoteDown(ctx, dto)
	if err != nil {
		handler.HandleResponse(ctx, err, schema.ErrTypeToast)
	} else {
		handler.HandleResponse(ctx, err, resp)
	}
}

// UserVotes godoc
// @Summary user's votes
// @Description user's vote
// @Tags Activity
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param page query int false "page size"
// @Param page_size query int false "page size"
// @Success 200 {object} handler.RespBody{data=pager.PageModel{list=[]schema.GetVoteWithPageResp}}
// @Router /answer/api/v1/personal/vote/page [get]
func (vc *VoteController) UserVotes(ctx *gin.Context) {
	req := schema.GetVoteWithPageReq{}
	req.UserID = middleware.GetLoginUserIDFromContext(ctx)
	if handler.BindAndCheck(ctx, &req) {
		return
	}
	if req.Page == 0 {
		req.Page = 1
	}
	if req.PageSize == 0 {
		req.PageSize = 30
	}

	resp, err := vc.VoteService.ListUserVotes(ctx, req)
	if err != nil {
		handler.HandleResponse(ctx, err, schema.ErrTypeModal)
	} else {
		handler.HandleResponse(ctx, err, resp)
	}
}
