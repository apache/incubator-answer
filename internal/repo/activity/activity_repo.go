package activity

import (
	"context"

	"github.com/answerdev/answer/internal/base/data"
	"github.com/answerdev/answer/internal/base/reason"
	"github.com/answerdev/answer/internal/entity"
	"github.com/answerdev/answer/internal/service/activity"
	"github.com/segmentfault/pacman/errors"
)

// activityRepo activity repository
type activityRepo struct {
	data *data.Data
}

// NewActivityRepo new repository
func NewActivityRepo(
	data *data.Data,
) activity.ActivityRepo {
	return &activityRepo{
		data: data,
	}
}

func (ar *activityRepo) GetObjectAllActivity(ctx context.Context, objectID string, showVote bool) (
	activityList []*entity.Activity, err error) {
	activityList = make([]*entity.Activity, 0)
	session := ar.data.DB.Desc("created_at") // TODO: if showVote is false do not show vote activity
	err = session.Find(&activityList, &entity.Activity{OriginalObjectID: objectID})
	if err != nil {
		return nil, errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return activityList, nil
}
