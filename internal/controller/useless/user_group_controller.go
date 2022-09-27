package useless

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/segmentfault/answer/internal/base/handler"
	"github.com/segmentfault/answer/internal/base/reason"
	"github.com/segmentfault/answer/internal/schema"
	"github.com/segmentfault/answer/internal/service"
	"github.com/segmentfault/pacman/errors"
	"github.com/segmentfault/pacman/log"
)

// UserGroupController userGroup controller
type UserGroupController struct {
	log              log.log
	userGroupService *service.UserGroupService
}

// NewUserGroupController new controller
func NewUserGroupController(userGroupService *service.UserGroupService) *UserGroupController {
	return &UserGroupController{userGroupService: userGroupService}
}

// AddUserGroup add user group
// @Summary add user group
// @Description add user group
// @Tags UserGroup
// @Accept json
// @Produce json
// @Param data body schema.AddUserGroupReq true "user group"
// @Success 200 {object} handler.RespBody
// Router /user-group [post]
func (uc *UserGroupController) AddUserGroup(ctx *gin.Context) {
	req := &schema.AddUserGroupReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}

	err := uc.userGroupService.AddUserGroup(ctx, req)
	handler.HandleResponse(ctx, err, nil)
}

// RemoveUserGroup delete user group
// @Summary delete user group
// @Description delete user group
// @Tags UserGroup
// @Accept json
// @Produce json
// @Param data body schema.RemoveUserGroupReq true "user group"
// @Success 200 {object} handler.RespBody
// Router /user-group [delete]
func (uc *UserGroupController) RemoveUserGroup(ctx *gin.Context) {
	req := &schema.RemoveUserGroupReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}

	err := uc.userGroupService.RemoveUserGroup(ctx, int(req.ID))
	handler.HandleResponse(ctx, err, nil)
}

// UpdateUserGroup update user group
// @Summary update user group
// @Description update user group
// @Tags UserGroup
// @Accept json
// @Produce json
// @Param data body schema.UpdateUserGroupReq true "user group"
// @Success 200 {object} handler.RespBody
// Router /user-group [put]
func (uc *UserGroupController) UpdateUserGroup(ctx *gin.Context) {
	req := &schema.UpdateUserGroupReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}

	err := uc.userGroupService.UpdateUserGroup(ctx, req)
	handler.HandleResponse(ctx, err, nil)
}

// GetUserGroup get user group one
// @Summary get user group one
// @Description get user group one
// @Tags UserGroup
// @Accept json
// @Produce json
// @Param id path int true "user groupid"
// @Success 200 {object} handler.RespBody{data=schema.GetUserGroupResp}
// Router /user-group/{id} [get]
func (uc *UserGroupController) GetUserGroup(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if id == 0 {
		handler.HandleResponse(ctx, errors.BadRequest(reason.RequestFormatError), nil)
		return
	}

	resp, err := uc.userGroupService.GetUserGroup(ctx, id)
	handler.HandleResponse(ctx, err, resp)
}

// GetUserGroupList get user group list
// @Summary get user group list
// @Description get user group list
// @Tags UserGroup
// @Produce json
// @Success 200 {object} handler.RespBody{data=[]schema.GetUserGroupResp}
// Router /user-groups [get]
func (uc *UserGroupController) GetUserGroupList(ctx *gin.Context) {
	req := &schema.GetUserGroupListReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}

	resp, err := uc.userGroupService.GetUserGroupList(ctx, req)
	handler.HandleResponse(ctx, err, resp)
}

// GetUserGroupWithPage get user group page
// @Summary get user group page
// @Description get user group page
// @Tags UserGroup
// @Produce json
// @Param page query int false "page size"
// @Param page_size query int false "page size"
// @Success 200 {object} handler.RespBody{data=pager.PageModel{list=[]schema.GetUserGroupResp}}
// Router /user-groups/page [get]
func (uc *UserGroupController) GetUserGroupWithPage(ctx *gin.Context) {
	req := &schema.GetUserGroupWithPageReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}

	resp, err := uc.userGroupService.GetUserGroupWithPage(ctx, req)
	handler.HandleResponse(ctx, err, resp)
}
