package permission

import (
	"context"

	"github.com/answerdev/answer/internal/schema"
)

// TODO: There is currently no permission management
func GetCommentPermission(ctx context.Context, userID string, commentCreatorUserID string) (
	actions []*schema.PermissionMemberAction) {
	actions = make([]*schema.PermissionMemberAction, 0)
	if len(userID) > 0 {
		actions = append(actions, &schema.PermissionMemberAction{
			Action: "report",
			Name:   "Flag",
			Type:   "reason",
		})
	}
	if userID != commentCreatorUserID {
		return actions
	}
	actions = append(actions, []*schema.PermissionMemberAction{
		{
			Action: "edit",
			Name:   "Edit",
			Type:   "edit",
		},
		{
			Action: "delete",
			Name:   "Delete",
			Type:   "reason",
		},
	}...)
	return actions
}

func GetTagPermission(ctx context.Context, userID string, tagCreatorUserID string) (
	actions []*schema.PermissionMemberAction) {
	if userID != tagCreatorUserID {
		return []*schema.PermissionMemberAction{}
	}
	return []*schema.PermissionMemberAction{
		{
			Action: "edit",
			Name:   "Edit",
			Type:   "edit",
		},
		{
			Action: "delete",
			Name:   "Delete",
			Type:   "reason",
		},
	}
}

func GetAnswerPermission(ctx context.Context, userID string, answerAuthID string, isAdmin bool) (
	actions []*schema.PermissionMemberAction) {
	actions = make([]*schema.PermissionMemberAction, 0)
	if !isAdmin {
		if len(userID) > 0 {
			actions = append(actions, &schema.PermissionMemberAction{
				Action: "report",
				Name:   "Flag",
				Type:   "reason",
			})
		}
		if userID != answerAuthID {
			return actions
		}
	}
	actions = append(actions, []*schema.PermissionMemberAction{
		{
			Action: "edit",
			Name:   "Edit",
			Type:   "edit",
		},
		{
			Action: "delete",
			Name:   "Delete",
			Type:   "confirm",
		},
	}...)
	return actions
}

func GetQuestionPermission(ctx context.Context, userID string, questionAuthID string, isAdmin bool) (actions []*schema.PermissionMemberAction) {
	actions = make([]*schema.PermissionMemberAction, 0)
	if !isAdmin {
		if len(userID) > 0 {
			actions = append(actions, &schema.PermissionMemberAction{
				Action: "report",
				Name:   "Flag",
				Type:   "reason",
			})
		}
		if userID != questionAuthID {
			return actions
		}
	}
	actions = append(actions, []*schema.PermissionMemberAction{
		{
			Action: "edit",
			Name:   "Edit",
			Type:   "edit",
		},
		{
			Action: "close",
			Name:   "Close",
			Type:   "confirm",
		},
		{
			Action: "delete",
			Name:   "Delete",
			Type:   "confirm",
		},
	}...)
	return actions
}
