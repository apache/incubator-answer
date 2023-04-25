package server

import (
	"html/template"
	"io/fs"

	brotli "github.com/anargu/gin-brotli"
	"github.com/answerdev/answer/internal/base/middleware"
	"github.com/answerdev/answer/internal/router"
	"github.com/answerdev/answer/plugin"
	"github.com/answerdev/answer/ui"
	"github.com/gin-gonic/gin"
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
	pluginAPIRouter *router.PluginAPIRouter,
) *gin.Engine {

	if debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(brotli.Brotli(brotli.DefaultCompression), middleware.ExtractAndSetAcceptLanguage)
	r.GET("/healthz", func(ctx *gin.Context) { ctx.String(200, "OK") })

	html, _ := fs.Sub(ui.Template, "template")
	htmlTemplate := template.Must(template.New("").Funcs(funcMap).ParseFS(html, "*"))
	r.SetHTMLTemplate(htmlTemplate)
	r.Use(middleware.HeadersByRequestURI())
	viewRouter.Register(r)

	rootGroup := r.Group("")
	swaggerRouter.Register(rootGroup)
	static := r.Group("")
	static.Use(avatarMiddleware.AvatarThumb())
	staticRouter.RegisterStaticRouter(static)

	// The route must be available without logging in
	mustUnAuthV1 := r.Group("/answer/api/v1")
	answerRouter.RegisterMustUnAuthAnswerAPIRouter(mustUnAuthV1)

	// register api that no need to login
	unAuthV1 := r.Group("/answer/api/v1")
	unAuthV1.Use(authUserMiddleware.Auth(), authUserMiddleware.EjectUserBySiteInfo())
	answerRouter.RegisterUnAuthAnswerAPIRouter(unAuthV1)

	// register api that must be authenticated
	authV1 := r.Group("/answer/api/v1")
	authV1.Use(authUserMiddleware.MustAuth())
	answerRouter.RegisterAnswerAPIRouter(authV1)

	adminauthV1 := r.Group("/answer/admin/api")
	adminauthV1.Use(authUserMiddleware.AdminAuth())
	answerRouter.RegisterAnswerAdminAPIRouter(adminauthV1)

	templateRouter.RegisterTemplateRouter(rootGroup)

	// plugin routes
	pluginAPIRouter.RegisterUnAuthConnectorRouter(mustUnAuthV1)
	pluginAPIRouter.RegisterAuthUserConnectorRouter(authV1)
	pluginAPIRouter.RegisterAuthAdminConnectorRouter(adminauthV1)

	_ = plugin.CallAgent(func(agent plugin.Agent) error {
		agent.RegisterUnAuthRouter(mustUnAuthV1)
		agent.RegisterAuthUserRouter(authV1)
		agent.RegisterAuthAdminRouter(adminauthV1)
		return nil
	})
	return r
}
