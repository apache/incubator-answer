package activity_common

import (
	"context"
	"time"

	"github.com/answerdev/answer/internal/entity"
	"github.com/answerdev/answer/internal/schema"
	"github.com/answerdev/answer/internal/service/activity_queue"
	"github.com/answerdev/answer/pkg/converter"
	"github.com/answerdev/answer/pkg/uid"
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
	GetUsersWhoHasGainedTheMostReputation(
		ctx context.Context, startTime, endTime time.Time, limit int) (rankStat []*entity.ActivityUserRankStat, err error)
	GetUsersWhoHasVoteMost(
		ctx context.Context, startTime, endTime time.Time, limit int) (voteStat []*entity.ActivityUserVoteStat, err error)
}

type ActivityCommon struct {
	activityRepo         ActivityRepo
	activityQueueService activity_queue.ActivityQueueService
}

// NewActivityCommon new activity common
func NewActivityCommon(
	activityRepo ActivityRepo,
	activityQueueService activity_queue.ActivityQueueService,
) *ActivityCommon {
	activity := &ActivityCommon{
		activityRepo:         activityRepo,
		activityQueueService: activityQueueService,
	}
	activity.activityQueueService.RegisterHandler(activity.HandleActivity)
	return activity
}

// HandleActivity handle activity message
func (ac *ActivityCommon) HandleActivity(ctx context.Context, msg *schema.ActivityMsg) error {
	activityType, err := ac.activityRepo.GetActivityTypeByConfigKey(context.Background(), string(msg.ActivityTypeKey))
	if err != nil {
		log.Errorf("error getting activity type %s, activity type is %d", err, activityType)
		return err
	}

	act := &entity.Activity{
		UserID:           msg.UserID,
		TriggerUserID:    msg.TriggerUserID,
		ObjectID:         uid.DeShortID(msg.ObjectID),
		OriginalObjectID: uid.DeShortID(msg.OriginalObjectID),
		ActivityType:     activityType,
		Cancelled:        entity.ActivityAvailable,
	}
	if len(msg.RevisionID) > 0 {
		act.RevisionID = converter.StringToInt64(msg.RevisionID)
	}
	if err := ac.activityRepo.AddActivity(ctx, act); err != nil {
		return err
	}
	return nil
}
