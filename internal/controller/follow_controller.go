package controller

import (
	"github.com/answerdev/answer/internal/base/handler"
	"github.com/answerdev/answer/internal/base/middleware"
	"github.com/answerdev/answer/internal/schema"
	"github.com/answerdev/answer/internal/service/follow"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

// FollowController activity controller
type FollowController struct {
	followService *follow.FollowService
}

// NewFollowController new controller
func NewFollowController(followService *follow.FollowService) *FollowController {
	return &FollowController{followService: followService}
}

// Follow godoc
// @Summary follow object or cancel follow operation
// @Description follow object or cancel follow operation
// @Tags Activity
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param data body schema.FollowReq true "follow"
// @Success 200 {object} handler.RespBody{data=schema.FollowResp}
// @Router /answer/api/v1/follow [post]
func (fc *FollowController) Follow(ctx *gin.Context) {
	req := &schema.FollowReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}

	dto := &schema.FollowDTO{}
	_ = copier.Copy(dto, req)
	dto.UserID = middleware.GetLoginUserIDFromContext(ctx)

	resp, err := fc.followService.Follow(ctx, dto)
	if err != nil {
		handler.HandleResponse(ctx, err, schema.ErrTypeToast)
	} else {
		handler.HandleResponse(ctx, err, resp)
	}
}

// UpdateFollowTags update user follow tags
// @Summary update user follow tags
// @Description update user follow tags
// @Tags Activity
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param data body schema.UpdateFollowTagsReq true "follow"
// @Success 200 {object} handler.RespBody{}
// @Router /answer/api/v1/follow/tags [put]
func (fc *FollowController) UpdateFollowTags(ctx *gin.Context) {
	req := &schema.UpdateFollowTagsReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}

	req.UserID = middleware.GetLoginUserIDFromContext(ctx)

	err := fc.followService.UpdateFollowTags(ctx, req)
	handler.HandleResponse(ctx, err, nil)
}
