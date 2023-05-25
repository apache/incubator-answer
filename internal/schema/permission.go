package schema

import (
	"strings"

	"github.com/answerdev/answer/internal/base/reason"
	"github.com/answerdev/answer/internal/base/translator"
	"github.com/answerdev/answer/internal/base/validator"
	"github.com/segmentfault/pacman/i18n"
)

// PermissionTrTplData template data as for translate permission message
type PermissionTrTplData struct {
	Rank int
}

// PermissionMemberAction permission member action
type PermissionMemberAction struct {
	Action string `json:"action"`
	Name   string `json:"name"`
	Type   string `json:"type"`
}

// GetPermissionReq get permission request
type GetPermissionReq struct {
	Action  string   `form:"action"`
	Actions []string `validate:"omitempty" form:"actions"`
}

func (r *GetPermissionReq) Check() (errField []*validator.FormErrorField, err error) {
	if len(r.Action) > 0 {
		r.Actions = strings.Split(r.Action, ",")
	}
	return nil, nil
}

// GetPermissionResp get permission response
type GetPermissionResp struct {
	HasPermission bool `json:"has_permission"`
	// only not allow, will return this tip
	NoPermissionTip string `json:"no_permission_tip"`
}

func (r *GetPermissionResp) TrTip(lang i18n.Language, requireRank int) {
	if r.HasPermission {
		return
	}
	if requireRank <= 0 {
		r.NoPermissionTip = translator.Tr(lang, reason.RankFailToMeetTheCondition)
	} else {
		r.NoPermissionTip = translator.TrWithData(
			lang, reason.NoEnoughRankToOperate, &PermissionTrTplData{Rank: requireRank})
	}
}
