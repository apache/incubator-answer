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
	"github.com/segmentfault/pacman/log"
)

const (
	ConnectorRedirectRouterPrefix = "/answer/api/v1/connector/redirect/"
	ConnectorLoginRouterPrefix    = "/answer/api/v1/connector/login/"
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

func (cc *ConnectorController) ConnectorRedirect(connector plugin.Connector) (fn func(ctx *gin.Context)) {
	return func(ctx *gin.Context) {
		general, err := cc.siteInfoService.GetSiteGeneral(ctx)
		if err != nil {
			ctx.Redirect(http.StatusFound, "/50x")
			return
		}

		receiverURL := fmt.Sprintf("%s%s%s", general.SiteUrl, ConnectorLoginRouterPrefix, connector.ConnectorSlugName())
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
		userInfo.Provider = connector.ConnectorSlugName()
		resp, err := cc.userExternalService.ExternalLogin(ctx, userInfo)
		if err != nil {
			log.Error(err)
			ctx.Redirect(http.StatusFound, "/50x")
			return
		}
		if len(resp.AccessToken) > 0 {
			ctx.Redirect(http.StatusFound, fmt.Sprintf("/index?token=%s", resp.AccessToken))
		} else {
			ctx.Redirect(http.StatusFound, fmt.Sprintf("/binding?binding_key=%s", resp.BindingKey))
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
			Link: fmt.Sprintf("%s%s%s", general.SiteUrl, ConnectorLoginRouterPrefix, fn.ConnectorSlugName()),
		})
		return nil
	})
	if err != nil {
		handler.HandleResponse(ctx, err, nil)
		return
	}
	handler.HandleResponse(ctx, nil, resp)
}

func (cc *ConnectorController) ExternalLoginBindingUserSendEmail(ctx *gin.Context) {
	req := &schema.ExternalLoginBindingUserSendEmailReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}
	resp, err := cc.userExternalService.ExternalLoginBindingUserSendEmail(ctx, req)
	handler.HandleResponse(ctx, err, resp)
}

func (cc *ConnectorController) ExternalLoginBindingUser(ctx *gin.Context) {
	req := &schema.ExternalLoginBindingUserReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}
	resp, err := cc.userExternalService.ExternalLoginBindingUser(ctx, req)
	handler.HandleResponse(ctx, err, resp)
}
