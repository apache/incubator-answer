package permission

import "github.com/answerdev/answer/internal/schema"

// TODO: There is currently no permission management
func GetCommentPermission(userID string, commentCreatorUserID string) (
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

func GetTagPermission(userID string, tagCreatorUserID string) (
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

func GetAnswerPermission(userID string, answerAuthID string) (
	actions []*schema.PermissionMemberAction) {
	actions = make([]*schema.PermissionMemberAction, 0)
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

func GetQuestionPermission(userID string, questionAuthID string) (
	actions []*schema.PermissionMemberAction) {
	actions = make([]*schema.PermissionMemberAction, 0)
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
