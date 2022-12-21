package controller_admin

import (
	"github.com/answerdev/answer/internal/base/handler"
	"github.com/answerdev/answer/internal/base/middleware"
	"github.com/answerdev/answer/internal/schema"
	"github.com/answerdev/answer/internal/service/user_admin"
	"github.com/gin-gonic/gin"
)

// UserAdminController user controller
type UserAdminController struct {
	userService *user_admin.UserAdminService
}

// NewUserAdminController new controller
func NewUserAdminController(userService *user_admin.UserAdminService) *UserAdminController {
	return &UserAdminController{userService: userService}
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
func (uc *UserAdminController) UpdateUserStatus(ctx *gin.Context) {
	req := &schema.UpdateUserStatusReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}

	err := uc.userService.UpdateUserStatus(ctx, req)
	handler.HandleResponse(ctx, err, nil)
}

// UpdateUserRole update user role
// @Summary update user role
// @Description update user role
// @Security ApiKeyAuth
// @Tags admin
// @Accept json
// @Produce json
// @Param data body schema.UpdateUserRoleReq true "user"
// @Success 200 {object} handler.RespBody
// @Router /answer/admin/api/user/role [put]
func (uc *UserAdminController) UpdateUserRole(ctx *gin.Context) {
	req := &schema.UpdateUserRoleReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}

	req.LoginUserID = middleware.GetLoginUserIDFromContext(ctx)

	err := uc.userService.UpdateUserRole(ctx, req)
	handler.HandleResponse(ctx, err, nil)
}

// AddUser add user
// @Summary add user
// @Description add user
// @Security ApiKeyAuth
// @Tags admin
// @Accept json
// @Produce json
// @Param data body schema.AddUserReq true "user"
// @Success 200 {object} handler.RespBody
// @Router /answer/admin/api/user [post]
func (uc *UserAdminController) AddUser(ctx *gin.Context) {
	req := &schema.AddUserReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}

	req.LoginUserID = middleware.GetLoginUserIDFromContext(ctx)

	err := uc.userService.AddUser(ctx, req)
	handler.HandleResponse(ctx, err, nil)
}

// UpdateUserPassword update user password
// @Summary update user password
// @Description update user password
// @Security ApiKeyAuth
// @Tags admin
// @Accept json
// @Produce json
// @Param data body schema.UpdateUserPasswordReq true "user"
// @Success 200 {object} handler.RespBody
// @Router /answer/admin/api/user/password [put]
func (uc *UserAdminController) UpdateUserPassword(ctx *gin.Context) {
	req := &schema.UpdateUserPasswordReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}

	req.LoginUserID = middleware.GetLoginUserIDFromContext(ctx)

	err := uc.userService.UpdateUserPassword(ctx, req)
	handler.HandleResponse(ctx, err, nil)
}

// GetUserPage get user page
// @Summary get user page
// @Description get user page
// @Security ApiKeyAuth
// @Tags admin
// @Produce json
// @Param page query int false "page size"
// @Param page_size query int false "page size"
// @Param query query string false "search query: email, username or id:[id]"
// @Param staff query bool false "staff user"
// @Param status query string false "user status" Enums(suspended, deleted, inactive)
// @Success 200 {object} handler.RespBody{data=pager.PageModel{records=[]schema.GetUserPageResp}}
// @Router /answer/admin/api/users/page [get]
func (uc *UserAdminController) GetUserPage(ctx *gin.Context) {
	req := &schema.GetUserPageReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}

	resp, err := uc.userService.GetUserPage(ctx, req)
	handler.HandleResponse(ctx, err, resp)
}
