package auth

import (
	"context"

	"github.com/answerdev/answer/internal/entity"
	"github.com/answerdev/answer/pkg/token"
	"github.com/answerdev/answer/plugin"
)

// AuthRepo auth repository
type AuthRepo interface {
	GetUserCacheInfo(ctx context.Context, accessToken string) (userInfo *entity.UserCacheInfo, err error)
	SetUserCacheInfo(ctx context.Context, accessToken, visitToken string, userInfo *entity.UserCacheInfo) error
	GetUserVisitCacheInfo(ctx context.Context, visitToken string) (accessToken string, err error)
	RemoveUserCacheInfo(ctx context.Context, accessToken string) (err error)
	SetUserStatus(ctx context.Context, userID string, userInfo *entity.UserCacheInfo) (err error)
	GetUserStatus(ctx context.Context, userID string) (userInfo *entity.UserCacheInfo, err error)
	RemoveUserStatus(ctx context.Context, userID string) (err error)
	GetAdminUserCacheInfo(ctx context.Context, accessToken string) (userInfo *entity.UserCacheInfo, err error)
	SetAdminUserCacheInfo(ctx context.Context, accessToken string, userInfo *entity.UserCacheInfo) error
	RemoveAdminUserCacheInfo(ctx context.Context, accessToken string) (err error)
	AddUserTokenMapping(ctx context.Context, userID, accessToken string) (err error)
	RemoveUserTokens(ctx context.Context, userID string, remainToken string)
}

// AuthService kit service
type AuthService struct {
	authRepo AuthRepo
}

// NewAuthService email service
func NewAuthService(authRepo AuthRepo) *AuthService {
	return &AuthService{
		authRepo: authRepo,
	}
}

func (as *AuthService) GetUserCacheInfo(ctx context.Context, accessToken string) (userInfo *entity.UserCacheInfo, err error) {
	userCacheInfo, err := as.authRepo.GetUserCacheInfo(ctx, accessToken)
	if err != nil {
		return nil, err
	}
	if userCacheInfo == nil {
		return nil, nil
	}
	cacheInfo, _ := as.authRepo.GetUserStatus(ctx, userCacheInfo.UserID)
	if cacheInfo != nil {
		userCacheInfo.UserStatus = cacheInfo.UserStatus
		userCacheInfo.EmailStatus = cacheInfo.EmailStatus
		userCacheInfo.RoleID = cacheInfo.RoleID
		// update current user cache info
		err := as.authRepo.SetUserCacheInfo(ctx, accessToken, userCacheInfo.VisitToken, userCacheInfo)
		if err != nil {
			return nil, err
		}
	}

	// try to get user status from user center
	uc, ok := plugin.GetUserCenter()
	if ok && len(userCacheInfo.ExternalID) > 0 {
		if userStatus := uc.UserStatus(userCacheInfo.ExternalID); userStatus != plugin.UserStatusAvailable {
			userCacheInfo.UserStatus = int(userStatus)
		}
	}
	return userCacheInfo, nil
}

func (as *AuthService) SetUserCacheInfo(ctx context.Context, userInfo *entity.UserCacheInfo) (
	accessToken string, visitToken string, err error) {
	accessToken = token.GenerateToken()
	visitToken = token.GenerateToken()
	err = as.authRepo.SetUserCacheInfo(ctx, accessToken, visitToken, userInfo)
	if err != nil {
		return "", "", err
	}
	return accessToken, visitToken, err
}

func (as *AuthService) CheckUserVisitToken(ctx context.Context, visitToken string) bool {
	accessToken, err := as.authRepo.GetUserVisitCacheInfo(ctx, visitToken)
	if err != nil {
		return false
	}
	if len(accessToken) == 0 {
		return false
	}
	return true
}

func (as *AuthService) SetUserStatus(ctx context.Context, userInfo *entity.UserCacheInfo) (err error) {
	return as.authRepo.SetUserStatus(ctx, userInfo.UserID, userInfo)
}

func (as *AuthService) RemoveUserCacheInfo(ctx context.Context, accessToken string) (err error) {
	return as.authRepo.RemoveUserCacheInfo(ctx, accessToken)
}

// AddUserTokenMapping add user token mapping
func (as *AuthService) AddUserTokenMapping(ctx context.Context, userID, accessToken string) (err error) {
	return as.authRepo.AddUserTokenMapping(ctx, userID, accessToken)
}

// RemoveUserAllTokens Log out all users under this user id
func (as *AuthService) RemoveUserAllTokens(ctx context.Context, userID string) {
	as.authRepo.RemoveUserTokens(ctx, userID, "")
}

// RemoveTokensExceptCurrentUser remove all tokens except the current user
func (as *AuthService) RemoveTokensExceptCurrentUser(ctx context.Context, userID string, accessToken string) {
	as.authRepo.RemoveUserTokens(ctx, userID, accessToken)
}

//Admin

func (as *AuthService) GetAdminUserCacheInfo(ctx context.Context, accessToken string) (userInfo *entity.UserCacheInfo, err error) {
	return as.authRepo.GetAdminUserCacheInfo(ctx, accessToken)
}

func (as *AuthService) SetAdminUserCacheInfo(ctx context.Context, accessToken string, userInfo *entity.UserCacheInfo) (err error) {
	err = as.authRepo.SetAdminUserCacheInfo(ctx, accessToken, userInfo)
	return err
}

func (as *AuthService) RemoveAdminUserCacheInfo(ctx context.Context, accessToken string) (err error) {
	return as.authRepo.RemoveAdminUserCacheInfo(ctx, accessToken)
}
