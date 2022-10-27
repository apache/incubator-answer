package controller

import (
	"github.com/answerdev/answer/internal/base/handler"
	"github.com/answerdev/answer/internal/schema"
	"github.com/answerdev/answer/internal/service"
	"github.com/gin-gonic/gin"
)

type SiteinfoController struct {
	siteInfoService *service.SiteInfoService
}

// NewSiteinfoController new siteinfo controller.
func NewSiteinfoController(siteInfoService *service.SiteInfoService) *SiteinfoController {
	return &SiteinfoController{
		siteInfoService: siteInfoService,
	}
}

// GetInfo godoc
// @Summary Get siteinfo
// @Description Get siteinfo
// @Tags site
// @Produce json
// @Success 200 {object} handler.RespBody{data=schema.SiteGeneralResp}
// @Router /answer/api/v1/siteinfo [get]
func (sc *SiteinfoController) GetInfo(ctx *gin.Context) {
	var (
		resp    = &schema.SiteInfoResp{}
		general schema.SiteGeneralResp
		face    schema.SiteInterfaceResp
		err     error
	)

	general, err = sc.siteInfoService.GetSiteGeneral(ctx)
	resp.General = &general
	if err != nil {
		handler.HandleResponse(ctx, err, resp)
		return
	}

	face, err = sc.siteInfoService.GetSiteInterface(ctx)
	resp.Face = &face

	handler.HandleResponse(ctx, err, resp)
}
