package router

import (
	"github.com/gin-gonic/gin"
)

type TemplateRouter struct {
}

func NewTemplateRouter() *TemplateRouter {
	return &TemplateRouter{}
}

// TemplateRouter template router
func (a *TemplateRouter) TemplateRouter(r *gin.RouterGroup) {
}
