package controller

import (
	"github.com/answerdev/answer/internal/base/handler"
	"github.com/answerdev/answer/internal/schema"
	"github.com/answerdev/answer/internal/service/siteinfo_common"
	"github.com/gin-gonic/gin"
	"github.com/segmentfault/pacman/log"
)

type SiteinfoController struct {
	siteInfoService *siteinfo_common.SiteInfoCommonService
}

// NewSiteinfoController new siteinfo controller.
func NewSiteinfoController(siteInfoService *siteinfo_common.SiteInfoCommonService) *SiteinfoController {
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
		log.Error(err)
	}
	resp.Interface, err = sc.siteInfoService.GetSiteInterface(ctx)
	if err != nil {
		log.Error(err)
	}

	resp.Branding, err = sc.siteInfoService.GetSiteBranding(ctx)
	if err != nil {
		log.Error(err)
	}
	handler.HandleResponse(ctx, nil, resp)
}

// GetSiteLegalInfo get site legal info
// @Summary get site legal info
// @Description get site legal info
// @Tags site
// @Param info_type query string true "legal information type" Enums(tos, privacy)
// @Produce json
// @Success 200 {object} handler.RespBody{data=schema.GetSiteLegalInfoResp}
// @Router /answer/api/v1/siteinfo/legal [get]
func (sc *SiteinfoController) GetSiteLegalInfo(ctx *gin.Context) {
	req := &schema.GetSiteLegalInfoReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}
	siteLegal, err := sc.siteInfoService.GetSiteLegal(ctx)
	if err != nil {
		handler.HandleResponse(ctx, err, nil)
		return
	}
	resp := &schema.GetSiteLegalInfoResp{}
	if req.IsTOS() {
		resp.TermsOfServiceOriginalText = siteLegal.TermsOfServiceOriginalText
		resp.TermsOfServiceParsedText = siteLegal.TermsOfServiceParsedText
	} else if req.IsPrivacy() {
		resp.PrivacyPolicyOriginalText = siteLegal.PrivacyPolicyOriginalText
		resp.PrivacyPolicyParsedText = siteLegal.PrivacyPolicyParsedText
	}
	handler.HandleResponse(ctx, nil, resp)
}
