package report_handle_admin

import (
	"context"

	"github.com/answerdev/answer/internal/service/config"

	"github.com/answerdev/answer/internal/base/constant"
	"github.com/answerdev/answer/internal/entity"
	"github.com/answerdev/answer/internal/schema"
	"github.com/answerdev/answer/internal/service/comment"
	"github.com/answerdev/answer/internal/service/notice_queue"
	questioncommon "github.com/answerdev/answer/internal/service/question_common"
	"github.com/answerdev/answer/pkg/obj"
)

type ReportHandle struct {
	questionCommon *questioncommon.QuestionCommon
	commentRepo    comment.CommentRepo
	configService  *config.ConfigService
}

func NewReportHandle(
	questionCommon *questioncommon.QuestionCommon,
	commentRepo comment.CommentRepo,
	configService *config.ConfigService) *ReportHandle {
	return &ReportHandle{
		questionCommon: questionCommon,
		commentRepo:    commentRepo,
		configService:  configService,
	}
}

// HandleObject this handle object status
func (rh *ReportHandle) HandleObject(ctx context.Context, reported *entity.Report, req schema.ReportHandleReq) (err error) {
	reasonDeleteCfg, err := rh.configService.GetConfigByKey(ctx, "reason.needs_delete")
	if err != nil {
		return err
	}
	reasonCloseCfg, err := rh.configService.GetConfigByKey(ctx, "reason.needs_close")
	if err != nil {
		return err
	}
	var (
		objectID       = reported.ObjectID
		reportedUserID = reported.ReportedUserID
		objectKey      string
	)

	objectKey, err = obj.GetObjectTypeStrByObjectID(objectID)
	if err != nil {
		return err
	}
	switch objectKey {
	case "question":
		switch req.FlaggedType {
		case reasonDeleteCfg.ID:
			err = rh.questionCommon.RemoveQuestion(ctx, &schema.RemoveQuestionReq{ID: objectID})
		case reasonCloseCfg.ID:
			err = rh.questionCommon.CloseQuestion(ctx, &schema.CloseQuestionReq{
				ID:        objectID,
				CloseType: req.FlaggedType,
				CloseMsg:  req.FlaggedContent,
			})
		}
	case "answer":
		switch req.FlaggedType {
		case reasonDeleteCfg.ID:
			err = rh.questionCommon.RemoveAnswer(ctx, objectID)
		}
	case "comment":
		switch req.FlaggedType {
		case reasonCloseCfg.ID:
			err = rh.commentRepo.RemoveComment(ctx, objectID)
			rh.sendNotification(ctx, reportedUserID, objectID, constant.NotificationYourCommentWasDeleted)
		}
	}
	return
}

// sendNotification send rank triggered notification
func (rh *ReportHandle) sendNotification(ctx context.Context, reportedUserID, objectID, notificationAction string) {
	msg := &schema.NotificationMsg{
		TriggerUserID:      reportedUserID,
		ReceiverUserID:     reportedUserID,
		Type:               schema.NotificationTypeInbox,
		ObjectID:           objectID,
		ObjectType:         constant.ReportObjectType,
		NotificationAction: notificationAction,
	}
	notice_queue.AddNotification(msg)
}
