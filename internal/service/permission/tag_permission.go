package permission

import (
	"context"

	"github.com/answerdev/answer/internal/schema"
)

// GetTagPermission get tag permission
func GetTagPermission(ctx context.Context, canEdit, canDelete bool) (
	actions []*schema.PermissionMemberAction) {
	actions = make([]*schema.PermissionMemberAction, 0)
	if canEdit {
		actions = append(actions, &schema.PermissionMemberAction{
			Action: "edit",
			Name:   "Edit",
			Type:   "edit",
		})
	}

	if canDelete {
		actions = append(actions, &schema.PermissionMemberAction{
			Action: "delete",
			Name:   "Delete",
			Type:   "reason",
		})
	}
	return actions
}

// GetTagSynonymPermission get tag synonym permission
func GetTagSynonymPermission(ctx context.Context, canEdit bool) (
	actions []*schema.PermissionMemberAction) {
	actions = make([]*schema.PermissionMemberAction, 0)
	if canEdit {
		actions = append(actions, &schema.PermissionMemberAction{
			Action: "edit",
			Name:   "Edit",
			Type:   "edit",
		})
	}
	return actions
}
