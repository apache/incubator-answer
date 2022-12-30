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

func (pr *PluginAPIRouter) RegisterConnector(r *gin.Engine) {
	pr.connectorController.ConnectorRedirectRegisterRouters(r)
}
