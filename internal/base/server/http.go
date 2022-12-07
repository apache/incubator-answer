package server

import (
	"html/template"
	"io/fs"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/answerdev/answer/internal/schema"
	"github.com/answerdev/answer/pkg/converter"
	"github.com/answerdev/answer/pkg/day"

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

	funcMap := template.FuncMap{
		"replaceHTMLTag": func(src string, tags ...string) string {
			p := `(?U)<(\d+)>.+</(\d+)>`

			re := regexp.MustCompile(p)
			ms := re.FindAllStringSubmatch(src, -1)
			for _, mi := range ms {
				if mi[1] == mi[2] {
					i, err := strconv.Atoi(mi[1])
					if err != nil || len(tags) < i {
						break
					}

					src = strings.ReplaceAll(src, mi[0], tags[i-1])
				}
			}

			return src
		},
		"join": func(sep string, elems ...string) string {
			return strings.Join(elems, sep)
		},
		"templateHTML": func(data string) template.HTML {
			return template.HTML(data)
		},
		"translator": func(la i18n.Language, data string, params ...interface{}) string {
			trans := translator.GlobalTrans.Tr(la, data)

			if len(params) > 0 && len(params)%2 == 0 {
				for i := 0; i < len(params); i += 2 {
					k := converter.InterfaceToString(params[i])
					v := converter.InterfaceToString(params[i+1])
					trans = strings.ReplaceAll(trans, "{{ "+k+" }}", v)
				}
			}

			return trans
		},
		"timeFormatISO": func(tz string, timestamp int64) string {
			_, _ = time.LoadLocation(tz)
			return time.Unix(timestamp, 0).Format("2006-01-02T15:04:05.000Z")
		},
		"translatorTimeFormatLongDate": func(la i18n.Language, tz string, timestamp int64) string {
			trans := translator.GlobalTrans.Tr(la, "ui.dates.long_date_with_time")
			return day.Format(timestamp, trans, tz)
		},
		"translatorTimeFormat": func(la i18n.Language, tz string, timestamp int64) string {
			var (
				now           = time.Now().Unix()
				between int64 = 0
				trans   string
			)
			_, _ = time.LoadLocation(tz)
			if now > timestamp {
				between = now - timestamp
			}

			if between <= 1 {
				return translator.GlobalTrans.Tr(la, "ui.dates.now")
			}

			if between > 1 && between < 60 {
				trans = translator.GlobalTrans.Tr(la, "ui.dates.x_seconds_ago")
				return strings.ReplaceAll(trans, "{{count}}", converter.IntToString(between))
			}

			if between >= 60 && between < 3600 {
				min := math.Floor(float64(between / 60))
				trans = translator.GlobalTrans.Tr(la, "ui.dates.x_minutes_ago")
				return strings.ReplaceAll(trans, "{{count}}", strconv.FormatFloat(min, 'f', 0, 64))
			}

			if between >= 3600 && between < 3600*24 {
				h := math.Floor(float64(between / 3600))
				trans = translator.GlobalTrans.Tr(la, "ui.dates.x_hours_ago")
				return strings.ReplaceAll(trans, "{{count}}", strconv.FormatFloat(h, 'f', 0, 64))
			}

			if between >= 3600*24 &&
				between < 3600*24*366 &&
				time.Unix(timestamp, 0).Format("2006") == time.Unix(now, 0).Format("2006") {
				trans = translator.GlobalTrans.Tr(la, "ui.dates.long_date")
				return day.Format(timestamp, trans, tz)
			}

			trans = translator.GlobalTrans.Tr(la, "ui.dates.long_date_with_year")
			return day.Format(timestamp, trans, tz)
		},
		"wrapComments": func(comments []*schema.GetCommentResp, la i18n.Language, tz string) map[string]interface{} {
			return map[string]interface{}{
				"comments": comments,
				"language": la,
				"timezone": tz,
			}
		},
	}
	r.SetFuncMap(funcMap)

	dev := os.Getenv("DEVCODE")
	if dev != "" {
		r.LoadHTMLGlob("../../ui/template/*")
	} else {
		html, _ := fs.Sub(ui.Template, "template")
		htmlTemplate := template.Must(template.New("").Funcs(funcMap).ParseFS(html, "*.html"))
		r.SetHTMLTemplate(htmlTemplate)
	}

	templateRouter.RegisterTemplateRouter(rootGroup)
	return r
}
