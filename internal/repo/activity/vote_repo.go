package activity

import (
	"context"
	"github.com/segmentfault/answer/pkg/converter"
	"strings"

	"github.com/segmentfault/answer/internal/base/pager"
	"github.com/segmentfault/answer/internal/service/config"
	"github.com/segmentfault/answer/internal/service/notice_queue"
	"github.com/segmentfault/answer/internal/service/rank"
	"github.com/segmentfault/answer/pkg/obj"

	"xorm.io/builder"

	"github.com/segmentfault/answer/internal/service/activity_common"
	"github.com/segmentfault/answer/internal/service/unique"

	"github.com/segmentfault/answer/internal/base/data"
	"github.com/segmentfault/answer/internal/base/reason"
	"github.com/segmentfault/answer/internal/entity"
	"github.com/segmentfault/answer/internal/schema"
	"github.com/segmentfault/answer/internal/service"
	"github.com/segmentfault/pacman/errors"
	"xorm.io/xorm"
)

// VoteRepo activity repository
type VoteRepo struct {
	data         *data.Data
	uniqueIDRepo unique.UniqueIDRepo
	configRepo   config.ConfigRepo
	activityRepo activity_common.ActivityRepo
	userRankRepo rank.UserRankRepo
	voteCommon   activity_common.VoteRepo
}

// NewVoteRepo new repository
func NewVoteRepo(
	data *data.Data,
	uniqueIDRepo unique.UniqueIDRepo,
	configRepo config.ConfigRepo,
	activityRepo activity_common.ActivityRepo,
	userRankRepo rank.UserRankRepo,
	voteCommon activity_common.VoteRepo) service.VoteRepo {
	return &VoteRepo{
		data:         data,
		uniqueIDRepo: uniqueIDRepo,
		configRepo:   configRepo,
		activityRepo: activityRepo,
		userRankRepo: userRankRepo,
		voteCommon:   voteCommon,
	}
}

var LimitUpActions = map[string][]string{
	"question": {"vote_up", "voted_up"},
	"answer":   {"vote_up", "voted_up"},
	"comment":  {"vote_up"},
}

var LimitDownActions = map[string][]string{
	"question": {"vote_down", "voted_down"},
	"answer":   {"vote_down", "voted_down"},
	"comment":  {"vote_down"},
}

func (vr *VoteRepo) vote(ctx context.Context, objectID string, userID, objectUserId string, actions []string) (resp *schema.VoteResp, err error) {
	resp = &schema.VoteResp{}
	_, err = vr.data.DB.Transaction(func(session *xorm.Session) (result any, err error) {
		result = nil
		for _, action := range actions {
			var (
				existsActivity entity.Activity
				insertActivity entity.Activity
				has            bool
				triggerUserID,
				activityUserId string
				activityType, deltaRank, hasRank int
			)

			activityUserId, activityType, deltaRank, hasRank, err = vr.CheckRank(ctx, objectID, objectUserId, userID, action)
			if err != nil {
				return
			}

			triggerUserID = userID
			if userID == activityUserId {
				triggerUserID = "0"
			}

			// check is voted up
			has, _ = session.
				Where(builder.Eq{"object_id": objectID}).
				And(builder.Eq{"user_id": activityUserId}).
				And(builder.Eq{"trigger_user_id": triggerUserID}).
				And(builder.Eq{"activity_type": activityType}).
				Get(&existsActivity)

			// is is voted,return
			if has && existsActivity.Cancelled == 0 {
				return
			}

			insertActivity = entity.Activity{
				ObjectID:      objectID,
				UserID:        activityUserId,
				TriggerUserID: converter.StringToInt64(triggerUserID),
				ActivityType:  activityType,
				Rank:          deltaRank,
				HasRank:       hasRank,
				Cancelled:     0,
			}

			// trigger user rank and send notification
			if hasRank != 0 {
				isReachStandard, err := vr.userRankRepo.TriggerUserRank(ctx, session, activityUserId, deltaRank, activityType)
				if err != nil {
					return nil, err
				}
				if isReachStandard {
					insertActivity.Rank = 0
				}

				vr.sendNotification(ctx, activityUserId, objectUserId, objectID)
			}

			if has {
				if _, err = session.Where("id = ?", existsActivity.ID).Cols("`cancelled`").
					Update(&entity.Activity{
						Cancelled: 0,
					}); err != nil {
					return
				}
			} else {
				_, err = session.Insert(&insertActivity)
				if err != nil {
					return nil, err
				}
			}

			// update votes
			if action == "vote_down" || action == "vote_up" {
				votes := 1
				if action == "vote_down" {
					votes = -1
				}
				err = vr.updateVotes(ctx, session, objectID, votes)
				if err != nil {
					return
				}
			}
		}
		return
	})
	if err != nil {
		return
	}

	resp, err = vr.GetVoteResultByObjectId(ctx, objectID)
	resp.VoteStatus = vr.voteCommon.GetVoteStatus(ctx, objectID, userID)

	return
}

