package controller

import (
	"fmt"
	"net/http"

	"github.com/answerdev/answer/internal/base/handler"
	"github.com/answerdev/answer/internal/plugin"
	_ "github.com/answerdev/answer/internal/plugin/connector"
	"github.com/answerdev/answer/internal/schema"
	"github.com/answerdev/answer/internal/service/export"
	"github.com/answerdev/answer/internal/service/siteinfo_common"
	"github.com/answerdev/answer/internal/service/user_external_login"
	"github.com/gin-gonic/gin"
	"github.com/segmentfault/pacman/log"
)

const (
	ConnectorLoginRouterPrefix    = "/answer/api/v1/connector/login/"
	ConnectorRedirectRouterPrefix = "/answer/api/v1/connector/redirect/"
)

// ConnectorController comment controller
type ConnectorController struct {
	siteInfoService     *siteinfo_common.SiteInfoCommonService
	userExternalService *user_external_login.UserExternalLoginService
	emailService        *export.EmailService
}

// NewConnectorController new controller
func NewConnectorController(
	siteInfoService *siteinfo_common.SiteInfoCommonService,
	emailService *export.EmailService,
	userExternalService *user_external_login.UserExternalLoginService,
) *ConnectorController {
	return &ConnectorController{
		siteInfoService:     siteInfoService,
		userExternalService: userExternalService,
		emailService:        emailService,
	}
}

func (cc *ConnectorController) ConnectorLogin(connector plugin.Connector) (fn func(ctx *gin.Context)) {
	return func(ctx *gin.Context) {
		general, err := cc.siteInfoService.GetSiteGeneral(ctx)
		if err != nil {
			log.Error(err)
			ctx.Redirect(http.StatusFound, "/50x")
			return
		}

		receiverURL := fmt.Sprintf("%s%s%s", general.SiteUrl, ConnectorRedirectRouterPrefix, connector.ConnectorSlugName())
		redirectURL := connector.ConnectorSender(ctx, receiverURL)
		if len(redirectURL) > 0 {
			ctx.Redirect(http.StatusFound, redirectURL)
		}
		return
	}
}

func (cc *ConnectorController) ConnectorRedirect(connector plugin.Connector) (fn func(ctx *gin.Context)) {
	return func(ctx *gin.Context) {
		siteGeneral, err := cc.siteInfoService.GetSiteGeneral(ctx)
		if err != nil {
			log.Errorf("get site info failed: %v", err)
			ctx.Redirect(http.StatusFound, "/50x")
			return
		}
		userInfo, err := connector.ConnectorReceiver(ctx)
		if err != nil {
			log.Errorf("connector received failed: %v", err)
			ctx.Redirect(http.StatusFound, "/50x")
			return
		}
		u := &schema.ExternalLoginUserInfoCache{
			Provider:   connector.ConnectorSlugName(),
			ExternalID: userInfo.ExternalID,
			Name:       userInfo.Name,
			Email:      userInfo.Email,
			MetaInfo:   userInfo.MetaInfo,
		}
		resp, err := cc.userExternalService.ExternalLogin(ctx, u)
		if err != nil {
			log.Errorf("external login failed: %v", err)
			ctx.Redirect(http.StatusFound, "/50x")
			return
		}
		if len(resp.AccessToken) > 0 {
			ctx.Redirect(http.StatusFound, fmt.Sprintf("%s/users/oauth?access_token=%s",
				siteGeneral.SiteUrl, resp.AccessToken))
		} else {
			ctx.Redirect(http.StatusFound, fmt.Sprintf("%s/users/confirm-email?binding_key=%s",
				siteGeneral.SiteUrl, resp.BindingKey))
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

// ExternalLoginBindingUserSendEmail external login binding user send email
// @Summary external login binding user send email
// @Description external login binding user send email
// @Tags Plugin
// @Accept json
// @Produce json
// @Param data body schema.ExternalLoginBindingUserSendEmailReq  true "external login binding user send email"
// @Success 200 {object} handler.RespBody{data=schema.ExternalLoginBindingUserSendEmailResp}
// @Router /answer/api/v1/connector/binding/email [post]
func (cc *ConnectorController) ExternalLoginBindingUserSendEmail(ctx *gin.Context) {
	req := &schema.ExternalLoginBindingUserSendEmailReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}

	resp, err := cc.userExternalService.ExternalLoginBindingUserSendEmail(ctx, req)
	handler.HandleResponse(ctx, err, resp)
}
