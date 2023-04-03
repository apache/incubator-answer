package user_external_login

import (
	"context"
	"encoding/json"
	"time"

	"github.com/answerdev/answer/internal/entity"
	"github.com/answerdev/answer/internal/schema"
	"github.com/answerdev/answer/internal/service/activity"
	usercommon "github.com/answerdev/answer/internal/service/user_common"
	"github.com/answerdev/answer/pkg/random"
	"github.com/answerdev/answer/plugin"
	"github.com/segmentfault/pacman/log"
)

// UserCenterLoginService user external login service
type UserCenterLoginService struct {
	userRepo              usercommon.UserRepo
	userExternalLoginRepo UserExternalLoginRepo
	userCommonService     *usercommon.UserCommon
	userActivity          activity.UserActiveActivityRepo
}

// NewUserCenterLoginService new user external login service
func NewUserCenterLoginService(
	userRepo usercommon.UserRepo,
	userCommonService *usercommon.UserCommon,
	userExternalLoginRepo UserExternalLoginRepo,
	userActivity activity.UserActiveActivityRepo,
) *UserCenterLoginService {
	return &UserCenterLoginService{
		userRepo:              userRepo,
		userCommonService:     userCommonService,
		userExternalLoginRepo: userExternalLoginRepo,
		userActivity:          userActivity,
	}
}

func (us *UserCenterLoginService) ExternalLogin(
	ctx context.Context, provider string, basicUserInfo *plugin.UserCenterBasicUserInfo) (
	resp *schema.UserExternalLoginResp, err error) {

	oldExternalLoginUserInfo, exist, err := us.userExternalLoginRepo.GetByExternalID(ctx,
		provider, basicUserInfo.ExternalID)
	if err != nil {
		return nil, err
	}
	if exist {
		// if user is already a member, login directly
		oldUserInfo, exist, err := us.userRepo.GetByUserID(ctx, oldExternalLoginUserInfo.UserID)
		if err != nil {
			return nil, err
		}
		if exist {
			if err := us.userRepo.UpdateLastLoginDate(ctx, oldUserInfo.ID); err != nil {
				log.Errorf("update user last login date failed: %v", err)
			}
			accessToken, _, err := us.userCommonService.CacheLoginUserInfo(
				ctx, oldUserInfo.ID, oldUserInfo.MailStatus, oldUserInfo.Status)
			return &schema.UserExternalLoginResp{AccessToken: accessToken}, err
		}
	}

	oldUserInfo, err := us.registerNewUser(ctx, provider, basicUserInfo)
	if err != nil {
		return nil, err
	}

	us.activeUser(ctx, oldUserInfo)

	accessToken, _, err := us.userCommonService.CacheLoginUserInfo(
		ctx, oldUserInfo.ID, oldUserInfo.MailStatus, oldUserInfo.Status)
	return &schema.UserExternalLoginResp{AccessToken: accessToken}, err
}

func (us *UserCenterLoginService) registerNewUser(ctx context.Context, provider string,
	basicUserInfo *plugin.UserCenterBasicUserInfo) (userInfo *entity.User, err error) {
	userInfo = &entity.User{}
	userInfo.EMail = basicUserInfo.Email
	userInfo.DisplayName = basicUserInfo.DisplayName

	userInfo.Username, err = us.userCommonService.MakeUsername(ctx, basicUserInfo.Username)
	if err != nil {
		log.Error(err)
		userInfo.Username = random.Username()
	}

	if len(basicUserInfo.Avatar) > 0 {
		avatarInfo := &schema.AvatarInfo{
			Type:   schema.AvatarTypeCustom,
			Custom: basicUserInfo.Avatar,
		}
		avatar, _ := json.Marshal(avatarInfo)
		userInfo.Avatar = string(avatar)
	}

	userInfo.MailStatus = entity.EmailStatusAvailable
	userInfo.Status = entity.UserStatusAvailable
	userInfo.LastLoginDate = time.Now()
	err = us.userRepo.AddUser(ctx, userInfo)
	if err != nil {
		return nil, err
	}

	metaInfo, _ := json.Marshal(basicUserInfo)
	newExternalUserInfo := &entity.UserExternalLogin{
		UserID:     userInfo.ID,
		Provider:   provider,
		ExternalID: basicUserInfo.ExternalID,
		MetaInfo:   string(metaInfo),
	}
	err = us.userExternalLoginRepo.AddUserExternalLogin(ctx, newExternalUserInfo)

	return userInfo, nil
}

func (us *UserCenterLoginService) activeUser(ctx context.Context, oldUserInfo *entity.User) {
	if err := us.userActivity.UserActive(ctx, oldUserInfo.ID); err != nil {
		log.Error(err)
	}
}

func (us *UserCenterLoginService) UserCenterUserSettings(ctx context.Context, userID string) (
	resp *schema.UserCenterUserSettingsResp, err error) {
	resp = &schema.UserCenterUserSettingsResp{}

	userCenter, ok := plugin.GetUserCenter()
	if !ok {
		return resp, nil
	}

	// get external login info
	externalLoginList, err := us.userExternalLoginRepo.GetUserExternalLoginList(ctx, userID)
	if err != nil {
		return nil, err
	}
	var externalInfo *entity.UserExternalLogin
	for _, t := range externalLoginList {
		if t.Provider == userCenter.Info().SlugName {
			externalInfo = t
		}
	}
	if externalInfo == nil {
		return resp, nil
	}

	settings, err := userCenter.UserSettings(externalInfo.ExternalID)
	if err != nil {
		log.Error(err)
		return resp, nil
	}

	if len(settings.AccountSettingRedirectURL) > 0 {
		resp.AccountSettingAgent = schema.UserSettingAgent{
			Enabled:     true,
			RedirectURL: settings.AccountSettingRedirectURL,
		}
	}
	if len(settings.ProfileSettingRedirectURL) > 0 {
		resp.ProfileSettingAgent = schema.UserSettingAgent{
			Enabled:     true,
			RedirectURL: settings.ProfileSettingRedirectURL,
		}
	}
	return resp, nil
}

func (us *UserCenterLoginService) UserCenterPersonalBranding(ctx context.Context, username string) (
	resp *schema.UserCenterPersonalBranding, err error) {
	resp = &schema.UserCenterPersonalBranding{
		PersonalBranding: make([]*schema.PersonalBranding, 0),
	}
	userCenter, ok := plugin.GetUserCenter()
	if !ok {
		return
	}

	userInfo, exist, err := us.userRepo.GetByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	if !exist {
		return resp, nil
	}

	// get external login info
	externalLoginList, err := us.userExternalLoginRepo.GetUserExternalLoginList(ctx, userInfo.ID)
	if err != nil {
		return nil, err
	}
	var externalInfo *entity.UserExternalLogin
	for _, t := range externalLoginList {
		if t.Provider == userCenter.Info().SlugName {
			externalInfo = t
		}
	}
	if externalInfo == nil {
		return resp, nil
	}

	resp.Enabled = true
	branding := userCenter.PersonalBranding(externalInfo.ExternalID)

	for _, t := range branding {
		resp.PersonalBranding = append(resp.PersonalBranding, &schema.PersonalBranding{
			Icon:  t.Icon,
			Name:  t.Name,
			Label: t.Label,
			Url:   t.Url,
		})
	}
	return resp, nil
}
