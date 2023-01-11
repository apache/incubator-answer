package controller_admin

import (
	"github.com/answerdev/answer/internal/base/handler"
	"github.com/answerdev/answer/internal/plugin"
	"github.com/answerdev/answer/internal/schema"
	service "github.com/answerdev/answer/internal/service/role"
	"github.com/gin-gonic/gin"
)

// PluginController role controller
type PluginController struct {
	roleService *service.RoleService
}

// NewPluginController new controller
func NewPluginController(roleService *service.RoleService) *PluginController {
	return &PluginController{roleService: roleService}
}

// GetPluginList get plugin list
// @Summary get plugin list
// @Description get plugin list
// @Tags admin
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Success 200 {object} handler.RespBody{data=[]schema.GetCommentResp}
// @Router /answer/admin/api/plugins [get]
func (pc *PluginController) GetPluginList(ctx *gin.Context) {
	resp := make([]*schema.GetPluginListResp, 0)
	_ = plugin.CallBase(func(base plugin.Base) error {
		info := base.Info()
		resp = append(resp, &schema.GetPluginListResp{
			Name:        info.Name,
			Description: info.Description,
			Version:     info.Version,
			Enabled:     plugin.StatusManager.IsEnabled(info.SlugName),
		})
		return nil
	})
	handler.HandleResponse(ctx, nil, resp)
}

// UpdatePluginStatus update plugin status
// @Summary update plugin status
// @Description update plugin status
// @Tags admin
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param data body schema.UpdatePluginStatusReq true "UpdatePluginStatusReq"
// @Success 200 {object} handler.RespBody
// @Router  /answer/api/v1/question/status [put]
func (pc *PluginController) UpdatePluginStatus(ctx *gin.Context) {
	req := &schema.UpdatePluginStatusReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}

	plugin.StatusManager.Enable(req.PluginSlugName, req.Enabled)
	handler.HandleResponse(ctx, nil, nil)
}

// GetPluginConfig get plugin config
func (pc *PluginController) GetPluginConfig(ctx *gin.Context) {
	req := &schema.GetPluginConfigReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}

	resp, err := pc.roleService.GetRoleList(ctx)
	handler.HandleResponse(ctx, err, resp)
}

// UpdatePluginConfig get plugin config
func (pc *PluginController) UpdatePluginConfig(ctx *gin.Context) {
	req := &schema.GetPluginConfigReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}

	resp, err := pc.roleService.GetRoleList(ctx)
	handler.HandleResponse(ctx, err, resp)
}
