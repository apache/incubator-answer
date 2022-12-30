package controller

import (
	"fmt"

	"github.com/answerdev/answer/internal/base/handler"
	"github.com/answerdev/answer/internal/plugin"
	_ "github.com/answerdev/answer/internal/plugin/connector"
	"github.com/answerdev/answer/internal/schema"
	"github.com/answerdev/answer/internal/service/siteinfo_common"
	"github.com/gin-gonic/gin"
)

const (
	oauthRedirectRouterPrefix = "/answer/api/v1/oauth/redirect/"
	oauthLoginRouterPrefix    = "/answer/api/v1/oauth/login/"
)

// ConnectorController comment controller
type ConnectorController struct {
	siteInfoService *siteinfo_common.SiteInfoCommonService
}

// NewConnectorController new controller
func NewConnectorController(
	siteInfoService *siteinfo_common.SiteInfoCommonService,
) *ConnectorController {
	return &ConnectorController{siteInfoService: siteInfoService}
}

func (cc *ConnectorController) ConnectorRedirectRegisterRouters(r *gin.Engine) {
	_ = plugin.CallConnector(func(fn plugin.Connector) error {
		// user login url
		r.GET(oauthLoginRouterPrefix+fn.ConnectorSlugName(), fn.ConnectorSender)
		// oauth redirect url
		r.GET(oauthRedirectRouterPrefix+fn.ConnectorSlugName(), fn.ConnectorReceiver)
		return nil
	})
	r.GET("/answer/api/v1/oauth/info", cc.ConnectorsInfo)
}

func (cc *ConnectorController) ConnectorsInfo(ctx *gin.Context) {
	general, err := cc.siteInfoService.GetSiteGeneral(ctx)
	if err != nil {
		handler.HandleResponse(ctx, err, nil)
		return
	}

	resp := make([]*schema.ConnectorInfoResp, 0)
	err = plugin.CallConnector(func(fn plugin.Connector) error {
		resp = append(resp, &schema.ConnectorInfoResp{
			Name: fn.ConnectorSlugName(),
			Icon: fn.ConnectorLogo(),
			Link: fmt.Sprintf("%s%s%s", general.SiteUrl, oauthLoginRouterPrefix, fn.ConnectorSlugName()),
		})
		return nil
	})
	if err != nil {
		handler.HandleResponse(ctx, err, nil)
		return
	}
	handler.HandleResponse(ctx, nil, resp)
}
