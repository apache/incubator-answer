package templaterender

import (
	"github.com/answerdev/answer/internal/base/pager"
	"github.com/answerdev/answer/internal/schema"
	"github.com/jinzhu/copier"
	"golang.org/x/net/context"
)

func (q *TemplateRenderController) TagList(ctx context.Context, req *schema.GetTagWithPageReq) (resp *pager.PageModel, err error) {
	resp, err = q.tagService.GetTagWithPage(ctx, req)
	if err != nil {
		return
	}
	return
}

func (q *TemplateRenderController) TagInfo(ctx context.Context, req *schema.GetTamplateTagInfoReq) (resp *schema.GetTagResp, questionList []*schema.QuestionInfo, questionCount int64, err error) {
	dto := &schema.GetTagInfoReq{}
	_ = copier.Copy(dto, req)
	resp, err = q.tagService.GetTagInfo(ctx, dto)
	searchQuestion := &schema.QuestionSearch{}
	searchQuestion.Page = req.Page
	searchQuestion.PageSize = req.PageSize
	searchQuestion.Order = "newest"
	searchQuestion.Tag = req.Name
	questionList, questionCount, err = q.questionService.SearchList(ctx, searchQuestion, "")
	if err != nil {
		return
	}
	return resp, questionList, questionCount, err
}
