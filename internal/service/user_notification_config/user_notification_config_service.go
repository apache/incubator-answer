package user_notification_config

import (
	"context"
	"github.com/answerdev/answer/internal/base/constant"
	"github.com/answerdev/answer/internal/entity"
	"github.com/answerdev/answer/internal/schema"
	usercommon "github.com/answerdev/answer/internal/service/user_common"
)

type UserNotificationConfigRepo interface {
	Save(ctx context.Context, uc *entity.UserNotificationConfig) (err error)
	GetByUserID(ctx context.Context, userID string) ([]*entity.UserNotificationConfig, error)
	GetBySource(ctx context.Context, source constant.NotificationSource) ([]*entity.UserNotificationConfig, error)
	GetByUserIDAndSource(ctx context.Context, userID string, source constant.NotificationSource) (
		conf *entity.UserNotificationConfig, exist bool, err error)
	GetByUsersAndSource(ctx context.Context, userIDs []string, source constant.NotificationSource) (
		[]*entity.UserNotificationConfig, error)
}

type UserNotificationConfigService struct {
	userRepo                   usercommon.UserRepo
	userNotificationConfigRepo UserNotificationConfigRepo
}

func NewUserNotificationConfigService(
	userRepo usercommon.UserRepo,
	userNotificationConfigRepo UserNotificationConfigRepo,
) *UserNotificationConfigService {
	return &UserNotificationConfigService{
		userRepo:                   userRepo,
		userNotificationConfigRepo: userNotificationConfigRepo,
	}
}

func (us *UserNotificationConfigService) GetUserNotificationConfig(ctx context.Context, userID string) (
	resp *schema.GetUserNotificationConfigResp, err error) {
	notificationConfigs, err := us.userNotificationConfigRepo.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	resp = &schema.GetUserNotificationConfigResp{}
	resp.NotificationConfig = schema.NewNotificationConfig(notificationConfigs)
	resp.Format()
	return resp, nil
}

func (us *UserNotificationConfigService) UpdateUserNotificationConfig(
	ctx context.Context, req *schema.UpdateUserNotificationConfigReq) (err error) {
	req.NotificationConfig.Format()

	err = us.userNotificationConfigRepo.Save(ctx,
		us.convertToEntity(ctx, req.UserID, constant.InboxSource, req.NotificationConfig.Inbox))
	if err != nil {
		return err
	}
	err = us.userNotificationConfigRepo.Save(ctx,
		us.convertToEntity(ctx, req.UserID, constant.AllNewQuestionSource, req.NotificationConfig.AllNewQuestion))
	if err != nil {
		return err
	}
	err = us.userNotificationConfigRepo.Save(ctx,
		us.convertToEntity(ctx, req.UserID, constant.AllNewQuestionForFollowingTagsSource,
			req.NotificationConfig.AllNewQuestionForFollowingTags))
	if err != nil {
		return err
	}
	return nil
}

func (us *UserNotificationConfigService) convertToEntity(ctx context.Context, userID string,
	source constant.NotificationSource, channels schema.NotificationChannels) (c *entity.UserNotificationConfig) {
	c = &entity.UserNotificationConfig{
		UserID:   userID,
		Source:   string(source),
		Channels: channels.ToJsonString(),
	}
	for _, ch := range channels {
		if ch.Enable {
			c.Enabled = true
			break
		}
	}
	return c
}

func (us *UserNotificationConfigService) CheckEnable(
	ctx context.Context, userID string, source constant.NotificationSource,
	channel constant.NotificationChannelKey) (enable bool, err error) {
	conf, exist, err := us.userNotificationConfigRepo.GetByUserIDAndSource(ctx, userID, source)
	if err != nil {
		return false, err
	}
	if !exist {
		return false, nil
	}
	notificationChannels := schema.NewNotificationChannelsFormJson(conf.Channels)
	return notificationChannels.CheckEnable(channel), nil
}
