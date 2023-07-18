package action

import (
	"context"
	"time"

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
		return cs.CaptchaActionEmail(ctx, unit, info)
	case entity.CaptchaActionPassword:
		return cs.CaptchaActionPassword(ctx, unit, info)
	case entity.CaptchaActionEditUserinfo:
		return cs.CaptchaActionEditUserinfo(ctx, unit, info)
	case entity.CaptchaActionQuestion:
		return cs.CaptchaActionQuestion(ctx, unit, info)
	case entity.CaptchaActionAnswer:
		return cs.CaptchaActionAnswer(ctx, unit, info)
	case entity.CaptchaActionComment:
		return cs.CaptchaActionComment(ctx, unit, info)
	case entity.CaptchaActionEdit:
		return cs.CaptchaActionEdit(ctx, unit, info)
	case entity.CaptchaActionInvitationAnswer:
		return cs.CaptchaActionInvitationAnswer(ctx, unit, info)
	case entity.CaptchaActionSearch:
		return cs.CaptchaActionSearch(ctx, unit, info)
	case entity.CaptchaActionReport:
		return cs.CaptchaActionReport(ctx, unit, info)
	case entity.CaptchaActionDelete:
		return cs.CaptchaActionDelete(ctx, unit, info)
	case entity.CaptchaActionVote:
		return cs.CaptchaActionVote(ctx, unit, info)

	}
	//actionType not found
	return false
}

func (cs *CaptchaService) CaptchaActionEmail(ctx context.Context, unit string, actioninfo *entity.ActionRecordInfo) bool {
	// You need a verification code every time
	spew.Dump("[CaptchaActionEmail]", actioninfo)
	return false
}

func (cs *CaptchaService) CaptchaActionPassword(ctx context.Context, unit string, actioninfo *entity.ActionRecordInfo) bool {
	spew.Dump("[CaptchaActionPassword]", actioninfo)
	setNum := 3
	setTime := int64(60 * 30) //seconds
	now := time.Now().Unix()
	if now-actioninfo.LastTime <= setTime || actioninfo.Num >= setNum {
		return false
	}
	if now-actioninfo.LastTime > setTime {
		cs.captchaRepo.SetActionType(ctx, unit, entity.CaptchaActionPassword, "", 0)
	}
	return true
}

func (cs *CaptchaService) CaptchaActionEditUserinfo(ctx context.Context, unit string, actioninfo *entity.ActionRecordInfo) bool {
	spew.Dump("[CaptchaActionEditUserinfo]", actioninfo)
	setNum := 3
	setTime := int64(60 * 30) //seconds
	now := time.Now().Unix()
	if now-actioninfo.LastTime <= setTime || actioninfo.Num >= setNum {
		return false
	}
	if now-actioninfo.LastTime > setTime {
		cs.captchaRepo.SetActionType(ctx, unit, entity.CaptchaActionEditUserinfo, "", 0)
	}
	return true
}

func (cs *CaptchaService) CaptchaActionQuestion(ctx context.Context, unit string, actioninfo *entity.ActionRecordInfo) bool {
	spew.Dump("[CaptchaActionQuestion]", actioninfo)
	setNum := 3
	setTime := int64(5) //seconds
	now := time.Now().Unix()
	if now-actioninfo.LastTime <= setTime || actioninfo.Num >= setNum {
		return false
	}
	return true
}

func (cs *CaptchaService) CaptchaActionAnswer(ctx context.Context, unit string, actioninfo *entity.ActionRecordInfo) bool {
	spew.Dump("[CaptchaActionAnswer]", actioninfo)
	setNum := 10
	setTime := int64(5) //seconds
	now := time.Now().Unix()
	if now-actioninfo.LastTime <= setTime || actioninfo.Num >= setNum {
		return false
	}
	return true
}

func (cs *CaptchaService) CaptchaActionComment(ctx context.Context, unit string, actioninfo *entity.ActionRecordInfo) bool {
	spew.Dump("[CaptchaActionComment]", actioninfo)
	setNum := 30
	setTime := int64(1) //seconds
	now := time.Now().Unix()
	if now-actioninfo.LastTime <= setTime || actioninfo.Num >= setNum {
		return false
	}
	return true
}

func (cs *CaptchaService) CaptchaActionEdit(ctx context.Context, unit string, actioninfo *entity.ActionRecordInfo) bool {
	spew.Dump("[CaptchaActionEdit]", actioninfo)
	setNum := 10
	// setTime := int64(5) //seconds
	// now := time.Now().Unix()
	if actioninfo.Num >= setNum {
		return false
	}
	return true
}

func (cs *CaptchaService) CaptchaActionInvitationAnswer(ctx context.Context, unit string, actioninfo *entity.ActionRecordInfo) bool {
	spew.Dump("[CaptchaActionInvitationAnswer]", actioninfo)
	setNum := 30
	// setTime := int64(5) //seconds
	// now := time.Now().Unix()
	if actioninfo.Num >= setNum {
		return false
	}
	return true
}

func (cs *CaptchaService) CaptchaActionSearch(ctx context.Context, unit string, actioninfo *entity.ActionRecordInfo) bool {
	now := time.Now().Unix()
	setNum := 20
	setTime := int64(60) //seconds
	spew.Dump("[CaptchaActionSearch]", unit, actioninfo, now-int64(actioninfo.LastTime))
	if now-int64(actioninfo.LastTime) <= setTime && actioninfo.Num >= setNum {
		spew.Dump("in1")
		return false
	}
	if now-actioninfo.LastTime > setTime {
		cs.captchaRepo.SetActionType(ctx, unit, entity.CaptchaActionSearch, "", 0)
	}
	return true
}

func (cs *CaptchaService) CaptchaActionReport(ctx context.Context, unit string, actioninfo *entity.ActionRecordInfo) bool {
	setNum := 30
	setTime := int64(1) //seconds
	now := time.Now().Unix()
	spew.Dump("[CaptchaActionReport]", actioninfo, now-int64(actioninfo.LastTime))
	if now-actioninfo.LastTime <= setTime || actioninfo.Num >= setNum {
		return false
	}
	return true
}

func (cs *CaptchaService) CaptchaActionDelete(ctx context.Context, unit string, actioninfo *entity.ActionRecordInfo) bool {
	spew.Dump("[CaptchaActionDelete]", actioninfo)
	setNum := 5
	setTime := int64(5) //seconds
	now := time.Now().Unix()
	if now-actioninfo.LastTime <= setTime || actioninfo.Num >= setNum {
		return false
	}
	return true
}

func (cs *CaptchaService) CaptchaActionVote(ctx context.Context, unit string, actioninfo *entity.ActionRecordInfo) bool {
	spew.Dump("[CaptchaActionVote]", actioninfo)
	setNum := 40
	// setTime := int64(5) //seconds
	if actioninfo.Num >= setNum {
		return false
	}
	return true
}
