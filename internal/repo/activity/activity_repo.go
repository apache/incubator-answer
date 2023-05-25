package activity

import (
	"context"
	"fmt"

	"github.com/answerdev/answer/internal/base/constant"
	"github.com/answerdev/answer/internal/base/data"
	"github.com/answerdev/answer/internal/base/reason"
	"github.com/answerdev/answer/internal/entity"
	"github.com/answerdev/answer/internal/service/activity"
	"github.com/answerdev/answer/internal/service/config"
	"github.com/segmentfault/pacman/errors"
	"github.com/segmentfault/pacman/log"
)

// activityRepo activity repository
type activityRepo struct {
	data          *data.Data
	configService *config.ConfigService
}

// NewActivityRepo new repository
func NewActivityRepo(
	data *data.Data,
	configService *config.ConfigService,
) activity.ActivityRepo {
	return &activityRepo{
		data:          data,
		configService: configService,
	}
}

func (ar *activityRepo) GetObjectAllActivity(ctx context.Context, objectID string, showVote bool) (
	activityList []*entity.Activity, err error) {
	activityList = make([]*entity.Activity, 0)
	session := ar.data.DB.Context(ctx).Desc("created_at")

	if !showVote {
		activityTypeNotShown := ar.getAllActivityType(ctx)
		session.NotIn("activity_type", activityTypeNotShown)
	}
	err = session.Find(&activityList, &entity.Activity{OriginalObjectID: objectID})
	if err != nil {
		return nil, errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return activityList, nil
}

func (ar *activityRepo) getAllActivityType(ctx context.Context) (activityTypes []int) {
	var activityTypeNotShown []int
	for _, obj := range []string{constant.AnswerObjectType, constant.QuestionObjectType, constant.CommentObjectType} {
		for _, act := range []string{
			constant.ActVotedDown,
			constant.ActVotedUp,
			constant.ActVoteDown,
			constant.ActVoteUp,
		} {
			id, err := ar.configService.GetIDByKey(ctx, fmt.Sprintf("%s.%s", obj, act))
			if err != nil {
				log.Error(err)
			} else {
				activityTypeNotShown = append(activityTypeNotShown, id)
			}
		}
	}
	return activityTypeNotShown
}
