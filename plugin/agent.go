package plugin

import (
	"github.com/gin-gonic/gin"
)

type Agent interface {
	Base
	RegisterUnAuthRouter(r *gin.RouterGroup)
	RegisterAuthUserRouter(r *gin.RouterGroup)
	RegisterAuthAdminRouter(r *gin.RouterGroup)
}

var (
	CallAgent,
	registerAgent = MakePlugin[Agent](true)
	siteURLFn func() string
)

// SiteURL The site url is the domain address of the current site. e.g. http://localhost:8080
// When some Agent plugins want to redirect to the origin site, it can use this function to get the site url.
func SiteURL() string {
	if siteURLFn != nil {
		return siteURLFn()
	}
	return ""
}

// RegisterGetSiteURLFunc Register a function to get the site url.
func RegisterGetSiteURLFunc(fn func() string) {
	siteURLFn = fn
}
