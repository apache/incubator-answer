package templaterender

import (
	"github.com/answerdev/answer/internal/schema"
	"golang.org/x/net/context"
)

func (q *TemplateRenderController) UserInfo(ctx context.Context, req *schema.GetOtherUserInfoByUsernameReq) (resp *schema.GetOtherUserInfoResp, err error) {
	return q.userService.GetOtherUserInfoByUsername(ctx, req.Username)
}
