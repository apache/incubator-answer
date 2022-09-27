package schema

const PermissionMemberActionTypeEdit = "edit"
const PermissionMemberActionTypeReason = "reason"

// PermissionMemberAction permission member action
type PermissionMemberAction struct {
	Action string `json:"action"`
	Name   string `json:"name"`
	Type   string `json:"type"`
}
