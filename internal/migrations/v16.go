package migrations

import (
	"context"
	"github.com/answerdev/answer/internal/base/constant"
	"github.com/answerdev/answer/internal/entity"
	"github.com/segmentfault/pacman/log"
	"xorm.io/xorm"
)

func setDefaultUserNotificationConfig(ctx context.Context, x *xorm.Engine) error {
	userIDs := make([]string, 0)
	err := x.Context(ctx).Table("user").Select("id").Find(&userIDs)
	if err != nil {
		return err
	}

	for _, id := range userIDs {
		bean := entity.UserNotificationConfig{UserID: id, Source: string(constant.InboxSource)}
		exist, err := x.Context(ctx).Get(&bean)
		if err != nil {
			log.Error(err)
		}
		if exist {
			continue
		}
		_, err = x.Context(ctx).Insert(&entity.UserNotificationConfig{
			UserID:   id,
			Source:   string(constant.InboxSource),
			Channels: `[{"key":"email","enable":true}]`,
			Enabled:  true,
		})
		if err != nil {
			log.Error(err)
		}
	}
	return nil
}
