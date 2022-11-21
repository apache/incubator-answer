package activity_common

import (
	"context"
	"fmt"

	"github.com/answerdev/answer/internal/entity"
	"github.com/answerdev/answer/internal/service/activity_common"
	"github.com/answerdev/answer/pkg/obj"
	"xorm.io/builder"
	"xorm.io/xorm"

	"github.com/answerdev/answer/internal/base/data"
	"github.com/answerdev/answer/internal/base/reason"
	"github.com/answerdev/answer/internal/service/config"
	"github.com/answerdev/answer/internal/service/unique"
	"github.com/segmentfault/pacman/errors"
)

// ActivityRepo activity repository
type ActivityRepo struct {
	data         *data.Data
	uniqueIDRepo unique.UniqueIDRepo
	configRepo   config.ConfigRepo
}

// NewActivityRepo new repository
func NewActivityRepo(
	data *data.Data,
	uniqueIDRepo unique.UniqueIDRepo,
	configRepo config.ConfigRepo,
) activity_common.ActivityRepo {
	return &ActivityRepo{
		data:         data,
		uniqueIDRepo: uniqueIDRepo,
		configRepo:   configRepo,
	}
}

func (ar *ActivityRepo) GetActivityTypeByObjID(ctx context.Context, objectID string, action string) (activityType, rank, hasRank int, err error) {
	objectKey, err := obj.GetObjectTypeStrByObjectID(objectID)
	if err != nil {
		return
	}

	confKey := fmt.Sprintf("%s.%s", objectKey, action)
	activityType, _ = ar.configRepo.GetConfigType(confKey)

	rank, err = ar.configRepo.GetInt(confKey)
	hasRank = 0
	if rank != 0 {
		hasRank = 1
	}
	return
}

func (ar *ActivityRepo) GetActivityTypeByObjKey(ctx context.Context, objectKey, action string) (activityType int, err error) {
	confKey := fmt.Sprintf("%s.%s", objectKey, action)
	activityType, err = ar.configRepo.GetConfigType(confKey)
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}

func (ar *ActivityRepo) GetActivityTypeByConfigKey(ctx context.Context, configKey string) (activityType int, err error) {
	activityType, err = ar.configRepo.GetConfigType(configKey)
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}

func (ar *ActivityRepo) GetActivity(ctx context.Context, session *xorm.Session,
	objectID, userID string, activityType int,
) (existsActivity *entity.Activity, exist bool, err error) {
	existsActivity = &entity.Activity{}
	exist, err = session.
		Where(builder.Eq{"object_id": objectID}).
		And(builder.Eq{"user_id": userID}).
		And(builder.Eq{"activity_type": activityType}).
		Get(existsActivity)
	return
}

func (ar *ActivityRepo) GetUserIDObjectIDActivitySum(ctx context.Context, userID, objectID string) (int, error) {
	sum := &entity.ActivityRankSum{}
	_, err := ar.data.DB.Table(entity.Activity{}.TableName()).
		Select("sum(rank) as rank").
		Where("user_id =?", userID).
		And("object_id = ?", objectID).
		And("cancelled =0").
		Get(sum)
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
		return 0, err
	}
	return sum.Rank, nil
}

// AddActivity add activity
func (ar *ActivityRepo) AddActivity(ctx context.Context, activity *entity.Activity) (err error) {
	_, err = ar.data.DB.Insert(activity)
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}
