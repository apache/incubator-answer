package templaterender

import (
	"github.com/answerdev/answer/internal/schema"
	"github.com/davecgh/go-spew/spew"
	"golang.org/x/net/context"
)

func (q *TemplateRenderController) UserInfo(ctx context.Context, req *schema.GetOtherUserInfoByUsernameReq) {

	resp, err := q.userService.GetOtherUserInfoByUsername(ctx, req.Username)
	spew.Dump(resp, err)
}
