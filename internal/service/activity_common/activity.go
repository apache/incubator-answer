package activity_common

import (
	"context"

	"github.com/answerdev/answer/internal/entity"
	"github.com/answerdev/answer/internal/service/activity_queue"
	"github.com/segmentfault/pacman/log"
	"xorm.io/xorm"
)

type ActivityRepo interface {
	GetActivityTypeByObjID(ctx context.Context, objectId string, action string) (activityType, rank int, hasRank int, err error)
	GetActivityTypeByObjKey(ctx context.Context, objectKey, action string) (activityType int, err error)
	GetActivity(ctx context.Context, session *xorm.Session, objectID, userID string, activityType int) (
		existsActivity *entity.Activity, exist bool, err error)
	GetUserIDObjectIDActivitySum(ctx context.Context, userID, objectID string) (int, error)
	GetActivityTypeByConfigKey(ctx context.Context, configKey string) (activityType int, err error)
	AddActivity(ctx context.Context, activity *entity.Activity) (err error)
}

type ActivityCommon struct {
	activityRepo ActivityRepo
}

// NewActivityCommon new activity common
func NewActivityCommon(
	activityRepo ActivityRepo,
) *ActivityCommon {
	activity := &ActivityCommon{
		activityRepo: activityRepo,
	}
	activity.HandleActivity()
	return activity
}

// HandleActivity handle activity message
func (ac *ActivityCommon) HandleActivity() {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Error(err)
			}
		}()

		for msg := range activity_queue.ActivityQueue {
			log.Debugf("received activity %+v", msg)

			activityType, err := ac.activityRepo.GetActivityTypeByConfigKey(context.Background(), string(msg.ActivityTypeKey))
			if err != nil {
				log.Errorf("error getting activity type %s, activity type is %s", err, activityType)
			}

			act := &entity.Activity{
				UserID:        msg.UserID,
				TriggerUserID: msg.TriggerUserID,
				ObjectID:      msg.ObjectID,
				ActivityType:  activityType,
				Cancelled:     entity.ActivityAvailable,
			}
			if err := ac.activityRepo.AddActivity(context.TODO(), act); err != nil {
				log.Error(err)
			}
		}
	}()
}
