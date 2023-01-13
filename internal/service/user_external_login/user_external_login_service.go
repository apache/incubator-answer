package user_external_login

import (
	"context"
	"fmt"
	"time"

	"github.com/answerdev/answer/internal/base/reason"
	"github.com/answerdev/answer/internal/entity"
	"github.com/answerdev/answer/internal/schema"
	"github.com/answerdev/answer/internal/service/activity"
	"github.com/answerdev/answer/internal/service/export"
	"github.com/answerdev/answer/internal/service/siteinfo_common"
	usercommon "github.com/answerdev/answer/internal/service/user_common"
	"github.com/answerdev/answer/pkg/random"
	"github.com/answerdev/answer/pkg/token"
	"github.com/google/uuid"
	"github.com/segmentfault/pacman/errors"
	"github.com/segmentfault/pacman/log"
)

type UserExternalLoginRepo interface {
	AddUserExternalLogin(ctx context.Context, user *entity.UserExternalLogin) (err error)
	UpdateInfo(ctx context.Context, userInfo *entity.UserExternalLogin) (err error)
	GetByExternalID(ctx context.Context, externalID string) (userInfo *entity.UserExternalLogin, exist bool, err error)
	GetUserExternalLoginList(ctx context.Context, userID string) (
		resp []*entity.UserExternalLogin, err error)
	DeleteUserExternalLogin(ctx context.Context, userID, externalID string) (err error)
	SetCacheUserExternalLoginInfo(ctx context.Context, key string, info *schema.ExternalLoginUserInfoCache) (err error)
	GetCacheUserExternalLoginInfo(ctx context.Context, key string) (info *schema.ExternalLoginUserInfoCache, err error)
}

// UserExternalLoginService user external login service
type UserExternalLoginService struct {
	userRepo              usercommon.UserRepo
	userExternalLoginRepo UserExternalLoginRepo
	userCommonService     *usercommon.UserCommon
	emailService          *export.EmailService
	siteInfoCommonService *siteinfo_common.SiteInfoCommonService
	userActivity          activity.UserActiveActivityRepo
}

// NewUserExternalLoginService new user external login service
func NewUserExternalLoginService(
	userRepo usercommon.UserRepo,
	userCommonService *usercommon.UserCommon,
	userExternalLoginRepo UserExternalLoginRepo,
	emailService *export.EmailService,
	siteInfoCommonService *siteinfo_common.SiteInfoCommonService,
	userActivity activity.UserActiveActivityRepo,
) *UserExternalLoginService {
	return &UserExternalLoginService{
		userRepo:              userRepo,
		userCommonService:     userCommonService,
		userExternalLoginRepo: userExternalLoginRepo,
		emailService:          emailService,
		siteInfoCommonService: siteInfoCommonService,
		userActivity:          userActivity,
	}
}

// ExternalLogin if user is already a member logged in
func (us *UserExternalLoginService) ExternalLogin(
	ctx context.Context, externalUserInfo *schema.ExternalLoginUserInfoCache) (
	resp *schema.UserExternalLoginResp, err error) {
	oldExternalLoginUserInfo, exist, err := us.userExternalLoginRepo.GetByExternalID(ctx, externalUserInfo.ExternalID)
	if err != nil {
		return nil, err
	}
	if exist {
		oldUserInfo, exist, err := us.userRepo.GetByUserID(ctx, oldExternalLoginUserInfo.UserID)
		if err != nil {
			return nil, err
		}
		if exist {
			accessToken, _, err := us.userCommonService.CacheLoginUserInfo(
				ctx, oldUserInfo.ID, oldUserInfo.MailStatus, oldUserInfo.Status)
			return &schema.UserExternalLoginResp{AccessToken: accessToken}, err
		}
	}

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
		oldUserInfo, err = us.registerNewUser(ctx, externalUserInfo)
		if err != nil {
			return nil, err
		}
	}
	err = us.bindOldUser(ctx, externalUserInfo, oldUserInfo)
	if err != nil {
		return nil, err
	}

	accessToken, _, err := us.userCommonService.CacheLoginUserInfo(
		ctx, oldUserInfo.ID, oldUserInfo.MailStatus, oldUserInfo.Status)
	return &schema.UserExternalLoginResp{AccessToken: accessToken}, err
}

