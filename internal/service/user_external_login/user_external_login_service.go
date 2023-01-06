package user_external_login

import (
	"context"
	"time"

	"github.com/answerdev/answer/internal/entity"
	"github.com/answerdev/answer/internal/plugin"
	"github.com/answerdev/answer/internal/schema"
	usercommon "github.com/answerdev/answer/internal/service/user_common"
)

type UserExternalLoginRepo interface {
	AddUserExternalLogin(ctx context.Context, user *entity.UserExternalLogin) (err error)
	UpdateInfo(ctx context.Context, userInfo *entity.UserExternalLogin) (err error)
	GetByExternalID(ctx context.Context, externalID string) (userInfo *entity.UserExternalLogin, exist bool, err error)
	SetCacheUserExternalLoginInfo(
		ctx context.Context, info plugin.ExternalLoginUserInfo) (err error)
	GetCacheUserExternalLoginInfo(
		ctx context.Context, externalID string) (info plugin.ExternalLoginUserInfo, err error)
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
	ctx context.Context, provider string, externalUserInfo plugin.ExternalLoginUserInfo) (
	resp *schema.UserExternalLoginResp, err error) {
	// cache external user info, waiting for user enter email address.
	if len(externalUserInfo.Email) == 0 {
		err = us.userExternalLoginRepo.SetCacheUserExternalLoginInfo(ctx, externalUserInfo)
		if err != nil {
			return nil, err
		}
		return &schema.UserExternalLoginResp{ExternalID: externalUserInfo.ExternalID}, nil
	}

	oldUserInfo, exist, err := us.userRepo.GetByEmail(ctx, externalUserInfo.Email)
	if err != nil {
		return nil, err
	}
	if !exist {
		oldUserInfo, err = us.RegisterNewUser(ctx, provider, externalUserInfo)
		if err != nil {
			return nil, err
		}
	}
	err = us.BindOldUser(ctx, provider, externalUserInfo, oldUserInfo)
	if err != nil {
		return nil, err
	}

	accessToken, _, err := us.userCommonService.CacheLoginUserInfo(
		ctx, oldUserInfo.ID, oldUserInfo.MailStatus, oldUserInfo.Status)
	return &schema.UserExternalLoginResp{AccessToken: accessToken}, err
}

func (us *UserExternalLoginService) RegisterNewUser(ctx context.Context, provider string,
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

func (us *UserExternalLoginService) BindOldUser(ctx context.Context, provider string,
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
			Provider:   provider,
			ExternalID: externalUserInfo.ExternalID,
			MetaInfo:   externalUserInfo.MetaInfo,
		}
		err = us.userExternalLoginRepo.AddUserExternalLogin(ctx, newExternalUserInfo)
	}
	return err
}
