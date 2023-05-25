package schema

import (
	"strings"

	"github.com/answerdev/answer/internal/base/validator"
)

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
