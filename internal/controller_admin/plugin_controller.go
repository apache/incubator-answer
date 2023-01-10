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
func (pc *PluginController) GetPluginList(ctx *gin.Context) {
	plugin.CallBase(func(base plugin.Base) error {
		base.Info()
		return nil
	})

	resp, err := pc.roleService.GetRoleList(ctx)
	handler.HandleResponse(ctx, err, resp)
}

// UpdatePluginStatus update plugin status
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
	resp, err := pc.roleService.GetRoleList(ctx)
	handler.HandleResponse(ctx, err, resp)
}

// UpdatePluginConfig get plugin config
func (pc *PluginController) UpdatePluginConfig(ctx *gin.Context) {
	resp, err := pc.roleService.GetRoleList(ctx)
	handler.HandleResponse(ctx, err, resp)
}
