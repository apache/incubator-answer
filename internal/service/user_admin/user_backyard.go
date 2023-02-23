package user_admin

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
	"github.com/answerdev/answer/internal/service/activity"
	"github.com/answerdev/answer/internal/service/auth"
	"github.com/answerdev/answer/internal/service/role"
	usercommon "github.com/answerdev/answer/internal/service/user_common"
	"github.com/jinzhu/copier"
	"github.com/segmentfault/pacman/errors"
	"github.com/segmentfault/pacman/log"
	"golang.org/x/crypto/bcrypt"
)

// UserAdminRepo user repository
type UserAdminRepo interface {
	UpdateUserStatus(ctx context.Context, userID string, userStatus, mailStatus int, email string) (err error)
	GetUserInfo(ctx context.Context, userID string) (user *entity.User, exist bool, err error)
	GetUserInfoByEmail(ctx context.Context, email string) (user *entity.User, exist bool, err error)
	GetUserPage(ctx context.Context, page, pageSize int, user *entity.User,
		usernameOrDisplayName string, isStaff bool) (users []*entity.User, total int64, err error)
	AddUser(ctx context.Context, user *entity.User) (err error)
	UpdateUserPassword(ctx context.Context, userID string, password string) (err error)
}

// UserAdminService user service
type UserAdminService struct {
	userRepo           UserAdminRepo
	userRoleRelService *role.UserRoleRelService
	authService        *auth.AuthService
	userCommonService  *usercommon.UserCommon
	userActivity       activity.UserActiveActivityRepo
}

// NewUserAdminService new user admin service
func NewUserAdminService(
	userRepo UserAdminRepo,
	userRoleRelService *role.UserRoleRelService,
	authService *auth.AuthService,
	userCommonService *usercommon.UserCommon,
	userActivity activity.UserActiveActivityRepo,
) *UserAdminService {
	return &UserAdminService{
		userRepo:           userRepo,
		userRoleRelService: userRoleRelService,
		authService:        authService,
		userCommonService:  userCommonService,
		userActivity:       userActivity,
	}
}

// UpdateUserStatus update user
func (us *UserAdminService) UpdateUserStatus(ctx context.Context, req *schema.UpdateUserStatusReq) (err error) {
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

	err = us.userRepo.UpdateUserStatus(ctx, userInfo.ID, userInfo.Status, userInfo.MailStatus, userInfo.EMail)
	if err != nil {
		return err
	}

	// if user reputation is zero means this user is inactive, so try to activate this user.
	if req.IsNormal() && userInfo.Rank == 0 {
		return us.userActivity.UserActive(ctx, userInfo.ID)
	}
	return nil
}

// UpdateUserRole update user role
func (us *UserAdminService) UpdateUserRole(ctx context.Context, req *schema.UpdateUserRoleReq) (err error) {
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

// AddUser add user
func (us *UserAdminService) AddUser(ctx context.Context, req *schema.AddUserReq) (err error) {
	_, has, err := us.userRepo.GetUserInfoByEmail(ctx, req.Email)
	if err != nil {
		return err
	}
	if has {
		return errors.BadRequest(reason.EmailDuplicate)
	}

	hashPwd, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	userInfo := &entity.User{}
	userInfo.EMail = req.Email
	userInfo.DisplayName = req.DisplayName
	userInfo.Pass = string(hashPwd)

	userInfo.Username, err = us.userCommonService.MakeUsername(ctx, userInfo.DisplayName)
	if err != nil {
		return err
	}
	userInfo.MailStatus = entity.EmailStatusAvailable
	userInfo.Status = entity.UserStatusAvailable
	userInfo.Rank = 1

	err = us.userRepo.AddUser(ctx, userInfo)
	if err != nil {
		return err
	}
	return
}

// UpdateUserPassword update user password
func (us *UserAdminService) UpdateUserPassword(ctx context.Context, req *schema.UpdateUserPasswordReq) (err error) {
	// Users cannot modify their password
	if req.UserID == req.LoginUserID {
		return errors.BadRequest(reason.AdminCannotUpdateTheirPassword)
	}
	userInfo, exist, err := us.userRepo.GetUserInfo(ctx, req.UserID)
	if err != nil {
		return err
	}
	if !exist {
		return errors.BadRequest(reason.UserNotFound)
	}

	hashPwd, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	err = us.userRepo.UpdateUserPassword(ctx, userInfo.ID, string(hashPwd))
	if err != nil {
		return err
	}
	// logout this user
	us.authService.RemoveAllUserTokens(ctx, req.UserID)
	return
}

// GetUserInfo get user one
func (us *UserAdminService) GetUserInfo(ctx context.Context, userID string) (resp *schema.GetUserInfoResp, err error) {
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
func (us *UserAdminService) GetUserPage(ctx context.Context, req *schema.GetUserPageReq) (pageModel *pager.PageModel, err error) {
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

func (us *UserAdminService) setUserRoleInfo(ctx context.Context, resp []*schema.GetUserPageResp) {
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
