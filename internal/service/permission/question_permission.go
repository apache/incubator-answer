package permission

import (
	"context"

	"github.com/answerdev/answer/internal/base/handler"
	"github.com/answerdev/answer/internal/base/translator"
	"github.com/answerdev/answer/internal/schema"
)

// GetQuestionPermission get question permission
func GetQuestionPermission(ctx context.Context, userID string, creatorUserID string,
	canEdit, canDelete, canClose, canReopen, canPin, canHide, CanUnPin, canShow bool) (
	actions []*schema.PermissionMemberAction) {
	lang := handler.GetLangByCtx(ctx)
	actions = make([]*schema.PermissionMemberAction, 0)
	if len(userID) > 0 {
		actions = append(actions, &schema.PermissionMemberAction{
			Action: "report",
			Name:   translator.Tr(lang, reportActionName),
			Type:   "reason",
		})
	}
	if canEdit || userID == creatorUserID {
		actions = append(actions, &schema.PermissionMemberAction{
			Action: "edit",
			Name:   translator.Tr(lang, editActionName),
			Type:   "edit",
		})
	}
	if canClose {
		actions = append(actions, &schema.PermissionMemberAction{
			Action: "close",
			Name:   translator.Tr(lang, closeActionName),
			Type:   "confirm",
		})
	}
	if canReopen {
		actions = append(actions, &schema.PermissionMemberAction{
			Action: "reopen",
			Name:   translator.Tr(lang, reopenActionName),
			Type:   "confirm",
		})
	}
	if canPin {
		actions = append(actions, &schema.PermissionMemberAction{
			Action: "pin",
			Name:   translator.Tr(lang, pinActionName),
			Type:   "confirm",
		})
	}
	if canHide {
		actions = append(actions, &schema.PermissionMemberAction{
			Action: "hide",
			Name:   translator.Tr(lang, hideActionName),
			Type:   "confirm",
		})
	}

	if CanUnPin {
		actions = append(actions, &schema.PermissionMemberAction{
			Action: "unpin",
			Name:   translator.Tr(lang, unpinActionName),
			Type:   "confirm",
		})
	}

	if canShow {
		actions = append(actions, &schema.PermissionMemberAction{
			Action: "show",
			Name:   translator.Tr(lang, showActionName),
			Type:   "confirm",
		})
	}
	if canDelete || userID == creatorUserID {
		actions = append(actions, &schema.PermissionMemberAction{
			Action: "delete",
			Name:   translator.Tr(lang, deleteActionName),
			Type:   "confirm",
		})
	}
	return actions
}

// GetQuestionExtendsPermission get question extends permission
func GetQuestionExtendsPermission(ctx context.Context, userID string, creatorUserID string,
	canInviteOtherToAnswer bool) (
	actions []*schema.PermissionMemberAction) {
	lang := handler.GetLangByCtx(ctx)
	actions = make([]*schema.PermissionMemberAction, 0)
	if canInviteOtherToAnswer || userID == creatorUserID {
		actions = append(actions, &schema.PermissionMemberAction{
			Action: "invite_other_to_answer",
			Name:   translator.Tr(lang, inviteSomeoneToAnswerActionName),
			Type:   "confirm",
		})
	}
	return actions
}
