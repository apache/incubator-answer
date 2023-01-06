package controller

import (
	"fmt"
	"net/http"

	"github.com/answerdev/answer/internal/base/handler"
	"github.com/answerdev/answer/internal/plugin"
	_ "github.com/answerdev/answer/internal/plugin/connector"
	"github.com/answerdev/answer/internal/schema"
	"github.com/answerdev/answer/internal/service/siteinfo_common"
	"github.com/answerdev/answer/internal/service/user_external_login"
	"github.com/gin-gonic/gin"
)

const (
	connectorRedirectRouterPrefix = "/answer/api/v1/connector/redirect/"
	connectorLoginRouterPrefix    = "/answer/api/v1/connector/login/"
)

// ConnectorController comment controller
type ConnectorController struct {
	siteInfoService     *siteinfo_common.SiteInfoCommonService
	userExternalService *user_external_login.UserExternalLoginService
}

// NewConnectorController new controller
func NewConnectorController(
	siteInfoService *siteinfo_common.SiteInfoCommonService,
	userExternalService *user_external_login.UserExternalLoginService,
) *ConnectorController {
	return &ConnectorController{
		siteInfoService:     siteInfoService,
		userExternalService: userExternalService,
	}
}

func (cc *ConnectorController) ConnectorRedirectRegisterRouters(r *gin.Engine) {
	_ = plugin.CallConnector(func(connector plugin.Connector) error {
		r.GET(connectorLoginRouterPrefix+connector.ConnectorSlugName(), cc.ConnectorRedirect(connector))
		r.GET(connectorRedirectRouterPrefix+connector.ConnectorSlugName(), cc.ConnectorLogin(connector))
		return nil
	})
	r.GET("/answer/api/v1/connector/info", cc.ConnectorsInfo)
}

func (cc *ConnectorController) ConnectorRedirect(connector plugin.Connector) (fn func(ctx *gin.Context)) {
	return func(ctx *gin.Context) {
		general, err := cc.siteInfoService.GetSiteGeneral(ctx)
		if err != nil {
			ctx.Redirect(http.StatusFound, "/50x")
			return
		}

		receiverURL := fmt.Sprintf("%s%s%s", general.SiteUrl, connectorLoginRouterPrefix, connector.ConnectorSlugName())
		redirectURL := connector.ConnectorSender(ctx, receiverURL)
		if len(redirectURL) > 0 {
			ctx.Redirect(http.StatusFound, redirectURL)
		}
		return
	}
}

func (cc *ConnectorController) ConnectorLogin(connector plugin.Connector) (fn func(ctx *gin.Context)) {
	return func(ctx *gin.Context) {
		userInfo, err := connector.ConnectorReceiver(ctx)
		if err != nil {
			ctx.Redirect(http.StatusFound, "/50x")
			return
		}
		resp, err := cc.userExternalService.ExternalLogin(ctx, connector.ConnectorSlugName(), userInfo)
		if err != nil {
			ctx.Redirect(http.StatusFound, "/50x")
			return
		}
		if len(resp.AccessToken) > 0 {
			ctx.Redirect(http.StatusFound, fmt.Sprintf("/index?token=%s", resp.AccessToken))
		} else {
			ctx.Redirect(http.StatusFound, fmt.Sprintf("/binding?external_id=%s", resp.ExternalID))
		}
	}
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
			Link: fmt.Sprintf("%s%s%s", general.SiteUrl, connectorLoginRouterPrefix, fn.ConnectorSlugName()),
		})
		return nil
	})
	if err != nil {
		handler.HandleResponse(ctx, err, nil)
		return
	}
	handler.HandleResponse(ctx, nil, resp)
}
