package notificationcommon

import (
	"context"
	"fmt"
	"time"

	"github.com/answerdev/answer/internal/base/constant"
	"github.com/answerdev/answer/internal/base/data"
	"github.com/answerdev/answer/internal/base/reason"
	"github.com/answerdev/answer/internal/entity"
	"github.com/answerdev/answer/internal/schema"
	"github.com/answerdev/answer/internal/service/activity_common"
	"github.com/answerdev/answer/internal/service/notice_queue"
	"github.com/answerdev/answer/internal/service/object_info"
	usercommon "github.com/answerdev/answer/internal/service/user_common"
	"github.com/answerdev/answer/pkg/uid"
	"github.com/answerdev/answer/plugin"
	"github.com/goccy/go-json"
	"github.com/jinzhu/copier"
	"github.com/segmentfault/pacman/errors"
	"github.com/segmentfault/pacman/log"
)

type NotificationRepo interface {
	AddNotification(ctx context.Context, notification *entity.Notification) (err error)
	GetNotificationPage(ctx context.Context, search *schema.NotificationSearch) ([]*entity.Notification, int64, error)
	ClearUnRead(ctx context.Context, userID string, notificationType int) (err error)
	ClearIDUnRead(ctx context.Context, userID string, id string) (err error)
	GetByUserIdObjectIdTypeId(ctx context.Context, userID, objectID string, notificationType int) (*entity.Notification, bool, error)
	UpdateNotificationContent(ctx context.Context, notification *entity.Notification) (err error)
	GetById(ctx context.Context, id string) (*entity.Notification, bool, error)
}

type NotificationCommon struct {
	data              *data.Data
	notificationRepo  NotificationRepo
	activityRepo      activity_common.ActivityRepo
	followRepo        activity_common.FollowRepo
	userCommon        *usercommon.UserCommon
	objectInfoService *object_info.ObjService
}

func NewNotificationCommon(
	data *data.Data,
	notificationRepo NotificationRepo,
	userCommon *usercommon.UserCommon,
	activityRepo activity_common.ActivityRepo,
	followRepo activity_common.FollowRepo,
	objectInfoService *object_info.ObjService,
) *NotificationCommon {
	notification := &NotificationCommon{
		data:              data,
		notificationRepo:  notificationRepo,
		activityRepo:      activityRepo,
		followRepo:        followRepo,
		userCommon:        userCommon,
		objectInfoService: objectInfoService,
	}
	notification.HandleNotification()
	return notification
}

func (ns *NotificationCommon) HandleNotification() {
	go func() {
		for msg := range notice_queue.NotificationQueue {
			log.Debugf("received notification %+v", msg)
			err := ns.AddNotification(context.TODO(), msg)
			if err != nil {
				log.Error(err)
			}
		}
	}()
}

