package router

import (
	"github.com/answerdev/answer/internal/controller"
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
	r.GET(controller.ConnectorLoginRouterPrefix+":name", connectorController.ConnectorLoginDispatcher)
	r.GET(controller.ConnectorRedirectRouterPrefix+":name", connectorController.ConnectorRedirectDispatcher)
	r.GET("/connector/info", connectorController.ConnectorsInfo)
	r.POST("/connector/binding/email", connectorController.ExternalLoginBindingUserSendEmail)
}

func (pr *PluginAPIRouter) RegisterAuthConnectorRouter(r *gin.RouterGroup) {
	connectorController := pr.connectorController
	r.GET("/connector/user/info", connectorController.ConnectorsUserInfo)
	r.DELETE("/connector/user/unbinding", connectorController.ExternalLoginUnbinding)
}
