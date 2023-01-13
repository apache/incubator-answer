package controller_admin

import (
	"encoding/json"

	"github.com/answerdev/answer/internal/base/handler"
	"github.com/answerdev/answer/internal/plugin"
	"github.com/answerdev/answer/internal/schema"
	"github.com/gin-gonic/gin"
)

// PluginController role controller
type PluginController struct {
}

// NewPluginController new controller
func NewPluginController() *PluginController {
	return &PluginController{}
}

// GetPluginList get plugin list
// @Summary get plugin list
// @Description get plugin list
// @Tags AdminPlugin
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param status query string false "status: active/inactive"
// @Param have_config query boolean false "have config"
// @Success 200 {object} handler.RespBody{data=[]schema.GetPluginListResp}
// @Router /answer/admin/api/plugins [get]
func (pc *PluginController) GetPluginList(ctx *gin.Context) {
	req := &schema.GetPluginListReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}

	pluginConfigMapping := make(map[string]bool)
	_ = plugin.CallConfig(func(fn plugin.Config) error {
		if len(fn.ConfigFields()) > 0 {
			pluginConfigMapping[fn.Info().SlugName] = true
		}
		return nil
	})

	resp := make([]*schema.GetPluginListResp, 0)
	_ = plugin.CallBase(func(base plugin.Base) error {
		info := base.Info()
		resp = append(resp, &schema.GetPluginListResp{
			Name:        info.Name,
			SlugName:    info.SlugName,
			Description: info.Description,
			Version:     info.Version,
			Enabled:     plugin.StatusManager.IsEnabled(info.SlugName),
			HaveConfig:  pluginConfigMapping[info.SlugName],
		})
		return nil
	})

	if len(req.Status) > 0 {
		resp = pc.filterPluginByStatus(resp, req.Status)
	}
	if req.HaveConfig {
		resp = pc.filterNoConfigPlugin(resp)
	}
	handler.HandleResponse(ctx, nil, resp)
}

func (pc *PluginController) filterNoConfigPlugin(list []*schema.GetPluginListResp) []*schema.GetPluginListResp {
	resp := make([]*schema.GetPluginListResp, 0)
	for _, t := range list {
		if t.HaveConfig {
			resp = append(resp, t)
		}
	}
	return resp
}

func (pc *PluginController) filterPluginByStatus(list []*schema.GetPluginListResp, status schema.PluginStatus,
) []*schema.GetPluginListResp {
	resp := make([]*schema.GetPluginListResp, 0)
	for _, t := range list {
		if status == schema.PluginStatusActive && t.Enabled {
			resp = append(resp, t)
		} else if status == schema.PluginStatusInactive && !t.Enabled {
			resp = append(resp, t)
		}
	}
	return resp
}

// UpdatePluginStatus update plugin status
// @Summary update plugin status
// @Description update plugin status
// @Tags AdminPlugin
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param data body schema.UpdatePluginStatusReq true "UpdatePluginStatusReq"
// @Success 200 {object} handler.RespBody
// @Router /answer/admin/api/plugin/status [put]
func (pc *PluginController) UpdatePluginStatus(ctx *gin.Context) {
	req := &schema.UpdatePluginStatusReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}

	plugin.StatusManager.Enable(req.PluginSlugName, req.Enabled)
	handler.HandleResponse(ctx, nil, nil)
}

// GetPluginConfig get plugin config
// @Summary get plugin config
// @Description get plugin config
// @Tags AdminPlugin
// @Security ApiKeyAuth
// @Produce  json
// @Param plugin_slug_name query string true "plugin_slug_name"
// @Success 200 {object} handler.RespBody{data=schema.GetPluginConfigResp}
// @Router /answer/admin/api/plugin/config [get]
func (pc *PluginController) GetPluginConfig(ctx *gin.Context) {
	req := &schema.GetPluginConfigReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}

	resp := &schema.GetPluginConfigResp{}
	_ = plugin.CallBase(func(base plugin.Base) error {
		if base.Info().SlugName != req.PluginSlugName {
			return nil
		}
		info := base.Info()
		resp.Name = info.Name
		resp.SlugName = info.SlugName
		resp.Description = info.Description
		resp.Version = info.Version
		return nil
	})

	_ = plugin.CallConfig(func(fn plugin.Config) error {
		if fn.Info().SlugName != req.PluginSlugName {
			return nil
		}
		// resp.ConfigFields = fn.ConfigFields()
		configFields := fn.ConfigFields()
		resp.ConfigFields = make([]*schema.ConfigField, 0)
		for _, field := range configFields {
			resp.ConfigFields = append(resp.ConfigFields, &schema.ConfigField{
				Name:        field.Items[0].Name,
				Type:        schema.TEXT,
				Title:       field.Name,
				Description: field.Description,
				Required:    field.Required,
				UIOptions: schema.UIOptions{
					Placeholder: field.Items[0].PlaceHolder,
					Type:        string(schema.TEXT),
				},
				Value: field.Items[0].Value,
			})
		}
		return nil
	})
	handler.HandleResponse(ctx, nil, resp)
}

// UpdatePluginConfig update plugin config
// @Summary update plugin config
// @Description update plugin config
// @Tags AdminPlugin
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param data body schema.UpdatePluginConfigReq true "UpdatePluginConfigReq"
// @Success 200 {object} handler.RespBody
// @Router /answer/admin/api/plugin/config [put]
func (pc *PluginController) UpdatePluginConfig(ctx *gin.Context) {
	req := &schema.UpdatePluginConfigReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}

	configFields, _ := json.Marshal(req.ConfigFields)
	err := plugin.CallConfig(func(fn plugin.Config) error {
		if fn.Info().SlugName == req.PluginSlugName {
			return fn.ConfigReceiver(configFields)
		}
		return nil
	})
	handler.HandleResponse(ctx, err, nil)
}
