package notification

import (
	"context"
	"github.com/answerdev/answer/internal/base/constant"
	"github.com/answerdev/answer/internal/schema"
	"github.com/segmentfault/pacman/i18n"
	"github.com/segmentfault/pacman/log"
	"time"
)

func (ns *ExternalNotificationService) handleInviteAnswerNotification(ctx context.Context,
	msg *schema.ExternalNotificationMsg) error {
	log.Debugf("try to send invite answer notification %+v", msg)

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
			ns.sendInviteAnswerNotificationEmail(ctx, msg.ReceiverUserID, msg.ReceiverEmail, msg.ReceiverLang, msg.NewInviteAnswerTemplateRawData)
		}
	}
	return nil
}

func (ns *ExternalNotificationService) sendInviteAnswerNotificationEmail(ctx context.Context,
	userID, email, lang string, rawData *schema.NewInviteAnswerTemplateRawData) {
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
	title, body, err := ns.emailService.NewInviteAnswerTemplate(ctx, rawData)
	if err != nil {
		log.Error(err)
		return
	}

	ns.emailService.SendAndSaveCodeWithTime(
		ctx, email, title, body, rawData.UnsubscribeCode, codeContent.ToJSONString(), 1*24*time.Hour)
}
