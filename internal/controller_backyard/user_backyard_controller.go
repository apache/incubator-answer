package controller_backyard

import (
	"github.com/gin-gonic/gin"
	"github.com/segmentfault/answer/internal/base/handler"
	"github.com/segmentfault/answer/internal/base/reason"
	"github.com/segmentfault/answer/internal/schema"
	"github.com/segmentfault/answer/internal/service/user_backyard"
	"github.com/segmentfault/pacman/errors"
)

// UserBackyardController user controller
type UserBackyardController struct {
	userService *user_backyard.UserBackyardService
}

// NewUserBackyardController new controller
func NewUserBackyardController(userService *user_backyard.UserBackyardService) *UserBackyardController {
	return &UserBackyardController{userService: userService}
}

// UpdateUserStatus update user
// @Summary update user
// @Description update user
// @Security ApiKeyAuth
// @Tags admin
// @Accept json
// @Produce json
// @Param data body schema.UpdateUserStatusReq true "user"
// @Success 200 {object} handler.RespBody
// @Router /answer/admin/api/user/status [put]
func (uc *UserBackyardController) UpdateUserStatus(ctx *gin.Context) {
	req := &schema.UpdateUserStatusReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}

	err := uc.userService.UpdateUserStatus(ctx, req)
	handler.HandleResponse(ctx, err, nil)
}

// GetUserInfo get user one
// @Summary get user one
// @Description get user one
// @Security ApiKeyAuth
// @Tags admin
// @Accept json
// @Produce json
// @Param id path int true "userid"
// @Success 200 {object} handler.RespBody{data=schema.GetUserInfoResp}
// Router /user/{id} [get]
func (uc *UserBackyardController) GetUserInfo(ctx *gin.Context) {
	userID := ctx.Query("user_id")
	if len(userID) == 0 {
		handler.HandleResponse(ctx, errors.BadRequest(reason.RequestFormatError), nil)
		return
	}

	resp, err := uc.userService.GetUserInfo(ctx, userID)
	handler.HandleResponse(ctx, err, resp)
}

// GetUserPage get user page
// @Summary get user page
// @Description get user page
// @Security ApiKeyAuth
// @Tags admin
// @Produce json
// @Param page query int false "page size"
// @Param page_size query int false "page size"
// @Param username query string false "username"
// @Param e_mail query string false "email"
// @Param status query string false "user status" Enums(normal, suspended, deleted, inactive)
// @Success 200 {object} handler.RespBody{data=pager.PageModel{records=[]schema.GetUserPageResp}}
// @Router /answer/admin/api/users/page [get]
func (uc *UserBackyardController) GetUserPage(ctx *gin.Context) {
	req := &schema.GetUserPageReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}

	resp, err := uc.userService.GetUserPage(ctx, req)
	handler.HandleResponse(ctx, err, resp)
}
