package controller_backyard

import (
	"github.com/gin-gonic/gin"
	"github.com/segmentfault/answer/internal/base/handler"
	"github.com/segmentfault/answer/internal/schema"
)

type ThemeController struct{}

// NewThemeController new theme controller.
func NewThemeController() *ThemeController {
	return &ThemeController{}
}

// GetThemeOptions godoc
// @Summary Get theme options
// @Description Get theme options
// @Security ApiKeyAuth
// @Tags admin
// @Produce json
// @Success 200 {object} handler.RespBody{}
// @Router /answer/admin/api/theme/options [get]
func (t *ThemeController) GetThemeOptions(ctx *gin.Context) {
	handler.HandleResponse(ctx, nil, schema.GetThemeOptions)
}
