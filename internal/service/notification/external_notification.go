package notification

import (
	"context"
	"github.com/answerdev/answer/internal/base/data"
	"github.com/answerdev/answer/internal/schema"
	"github.com/answerdev/answer/internal/service/activity_common"
	"github.com/answerdev/answer/internal/service/export"
	"github.com/answerdev/answer/internal/service/notice_queue"
	usercommon "github.com/answerdev/answer/internal/service/user_common"
	"github.com/answerdev/answer/internal/service/user_notification_config"
	"github.com/segmentfault/pacman/log"
)

type ExternalNotificationService struct {
	data                       *data.Data
	userNotificationConfigRepo user_notification_config.UserNotificationConfigRepo
	followRepo                 activity_common.FollowRepo
	emailService               *export.EmailService
	userRepo                   usercommon.UserRepo
	notificationQueueService   notice_queue.ExternalNotificationQueueService
}

func NewExternalNotificationService(
	data *data.Data,
	userNotificationConfigRepo user_notification_config.UserNotificationConfigRepo,
	followRepo activity_common.FollowRepo,
	emailService *export.EmailService,
	userRepo usercommon.UserRepo,
	notificationQueueService notice_queue.ExternalNotificationQueueService,
) *ExternalNotificationService {
	n := &ExternalNotificationService{
		data:                       data,
		userNotificationConfigRepo: userNotificationConfigRepo,
		followRepo:                 followRepo,
		emailService:               emailService,
		userRepo:                   userRepo,
		notificationQueueService:   notificationQueueService,
	}
	notificationQueueService.RegisterHandler(n.Handler)
	return n
}

func (ns *ExternalNotificationService) Handler(ctx context.Context, msg *schema.ExternalNotificationMsg) error {
	log.Debugf("try to send external notification %+v", msg)

	if msg.NewQuestionTemplateRawData != nil {
		return ns.handleNewQuestionNotification(ctx, msg)
	}
	if msg.NewCommentTemplateRawData != nil {
		return ns.handleNewCommentNotification(ctx, msg)
	}
	if msg.NewAnswerTemplateRawData != nil {
		return ns.handleNewAnswerNotification(ctx, msg)
	}
	if msg.NewInviteAnswerTemplateRawData != nil {
		return ns.handleInviteAnswerNotification(ctx, msg)
	}
	log.Errorf("unknown notification message: %+v", msg)
	return nil
}
