package activity

import (
	"context"

	"github.com/segmentfault/answer/internal/service/activity_common"
	"github.com/segmentfault/answer/internal/service/follow"
	"github.com/segmentfault/answer/pkg/obj"
	"github.com/segmentfault/pacman/log"
	"xorm.io/builder"

	"github.com/segmentfault/answer/internal/base/data"
	"github.com/segmentfault/answer/internal/base/reason"
	"github.com/segmentfault/answer/internal/entity"
	"github.com/segmentfault/answer/internal/service/unique"
	"github.com/segmentfault/pacman/errors"
	"xorm.io/xorm"
)

// FollowRepo activity repository
type FollowRepo struct {
	data         *data.Data
	uniqueIDRepo unique.UniqueIDRepo
	activityRepo activity_common.ActivityRepo
}

// NewFollowRepo new repository
func NewFollowRepo(
	data *data.Data,
	uniqueIDRepo unique.UniqueIDRepo,
	activityRepo activity_common.ActivityRepo,
) follow.FollowRepo {
	return &FollowRepo{
		data:         data,
		uniqueIDRepo: uniqueIDRepo,
		activityRepo: activityRepo,
	}
}

func (ar *FollowRepo) Follow(ctx context.Context, objectId, userId string) error {
	activityType, _, _, err := ar.activityRepo.GetActivityTypeByObjID(ctx, objectId, "follow")
	if err != nil {
		return err
	}

	_, err = ar.data.DB.Transaction(func(session *xorm.Session) (result any, err error) {
		var (
			existsActivity entity.Activity
			has            bool
		)
		result = nil

		has, err = session.Where(builder.Eq{"activity_type": activityType}).
			And(builder.Eq{"user_id": userId}).
			And(builder.Eq{"object_id": objectId}).
			Get(&existsActivity)

		if err != nil {
			return
		}

		if has && existsActivity.Cancelled == 0 {
			return
		}

		if has {
			_, err = session.Where(builder.Eq{"id": existsActivity.ID}).
				Cols(`cancelled`).
				Update(&entity.Activity{
					Cancelled: 0,
				})
		} else {
			// update existing activity with new user id and u object id
			_, err = session.Insert(&entity.Activity{
				UserID:       userId,
				ObjectID:     objectId,
				ActivityType: activityType,
				Cancelled:    0,
				Rank:         0,
				HasRank:      0,
			})
		}

		if err != nil {
			log.Error(err)
			return
		}

		// start update followers when everything is fine
		err = ar.updateFollows(ctx, session, objectId, 1)
		if err != nil {
			log.Error(err)
		}

		return
	})

	return err
}

func (ar *FollowRepo) FollowCancel(ctx context.Context, objectId, userId string) error {
	activityType, _, _, err := ar.activityRepo.GetActivityTypeByObjID(nil, objectId, "follow")
	if err != nil {
		return err
	}

	_, err = ar.data.DB.Transaction(func(session *xorm.Session) (result any, err error) {
		var (
			existsActivity entity.Activity
			has            bool
		)
		result = nil

		has, err = session.Where(builder.Eq{"activity_type": activityType}).
			And(builder.Eq{"user_id": userId}).
			And(builder.Eq{"object_id": objectId}).
			Get(&existsActivity)

		if err != nil || !has {
			return
		}

		if has && existsActivity.Cancelled == 1 {
			return
		}
		if _, err = session.Where("id = ?", existsActivity.ID).
			Cols("cancelled").
			Update(&entity.Activity{
				Cancelled: 1,
			}); err != nil {
			return
		}
		err = ar.updateFollows(ctx, session, objectId, -1)
		return
	})
	return err
}

func (ar *FollowRepo) updateFollows(ctx context.Context, session *xorm.Session, objectId string, follows int) error {
	objectType, err := obj.GetObjectTypeStrByObjectID(objectId)
	switch objectType {
	case "question":
		_, err = session.Where("id = ?", objectId).Incr("follow_count", follows).Update(&entity.Question{})
	case "user":
		_, err = session.Where("id = ?", objectId).Incr("follow_count", follows).Update(&entity.User{})
	case "tag":
		_, err = session.Where("id = ?", objectId).Incr("follow_count", follows).Update(&entity.Tag{})
	default:
		err = errors.InternalServer(reason.DisallowFollow).WithMsg("this object can't be followed")
	}
	return err
}
