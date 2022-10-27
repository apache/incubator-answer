package notification

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/answerdev/answer/internal/base/constant"
	"github.com/answerdev/answer/internal/base/data"
	"github.com/answerdev/answer/internal/base/pager"
	"github.com/answerdev/answer/internal/base/translator"
	"github.com/answerdev/answer/internal/schema"
	notficationcommon "github.com/answerdev/answer/internal/service/notification_common"
	"github.com/segmentfault/pacman/i18n"
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

func (ns *NotificationService) GetList(ctx context.Context, search *schema.NotificationSearch) (
	pageModel *pager.PageModel, err error) {
	resp := make([]*schema.NotificationContent, 0)
	searchType, ok := schema.NotificationType[search.TypeStr]
	if !ok {
		return pager.NewPageModel(0, resp), nil
	}
	search.Type = searchType
	notifications, count, err := ns.notificationRepo.SearchList(ctx, search)
	if err != nil {
		return nil, err
	}
	for _, notificationInfo := range notifications {
		item := &schema.NotificationContent{}
		err := json.Unmarshal([]byte(notificationInfo.Content), item)
		if err != nil {
			log.Error("NotificationContent Unmarshal Error", err.Error())
			continue
		}
		lang, _ := ctx.Value(constant.AcceptLanguageFlag).(i18n.Language)
		item.NotificationAction = translator.GlobalTrans.Tr(lang, item.NotificationAction)
		item.ID = notificationInfo.ID
		item.UpdateTime = notificationInfo.UpdatedAt.Unix()
		if notificationInfo.IsRead == schema.NotificationRead {
			item.IsRead = true
		}
		resp = append(resp, item)
	}
	return pager.NewPageModel(count, resp), nil
}
