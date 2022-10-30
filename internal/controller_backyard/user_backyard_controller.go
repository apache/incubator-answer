package controller_backyard

import (
	"github.com/answerdev/answer/internal/base/handler"
	"github.com/answerdev/answer/internal/schema"
	"github.com/answerdev/answer/internal/service/user_backyard"
	"github.com/gin-gonic/gin"
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
// @Summary     update user
// @Description update user
// @Security    ApiKeyAuth
// @Tags        admin
// @Accept      json
// @Produce     json
// @Param       data body     schema.UpdateUserStatusReq true "user"
// @Success     200  {object} handler.RespBody
// @Router      /answer/admin/api/user/status [put]
func (uc *UserBackyardController) UpdateUserStatus(ctx *gin.Context) {
	req := &schema.UpdateUserStatusReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}

	err := uc.userService.UpdateUserStatus(ctx, req)
	handler.HandleResponse(ctx, err, nil)
}

// GetUserPage get user page
// @Summary     get user page
// @Description get user page
// @Security    ApiKeyAuth
// @Tags        admin
// @Produce     json
// @Param       page      query    int    false "page size"
// @Param       page_size query    int    false "page size"
// @Param       username  query    string false "username"
// @Param       e_mail    query    string false "email"
// @Param       status    query    string false "user status" Enums(normal, suspended, deleted, inactive)
// @Success     200       {object} handler.RespBody{data=pager.PageModel{records=[]schema.GetUserPageResp}}
// @Router      /answer/admin/api/users/page [get]
func (uc *UserBackyardController) GetUserPage(ctx *gin.Context) {
	req := &schema.GetUserPageReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}

	resp, err := uc.userService.GetUserPage(ctx, req)
	handler.HandleResponse(ctx, err, resp)
}
