package server

import (
	"html/template"
	"io/fs"
	"os"
	"time"

	brotli "github.com/anargu/gin-brotli"
	"github.com/answerdev/answer/internal/base/middleware"
	"github.com/answerdev/answer/internal/base/translator"
	"github.com/answerdev/answer/internal/router"
	"github.com/answerdev/answer/ui"
	"github.com/gin-gonic/gin"
	"github.com/segmentfault/pacman/i18n"
)

// NewHTTPServer new http server.
func NewHTTPServer(debug bool,
	staticRouter *router.StaticRouter,
	answerRouter *router.AnswerAPIRouter,
	swaggerRouter *router.SwaggerRouter,
	viewRouter *router.UIRouter,
	authUserMiddleware *middleware.AuthUserMiddleware,
	avatarMiddleware *middleware.AvatarMiddleware,
	templateRouter *router.TemplateRouter,
) *gin.Engine {

	if debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(brotli.Brotli(brotli.DefaultCompression))
	r.GET("/healthz", func(ctx *gin.Context) { ctx.String(200, "OK") })

	viewRouter.Register(r)

	rootGroup := r.Group("")
	swaggerRouter.Register(rootGroup)
	static := r.Group("")
	static.Use(avatarMiddleware.AvatarThumb())
	staticRouter.RegisterStaticRouter(static)

	// register api that no need to login
	unAuthV1 := r.Group("/answer/api/v1")
	unAuthV1.Use(authUserMiddleware.Auth())
	answerRouter.RegisterUnAuthAnswerAPIRouter(unAuthV1)

	// register api that must be authenticated
	authV1 := r.Group("/answer/api/v1")
	authV1.Use(authUserMiddleware.MustAuth())
	answerRouter.RegisterAnswerAPIRouter(authV1)

	cmsauthV1 := r.Group("/answer/admin/api")
	cmsauthV1.Use(authUserMiddleware.CmsAuth())
	answerRouter.RegisterAnswerCmsAPIRouter(cmsauthV1)

	r.SetFuncMap(template.FuncMap{
		"templateHTML": func(data string) template.HTML {
			return template.HTML(data)
		},
		"translator": func(la i18n.Language, data string) string {
			return translator.GlobalTrans.Tr(la, data)
		},
		"translatorTimeFormat": func(la i18n.Language, timestamp int64) string {
			return time.Unix(timestamp, 0).Format("2006-01-02 15:04:05")
		},
	})

	dev := os.Getenv("DEVCODE")
	if dev != "" {
		r.LoadHTMLGlob("../../ui/template/*")
	} else {
		html, _ := fs.Sub(ui.Template, "template")
		htmlTemplate := template.Must(template.New("").ParseFS(html, "*.html"))
		r.SetHTMLTemplate(htmlTemplate)
	}

	templateRouter.RegisterTemplateRouter(rootGroup)
	return r
}
