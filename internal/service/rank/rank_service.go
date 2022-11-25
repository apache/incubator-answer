package rank

import (
	"context"

	"github.com/answerdev/answer/internal/base/pager"
	"github.com/answerdev/answer/internal/base/reason"
	"github.com/answerdev/answer/internal/entity"
	"github.com/answerdev/answer/internal/schema"
	"github.com/answerdev/answer/internal/service/activity_type"
	"github.com/answerdev/answer/internal/service/config"
	"github.com/answerdev/answer/internal/service/object_info"
	usercommon "github.com/answerdev/answer/internal/service/user_common"
	"github.com/segmentfault/pacman/errors"
	"github.com/segmentfault/pacman/log"
	"xorm.io/xorm"
)

const (
	QuestionAddRank               = "rank.question.add"
	QuestionEditRank              = "rank.question.edit"
	QuestionEditWithoutReviewRank = "rank.question.edit_without_review"
	QuestionDeleteRank            = "rank.question.delete"
	QuestionVoteUpRank            = "rank.question.vote_up"
	QuestionVoteDownRank          = "rank.question.vote_down"
	AnswerAddRank                 = "rank.answer.add"
	AnswerEditRank                = "rank.answer.edit"
	AnswerEditWithoutReviewRank   = "rank.answer.edit_without_review"
	AnswerDeleteRank              = "rank.answer.delete"
	AnswerAcceptRank              = "rank.answer.accept"
	AnswerVoteUpRank              = "rank.answer.vote_up"
	AnswerVoteDownRank            = "rank.answer.vote_down"
	CommentAddRank                = "rank.comment.add"
	CommentEditRank               = "rank.comment.edit"
	CommentDeleteRank             = "rank.comment.delete"
	CommentVoteUpRank             = "rank.comment.vote_up"
	CommentVoteDownRank           = "rank.comment.vote_down"
	ReportAddRank                 = "rank.report.add"
	TagAddRank                    = "rank.tag.add"
	TagEditRank                   = "rank.tag.edit"
	TagEditWithoutReviewRank      = "rank.tag.edit_without_review"
	TagDeleteRank                 = "rank.tag.delete"
	TagSynonymRank                = "rank.tag.synonym"
	LinkUrlLimitRank              = "rank.link.url_limit"
	VoteDetailRank                = "rank.vote.detail"
	AnswerAuditRank               = "rank.answer.audit"
	QuestionAuditRank             = "rank.question.audit"
	TagAuditRank                  = "rank.tag.audit"
)

type UserRankRepo interface {
	TriggerUserRank(ctx context.Context, session *xorm.Session, userId string, rank int, activityType int) (isReachStandard bool, err error)
	UserRankPage(ctx context.Context, userId string, page, pageSize int) (rankPage []*entity.Activity, total int64, err error)
}

// RankService rank service
type RankService struct {
	userCommon        *usercommon.UserCommon
	configRepo        config.ConfigRepo
	userRankRepo      UserRankRepo
	objectInfoService *object_info.ObjService
}

// NewRankService new rank service
func NewRankService(
	userCommon *usercommon.UserCommon,
	userRankRepo UserRankRepo,
	objectInfoService *object_info.ObjService,
	configRepo config.ConfigRepo) *RankService {
	return &RankService{
		userCommon:        userCommon,
		configRepo:        configRepo,
		userRankRepo:      userRankRepo,
		objectInfoService: objectInfoService,
	}
}

// CheckOperationPermission verify that the user has operation
func (rs *RankService) CheckOperationPermission(ctx context.Context, userID string, action string, objectID string) (
	can bool, err error) {
	if len(userID) == 0 {
		return false, nil
	}

	// get the rank of the current user
	userInfo, exist, err := rs.userCommon.GetUserBasicInfoByID(ctx, userID)
	if err != nil {
		return false, err
	}
	if !exist {
		return false, nil
	}
	// administrator have all permissions
	if userInfo.IsAdmin {
		return true, nil
	}

	if len(objectID) > 0 {
		objectInfo, err := rs.objectInfoService.GetInfo(ctx, objectID)
		if err != nil {
			log.Error(err)
		}
		// if the user is this object creator, the user can operate this object.
		if objectInfo != nil &&
			objectInfo.ObjectCreatorUserID == userID {
			return true, nil
		}
	}

	return rs.checkUserRank(ctx, userInfo.ID, userInfo.Rank, action)
}

