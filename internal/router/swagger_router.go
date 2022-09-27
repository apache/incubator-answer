package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/segmentfault/answer/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// SwaggerRouter swagger api router
type SwaggerRouter struct {
	config *SwaggerConfig
}

// NewSwaggerRouter new swagger api router
func NewSwaggerRouter(config *SwaggerConfig) *SwaggerRouter {
	return &SwaggerRouter{
		config: config,
	}
}

// Register register swagger api router
func (a *SwaggerRouter) Register(r *gin.RouterGroup) {
	if a.config.Show {
		a.InitSwaggerDocs()
		gofmt := fmt.Sprintf("%s://%s%s/swagger/doc.json", a.config.Protocol, a.config.Host, a.config.Address)
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, ginSwagger.URL(gofmt)))
	}
}

// InitSwaggerDocs init swagger docs
func (a *SwaggerRouter) InitSwaggerDocs() {
	docs.SwaggerInfo.Title = "answer"
	docs.SwaggerInfo.Description = "answer api"
	docs.SwaggerInfo.Version = "v0.0.1"
	docs.SwaggerInfo.Host = fmt.Sprintf("%s%s", a.config.Host, a.config.Address)
	docs.SwaggerInfo.BasePath = "/"
}
