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

func (pr *PluginAPIRouter) RegisterUnAuthConnectorRouter(r *gin.RouterGroup) {
	connectorController := pr.connectorController
	_ = plugin.CallConnector(func(connector plugin.Connector) error {
		connectorSlugName := connector.ConnectorSlugName()
		r.GET(controller.ConnectorLoginRouterPrefix+connectorSlugName, connectorController.ConnectorLogin(connector))
		r.GET(controller.ConnectorRedirectRouterPrefix+connectorSlugName, connectorController.ConnectorRedirect(connector))
		return nil
	})
	r.GET("/connector/info", connectorController.ConnectorsInfo)
	r.POST("/connector/binding/email", connectorController.ExternalLoginBindingUserSendEmail)
}

func (pr *PluginAPIRouter) RegisterAuthConnectorRouter(r *gin.RouterGroup) {
	connectorController := pr.connectorController
	r.GET("/connector/user/info", connectorController.ConnectorsUserInfo)
	r.DELETE("/connector/user/unbinding", connectorController.ExternalLoginUnbinding)
}
