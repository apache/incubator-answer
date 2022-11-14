package controller

import (
	"github.com/answerdev/answer/internal/base/handler"
	"github.com/answerdev/answer/internal/service/dashboard"
	"github.com/gin-gonic/gin"
)

type DashboardController struct {
	dashboardService *dashboard.DashboardService
}

// NewDashboardController new controller
func NewDashboardController(
	dashboardService *dashboard.DashboardService,
) *DashboardController {
	return &DashboardController{
		dashboardService: dashboardService,
	}
}

// DashboardInfo godoc
// @Summary DashboardInfo
// @Description DashboardInfo
// @Tags admin
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Router /answer/admin/api/dashboard [get]
// @Success 200 {object} handler.RespBody
func (ac *DashboardController) DashboardInfo(ctx *gin.Context) {
	info, err := ac.dashboardService.StatisticalByCache(ctx)
	handler.HandleResponse(ctx, err, gin.H{
		"info": info,
	})
}
