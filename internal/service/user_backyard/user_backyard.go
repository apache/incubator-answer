package user_backyard

import (
	"context"
	"fmt"
	"net/mail"
	"strings"
	"time"
	"unicode"

	"github.com/answerdev/answer/internal/base/pager"
	"github.com/answerdev/answer/internal/base/reason"
	"github.com/answerdev/answer/internal/entity"
	"github.com/answerdev/answer/internal/schema"
	"github.com/answerdev/answer/internal/service/auth"
	"github.com/answerdev/answer/internal/service/role"
	"github.com/jinzhu/copier"
	"github.com/segmentfault/pacman/errors"
	"github.com/segmentfault/pacman/log"
)

// UserBackyardRepo user repository
type UserBackyardRepo interface {
	UpdateUserStatus(ctx context.Context, userID string, userStatus, mailStatus int, email string) (err error)
	GetUserInfo(ctx context.Context, userID string) (user *entity.User, exist bool, err error)
	GetUserPage(ctx context.Context, page, pageSize int, user *entity.User,
		usernameOrDisplayName string, isStaff bool) (users []*entity.User, total int64, err error)
}

// UserBackyardService user service
type UserBackyardService struct {
	userRepo           UserBackyardRepo
	userRoleRelService *role.UserRoleRelService
	authService        *auth.AuthService
}

// NewUserBackyardService new user backyard service
func NewUserBackyardService(
	userRepo UserBackyardRepo,
	userRoleRelService *role.UserRoleRelService,
	authService *auth.AuthService,
) *UserBackyardService {
	return &UserBackyardService{
		userRepo:           userRepo,
		userRoleRelService: userRoleRelService,
		authService:        authService,
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

// UpdateUserRole update user role
func (us *UserBackyardService) UpdateUserRole(ctx context.Context, req *schema.UpdateUserRoleReq) (err error) {
	// Users cannot modify their roles
	if req.UserID == req.LoginUserID {
		return errors.BadRequest(reason.UserCannotUpdateYourRole)
	}

	err = us.userRoleRelService.SaveUserRole(ctx, req.UserID, req.RoleID)
	if err != nil {
		return err
	}

	us.authService.RemoveAllUserTokens(ctx, req.UserID)
	return
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

	if len(req.Query) > 0 {
		if email, e := mail.ParseAddress(req.Query); e == nil {
			user.EMail = email.Address
			req.Query = ""
		} else if strings.HasPrefix(req.Query, "user:") {
			id := strings.TrimSpace(strings.TrimPrefix(req.Query, "user:"))
			idSearch := true
			for _, r := range id {
				if !unicode.IsDigit(r) {
					idSearch = false
					break
				}
			}
			if idSearch {
				user.ID = id
				req.Query = ""
			} else {
				req.Query = id
			}
		}
	}

	users, total, err := us.userRepo.GetUserPage(ctx, req.Page, req.PageSize, user, req.Query, req.Staff)
	if err != nil {
		return
	}

	resp := make([]*schema.GetUserPageResp, 0)
	for _, u := range users {
		avatar := schema.FormatAvatarInfo(u.Avatar)
		t := &schema.GetUserPageResp{
			UserID:      u.ID,
			CreatedAt:   u.CreatedAt.Unix(),
			Username:    u.Username,
			EMail:       u.EMail,
			Rank:        u.Rank,
			DisplayName: u.DisplayName,
			Avatar:      avatar,
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
	us.setUserRoleInfo(ctx, resp)
	return pager.NewPageModel(total, resp), nil
}

func (us *UserBackyardService) setUserRoleInfo(ctx context.Context, resp []*schema.GetUserPageResp) {
	var userIDs []string
	for _, u := range resp {
		userIDs = append(userIDs, u.UserID)
	}

	userRoleMapping, err := us.userRoleRelService.GetUserRoleMapping(ctx, userIDs)
	if err != nil {
		log.Error(err)
		return
	}

	for _, u := range resp {
		r := userRoleMapping[u.UserID]
		if r == nil {
			continue
		}
		u.RoleID = r.ID
		u.RoleName = r.Name
	}
}
