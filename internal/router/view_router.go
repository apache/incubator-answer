package router

import (
	"embed"
	"errors"
	"io/fs"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/segmentfault/answer/web"
)

// RegisterViewRouter
type ViewRouter struct {
}

// NewRegisterViewRouter
func NewViewRouter() *ViewRouter {
	return &ViewRouter{}
}

type Resource struct {
	fs   embed.FS
	path string
}

func NewResource() *Resource {
	return &Resource{
		fs:   web.Static,
		path: "html",
	}
}

func (r *Resource) Open(name string) (fs.File, error) {
	if filepath.Separator != '/' && strings.ContainsRune(name, filepath.Separator) {
		return nil, errors.New("http: invalid character in file path")
	}
	fullName := filepath.Join(r.path, filepath.FromSlash(path.Clean("/static/"+name)))
	file, err := r.fs.Open(fullName)
	return file, err
}

func (a *ViewRouter) RegisterViewRouter(r *gin.Engine) {
	//export answer_html_static_path="../../web/static"
	//export answer_html_page_path="../../web"
	static := os.Getenv("answer_html_static_path")
	index := os.Getenv("answer_html_page_path")
	if len(static) > 0 && len(index) > 0 {
		r.LoadHTMLGlob(index + "/*.html")
		r.Static("/static", static)
		r.NoRoute(func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.html", gin.H{})
		})
		return
	} else {
		r.StaticFS("/static", http.FS(NewResource()))
		r.NoRoute(func(c *gin.Context) {
			c.Header("content-type", "text/html;charset=utf-8")
			c.String(200, string(web.Html))
		})
	}

}
