package activity

import (
	"context"

	"github.com/answerdev/answer/internal/base/data"
	"github.com/answerdev/answer/internal/base/reason"
	"github.com/answerdev/answer/internal/entity"
	"github.com/answerdev/answer/internal/service/activity"
	"github.com/answerdev/answer/internal/service/activity_common"
	"github.com/answerdev/answer/internal/service/config"
	"github.com/answerdev/answer/internal/service/rank"
	"github.com/segmentfault/pacman/errors"
	"xorm.io/xorm"
)

// UserActiveActivityRepo answer accepted
type UserActiveActivityRepo struct {
	data         *data.Data
	activityRepo activity_common.ActivityRepo
	userRankRepo rank.UserRankRepo
	configRepo   config.ConfigRepo
}

const (
	UserActivated = "user.activated"
)

// NewUserActiveActivityRepo new repository
func NewUserActiveActivityRepo(
	data *data.Data,
	activityRepo activity_common.ActivityRepo,
	userRankRepo rank.UserRankRepo,
	configRepo config.ConfigRepo,
) activity.UserActiveActivityRepo {
	return &UserActiveActivityRepo{
		data:         data,
		activityRepo: activityRepo,
		userRankRepo: userRankRepo,
		configRepo:   configRepo,
	}
}

// UserActive accept other answer
func (ar *UserActiveActivityRepo) UserActive(ctx context.Context, userID string) (err error) {
	_, err = ar.data.DB.Transaction(func(session *xorm.Session) (result any, err error) {
		session = session.Context(ctx)
		activityType, err := ar.configRepo.GetConfigType(UserActivated)
		if err != nil {
			return nil, err
		}
		deltaRank, err := ar.configRepo.GetInt(UserActivated)
		if err != nil {
			return nil, err
		}

		addActivity := &entity.Activity{
			UserID:           userID,
			ObjectID:         "0",
			OriginalObjectID: "0",
			ActivityType:     activityType,
			Rank:             deltaRank,
			HasRank:          1,
		}
		_, exists, err := ar.activityRepo.GetActivity(ctx, session, "0", addActivity.UserID, activityType)
		if err != nil {
			return nil, errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
		}
		if exists {
			return nil, nil
		}

		_, err = ar.userRankRepo.TriggerUserRank(ctx, session, addActivity.UserID, addActivity.Rank, activityType)
		if err != nil {
			return nil, errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
		}
		_, err = session.Insert(addActivity)
		if err != nil {
			return nil, errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
		}
		return nil, nil
	})
	return err
}
