package router

import (
	"github.com/answerdev/answer/internal/controller"
	"github.com/gin-gonic/gin"
)

type TemplateRouter struct {
	templateController *controller.TemplateController
}

func NewTemplateRouter(
	templateController *controller.TemplateController,
) *TemplateRouter {
	return &TemplateRouter{
		templateController: templateController,
	}
}

// TemplateRouter template router
func (a *TemplateRouter) RegisterTemplateRouter(r *gin.RouterGroup) {

	r.GET("/", a.templateController.Index)
	r.GET("/index", a.templateController.Index)
	r.GET("/questions", a.templateController.Index)
	r.GET("/questions/:id/", a.templateController.QuestionInfo)
	r.GET("/questions/:id/:title/", a.templateController.QuestionInfo)
	r.GET("/questions/:id/:title/:answerid", a.templateController.QuestionInfo)
	r.GET("/tags", a.templateController.TagList)
	r.GET("/tags/:tag", a.templateController.TagInfo)
	r.GET("/users/:username", a.templateController.UserInfo)
}
