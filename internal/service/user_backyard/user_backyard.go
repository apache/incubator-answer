package user_backyard

import (
	"context"
	"fmt"
	"time"

	"github.com/jinzhu/copier"
	"github.com/segmentfault/answer/internal/base/pager"
	"github.com/segmentfault/answer/internal/base/reason"
	"github.com/segmentfault/answer/internal/entity"
	"github.com/segmentfault/answer/internal/schema"
	"github.com/segmentfault/pacman/errors"
)

// UserBackyardRepo user repository
type UserBackyardRepo interface {
	UpdateUserStatus(ctx context.Context, userID string, userStatus, mailStatus int, email string) (err error)
	GetUserInfo(ctx context.Context, userID string) (user *entity.User, exist bool, err error)
	GetUserPage(ctx context.Context, page, pageSize int, user *entity.User) (users []*entity.User, total int64, err error)
}

// UserBackyardService user service
type UserBackyardService struct {
	userRepo UserBackyardRepo
}

func NewUserBackyardService(userRepo UserBackyardRepo) *UserBackyardService {
	return &UserBackyardService{
		userRepo: userRepo,
	}
}

// UpdateUserStatus update user
func (us *UserBackyardService) UpdateUserStatus(ctx context.Context, req *schema.UpdateUserStatusReq) (err error) {
	userInfo, exist, err := us.userRepo.GetUserInfo(ctx, req.UserID)
	if err != nil {
		return
	}
	if !exist {
		return errors.BadRequest(reason.UserNotFound)
	}
	// if user status is deleted
	if userInfo.Status == entity.UserStatusDeleted {
		return nil
	}

	if req.IsInactive() {
		userInfo.MailStatus = entity.EmailStatusToBeVerified
	}
	if req.IsDeleted() {
		userInfo.Status = entity.UserStatusDeleted
		userInfo.EMail = fmt.Sprintf("%s.%d", userInfo.EMail, time.Now().UnixNano())
	}
	if req.IsSuspended() {
		userInfo.Status = entity.UserStatusSuspended
	}
	if req.IsNormal() {
		userInfo.Status = entity.UserStatusAvailable
		userInfo.MailStatus = entity.EmailStatusAvailable
	}
	return us.userRepo.UpdateUserStatus(ctx, userInfo.ID, userInfo.Status, userInfo.MailStatus, userInfo.EMail)
}

// GetUserInfo get user one
func (us *UserBackyardService) GetUserInfo(ctx context.Context, userID string) (resp *schema.GetUserInfoResp, err error) {
	user, exist, err := us.userRepo.GetUserInfo(ctx, userID)
	if err != nil {
		return
	}
	if !exist {
		return nil, errors.BadRequest(reason.UserNotFound)
	}

	resp = &schema.GetUserInfoResp{}
	_ = copier.Copy(resp, user)
	return resp, nil
}

// GetUserPage get user list page
func (us *UserBackyardService) GetUserPage(ctx context.Context, req *schema.GetUserPageReq) (pageModel *pager.PageModel, err error) {
	user := &entity.User{}
	_ = copier.Copy(user, req)

	if req.IsInactive() {
		user.MailStatus = entity.EmailStatusToBeVerified
		user.Status = entity.UserStatusAvailable
	} else if req.IsSuspended() {
		user.Status = entity.UserStatusSuspended
	} else if req.IsDeleted() {
		user.Status = entity.UserStatusDeleted
	}

	users, total, err := us.userRepo.GetUserPage(ctx, req.Page, req.PageSize, user)
	if err != nil {
		return
	}

	resp := make([]*schema.GetUserPageResp, 0)
	for _, u := range users {
		t := &schema.GetUserPageResp{
			UserID:      u.ID,
			CreatedAt:   u.CreatedAt.Unix(),
			Username:    u.Username,
			EMail:       u.EMail,
			Rank:        u.Rank,
			DisplayName: u.DisplayName,
			Avatar:      u.Avatar,
		}
		if u.Status == entity.UserStatusDeleted {
			t.Status = schema.UserDeleted
			t.DeletedAt = u.DeletedAt.Unix()
		} else if u.Status == entity.UserStatusSuspended {
			t.Status = schema.UserSuspended
			t.SuspendedAt = u.SuspendedAt.Unix()
		} else if u.MailStatus == entity.EmailStatusToBeVerified {
			t.Status = schema.UserInactive
		} else {
			t.Status = schema.UserNormal
		}
		resp = append(resp, t)
	}
	return pager.NewPageModel(total, resp), nil
}
