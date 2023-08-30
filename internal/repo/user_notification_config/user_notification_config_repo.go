package user_notification_config

import (
	"context"
	"github.com/answerdev/answer/internal/base/constant"
	"github.com/answerdev/answer/internal/base/data"
	"github.com/answerdev/answer/internal/base/reason"
	"github.com/answerdev/answer/internal/entity"
	"github.com/answerdev/answer/internal/service/user_notification_config"
	"github.com/segmentfault/pacman/errors"
)

// userNotificationConfigRepo notification repository
type userNotificationConfigRepo struct {
	data *data.Data
}

// NewUserNotificationConfigRepo new repository
func NewUserNotificationConfigRepo(data *data.Data) user_notification_config.UserNotificationConfigRepo {
	return &userNotificationConfigRepo{
		data: data,
	}
}

// Add add notification config
func (ur *userNotificationConfigRepo) Add(ctx context.Context, userIDs []string, source, channels string) (err error) {
	var configs []*entity.UserNotificationConfig
	for _, userID := range userIDs {
		configs = append(configs, &entity.UserNotificationConfig{
			UserID:   userID,
			Source:   source,
			Channels: channels,
			Enabled:  true,
		})
	}
	_, err = ur.data.DB.Context(ctx).Insert(configs)
	if err != nil {
		return errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return nil
}

// Save save notification config, if existed, update, if not exist, insert
func (ur *userNotificationConfigRepo) Save(ctx context.Context, uc *entity.UserNotificationConfig) (err error) {
	old := &entity.UserNotificationConfig{UserID: uc.UserID, Source: uc.Source}
	exist, err := ur.data.DB.Context(ctx).Get(old)
	if err != nil {
		return errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	if exist {
		old.Channels = uc.Channels
		old.Enabled = uc.Enabled
		_, err = ur.data.DB.Context(ctx).ID(old.ID).UseBool("enabled").Cols("channels", "enabled").Update(old)
	} else {
		_, err = ur.data.DB.Context(ctx).Insert(uc)
	}
	if err != nil {
		return errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return nil
}

// GetByUserID get notification config by user id
func (ur *userNotificationConfigRepo) GetByUserID(ctx context.Context, userID string) (
	[]*entity.UserNotificationConfig, error) {
	var configs []*entity.UserNotificationConfig
	err := ur.data.DB.Context(ctx).Where("user_id = ?", userID).Find(&configs)
	if err != nil {
		return nil, errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return configs, nil
}

// GetBySource get notification config by source
func (ur *userNotificationConfigRepo) GetBySource(ctx context.Context, source constant.NotificationSource) (
	[]*entity.UserNotificationConfig, error) {
	var configs []*entity.UserNotificationConfig
	err := ur.data.DB.Context(ctx).UseBool("enabled").
		Find(&configs, &entity.UserNotificationConfig{Source: string(source), Enabled: true})
	if err != nil {
		return nil, errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return configs, nil
}

// GetByUserIDAndSource get notification config by user id and source
func (ur *userNotificationConfigRepo) GetByUserIDAndSource(ctx context.Context, userID string, source constant.NotificationSource) (
	conf *entity.UserNotificationConfig, exist bool, err error) {
	config := &entity.UserNotificationConfig{UserID: userID, Source: string(source)}
	exist, err = ur.data.DB.Context(ctx).Get(config)
	if err != nil {
		return nil, false, errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return config, exist, nil
}

// GetByUsersAndSource get notification config by user ids and source
func (ur *userNotificationConfigRepo) GetByUsersAndSource(
	ctx context.Context, userIDs []string, source constant.NotificationSource) (
	[]*entity.UserNotificationConfig, error) {
	var configs []*entity.UserNotificationConfig
	err := ur.data.DB.Context(ctx).UseBool("enabled").In("user_id", userIDs).
		Find(&configs, &entity.UserNotificationConfig{Source: string(source), Enabled: true})
	if err != nil {
		return nil, errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return configs, nil
}
