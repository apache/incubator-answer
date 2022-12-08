package rank

import (
	"context"

	"github.com/answerdev/answer/internal/base/constant"
	"github.com/answerdev/answer/internal/base/pager"
	"github.com/answerdev/answer/internal/base/reason"
	"github.com/answerdev/answer/internal/entity"
	"github.com/answerdev/answer/internal/schema"
	"github.com/answerdev/answer/internal/service/activity_type"
	"github.com/answerdev/answer/internal/service/config"
	"github.com/answerdev/answer/internal/service/object_info"
	"github.com/answerdev/answer/internal/service/permission"
	"github.com/answerdev/answer/internal/service/role"
	usercommon "github.com/answerdev/answer/internal/service/user_common"
	"github.com/segmentfault/pacman/errors"
	"github.com/segmentfault/pacman/log"
	"xorm.io/xorm"
)

const (
	PermissionPrefix = "rank."
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
	roleService       *role.UserRoleRelService
	rolePowerService  *role.RolePowerRelService
}

// NewRankService new rank service
func NewRankService(
	userCommon *usercommon.UserCommon,
	userRankRepo UserRankRepo,
	objectInfoService *object_info.ObjService,
	roleService *role.UserRoleRelService,
	rolePowerService *role.RolePowerRelService,
	configRepo config.ConfigRepo) *RankService {
	return &RankService{
		userCommon:        userCommon,
		configRepo:        configRepo,
		userRankRepo:      userRankRepo,
		objectInfoService: objectInfoService,
		roleService:       roleService,
		rolePowerService:  rolePowerService,
	}
}

// CheckOperationPermission verify that the user has permission
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
	powerMapping := rs.getUserPowerMapping(ctx, userID)
	if powerMapping[action] {
		return true, nil
	}

	if len(objectID) > 0 {
		objectInfo, err := rs.objectInfoService.GetInfo(ctx, objectID)
		if err != nil {
			return can, err
		}
		// if the user is this object creator, the user can operate this object.
		if objectInfo != nil &&
			objectInfo.ObjectCreatorUserID == userID {
			return true, nil
		}
	}

	can = rs.checkUserRank(ctx, userInfo.ID, userInfo.Rank, PermissionPrefix+action)
	return can, nil
}

// CheckOperationPermissions verify that the user has permission
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
			return can, err
		}
		// if the user is this object creator, the user can operate this object.
		if objectInfo != nil &&
			objectInfo.ObjectCreatorUserID == userID {
			objectOwner = true
		}
	}

	powerMapping := rs.getUserPowerMapping(ctx, userID)

	for idx, action := range actions {
		if powerMapping[action] || objectOwner {
			can[idx] = true
			continue
		}
		meetRank := rs.checkUserRank(ctx, userInfo.ID, userInfo.Rank, PermissionPrefix+action)
		can[idx] = meetRank
	}
	return can, nil
}

// CheckVotePermission verify that the user has vote permission
func (rs *RankService) CheckVotePermission(ctx context.Context, userID, objectID string, voteUp bool) (
	can bool, err error) {
	if len(userID) == 0 || len(objectID) == 0 {
		return false, nil
	}

	// get the rank of the current user
	userInfo, exist, err := rs.userCommon.GetUserBasicInfoByID(ctx, userID)
	if err != nil {
		return can, err
	}
	if !exist {
		return can, nil
	}
	objectInfo, err := rs.objectInfoService.GetInfo(ctx, objectID)
	if err != nil {
		return can, err
	}
	action := ""
	switch objectInfo.ObjectType {
	case constant.QuestionObjectType:
		if voteUp {
			action = permission.QuestionVoteUp
		} else {
			action = permission.QuestionVoteDown
		}
	case constant.AnswerObjectType:
		if voteUp {
			action = permission.AnswerVoteUp
		} else {
			action = permission.AnswerVoteDown
		}
	case constant.CommentObjectType:
		if voteUp {
			action = permission.CommentVoteUp
		} else {
			action = permission.CommentVoteDown
		}
	}
	powerMapping := rs.getUserPowerMapping(ctx, userID)
	if powerMapping[action] {
		return true, nil
	}

	meetRank := rs.checkUserRank(ctx, userInfo.ID, userInfo.Rank, PermissionPrefix+action)
	return meetRank, nil
}

// getUserPowerMapping get user power mapping
func (rs *RankService) getUserPowerMapping(ctx context.Context, userID string) (powerMapping map[string]bool) {
	powerMapping = make(map[string]bool, 0)
	userRole, err := rs.roleService.GetUserRole(ctx, userID)
	if err != nil {
		log.Error(err)
		return powerMapping
	}
	powers, err := rs.rolePowerService.GetRolePowerList(ctx, userRole)
	if err != nil {
		log.Error(err)
		return powerMapping
	}

	for _, power := range powers {
		powerMapping[power] = true
	}
	return powerMapping
}

// CheckRankPermission verify that the user meets the prestige criteria
func (rs *RankService) checkUserRank(ctx context.Context, userID string, userRank int, action string) (
	can bool) {
	// get the amount of rank required for the current operation
	requireRank, err := rs.configRepo.GetInt(action)
	if err != nil {
		log.Error(err)
		return false
	}
	if userRank < requireRank || requireRank < 0 {
		log.Debugf("user %s want to do action %s, but rank %d < %d",
			userID, action, userRank, requireRank)
		return false
	}
	return true
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
