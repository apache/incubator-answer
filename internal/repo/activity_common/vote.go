package activity_common

import (
	"context"

	"github.com/answerdev/answer/internal/base/data"
	"github.com/answerdev/answer/internal/entity"
	"github.com/answerdev/answer/internal/service/activity_common"
)

// VoteRepo activity repository
type VoteRepo struct {
	data         *data.Data
	activityRepo activity_common.ActivityRepo
}

// NewVoteRepo new repository
func NewVoteRepo(data *data.Data, activityRepo activity_common.ActivityRepo) activity_common.VoteRepo {
	return &VoteRepo{
		data:         data,
		activityRepo: activityRepo,
	}
}

func (vr *VoteRepo) GetVoteStatus(ctx context.Context, objectId, userId string) (status string) {
	for _, action := range []string{"vote_up", "vote_down"} {
		at := &entity.Activity{}
		activityType, _, _, err := vr.activityRepo.GetActivityTypeByObjID(ctx, objectId, action)
		if err != nil {
			return ""
		}
		has, err := vr.data.DB.Where("object_id =? AND cancelled=0 AND activity_type=? AND user_id=?", objectId, activityType, userId).Get(at)
		if err != nil {
			return ""
		}
		if has {
			return action
		}
	}
	return ""
}
