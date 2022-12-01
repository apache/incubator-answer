package router

import (
	"github.com/answerdev/answer/internal/controller"
	templaterender "github.com/answerdev/answer/internal/controller/template_render"
	"github.com/gin-gonic/gin"
)

type TemplateRouter struct {
	templateController       *controller.TemplateController
	templateRenderController *templaterender.TemplateRenderController
}

func NewTemplateRouter(
	templateController *controller.TemplateController,
	templateRenderController *templaterender.TemplateRenderController,
) *TemplateRouter {
	return &TemplateRouter{
		templateController:       templateController,
		templateRenderController: templateRenderController,
	}
}

// TemplateRouter template router
func (a *TemplateRouter) RegisterTemplateRouter(r *gin.RouterGroup) {

	r.GET("/", a.templateController.Index)
	r.GET("/index", a.templateController.Index)

	r.GET("/questions", a.templateController.QuestionList)
	r.GET("/questions/:id/", a.templateController.QuestionInfo)
	r.GET("/questions/:id/:title/", a.templateController.QuestionInfo)
	r.GET("/questions/:id/:title/:answerid", a.templateController.QuestionList)

	r.GET("/tags", a.templateController.TagList)
	r.GET("/tags/:tag", a.templateController.TagInfo)
	r.GET("/users/:username", a.templateController.UserInfo)
	r.GET("/404", a.templateController.Page404)
}
