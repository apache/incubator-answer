package notification

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/segmentfault/answer/internal/base/data"
	"github.com/segmentfault/answer/internal/schema"
	notficationcommon "github.com/segmentfault/answer/internal/service/notification_common"
	"github.com/segmentfault/pacman/log"
)

// NotificationService user service
type NotificationService struct {
	data               *data.Data
	notificationRepo   notficationcommon.NotificationRepo
	notificationCommon *notficationcommon.NotificationCommon
}

func NewNotificationService(
	data *data.Data,
	notificationRepo notficationcommon.NotificationRepo,
	notificationCommon *notficationcommon.NotificationCommon,
) *NotificationService {
	return &NotificationService{
		data:               data,
		notificationRepo:   notificationRepo,
		notificationCommon: notificationCommon,
	}
}

func (ns *NotificationService) GetRedDot(ctx context.Context, userID string) (*schema.RedDot, error) {
	redBot := &schema.RedDot{}
	inboxKey := fmt.Sprintf("answer_RedDot_%d_%s", schema.NotificationTypeInbox, userID)
	achievementKey := fmt.Sprintf("answer_RedDot_%d_%s", schema.NotificationTypeAchievement, userID)
	inboxValue, err := ns.data.Cache.GetInt64(ctx, inboxKey)
	if err != nil {
		redBot.Inbox = 0
	} else {
		redBot.Inbox = inboxValue
	}
	achievementValue, err := ns.data.Cache.GetInt64(ctx, achievementKey)
	if err != nil {
		redBot.Achievement = 0
	} else {
		redBot.Achievement = achievementValue
	}
	return redBot, nil
}

func (ns *NotificationService) ClearRedDot(ctx context.Context, userID string, botTypeStr string) (*schema.RedDot, error) {
	botType, ok := schema.NotificationType[botTypeStr]
	if ok {
		key := fmt.Sprintf("answer_RedDot_%d_%s", botType, userID)
		err := ns.data.Cache.Del(ctx, key)
		if err != nil {
			log.Error("ClearRedDot del cache error", err.Error())
		}
	}
	return ns.GetRedDot(ctx, userID)
}

func (ns *NotificationService) ClearUnRead(ctx context.Context, userID string, botTypeStr string) error {
	botType, ok := schema.NotificationType[botTypeStr]
	if ok {
		err := ns.notificationRepo.ClearUnRead(ctx, userID, botType)
		if err != nil {
			return err
		}
	}
	return nil
}

func (ns *NotificationService) ClearIDUnRead(ctx context.Context, userID string, id string) error {
	notificationInfo, exist, err := ns.notificationRepo.GetById(ctx, id)
	if err != nil {
		log.Error("notificationRepo.GetById error", err.Error())
		return nil
	}
	if !exist {
		return nil
	}
	if notificationInfo.UserID == userID && notificationInfo.IsRead == schema.NotificationNotRead {
		err := ns.notificationRepo.ClearIDUnRead(ctx, userID, id)
		if err != nil {
			return err
		}
	}

	return nil
}

func (ns *NotificationService) GetList(ctx context.Context, search *schema.NotificationSearch) ([]*schema.NotificationContent, int64, error) {
	list := make([]*schema.NotificationContent, 0)
	searchType, ok := schema.NotificationType[search.TypeStr]
	if !ok {
		return list, 0, nil
	}
	search.Type = searchType
	dblist, count, err := ns.notificationRepo.SearchList(ctx, search)
	if err != nil {
		return list, count, err
	}
	for _, dbitem := range dblist {
		item := &schema.NotificationContent{}
		err := json.Unmarshal([]byte(dbitem.Content), item)
		if err != nil {
			log.Error("NotificationContent Unmarshal Error", err.Error())
			continue
		}
		item.ID = dbitem.ID
		item.UpdateTime = dbitem.UpdatedAt.Unix()
		if dbitem.IsRead == schema.NotificationRead {
			item.IsRead = true
		}
		list = append(list, item)
	}
	return list, count, nil
}