func (us *UserExternalLoginService) registerNewUser(ctx context.Context,
	externalUserInfo *schema.ExternalLoginUserInfoCache) (userInfo *entity.User, err error) {
	userInfo = &entity.User{}
	userInfo.EMail = externalUserInfo.Email
	userInfo.DisplayName = externalUserInfo.Name
	userInfo.Username, err = us.userCommonService.MakeUsername(ctx, externalUserInfo.Name)
	if err != nil {
		userInfo.Username = random.Username()
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

func (us *UserExternalLoginService) bindOldUser(ctx context.Context,
	externalUserInfo *schema.ExternalLoginUserInfoCache, oldUserInfo *entity.User) (err error) {
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

// ExternalLoginBindingUserSendEmail Send an email for third-party account login for binding user
func (us *UserExternalLoginService) ExternalLoginBindingUserSendEmail(
	ctx context.Context, req *schema.ExternalLoginBindingUserSendEmailReq) (
	resp *schema.ExternalLoginBindingUserSendEmailResp, err error) {
	siteGeneral, err := us.siteInfoCommonService.GetSiteGeneral(ctx)
	if err != nil {
		return nil, err
	}
	resp = &schema.ExternalLoginBindingUserSendEmailResp{}
	externalLoginInfo, err := us.userExternalLoginRepo.GetCacheUserExternalLoginInfo(ctx, req.BindingKey)
	if err != nil || len(externalLoginInfo.ExternalID) == 0 {
		return nil, errors.BadRequest(reason.UserNotFound)
	}
	if len(externalLoginInfo.Email) > 0 {
		log.Warnf("the binding email has been sent %s", req.BindingKey)
		return &schema.ExternalLoginBindingUserSendEmailResp{}, nil
	}

	userInfo, exist, err := us.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	if exist && !req.Must {
		resp.EmailExistAndMustBeConfirmed = true
		return resp, nil
	}

	if !exist {
		externalLoginInfo.Email = req.Email
		userInfo, err = us.registerNewUser(ctx, externalLoginInfo)
		if err != nil {
			return nil, err
		}
		resp.AccessToken, _, err = us.userCommonService.CacheLoginUserInfo(
			ctx, userInfo.ID, userInfo.MailStatus, userInfo.Status)
		if err != nil {
			log.Error(err)
		}
	}
	err = us.userExternalLoginRepo.SetCacheUserExternalLoginInfo(ctx, req.BindingKey, externalLoginInfo)
	if err != nil {
		return nil, err
	}

	// send bind confirmation email
	data := &schema.EmailCodeContent{
		SourceType: schema.BindingSourceType,
		Email:      req.Email,
		UserID:     userInfo.ID,
		BindingKey: req.BindingKey,
	}
	code := uuid.NewString()
	verifyEmailURL := fmt.Sprintf("%s/users/account-activation?code=%s", siteGeneral.SiteUrl, code)
	title, body, err := us.emailService.RegisterTemplate(ctx, verifyEmailURL)
	if err != nil {
		return nil, err
	}
	go us.emailService.SendAndSaveCode(ctx, userInfo.EMail, title, body, code, data.ToJSONString())
	return resp, nil
}

// ExternalLoginBindingUser
// The user clicks on the email link of the bound account and requests the API to bind the user officially
func (us *UserExternalLoginService) ExternalLoginBindingUser(
	ctx context.Context, bindingKey string, oldUserInfo *entity.User) (err error) {
	externalLoginInfo, err := us.userExternalLoginRepo.GetCacheUserExternalLoginInfo(ctx, bindingKey)
	if err != nil || len(externalLoginInfo.ExternalID) == 0 {
		return errors.BadRequest(reason.UserNotFound)
	}
	return us.bindOldUser(ctx, externalLoginInfo, oldUserInfo)
}

// GetExternalLoginUserInfoList get external login user info list
func (us *UserExternalLoginService) GetExternalLoginUserInfoList(
	ctx context.Context, userID string) (resp []*entity.UserExternalLogin, err error) {
	return us.userExternalLoginRepo.GetUserExternalLoginList(ctx, userID)
}

// ExternalLoginUnbinding external login unbinding
func (us *UserExternalLoginService) ExternalLoginUnbinding(
	ctx context.Context, req *schema.ExternalLoginUnbindingReq) (err error) {
	return us.userExternalLoginRepo.DeleteUserExternalLogin(ctx, req.UserID, req.ExternalID)
}