func (vr *VoteRepo) voteCancel(ctx context.Context, objectID string, userID, objectUserId string, actions []string) (resp *schema.VoteResp, err error) {
	resp = &schema.VoteResp{}
	_, err = vr.data.DB.Transaction(func(session *xorm.Session) (result any, err error) {
		for _, action := range actions {
			var (
				existsActivity entity.Activity
				has            bool
				triggerUserID,
				activityUserId string
				activityType,
				deltaRank, hasRank int
			)
			result = nil

			activityUserId, activityType, deltaRank, hasRank, err = vr.CheckRank(ctx, objectID, objectUserId, userID, action)
			if err != nil {
				return
			}

			triggerUserID = userID
			if userID == activityUserId {
				triggerUserID = "0"
			}

			has, err = session.
				Where(builder.Eq{"user_id": activityUserId}).
				And(builder.Eq{"trigger_user_id": triggerUserID}).
				And(builder.Eq{"activity_type": activityType}).
				And(builder.Eq{"object_id": objectID}).
				Get(&existsActivity)

			if !has {
				return
			}

			if existsActivity.Cancelled == 1 {
				return
			}

			if _, err = session.Where("id = ?", existsActivity.ID).Cols("`cancelled`").
				Update(&entity.Activity{
					Cancelled: 1,
				}); err != nil {
				return
			}

			// trigger user rank and send notification
			if hasRank != 0 {
				_, err = vr.userRankRepo.TriggerUserRank(ctx, session, activityUserId, -deltaRank, activityType)
				if err != nil {
					return
				}

				vr.sendNotification(ctx, activityUserId, objectUserId, objectID)
			}

			// update votes
			if action == "vote_down" || action == "vote_up" {
				votes := -1
				if action == "vote_down" {
					votes = 1
				}
				err = vr.updateVotes(ctx, session, objectID, votes)
				if err != nil {
					return
				}
			}
		}

		return
	})
	if err != nil {
		return
	}
	resp, err = vr.GetVoteResultByObjectId(ctx, objectID)
	resp.VoteStatus = vr.voteCommon.GetVoteStatus(ctx, objectID, userID)
	return
}

func (vr *VoteRepo) VoteUp(ctx context.Context, objectID string, userID, objectUserId string) (resp *schema.VoteResp, err error) {
	resp = &schema.VoteResp{}
	objectType, err := obj.GetObjectTypeStrByObjectID(objectID)
	if err != nil {
		err = errors.BadRequest(reason.ObjectNotFound)
		return
	}

	actions, ok := LimitUpActions[objectType]
	if !ok {
		err = errors.BadRequest(reason.DisallowVote)
		return
	}

	_, _ = vr.VoteDownCancel(ctx, objectID, userID, objectUserId)
	return vr.vote(ctx, objectID, userID, objectUserId, actions)
}

func (vr *VoteRepo) VoteDown(ctx context.Context, objectID string, userID, objectUserId string) (resp *schema.VoteResp, err error) {
	resp = &schema.VoteResp{}
	objectType, err := obj.GetObjectTypeStrByObjectID(objectID)
	if err != nil {
		err = errors.BadRequest(reason.ObjectNotFound)
		return
	}
	actions, ok := LimitDownActions[objectType]
	if !ok {
		err = errors.BadRequest(reason.DisallowVote)
		return
	}

	_, _ = vr.VoteUpCancel(ctx, objectID, userID, objectUserId)
	return vr.vote(ctx, objectID, userID, objectUserId, actions)
}

func (vr *VoteRepo) VoteUpCancel(ctx context.Context, objectID string, userID, objectUserId string) (resp *schema.VoteResp, err error) {
	var (
		objectType string
	)
	resp = &schema.VoteResp{}

	objectType, err = obj.GetObjectTypeStrByObjectID(objectID)
	if err != nil {
		err = errors.BadRequest(reason.ObjectNotFound)
		return
	}
	actions, ok := LimitUpActions[objectType]
	if !ok {
		err = errors.BadRequest(reason.DisallowVote)
		return
	}

	return vr.voteCancel(ctx, objectID, userID, objectUserId, actions)
}

