package notification

import (
	"context"
	"github.com/answerdev/answer/internal/base/constant"
	"github.com/answerdev/answer/internal/schema"
	"github.com/answerdev/answer/pkg/token"
	"github.com/segmentfault/pacman/i18n"
	"github.com/segmentfault/pacman/log"
	"time"
)

type NewQuestionSubscriber struct {
	UserID   string                      `json:"user_id"`
	Channels schema.NotificationChannels `json:"channels"`
}

func (ns *ExternalNotificationService) handleNewQuestionNotification(ctx context.Context,
	msg *schema.ExternalNotificationMsg) error {
	log.Debugf("try to send new question notification %+v", msg)
	subscribers, err := ns.getNewQuestionSubscribers(ctx, msg)
	if err != nil {
		return err
	}
	log.Debugf("get subscribers %d for question %s", len(subscribers), msg.NewQuestionTemplateRawData.QuestionID)

	for _, subscriber := range subscribers {
		for _, channel := range subscriber.Channels {
			if !channel.Enable {
				continue
			}
			switch channel.Key {
			case constant.EmailChannel:
				ns.sendNewQuestionNotificationEmail(ctx, subscriber.UserID, &schema.NewQuestionTemplateRawData{
					QuestionTitle:   msg.NewQuestionTemplateRawData.QuestionTitle,
					QuestionID:      msg.NewQuestionTemplateRawData.QuestionID,
					UnsubscribeCode: token.GenerateToken(),
					Tags:            msg.NewQuestionTemplateRawData.Tags,
					TagIDs:          msg.NewQuestionTemplateRawData.TagIDs,
				})
			}
		}
	}
	return nil
}

func (ns *ExternalNotificationService) getNewQuestionSubscribers(ctx context.Context, msg *schema.ExternalNotificationMsg) (
	subscribers []*NewQuestionSubscriber, err error) {
	subscribersMapping := make(map[string]*NewQuestionSubscriber)

	// 1. get all this new question's tags followers
	tagsFollowerIDs := make([]string, 0)
	followerMapping := make(map[string]bool)
	for _, tagID := range msg.NewQuestionTemplateRawData.TagIDs {
		userIDs, err := ns.followRepo.GetFollowUserIDs(ctx, tagID)
		if err != nil {
			log.Error(err)
			continue
		}
		for _, userID := range userIDs {
			if _, ok := followerMapping[userID]; ok {
				continue
			}
			followerMapping[userID] = true
			tagsFollowerIDs = append(tagsFollowerIDs, userID)
		}
	}
	userNotificationConfigs, err := ns.userNotificationConfigRepo.GetByUsersAndSource(
		ctx, tagsFollowerIDs, constant.AllNewQuestionForFollowingTagsSource)
	if err != nil {
		return nil, err
	}
	for _, userNotificationConfig := range userNotificationConfigs {
		if _, ok := subscribersMapping[userNotificationConfig.UserID]; ok {
			continue
		}
		subscribersMapping[userNotificationConfig.UserID] = &NewQuestionSubscriber{
			UserID:   userNotificationConfig.UserID,
			Channels: schema.NewNotificationChannelsFormJson(userNotificationConfig.Channels),
		}
		subscribers = append(subscribers, subscribersMapping[userNotificationConfig.UserID])
	}
	log.Debugf("get %d subscribers from tags", len(subscribersMapping))

	// 2. get all new question's followers
	notificationConfigs, err := ns.userNotificationConfigRepo.GetBySource(ctx, constant.AllNewQuestionSource)
	if err != nil {
		return nil, err
	}
	for _, notificationConfig := range notificationConfigs {
		if _, ok := subscribersMapping[notificationConfig.UserID]; ok {
			continue
		}
		if ns.checkSendNewQuestionNotificationEmailLimit(ctx, notificationConfig.UserID) {
			continue
		}
		subscribersMapping[notificationConfig.UserID] = &NewQuestionSubscriber{
			UserID:   notificationConfig.UserID,
			Channels: schema.NewNotificationChannelsFormJson(notificationConfig.Channels),
		}
		subscribers = append(subscribers, subscribersMapping[notificationConfig.UserID])
	}
	log.Debugf("get %d subscribers from all new question config", len(subscribers))
	return subscribers, nil
}

func (ns *ExternalNotificationService) checkSendNewQuestionNotificationEmailLimit(ctx context.Context, userID string) bool {
	// TODO: check if reach send limit
	return false
}

func (ns *ExternalNotificationService) sendNewQuestionNotificationEmail(ctx context.Context,
	userID string, rawData *schema.NewQuestionTemplateRawData) {
	userInfo, exist, err := ns.userRepo.GetByUserID(ctx, userID)
	if err != nil {
		log.Error(err)
		return
	}
	if !exist {
		log.Errorf("user %s not exist", userID)
		return
	}
	// If receiver has set language, use it to send email.
	if len(userInfo.Language) > 0 {
		ctx = context.WithValue(ctx, constant.AcceptLanguageFlag, i18n.Language(userInfo.Language))
	}
	title, body, err := ns.emailService.NewQuestionTemplate(ctx, rawData)
	if err != nil {
		log.Error(err)
		return
	}

	codeContent := &schema.EmailCodeContent{
		SourceType: schema.UnsubscribeSourceType,
		Email:      userInfo.EMail,
		UserID:     userID,
		NotificationSources: []constant.NotificationSource{
			constant.AllNewQuestionSource,
			constant.AllNewQuestionForFollowingTagsSource,
		},
	}
	ns.emailService.SendAndSaveCodeWithTime(
		ctx, userInfo.EMail, title, body, rawData.UnsubscribeCode, codeContent.ToJSONString(), 1*24*time.Hour)
}
