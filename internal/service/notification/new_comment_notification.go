package notification

import (
	"context"
	"github.com/answerdev/answer/internal/base/constant"
	"github.com/answerdev/answer/internal/schema"
	"github.com/segmentfault/pacman/i18n"
	"github.com/segmentfault/pacman/log"
	"time"
)

func (ns *ExternalNotificationService) handleNewCommentNotification(ctx context.Context,
	msg *schema.ExternalNotificationMsg) error {
	log.Debugf("try to send new comment notification %+v", msg)

	notificationConfig, exist, err := ns.userNotificationConfigRepo.GetByUserIDAndSource(ctx, msg.ReceiverUserID, constant.InboxSource)
	if err != nil {
		return err
	}
	if !exist {
		return nil
	}
	channels := schema.NewNotificationChannelsFormJson(notificationConfig.Channels)
	for _, channel := range channels {
		if !channel.Enable {
			continue
		}
		switch channel.Key {
		case constant.EmailChannel:
			ns.sendNewCommentNotificationEmail(ctx, msg.ReceiverUserID, msg.ReceiverEmail, msg.ReceiverLang, msg.NewCommentTemplateRawData)
		}
	}

	return nil
}

func (ns *ExternalNotificationService) sendNewCommentNotificationEmail(ctx context.Context,
	userID, email, lang string, rawData *schema.NewCommentTemplateRawData) {
	codeContent := &schema.EmailCodeContent{
		SourceType: schema.UnsubscribeSourceType,
		NotificationSources: []constant.NotificationSource{
			constant.InboxSource,
		},
		Email:  email,
		UserID: userID,
	}
	// If receiver has set language, use it to send email.
	if len(lang) > 0 {
		ctx = context.WithValue(ctx, constant.AcceptLanguageFlag, i18n.Language(lang))
	}
	title, body, err := ns.emailService.NewCommentTemplate(ctx, rawData)
	if err != nil {
		log.Error(err)
		return
	}

	ns.emailService.SendAndSaveCodeWithTime(
		ctx, email, title, body, rawData.UnsubscribeCode, codeContent.ToJSONString(), 1*24*time.Hour)
}