func (vr *VoteRepo) VoteDownCancel(ctx context.Context, objectID string, userID, objectUserId string) (resp *schema.VoteResp, err error) {
	var (
		objectType string
	)
	resp = &schema.VoteResp{}

	objectType, err = obj.GetObjectTypeStrByObjectID(objectID)
	if err != nil {
		err = errors.BadRequest(reason.ObjectNotFound)
		return
	}
	actions, ok := LimitDownActions[objectType]
	if !ok {
		err = errors.BadRequest(reason.DisallowVote)
		return
	}

	return vr.voteCancel(ctx, objectID, userID, objectUserId, actions)
}

func (vr *VoteRepo) CheckRank(ctx context.Context, objectID, objectUserId, userID string, action string) (activityUserId string, activityType, rank, hasRank int, err error) {
	activityType, rank, hasRank, err = vr.activityRepo.GetActivityTypeByObjID(ctx, objectID, action)

	if err != nil {
		return
	}

	activityUserId = userID
	if strings.Contains(action, "voted") {
		activityUserId = objectUserId
	}

	return activityUserId, activityType, rank, hasRank, nil
}

func (vr *VoteRepo) GetVoteResultByObjectId(ctx context.Context, objectID string) (resp *schema.VoteResp, err error) {
	resp = &schema.VoteResp{}
	for _, action := range []string{"vote_up", "vote_down"} {
		var (
			activity     entity.Activity
			votes        int64
			activityType int
		)

		activityType, _, _, err = vr.activityRepo.GetActivityTypeByObjID(ctx, objectID, action)

		votes, err = vr.data.DB.Where(builder.Eq{"object_id": objectID}).
			And(builder.Eq{"activity_type": activityType}).
			And(builder.Eq{"cancelled": 0}).
			Count(&activity)

		if err != nil {
			return
		}

		if action == "vote_up" {
			resp.UpVotes = int(votes)
		} else {
			resp.DownVotes = int(votes)
		}
	}

	resp.Votes = resp.UpVotes - resp.DownVotes

	return resp, nil
}

func (vr *VoteRepo) ListUserVotes(
	ctx context.Context,
	userID string,
	req schema.GetVoteWithPageReq,
	activityTypes []int,
) (voteList []entity.Activity, total int64, err error) {
	session := vr.data.DB.NewSession()
	cond := builder.
		And(
			builder.Eq{"user_id": userID},
			builder.Eq{"cancelled": 0},
			builder.In("activity_type", activityTypes),
		)

	session.Where(cond).OrderBy("updated_at desc")

	total, err = pager.Help(req.Page, req.PageSize, &voteList, &entity.Activity{}, session)
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}

// updateVotes
// if votes < 0 Decr object vote_count,otherwise Incr object vote_count
func (vr *VoteRepo) updateVotes(ctx context.Context, session *xorm.Session, objectID string, votes int) (err error) {
	var (
		objectType string
		e          error
	)

	objectType, err = obj.GetObjectTypeStrByObjectID(objectID)
	switch objectType {
	case "question":
		_, err = session.Where("id = ?", objectID).Incr("vote_count", votes).Update(&entity.Question{})
	case "answer":
		_, err = session.Where("id = ?", objectID).Incr("vote_count", votes).Update(&entity.Answer{})
	case "comment":
		_, err = session.Where("id = ?", objectID).Incr("vote_count", votes).Update(&entity.Comment{})
	default:
		e = errors.BadRequest(reason.DisallowVote)
	}

	if e != nil {
		err = e
	} else if err != nil {
		err = errors.BadRequest(reason.DatabaseError).WithError(err).WithStack()
	}

	return
}

// sendNotification send rank triggered notification
func (vr *VoteRepo) sendNotification(ctx context.Context, activityUserId, objectUserId, objectID string) {
	objectType, err := obj.GetObjectTypeStrByObjectID(objectID)
	if err != nil {
		return
	}

	msg := &schema.NotificationMsg{
		ReceiverUserID: activityUserId,
		TriggerUserID:  objectUserId,
		Type:           schema.NotificationTypeAchievement,
		ObjectID:       objectID,
		ObjectType:     objectType,
	}
	notice_queue.AddNotification(msg)
}
