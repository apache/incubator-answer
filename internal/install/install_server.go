package install

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

// NewInstallHTTPServer new install http server.
func NewInstallHTTPServer() *gin.Engine {
	r := gin.New()
	gin.SetMode(gin.DebugMode)
	r.GET("/healthz", func(ctx *gin.Context) { ctx.String(200, "OK") })
	r.StaticFS("/static", http.FS(&_resource{
		fs: ui.Build,
	}))

	installApi := r.Group("")
	installApi.GET("/install", WebPage)

	installApi.GET("/installation/language/options", LangOptions)

	installApi.POST("/installation/db/check", CheckDatabase)

	installApi.POST("/installation/config-file/check", CheckConfigFile)

	installApi.POST("/installation/init", InitEnvironment)

	installApi.POST("/installation/base-info", InitBaseInfo)
	return r
}

func WebPage(c *gin.Context) {
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
