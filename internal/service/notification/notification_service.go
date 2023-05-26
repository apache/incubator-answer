package notification

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/answerdev/answer/internal/base/constant"
	"github.com/answerdev/answer/internal/base/data"
	"github.com/answerdev/answer/internal/base/handler"
	"github.com/answerdev/answer/internal/base/pager"
	"github.com/answerdev/answer/internal/base/translator"
	"github.com/answerdev/answer/internal/entity"
	"github.com/answerdev/answer/internal/schema"
	notficationcommon "github.com/answerdev/answer/internal/service/notification_common"
	"github.com/answerdev/answer/internal/service/revision_common"
	"github.com/answerdev/answer/pkg/uid"
	"github.com/jinzhu/copier"
	"github.com/segmentfault/pacman/log"
)

// NotificationService user service
type NotificationService struct {
	data               *data.Data
	notificationRepo   notficationcommon.NotificationRepo
	notificationCommon *notficationcommon.NotificationCommon
	revisionService    *revision_common.RevisionService
}

func NewNotificationService(
	data *data.Data,
	notificationRepo notficationcommon.NotificationRepo,
	notificationCommon *notficationcommon.NotificationCommon,
	revisionService *revision_common.RevisionService,

) *NotificationService {
	return &NotificationService{
		data:               data,
		notificationRepo:   notificationRepo,
		notificationCommon: notificationCommon,
		revisionService:    revisionService,
	}
}

func (ns *NotificationService) GetRedDot(ctx context.Context, req *schema.GetRedDot) (*schema.RedDot, error) {
	redBot := &schema.RedDot{}
	inboxKey := fmt.Sprintf("answer_RedDot_%d_%s", schema.NotificationTypeInbox, req.UserID)
	achievementKey := fmt.Sprintf("answer_RedDot_%d_%s", schema.NotificationTypeAchievement, req.UserID)
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
	revisionCount := &schema.RevisionSearch{}
	_ = copier.Copy(revisionCount, req)
	if req.CanReviewAnswer || req.CanReviewQuestion || req.CanReviewTag {
		redBot.CanRevision = true
		revisionCountNum, err := ns.revisionService.GetUnreviewedRevisionCount(ctx, revisionCount)
		if err != nil {
			return redBot, err
		}
		redBot.Revision = revisionCountNum
	}

	return redBot, nil
}

func (ns *NotificationService) ClearRedDot(ctx context.Context, req *schema.NotificationClearRequest) (*schema.RedDot, error) {
	botType, ok := schema.NotificationType[req.TypeStr]
	if ok {
		key := fmt.Sprintf("answer_RedDot_%d_%s", botType, req.UserID)
		err := ns.data.Cache.Del(ctx, key)
		if err != nil {
			log.Error("ClearRedDot del cache error", err.Error())
		}
	}
	getRedDotreq := &schema.GetRedDot{}
	_ = copier.Copy(getRedDotreq, req)
	return ns.GetRedDot(ctx, getRedDotreq)
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

func (ns *NotificationService) GetNotificationPage(ctx context.Context, searchCond *schema.NotificationSearch) (
	pageModel *pager.PageModel, err error) {
	resp := make([]*schema.NotificationContent, 0)
	searchType, ok := schema.NotificationType[searchCond.TypeStr]
	if !ok {
		return pager.NewPageModel(0, resp), nil
	}
	searchInboxType := schema.NotificationInboxTypeAll
	if searchType == schema.NotificationTypeInbox {
		_, ok = schema.NotificationInboxType[searchCond.InboxTypeStr]
		if ok {
			searchInboxType = schema.NotificationInboxType[searchCond.InboxTypeStr]
		}
	}
	searchCond.Type = searchType
	searchCond.InboxType = searchInboxType
	notifications, total, err := ns.notificationRepo.GetNotificationPage(ctx, searchCond)
	if err != nil {
		return nil, err
	}
	resp, err = ns.formatNotificationPage(ctx, notifications)
	if err != nil {
		return nil, err
	}
	return pager.NewPageModel(total, resp), nil
}

func (ns *NotificationService) formatNotificationPage(ctx context.Context, notifications []*entity.Notification) (
	resp []*schema.NotificationContent, err error) {
	lang := handler.GetLangByCtx(ctx)
	for _, notificationInfo := range notifications {
		item := &schema.NotificationContent{}
		if err := json.Unmarshal([]byte(notificationInfo.Content), item); err != nil {
			log.Error("NotificationContent Unmarshal Error", err.Error())
			continue
		}
		// If notification is downvote, the user info is not needed.
		if item.NotificationAction == constant.NotificationDownVotedTheQuestion ||
			item.NotificationAction == constant.NotificationDownVotedTheAnswer {
			item.UserInfo = nil
		}

		item.ID = notificationInfo.ID
		item.NotificationAction = translator.Tr(lang, item.NotificationAction)
		item.UpdateTime = notificationInfo.UpdatedAt.Unix()
		item.IsRead = notificationInfo.IsRead == schema.NotificationRead

		if answerID, ok := item.ObjectInfo.ObjectMap["answer"]; ok {
			if item.ObjectInfo.ObjectID == answerID {
				item.ObjectInfo.ObjectID = uid.EnShortID(item.ObjectInfo.ObjectMap["answer"])
			}
			item.ObjectInfo.ObjectMap["answer"] = uid.EnShortID(item.ObjectInfo.ObjectMap["answer"])
		}
		if questionID, ok := item.ObjectInfo.ObjectMap["question"]; ok {
			if item.ObjectInfo.ObjectID == questionID {
				item.ObjectInfo.ObjectID = uid.EnShortID(item.ObjectInfo.ObjectMap["question"])
			}
			item.ObjectInfo.ObjectMap["question"] = uid.EnShortID(item.ObjectInfo.ObjectMap["question"])
		}

		resp = append(resp, item)
	}
	return resp, nil
}
