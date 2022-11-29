package server

import (
	brotli "github.com/anargu/gin-brotli"
	"github.com/answerdev/answer/internal/base/middleware"
	"github.com/answerdev/answer/internal/router"
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
	templateRouter.RegisterTemplateRouter(rootGroup)
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

	return r
}
