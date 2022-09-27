package report_handle_backyard

import (
	"context"
	"github.com/segmentfault/answer/internal/service/config"

	"github.com/segmentfault/answer/internal/base/constant"
	"github.com/segmentfault/answer/internal/entity"
	"github.com/segmentfault/answer/internal/schema"
	"github.com/segmentfault/answer/internal/service/comment"
	"github.com/segmentfault/answer/internal/service/notice_queue"
	questioncommon "github.com/segmentfault/answer/internal/service/question_common"
	"github.com/segmentfault/answer/pkg/obj"
)

type ReportHandle struct {
	questionCommon *questioncommon.QuestionCommon
	commentRepo    comment.CommentRepo
	configRepo     config.ConfigRepo
}

func NewReportHandle(
	questionCommon *questioncommon.QuestionCommon,
	commentRepo comment.CommentRepo,
	configRepo config.ConfigRepo) *ReportHandle {
	return &ReportHandle{
		questionCommon: questionCommon,
		commentRepo:    commentRepo,
		configRepo:     configRepo,
	}
}

// HandleObject this handle object status
func (rh *ReportHandle) HandleObject(ctx context.Context, reported entity.Report, req schema.ReportHandleReq) (err error) {
	var (
		objectID        = reported.ObjectID
		reportedUserID  = reported.ReportedUserID
		objectKey       string
		reasonDelete, _ = rh.configRepo.GetConfigType("reason.needs_delete")
		reasonClose, _  = rh.configRepo.GetConfigType("reason.needs_close")
	)

	objectKey, err = obj.GetObjectTypeStrByObjectID(objectID)
	if err != nil {
		return err
	}
	switch objectKey {
	case "question":
		switch req.FlagedType {
		case reasonDelete:
			err = rh.questionCommon.RemoveQuestion(ctx, &schema.RemoveQuestionReq{ID: objectID})
		case reasonClose:
			err = rh.questionCommon.CloseQuestion(ctx, &schema.CloseQuestionReq{
				ID:        objectID,
				CloseType: req.FlagedType,
				CloseMsg:  req.FlagedContent,
			})
		}
	case "answer":
		switch req.FlagedType {
		case reasonDelete:
			err = rh.questionCommon.RemoveAnswer(ctx, objectID)
		}
	case "comment":
		switch req.FlagedType {
		case reasonDelete:
			err = rh.commentRepo.RemoveComment(ctx, objectID)
			rh.sendNotification(ctx, reportedUserID, objectID, constant.YourCommentWasDeleted)
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
