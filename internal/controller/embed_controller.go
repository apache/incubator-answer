package controller

import (
	"github.com/apache/incubator-answer/internal/base/handler"
	"github.com/apache/incubator-answer/internal/schema"
	"github.com/apache/incubator-answer/plugin"
	"github.com/gin-gonic/gin"
)

type EmbedController struct {
}

func NewEmbedController() *EmbedController {
	return &EmbedController{}
}

// GetEmbedConfig godoc
// @Summary GetEmbedConfig
// @Description GetEmbedConfig
// @Tags PluginEmbed
// @Accept json
// @Produce json
// @Router /answer/api/v1/embed/config [get]
// @Success 200 {object} handler.RespBody{data=[]schema.GetEmbedOptionResp}
func (c *EmbedController) GetEmbedConfig(ctx *gin.Context) {
	resp := make([]*schema.GetEmbedOptionResp, 0)
	var slugName string

	_ = plugin.CallEmbed(func(base plugin.Embed) error {
		slugName = base.Info().SlugName
		return nil
	})

	_ = plugin.CallConfig(func(fn plugin.Config) error {
		if fn.Info().SlugName == slugName {
			for _, field := range fn.ConfigFields() {
				resp = append(resp, &schema.GetEmbedOptionResp{
					Platform: field.Name,
					Enable:   field.Value.(bool),
				})
			}
			return nil
		}
		return nil
	})
	handler.HandleResponse(ctx, nil, resp)
}
