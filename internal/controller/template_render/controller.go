package templaterender

import (
	"github.com/answerdev/answer/internal/service"
	"github.com/google/wire"
)

// ProviderSetTemplateRenderController is template render controller providers.
var ProviderSetTemplateRenderController = wire.NewSet(
	NewTemplateRenderController,
)

type TemplateRenderController struct {
	questionService *service.QuestionService
	userService     *service.UserService
}

func NewTemplateRenderController(
	questionService *service.QuestionService,
	userService *service.UserService,

) *TemplateRenderController {
	return &TemplateRenderController{
		questionService: questionService,
		userService:     userService,
	}
}
