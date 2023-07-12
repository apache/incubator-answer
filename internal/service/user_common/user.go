package usercommon

import (
	"context"
	"strings"

	"github.com/Chain-Zhang/pinyin"
	"github.com/answerdev/answer/internal/base/reason"
	"github.com/answerdev/answer/internal/entity"
	"github.com/answerdev/answer/internal/schema"
	"github.com/answerdev/answer/internal/service/auth"
	"github.com/answerdev/answer/internal/service/role"
	"github.com/answerdev/answer/internal/service/siteinfo_common"
	"github.com/answerdev/answer/pkg/checker"
	"github.com/answerdev/answer/pkg/random"
	"github.com/segmentfault/pacman/errors"
	"github.com/segmentfault/pacman/log"
)

type UserRepo interface {
	AddUser(ctx context.Context, user *entity.User) (err error)
	IncreaseAnswerCount(ctx context.Context, userID string, amount int) (err error)
	IncreaseQuestionCount(ctx context.Context, userID string, amount int) (err error)
	UpdateQuestionCount(ctx context.Context, userID string, count int64) (err error)
	UpdateAnswerCount(ctx context.Context, userID string, count int) (err error)
	UpdateLastLoginDate(ctx context.Context, userID string) (err error)
	UpdateEmailStatus(ctx context.Context, userID string, emailStatus int) error
	UpdateNoticeStatus(ctx context.Context, userID string, noticeStatus int) error
	UpdateEmail(ctx context.Context, userID, email string) error
	UpdateLanguage(ctx context.Context, userID, language string) error
	UpdatePass(ctx context.Context, userID, pass string) error
	UpdateInfo(ctx context.Context, userInfo *entity.User) (err error)
	GetByUserID(ctx context.Context, userID string) (userInfo *entity.User, exist bool, err error)
	BatchGetByID(ctx context.Context, ids []string) ([]*entity.User, error)
	GetByUsername(ctx context.Context, username string) (userInfo *entity.User, exist bool, err error)
	GetByUsernames(ctx context.Context, usernames []string) ([]*entity.User, error)
	GetByEmail(ctx context.Context, email string) (userInfo *entity.User, exist bool, err error)
	GetUserCount(ctx context.Context) (count int64, err error)
	SearchUserListByName(ctx context.Context, name string) (userList []*entity.User, err error)
}

// UserCommon user service
type UserCommon struct {
	userRepo              UserRepo
	userRoleService       *role.UserRoleRelService
	authService           *auth.AuthService
	siteInfoCommonService *siteinfo_common.SiteInfoCommonService
}

func NewUserCommon(
	userRepo UserRepo,
	userRoleService *role.UserRoleRelService,
	authService *auth.AuthService,
	siteInfoCommonService *siteinfo_common.SiteInfoCommonService,
) *UserCommon {
	return &UserCommon{
		userRepo:              userRepo,
		userRoleService:       userRoleService,
		authService:           authService,
		siteInfoCommonService: siteInfoCommonService,
	}
}

func (us *UserCommon) GetUserBasicInfoByID(ctx context.Context, ID string) (
	userBasicInfo *schema.UserBasicInfo, exist bool, err error) {
	userInfo, exist, err := us.userRepo.GetByUserID(ctx, ID)
	if err != nil {
		return nil, exist, err
	}
	info := us.FormatUserBasicInfo(ctx, userInfo)
	info.Avatar = us.siteInfoCommonService.FormatAvatar(ctx, userInfo.Avatar, userInfo.EMail).GetURL()
	return info, exist, nil
}

func (us *UserCommon) GetUserBasicInfoByUserName(ctx context.Context, username string) (*schema.UserBasicInfo, bool, error) {
	userInfo, exist, err := us.userRepo.GetByUsername(ctx, username)
	if err != nil {
		return nil, exist, err
	}
	info := us.FormatUserBasicInfo(ctx, userInfo)
	info.Avatar = us.siteInfoCommonService.FormatAvatar(ctx, userInfo.Avatar, userInfo.EMail).GetURL()
	return info, exist, nil
}

