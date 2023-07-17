package action

import (
	"context"

	"github.com/answerdev/answer/internal/entity"
	"github.com/davecgh/go-spew/spew"
)

// ValidationStrategy
// true pass
// false need captcha
func (cs *CaptchaService) ValidationStrategy(ctx context.Context, unit, actionType string) bool {
	info, err := cs.captchaRepo.GetActionType(ctx, unit, actionType)
	spew.Dump("[ValidationStrategy=验证策略]", unit, actionType, info, err)
	if err != nil {
		//No record, no processing
		//
	}
	switch actionType {
	case entity.CaptchaActionEmail:
		return cs.CaptchaActionEmail(ctx, info)
	case entity.CaptchaActionPassword:
		return cs.CaptchaActionPassword(ctx, info)
	case entity.CaptchaActionEditUserinfo:
		return cs.CaptchaActionEditUserinfo(ctx, info)
	case entity.CaptchaActionQuestion:
		return cs.CaptchaActionQuestion(ctx, info)
	case entity.CaptchaActionAnswer:
		return cs.CaptchaActionAnswer(ctx, info)
	case entity.CaptchaActionComment:
		return cs.CaptchaActionComment(ctx, info)
	case entity.CaptchaActionEdit:
		return cs.CaptchaActionEdit(ctx, info)
	case entity.CaptchaActionInvitationAnswer:
		return cs.CaptchaActionInvitationAnswer(ctx, info)
	case entity.CaptchaActionSearch:
		return cs.CaptchaActionSearch(ctx, info)
	case entity.CaptchaActionReport:
		return cs.CaptchaActionReport(ctx, info)
	case entity.CaptchaActionDelete:
		return cs.CaptchaActionDelete(ctx, info)
	case entity.CaptchaActionVote:
		return cs.CaptchaActionVote(ctx, info)

	}
	//actionType not found
	return false
}

func (cs *CaptchaService) CaptchaActionEmail(ctx context.Context, actioninfo *entity.ActionRecordInfo) bool {
	// setNum := 0
	// setTime := 0 //seconds
	// You need a verification code every time
	spew.Dump("[CaptchaActionEmail]", actioninfo)
	return false
}

func (cs *CaptchaService) CaptchaActionPassword(ctx context.Context, actioninfo *entity.ActionRecordInfo) bool {
	spew.Dump("[CaptchaActionPassword]", actioninfo)
	// setNum := 0
	// setTime := 0 //seconds
	return false
}

func (cs *CaptchaService) CaptchaActionEditUserinfo(ctx context.Context, actioninfo *entity.ActionRecordInfo) bool {
	spew.Dump("[CaptchaActionEditUserinfo]", actioninfo)
	// setNum := 0
	// setTime := 0 //seconds
	return false
}

func (cs *CaptchaService) CaptchaActionQuestion(ctx context.Context, actioninfo *entity.ActionRecordInfo) bool {
	spew.Dump("[CaptchaActionQuestion]", actioninfo)
	// setNum := 0
	// setTime := 0 //seconds
	return false
}

func (cs *CaptchaService) CaptchaActionAnswer(ctx context.Context, actioninfo *entity.ActionRecordInfo) bool {
	spew.Dump("[CaptchaActionAnswer]", actioninfo)
	// setNum := 0
	// setTime := 0 //seconds
	return false
}

func (cs *CaptchaService) CaptchaActionComment(ctx context.Context, actioninfo *entity.ActionRecordInfo) bool {
	spew.Dump("[CaptchaActionComment]", actioninfo)
	// setNum := 0
	// setTime := 0 //seconds
	return false
}

func (cs *CaptchaService) CaptchaActionEdit(ctx context.Context, actioninfo *entity.ActionRecordInfo) bool {
	spew.Dump("[CaptchaActionEdit]", actioninfo)
	// setNum := 0
	// setTime := 0 //seconds
	return false
}

func (cs *CaptchaService) CaptchaActionInvitationAnswer(ctx context.Context, actioninfo *entity.ActionRecordInfo) bool {
	spew.Dump("[CaptchaActionInvitationAnswer]", actioninfo)
	// setNum := 0
	// setTime := 0 //seconds
	return false
}

func (cs *CaptchaService) CaptchaActionSearch(ctx context.Context, actioninfo *entity.ActionRecordInfo) bool {
	spew.Dump("[CaptchaActionSearch]", actioninfo)
	// setNum := 0
	// setTime := 0 //seconds
	return false
}

func (cs *CaptchaService) CaptchaActionReport(ctx context.Context, actioninfo *entity.ActionRecordInfo) bool {
	spew.Dump("[CaptchaActionReport]", actioninfo)
	// setNum := 0
	// setTime := 0 //seconds
	return false
}

func (cs *CaptchaService) CaptchaActionDelete(ctx context.Context, actioninfo *entity.ActionRecordInfo) bool {
	spew.Dump("[CaptchaActionDelete]", actioninfo)
	// setNum := 0
	// setTime := 0 //seconds
	return false
}

func (cs *CaptchaService) CaptchaActionVote(ctx context.Context, actioninfo *entity.ActionRecordInfo) bool {
	spew.Dump("[CaptchaActionVote]", actioninfo)
	// setNum := 0
	// setTime := 0 //seconds
	return false
}
