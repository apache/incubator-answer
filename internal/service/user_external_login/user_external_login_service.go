package user_external_login

import (
	"context"
	"time"

	"github.com/answerdev/answer/internal/base/reason"
	"github.com/answerdev/answer/internal/entity"
	"github.com/answerdev/answer/internal/plugin"
	"github.com/answerdev/answer/internal/schema"
	usercommon "github.com/answerdev/answer/internal/service/user_common"
	"github.com/answerdev/answer/pkg/token"
	"github.com/segmentfault/pacman/errors"
)

type UserExternalLoginRepo interface {
	AddUserExternalLogin(ctx context.Context, user *entity.UserExternalLogin) (err error)
	UpdateInfo(ctx context.Context, userInfo *entity.UserExternalLogin) (err error)
	GetByExternalID(ctx context.Context, externalID string) (userInfo *entity.UserExternalLogin, exist bool, err error)
	SetCacheUserExternalLoginInfo(
		ctx context.Context, key string, info plugin.ExternalLoginUserInfo) (err error)
	GetCacheUserExternalLoginInfo(
		ctx context.Context, key string) (info plugin.ExternalLoginUserInfo, err error)
}

// UserExternalLoginService user external login service
type UserExternalLoginService struct {
	userRepo              usercommon.UserRepo
	userExternalLoginRepo UserExternalLoginRepo
	userCommonService     *usercommon.UserCommon
}

// NewUserExternalLoginService new user external login service
func NewUserExternalLoginService(
	userRepo usercommon.UserRepo,
	userCommonService *usercommon.UserCommon,
	userExternalLoginRepo UserExternalLoginRepo,
) *UserExternalLoginService {
	return &UserExternalLoginService{
		userRepo:              userRepo,
		userCommonService:     userCommonService,
		userExternalLoginRepo: userExternalLoginRepo,
	}
}

// ExternalLogin if user is already a member logged in
func (us *UserExternalLoginService) ExternalLogin(
	ctx context.Context, externalUserInfo plugin.ExternalLoginUserInfo) (
	resp *schema.UserExternalLoginResp, err error) {
	// cache external user info, waiting for user enter email address.
	if len(externalUserInfo.Email) == 0 {
		bindingKey := token.GenerateToken()
		err = us.userExternalLoginRepo.SetCacheUserExternalLoginInfo(ctx, bindingKey, externalUserInfo)
		if err != nil {
			return nil, err
		}
		return &schema.UserExternalLoginResp{BindingKey: bindingKey}, nil
	}

	oldUserInfo, exist, err := us.userRepo.GetByEmail(ctx, externalUserInfo.Email)
	if err != nil {
		return nil, err
	}
	if !exist {
		oldUserInfo, err = us.RegisterNewUser(ctx, externalUserInfo)
		if err != nil {
			return nil, err
		}
	}
	err = us.BindOldUser(ctx, externalUserInfo, oldUserInfo)
	if err != nil {
		return nil, err
	}

	accessToken, _, err := us.userCommonService.CacheLoginUserInfo(
		ctx, oldUserInfo.ID, oldUserInfo.MailStatus, oldUserInfo.Status)
	return &schema.UserExternalLoginResp{AccessToken: accessToken}, err
}

func (us *UserExternalLoginService) RegisterNewUser(ctx context.Context,
	externalUserInfo plugin.ExternalLoginUserInfo) (userInfo *entity.User, err error) {
	userInfo = &entity.User{}
	userInfo.EMail = externalUserInfo.Email
	userInfo.DisplayName = externalUserInfo.Name
	userInfo.Username, err = us.userCommonService.MakeUsername(ctx, externalUserInfo.Name)
	if err != nil {
		userInfo.Username = "" // TODO random
	}
	userInfo.MailStatus = entity.EmailStatusToBeVerified
	userInfo.Status = entity.UserStatusAvailable
	userInfo.LastLoginDate = time.Now()
	err = us.userRepo.AddUser(ctx, userInfo)
	if err != nil {
		return nil, err
	}
	return userInfo, nil
}

func (us *UserExternalLoginService) BindOldUser(ctx context.Context,
	externalUserInfo plugin.ExternalLoginUserInfo, oldUserInfo *entity.User) (err error) {
	oldExternalUserInfo, exist, err := us.userExternalLoginRepo.GetByExternalID(ctx, externalUserInfo.ExternalID)
	if err != nil {
		return err
	}
	if exist {
		oldExternalUserInfo.MetaInfo = externalUserInfo.MetaInfo
		oldExternalUserInfo.UserID = oldUserInfo.ID
		err = us.userExternalLoginRepo.UpdateInfo(ctx, oldExternalUserInfo)
	} else {
		newExternalUserInfo := &entity.UserExternalLogin{
			UserID:     oldUserInfo.ID,
			Provider:   externalUserInfo.Provider,
			ExternalID: externalUserInfo.ExternalID,
			MetaInfo:   externalUserInfo.MetaInfo,
		}
		err = us.userExternalLoginRepo.AddUserExternalLogin(ctx, newExternalUserInfo)
	}
	return err
}

func (us *UserExternalLoginService) ExternalLoginBindingUserSendEmail(
	ctx context.Context, req *schema.ExternalLoginBindingUserSendEmailReq) (
	resp *schema.ExternalLoginBindingUserSendEmailResp, err error) {
	resp = &schema.ExternalLoginBindingUserSendEmailResp{}
	externalLoginInfo, err := us.userExternalLoginRepo.GetCacheUserExternalLoginInfo(ctx, req.BindingKey)
	if err != nil || len(externalLoginInfo.ExternalID) == 0 {
		return nil, errors.BadRequest(reason.UserNotFound)
	}

	_, exist, err := us.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	if exist && !req.Must {
		resp.EmailExistAndMustBeConfirmed = true
		return resp, nil
	}

	if !exist {
		externalLoginInfo.Email = req.Email
		_, err = us.RegisterNewUser(ctx, externalLoginInfo)
		if err != nil {
			return nil, err
		}
	}

	// TODO send bind confirmation email
	return resp, nil
}

func (us *UserExternalLoginService) ExternalLoginBindingUser(
	ctx context.Context, req *schema.ExternalLoginBindingUserReq) (
	resp *schema.ExternalLoginBindingUserResp, err error) {
	return
}
