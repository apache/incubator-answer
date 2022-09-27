package router

import (
	"github.com/gin-gonic/gin"
	"github.com/segmentfault/answer/internal/service/service_config"
)

// StaticRouter static api router
type StaticRouter struct {
	serviceConfig *service_config.ServiceConfig
}

// NewStaticRouter new static api router
func NewStaticRouter(serviceConfig *service_config.ServiceConfig) *StaticRouter {
	return &StaticRouter{
		serviceConfig: serviceConfig,
	}
}

// RegisterStaticRouter register static api router
func (a *StaticRouter) RegisterStaticRouter(r *gin.RouterGroup) {
	r.Static("/uploads", a.serviceConfig.UploadPath)
}