// AddNotification
// need set
// LoginUserID
// Type  1 inbox 2 achievement
// [inbox] Activity
// [achievement] Rank
// ObjectInfo.Title
// ObjectInfo.ObjectID
// ObjectInfo.ObjectType
func (ns *NotificationCommon) AddNotification(ctx context.Context, msg *schema.NotificationMsg) error {
	if msg.Type == schema.NotificationTypeAchievement && plugin.RankAgentEnabled() {
		return nil
	}
	req := &schema.NotificationContent{
		TriggerUserID:  msg.TriggerUserID,
		ReceiverUserID: msg.ReceiverUserID,
		ObjectInfo: schema.ObjectInfo{
			Title:      msg.Title,
			ObjectID:   uid.DeShortID(msg.ObjectID),
			ObjectType: msg.ObjectType,
		},
		NotificationAction: msg.NotificationAction,
		Type:               msg.Type,
	}
	var questionID string // just for notify all followers
	objInfo, err := ns.objectInfoService.GetInfo(ctx, req.ObjectInfo.ObjectID)
	if err != nil {
		log.Error(err)
	} else {
		req.ObjectInfo.Title = objInfo.Title
		questionID = objInfo.QuestionID
		objectMap := make(map[string]string)
		objectMap["question"] = uid.DeShortID(objInfo.QuestionID)
		objectMap["answer"] = uid.DeShortID(objInfo.AnswerID)
		objectMap["comment"] = objInfo.CommentID
		req.ObjectInfo.ObjectMap = objectMap
	}

	if msg.Type == schema.NotificationTypeAchievement {
		notificationInfo, exist, err := ns.notificationRepo.GetByUserIdObjectIdTypeId(ctx, req.ReceiverUserID, req.ObjectInfo.ObjectID, req.Type)
		if err != nil {
			return fmt.Errorf("get by user id object id type id error: %w", err)
		}
		rank, err := ns.activityRepo.GetUserIDObjectIDActivitySum(ctx, req.ReceiverUserID, req.ObjectInfo.ObjectID)
		if err != nil {
			return fmt.Errorf("get user id object id activity sum error: %w", err)
		}
		req.Rank = rank
		if exist {
			//modify notification
			updateContent := &schema.NotificationContent{}
			err := json.Unmarshal([]byte(notificationInfo.Content), updateContent)
			if err != nil {
				return fmt.Errorf("unmarshal notification content error: %w", err)
			}
			updateContent.Rank = rank
			content, _ := json.Marshal(updateContent)
			notificationInfo.Content = string(content)
			err = ns.notificationRepo.UpdateNotificationContent(ctx, notificationInfo)
			if err != nil {
				return fmt.Errorf("update notification content error: %w", err)
			}
			return nil
		}
	}

	info := &entity.Notification{}
	now := time.Now()
	info.UserID = req.ReceiverUserID
	info.Type = req.Type
	info.IsRead = schema.NotificationNotRead
	info.Status = schema.NotificationStatusNormal
	info.CreatedAt = now
	info.UpdatedAt = now
	info.ObjectID = req.ObjectInfo.ObjectID

	userBasicInfo, exist, err := ns.userCommon.GetUserBasicInfoByID(ctx, req.TriggerUserID)
	if err != nil {
		return fmt.Errorf("get user basic info error: %w", err)
	}
	if !exist {
		return fmt.Errorf("user not exist: %s", req.TriggerUserID)
	}
	req.UserInfo = userBasicInfo
	content, _ := json.Marshal(req)
	_, ok := constant.NotificationMsgTypeMapping[req.NotificationAction]
	if ok {
		info.MsgType = constant.NotificationMsgTypeMapping[req.NotificationAction]
	}
	info.Content = string(content)
	err = ns.notificationRepo.AddNotification(ctx, info)
	if err != nil {
		return fmt.Errorf("add notification error: %w", err)
	}
	err = ns.addRedDot(ctx, info.UserID, info.Type)
	if err != nil {
		log.Error("addRedDot Error", err.Error())
	}

	go ns.SendNotificationToAllFollower(context.Background(), msg, questionID)
	return nil
}

func (ns *NotificationCommon) addRedDot(ctx context.Context, userID string, botType int) error {
	key := fmt.Sprintf("answer_RedDot_%d_%s", botType, userID)
	err := ns.data.Cache.SetInt64(ctx, key, 1, 30*24*time.Hour) //Expiration time is one month.
	if err != nil {
		return errors.InternalServer(reason.UnknownError).WithError(err).WithStack()
	}
	return nil
}

// SendNotificationToAllFollower send notification to all followers
func (ns *NotificationCommon) SendNotificationToAllFollower(ctx context.Context, msg *schema.NotificationMsg,
	questionID string) {
	if msg.NoNeedPushAllFollow {
		return
	}
	if msg.NotificationAction != constant.NotificationUpdateQuestion &&
		msg.NotificationAction != constant.NotificationAnswerTheQuestion &&
		msg.NotificationAction != constant.NotificationUpdateAnswer &&
		msg.NotificationAction != constant.NotificationAcceptAnswer {
		return
	}
	condObjectID := msg.ObjectID
	if len(questionID) > 0 {
		condObjectID = uid.DeShortID(questionID)
	}
	userIDs, err := ns.followRepo.GetFollowUserIDs(ctx, condObjectID)
	if err != nil {
		log.Error(err)
		return
	}
	log.Infof("send notification to all followers: %s %d", condObjectID, len(userIDs))
	for _, userID := range userIDs {
		t := &schema.NotificationMsg{}
		_ = copier.Copy(t, msg)
		t.ReceiverUserID = userID
		t.TriggerUserID = msg.TriggerUserID
		t.NoNeedPushAllFollow = true
		notice_queue.AddNotification(t)
	}
}
