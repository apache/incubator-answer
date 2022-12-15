package templaterender

import (
	"context"
	"github.com/answerdev/answer/internal/base/pager"
	"github.com/answerdev/answer/internal/schema"
)

func (t *TemplateRenderController) CommentList(
	ctx context.Context,
	objectIDs []string,
) (
	comments map[string][]*schema.GetCommentResp,
	err error,
) {

	comments = make(map[string][]*schema.GetCommentResp, len(objectIDs))

	for _, objectID := range objectIDs {
		var (
			req = &schema.GetCommentWithPageReq{
				Page:      1,
				PageSize:  3,
				ObjectID:  objectID,
				QueryCond: "vote",
				UserID:    "",
			}
			pageModel *pager.PageModel
		)
		pageModel, err = t.commentService.GetCommentWithPage(ctx, req)
		if err != nil {
			return
		}
		li := pageModel.List
		comments[objectID] = li.([]*schema.GetCommentResp)
	}
	return
}
