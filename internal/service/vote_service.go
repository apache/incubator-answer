package service

import (
	"context"

	"github.com/answerdev/answer/internal/base/pager"
	"github.com/answerdev/answer/internal/entity"
	"github.com/answerdev/answer/internal/service/activity_type"
	"github.com/answerdev/answer/internal/service/comment_common"
	"github.com/answerdev/answer/internal/service/config"
	"github.com/answerdev/answer/internal/service/object_info"
	"github.com/answerdev/answer/pkg/obj"
	"github.com/segmentfault/pacman/log"

	"github.com/answerdev/answer/internal/base/reason"
	"github.com/answerdev/answer/internal/schema"
	answercommon "github.com/answerdev/answer/internal/service/answer_common"
	questioncommon "github.com/answerdev/answer/internal/service/question_common"
	"github.com/answerdev/answer/internal/service/unique"
	"github.com/segmentfault/pacman/errors"
)

// VoteRepo activity repository
type VoteRepo interface {
	VoteUp(ctx context.Context, objectID string, userID, objectUserID string) (resp *schema.VoteResp, err error)
	VoteDown(ctx context.Context, objectID string, userID, objectUserID string) (resp *schema.VoteResp, err error)
	VoteUpCancel(ctx context.Context, objectID string, userID, objectUserID string) (resp *schema.VoteResp, err error)
	VoteDownCancel(ctx context.Context, objectID string, userID, objectUserID string) (resp *schema.VoteResp, err error)
	GetVoteResultByObjectId(ctx context.Context, objectID string) (resp *schema.VoteResp, err error)
	ListUserVotes(ctx context.Context, userID string, req schema.GetVoteWithPageReq, activityTypes []int) (voteList []entity.Activity, total int64, err error)
}

// VoteService user service
type VoteService struct {
	voteRepo          VoteRepo
	UniqueIDRepo      unique.UniqueIDRepo
	configRepo        config.ConfigRepo
	questionRepo      questioncommon.QuestionRepo
	answerRepo        answercommon.AnswerRepo
	commentCommonRepo comment_common.CommentCommonRepo
	objectService     *object_info.ObjService
}

func NewVoteService(
	VoteRepo VoteRepo,
	uniqueIDRepo unique.UniqueIDRepo,
	configRepo config.ConfigRepo,
	questionRepo questioncommon.QuestionRepo,
	answerRepo answercommon.AnswerRepo,
	commentCommonRepo comment_common.CommentCommonRepo,
	objectService *object_info.ObjService,
) *VoteService {
	return &VoteService{
		voteRepo:          VoteRepo,
		UniqueIDRepo:      uniqueIDRepo,
		configRepo:        configRepo,
		questionRepo:      questionRepo,
		answerRepo:        answerRepo,
		commentCommonRepo: commentCommonRepo,
		objectService:     objectService,
	}
}

// VoteUp vote up
func (as *VoteService) VoteUp(ctx context.Context, dto *schema.VoteDTO) (voteResp *schema.VoteResp, err error) {
	voteResp = &schema.VoteResp{}

	var objectUserID string

	objectUserID, err = as.GetObjectUserId(ctx, dto.ObjectID)
	if err != nil {
		return
	}

	// check user is voting self or not
	if objectUserID == dto.UserID {
		err = errors.BadRequest(reason.DisallowVoteYourSelf)
		return
	}

	if dto.IsCancel {
		return as.voteRepo.VoteUpCancel(ctx, dto.ObjectID, dto.UserID, objectUserID)
	} else {
		return as.voteRepo.VoteUp(ctx, dto.ObjectID, dto.UserID, objectUserID)
	}
}

// VoteDown vote down
func (as *VoteService) VoteDown(ctx context.Context, dto *schema.VoteDTO) (voteResp *schema.VoteResp, err error) {
	voteResp = &schema.VoteResp{}

	var objectUserID string

	objectUserID, err = as.GetObjectUserId(ctx, dto.ObjectID)
	if err != nil {
		return
	}

	// check user is voting self or not
	if objectUserID == dto.UserID {
		err = errors.BadRequest(reason.DisallowVoteYourSelf)
		return
	}

	if dto.IsCancel {
		return as.voteRepo.VoteDownCancel(ctx, dto.ObjectID, dto.UserID, objectUserID)
	} else {
		return as.voteRepo.VoteDown(ctx, dto.ObjectID, dto.UserID, objectUserID)
	}
}

func (vs *VoteService) GetObjectUserId(ctx context.Context, objectID string) (userID string, err error) {
	var objectKey string
	objectKey, err = obj.GetObjectTypeStrByObjectID(objectID)

	if err != nil {
		err = nil
		return
	}

	switch objectKey {
	case "question":
		object, has, e := vs.questionRepo.GetQuestion(ctx, objectID)
		if e != nil || !has {
			err = errors.BadRequest(reason.QuestionNotFound).WithError(e).WithStack()
			return
		}
		userID = object.UserID
	case "answer":
		object, has, e := vs.answerRepo.GetAnswer(ctx, objectID)
		if e != nil || !has {
			err = errors.BadRequest(reason.AnswerNotFound).WithError(e).WithStack()
			return
		}
		userID = object.UserID
	case "comment":
		object, has, e := vs.commentCommonRepo.GetComment(ctx, objectID)
		if e != nil || !has {
			err = errors.BadRequest(reason.CommentNotFound).WithError(e).WithStack()
			return
		}
		userID = object.UserID
	default:
		err = errors.BadRequest(reason.DisallowVote).WithError(err).WithStack()
		return
	}

	return
}

// ListUserVotes list user's votes
func (vs *VoteService) ListUserVotes(ctx context.Context, req schema.GetVoteWithPageReq) (model *pager.PageModel, err error) {
	var (
		resp     []schema.GetVoteWithPageResp
		typeKeys = []string{
			"question.vote_up",
			"question.vote_down",
			"answer.vote_up",
			"answer.vote_down",
		}
		activityTypes []int
	)

	for _, typeKey := range typeKeys {
		t, err := vs.configRepo.GetConfigType(typeKey)
		if err != nil {
			continue
		}
		activityTypes = append(activityTypes, t)
	}

	voteList, total, err := vs.voteRepo.ListUserVotes(ctx, req.UserID, req, activityTypes)
	if err != nil {
		return
	}

	for _, voteInfo := range voteList {
		objInfo, err := vs.objectService.GetInfo(ctx, voteInfo.ObjectID)
		if err != nil {
			log.Error(err)
		}

		item := schema.GetVoteWithPageResp{
			CreatedAt:  voteInfo.CreatedAt.Unix(),
			ObjectID:   objInfo.ObjectID,
			QuestionID: objInfo.QuestionID,
			AnswerID:   objInfo.AnswerID,
			ObjectType: objInfo.ObjectType,
			Title:      objInfo.Title,
			Content:    objInfo.Content,
			VoteType:   activity_type.Format(voteInfo.ActivityType),
		}
		resp = append(resp, item)
	}

	return pager.NewPageModel(total, resp), err
}
