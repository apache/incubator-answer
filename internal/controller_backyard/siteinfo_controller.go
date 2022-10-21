package controller_backyard

import (
	"github.com/gin-gonic/gin"
	"github.com/segmentfault/answer/internal/base/handler"
	"github.com/segmentfault/answer/internal/schema"
	"github.com/segmentfault/answer/internal/service"
)

type SiteInfoController struct {
	siteInfoService *service.SiteInfoService
}

// NewSiteInfoController new siteinfo controller.
func NewSiteInfoController(siteInfoService *service.SiteInfoService) *SiteInfoController {
	return &SiteInfoController{
		siteInfoService: siteInfoService,
	}
}

// GetGeneral godoc
// @Summary Get siteinfo general
// @Description Get siteinfo general
// @Security ApiKeyAuth
// @Tags admin
// @Produce json
// @Success 200 {object} handler.RespBody{data=schema.SiteGeneralResp}
// @Router /answer/admin/api/siteinfo/general [get]
func (sc *SiteInfoController) GetGeneral(ctx *gin.Context) {
	resp, err := sc.siteInfoService.GetSiteGeneral(ctx)
	handler.HandleResponse(ctx, err, resp)
}

// GetInterface godoc
// @Summary Get siteinfo interface
// @Description Get siteinfo interface
// @Security ApiKeyAuth
// @Tags admin
// @Produce json
// @Success 200 {object} handler.RespBody{data=schema.SiteInterfaceResp}
// @Router /answer/admin/api/siteinfo/interface [get]
// @Param data body schema.AddCommentReq true "general"
func (sc *SiteInfoController) GetInterface(ctx *gin.Context) {
	resp, err := sc.siteInfoService.GetSiteInterface(ctx)
	handler.HandleResponse(ctx, err, resp)
}

// UpdateGeneral godoc
// @Summary Get siteinfo interface
// @Description Get siteinfo interface
// @Security ApiKeyAuth
// @Tags admin
// @Produce json
// @Param data body schema.SiteGeneralReq true "general"
// @Success 200 {object} handler.RespBody{}
// @Router /answer/admin/api/siteinfo/general [put]
func (sc *SiteInfoController) UpdateGeneral(ctx *gin.Context) {
	req := schema.SiteGeneralReq{}
	if handler.BindAndCheck(ctx, &req) {
		return
	}
	err := sc.siteInfoService.SaveSiteGeneral(ctx, req)
	handler.HandleResponse(ctx, err, nil)
}

// UpdateInterface godoc
// @Summary Get siteinfo interface
// @Description Get siteinfo interface
// @Security ApiKeyAuth
// @Tags admin
// @Produce json
// @Param data body schema.SiteInterfaceReq true "general"
// @Success 200 {object} handler.RespBody{}
// @Router /answer/admin/api/siteinfo/interface [put]
func (sc *SiteInfoController) UpdateInterface(ctx *gin.Context) {
	req := schema.SiteInterfaceReq{}
	if handler.BindAndCheck(ctx, &req) {
		return
	}
	err := sc.siteInfoService.SaveSiteInterface(ctx, req)
	handler.HandleResponse(ctx, err, nil)
}

// GetSMTPConfig get smtp config
// @Summary GetSMTPConfig get smtp config
// @Description GetSMTPConfig get smtp config
// @Security ApiKeyAuth
// @Tags admin
// @Produce json
// @Success 200 {object} handler.RespBody{data=schema.SiteInterfaceResp}
// @Router /answer/admin/api/setting/smtp [get]
func (sc *SiteInfoController) GetSMTPConfig(ctx *gin.Context) {
	resp, err := sc.siteInfoService.GetSMTPConfig(ctx)
	handler.HandleResponse(ctx, err, resp)
}

// UpdateSMTPConfig update smtp config
// @Summary update smtp config
// @Description update smtp config
// @Security ApiKeyAuth
// @Tags admin
// @Produce json
// @Param data body schema.SiteInterfaceReq true "general"
// @Success 200 {object} handler.RespBody{}
// @Router /answer/admin/api/setting/smtp [put]
func (sc *SiteInfoController) UpdateSMTPConfig(ctx *gin.Context) {
	req := &schema.UpdateSMTPConfigReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}
	err := sc.siteInfoService.UpdateSMTPConfig(ctx, req)
	handler.HandleResponse(ctx, err, nil)
}
