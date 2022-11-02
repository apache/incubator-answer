package server

import (
	"embed"
	"net/http"

	"github.com/answerdev/answer/ui"
	"github.com/gin-gonic/gin"
)

type _resource struct {
	fs embed.FS
}

// NewHTTPServer new http server.
func NewInstallHTTPServer() *gin.Engine {
	r := gin.New()
	gin.SetMode(gin.DebugMode)

	r.GET("/healthz", func(ctx *gin.Context) { ctx.String(200, "OK??") })

	// gin.SetMode(gin.ReleaseMode)
	r.StaticFS("/static", http.FS(ui.Build))

	return r
}
