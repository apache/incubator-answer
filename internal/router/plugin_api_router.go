package router

import (
	"github.com/answerdev/answer/internal/controller"
	"github.com/answerdev/answer/internal/plugin"
	"github.com/gin-gonic/gin"
)

type PluginAPIRouter struct {
	connectorController *controller.ConnectorController
}

func NewPluginAPIRouter(
	connectorController *controller.ConnectorController,
) *PluginAPIRouter {
	return &PluginAPIRouter{
		connectorController: connectorController,
	}
}

func (pr *PluginAPIRouter) RegisterConnector(r *gin.Engine) {
	connectorController := pr.connectorController
	_ = plugin.CallConnector(func(connector plugin.Connector) error {
		connectorSlugName := connector.ConnectorSlugName()
		r.GET(controller.ConnectorLoginRouterPrefix+connectorSlugName, connectorController.ConnectorRedirect(connector))
		r.GET(controller.ConnectorRedirectRouterPrefix+connectorSlugName, connectorController.ConnectorLogin(connector))
		return nil
	})
	r.GET("/answer/api/v1/connector/info", connectorController.ConnectorsInfo)
	r.POST("/answer/api/v1/connector/binding/email", connectorController.ExternalLoginBindingUserSendEmail)
	r.POST("/answer/api/v1/connector/binding", connectorController.ExternalLoginBindingUser)
}
