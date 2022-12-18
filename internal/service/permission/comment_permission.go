package permission

import (
	"context"

	"github.com/answerdev/answer/internal/schema"
)

// GetCommentPermission get comment permission
func GetCommentPermission(ctx context.Context, userID string, creatorUserID string, canEdit, canDelete bool) (
	actions []*schema.PermissionMemberAction) {
	actions = make([]*schema.PermissionMemberAction, 0)
	if len(userID) > 0 {
		actions = append(actions, &schema.PermissionMemberAction{
			Action: "report",
			Name:   "Flag",
			Type:   "reason",
		})
	}
	if canEdit || userID == creatorUserID {
		actions = append(actions, &schema.PermissionMemberAction{
			Action: "edit",
			Name:   "Edit",
			Type:   "edit",
		})
	}

	if canDelete || userID == creatorUserID {
		actions = append(actions, &schema.PermissionMemberAction{
			Action: "delete",
			Name:   "Delete",
			Type:   "reason",
		})
	}
	return actions
}
