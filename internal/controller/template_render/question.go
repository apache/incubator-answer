package templaterender

import (
	"github.com/answerdev/answer/internal/schema"
	"github.com/gin-gonic/gin"
)

func (t *TemplateRenderController) Index(ctx *gin.Context, req *schema.QuestionSearch) ([]*schema.QuestionInfo, int64, error) {
	return t.questionService.SearchList(ctx, req, req.UserID)
}

func (t *TemplateRenderController) QuestionDetail(ctx *gin.Context, id string) (resp *schema.QuestionInfo, err error) {
	return t.questionService.GetQuestion(ctx, id, "", true)
}
