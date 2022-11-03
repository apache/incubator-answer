package server

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"

	"github.com/answerdev/answer/ui"
	"github.com/gin-gonic/gin"
	"github.com/segmentfault/pacman/log"
)

const UIStaticPath = "build/static"

type _resource struct {
	fs embed.FS
}

// Open to implement the interface by http.FS required
func (r *_resource) Open(name string) (fs.File, error) {
	name = fmt.Sprintf(UIStaticPath+"/%s", name)
	log.Debugf("open static path %s", name)
	return r.fs.Open(name)
}

// NewHTTPServer new http server.
func NewInstallHTTPServer() *gin.Engine {
	r := gin.New()
	gin.SetMode(gin.DebugMode)

	r.GET("/healthz", func(ctx *gin.Context) { ctx.String(200, "OK??") })

	// gin.SetMode(gin.ReleaseMode)

	r.StaticFS("/static", http.FS(&_resource{
		fs: ui.Build,
	}))

	installApi := r.Group("")
	installApi.GET("/install", Install)

	return r
}

func Install(c *gin.Context) {
	filePath := ""
	var file []byte
	var err error
	filePath = "build/index.html"
	c.Header("content-type", "text/html;charset=utf-8")
	file, err = ui.Build.ReadFile(filePath)
	if err != nil {
		log.Error(err)
		c.Status(http.StatusNotFound)
		return
	}
	c.String(http.StatusOK, string(file))
}
