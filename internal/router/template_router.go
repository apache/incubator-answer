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

	r.GET("/questions", a.templateRenderController.Index)
	r.GET("/questions/:id/", a.templateRenderController.QuestionDetail)
	r.GET("/questions/:id/:title/", a.templateRenderController.QuestionDetail)
	r.GET("/questions/:id/:title/:answerid", a.templateRenderController.AnswerDetail)

	r.GET("/tags", a.templateController.TagList)
	r.GET("/tags/:tag", a.templateController.TagInfo)
	r.GET("/users/:username", a.templateController.UserInfo)
}