func (us *UserCommon) BatchGetUserBasicInfoByUserNames(ctx context.Context, usernames []string) (map[string]*schema.UserBasicInfo, error) {
	infomap := make(map[string]*schema.UserBasicInfo)
	list, err := us.userRepo.GetByUsernames(ctx, usernames)
	if err != nil {
		return infomap, err
	}
	avatarMapping := us.siteInfoCommonService.FormatListAvatar(ctx, list)
	for _, user := range list {
		info := us.FormatUserBasicInfo(ctx, user)
		info.Avatar = avatarMapping[user.ID].GetURL()
		infomap[user.Username] = info
	}
	return infomap, nil
}

func (us *UserCommon) UpdateAnswerCount(ctx context.Context, userID string, num int) error {
	return us.userRepo.UpdateAnswerCount(ctx, userID, num)
}

func (us *UserCommon) UpdateQuestionCount(ctx context.Context, userID string, num int64) error {
	return us.userRepo.UpdateQuestionCount(ctx, userID, num)
}

func (us *UserCommon) BatchUserBasicInfoByID(ctx context.Context, IDs []string) (map[string]*schema.UserBasicInfo, error) {
	userMap := make(map[string]*schema.UserBasicInfo)
	userList, err := us.userRepo.BatchGetByID(ctx, IDs)
	if err != nil {
		return userMap, err
	}
	avatarMapping := us.siteInfoCommonService.FormatListAvatar(ctx, userList)
	for _, user := range userList {
		info := us.FormatUserBasicInfo(ctx, user)
		info.Avatar = avatarMapping[user.ID].GetURL()
		userMap[user.ID] = info
	}
	return userMap, nil
}

// FormatUserBasicInfo format user basic info
func (us *UserCommon) FormatUserBasicInfo(ctx context.Context, userInfo *entity.User) *schema.UserBasicInfo {
	userBasicInfo := &schema.UserBasicInfo{}
	userBasicInfo.ID = userInfo.ID
	userBasicInfo.Username = userInfo.Username
	userBasicInfo.Rank = userInfo.Rank
	userBasicInfo.DisplayName = userInfo.DisplayName
	userBasicInfo.Website = userInfo.Website
	userBasicInfo.Location = userInfo.Location
	userBasicInfo.IPInfo = userInfo.IPInfo
	userBasicInfo.Status = schema.UserStatusShow[userInfo.Status]
	if userBasicInfo.Status == schema.UserDeleted {
		userBasicInfo.Avatar = ""
		userBasicInfo.DisplayName = "Anonymous"
	}
	return userBasicInfo
}

// MakeUsername
// Generate a unique Username based on the displayName
func (us *UserCommon) MakeUsername(ctx context.Context, displayName string) (username string, err error) {
	// Chinese processing
	if has := checker.IsChinese(displayName); has {
		str, err := pinyin.New(displayName).Split("").Mode(pinyin.WithoutTone).Convert()
		if err != nil {
			return "", errors.BadRequest(reason.UsernameInvalid)
		} else {
			displayName = str
		}
	}

	username = strings.ReplaceAll(displayName, " ", "_")
	username = strings.ToLower(username)
	suffix := ""

	if checker.IsInvalidUsername(username) {
		return "", errors.BadRequest(reason.UsernameInvalid)
	}

	if checker.IsReservedUsername(username) {
		return "", errors.BadRequest(reason.UsernameInvalid)
	}

	for {
		_, has, err := us.userRepo.GetByUsername(ctx, username+suffix)
		if err != nil {
			return "", err
		}
		if !has {
			break
		}
		suffix = random.UsernameSuffix()
	}
	return username + suffix, nil
}

func (us *UserCommon) CacheLoginUserInfo(ctx context.Context, userID string, userStatus, emailStatus int, externalID string) (
	accessToken string, userCacheInfo *entity.UserCacheInfo, err error) {
	roleID, err := us.userRoleService.GetUserRole(ctx, userID)
	if err != nil {
		log.Error(err)
	}

	userCacheInfo = &entity.UserCacheInfo{
		UserID:      userID,
		EmailStatus: emailStatus,
		UserStatus:  userStatus,
		RoleID:      roleID,
		ExternalID:  externalID,
	}

	accessToken, err = us.authService.SetUserCacheInfo(ctx, userCacheInfo)
	if err != nil {
		return "", nil, err
	}
	if userCacheInfo.RoleID == role.RoleAdminID {
		if err = us.authService.SetAdminUserCacheInfo(ctx, accessToken, &entity.UserCacheInfo{UserID: userID}); err != nil {
			return "", nil, err
		}
	}
	return accessToken, userCacheInfo, nil
}
