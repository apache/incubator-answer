package templaterender

import (
	"github.com/answerdev/answer/internal/schema"
	"github.com/gin-gonic/gin"
)

func (t *TemplateRenderController) Index(ctx *gin.Context, req *schema.QuestionSearch) ([]*schema.QuestionInfo, int64, error) {
	return t.questionService.SearchList(ctx, req, req.UserID)
}

func (t *TemplateRenderController) QuestionDetail(ctx *gin.Context) {

}

func (t *TemplateRenderController) AnswerDetail(ctx *gin.Context) {}