// CheckOperationPermissions verify that the user has operation
func (rs *RankService) CheckOperationPermissions(ctx context.Context, userID string, actions []string, objectID string) (
	can []bool, err error) {
	can = make([]bool, len(actions))
	if len(userID) == 0 {
		return can, nil
	}

	// get the rank of the current user
	userInfo, exist, err := rs.userCommon.GetUserBasicInfoByID(ctx, userID)
	if err != nil {
		return can, err
	}
	if !exist {
		return can, nil
	}

	objectOwner := false
	if len(objectID) > 0 {
		objectInfo, err := rs.objectInfoService.GetInfo(ctx, objectID)
		if err != nil {
			log.Error(err)
		}
		// if the user is this object creator, the user can operate this object.
		if objectInfo != nil &&
			objectInfo.ObjectCreatorUserID == userID {
			objectOwner = true
		}
	}

	for idx, action := range actions {
		if userInfo.IsAdmin || objectOwner {
			can[idx] = true
			continue
		}
		meetRank, err := rs.checkUserRank(ctx, userInfo.ID, userInfo.Rank, action)
		if err != nil {
			log.Error(err)
		}
		can[idx] = meetRank
	}
	return can, nil
}

// CheckRankPermission verify that the user meets the prestige criteria
func (rs *RankService) checkUserRank(ctx context.Context, userID string, userRank int, action string) (
	can bool, err error) {
	// get the amount of rank required for the current operation
	requireRank, err := rs.configRepo.GetInt(action)
	if err != nil {
		return false, err
	}
	currentUserRank := userRank
	if currentUserRank < requireRank {
		log.Debugf("user %s want to do action %s, but rank %d < %d",
			userID, action, currentUserRank, requireRank)
		return false, nil
	}
	return true, nil
}

// GetRankPersonalWithPage get personal comment list page
func (rs *RankService) GetRankPersonalWithPage(ctx context.Context, req *schema.GetRankPersonalWithPageReq) (
	pageModel *pager.PageModel, err error) {
	if len(req.Username) > 0 {
		userInfo, exist, err := rs.userCommon.GetUserBasicInfoByUserName(ctx, req.Username)
		if err != nil {
			return nil, err
		}
		if !exist {
			return nil, errors.BadRequest(reason.UserNotFound)
		}
		req.UserID = userInfo.ID
	}
	if len(req.UserID) == 0 {
		return nil, errors.BadRequest(reason.UserNotFound)
	}

	userRankPage, total, err := rs.userRankRepo.UserRankPage(ctx, req.UserID, req.Page, req.PageSize)
	if err != nil {
		return nil, err
	}
	resp := make([]*schema.GetRankPersonalWithPageResp, 0)
	for _, userRankInfo := range userRankPage {
		commentResp := &schema.GetRankPersonalWithPageResp{
			CreatedAt:  userRankInfo.CreatedAt.Unix(),
			ObjectID:   userRankInfo.ObjectID,
			Reputation: userRankInfo.Rank,
		}
		if len(userRankInfo.ObjectID) > 0 {
			objInfo, err := rs.objectInfoService.GetInfo(ctx, userRankInfo.ObjectID)
			if err != nil {
				log.Error(err)
			} else {
				commentResp.RankType = activity_type.Format(userRankInfo.ActivityType)
				commentResp.ObjectType = objInfo.ObjectType
				commentResp.Title = objInfo.Title
				commentResp.Content = objInfo.Content
				commentResp.QuestionID = objInfo.QuestionID
				commentResp.AnswerID = objInfo.AnswerID
			}
		}
		resp = append(resp, commentResp)
	}
	return pager.NewPageModel(total, resp), nil
}
