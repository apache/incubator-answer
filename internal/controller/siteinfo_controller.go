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

// GetSiteInfo get site info
// @Summary get site info
// @Description get site info
// @Tags site
// @Produce json
// @Success 200 {object} handler.RespBody{data=schema.SiteGeneralResp}
// @Router /answer/api/v1/siteinfo [get]
func (sc *SiteinfoController) GetSiteInfo(ctx *gin.Context) {
	var err error
	resp := &schema.SiteInfoResp{}
	resp.General, err = sc.siteInfoService.GetSiteGeneral(ctx)
	if err != nil {
		handler.HandleResponse(ctx, err, resp)
		return
	}
	resp.Face, err = sc.siteInfoService.GetSiteInterface(ctx)
	handler.HandleResponse(ctx, err, resp)
}
