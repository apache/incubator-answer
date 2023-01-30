package permission

import (
	"context"
	"time"

	"github.com/answerdev/answer/internal/base/constant"
	"github.com/answerdev/answer/internal/schema"
)

// GetCommentPermission get comment permission
func GetCommentPermission(ctx context.Context, userID string, creatorUserID string,
	createdAt time.Time, canEdit, canDelete bool) (actions []*schema.PermissionMemberAction) {
	actions = make([]*schema.PermissionMemberAction, 0)
	if len(userID) > 0 {
		actions = append(actions, &schema.PermissionMemberAction{
			Action: "report",
			Name:   "Flag",
			Type:   "reason",
		})
	}
	deadline := createdAt.Add(constant.CommentEditDeadline)
	if canEdit || (userID == creatorUserID && time.Now().Before(deadline)) {
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
