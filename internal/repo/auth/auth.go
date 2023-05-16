package auth

import (
	"context"
	"encoding/json"

	"github.com/answerdev/answer/internal/base/constant"
	"github.com/answerdev/answer/internal/base/data"
	"github.com/answerdev/answer/internal/base/reason"
	"github.com/answerdev/answer/internal/entity"
	"github.com/answerdev/answer/internal/service/auth"
	"github.com/segmentfault/pacman/errors"
	"github.com/segmentfault/pacman/log"
)

// authRepo auth repository
type authRepo struct {
	data *data.Data
}

// GetUserCacheInfo get user cache info
func (ar *authRepo) GetUserCacheInfo(ctx context.Context, accessToken string) (userInfo *entity.UserCacheInfo, err error) {
	userInfoCache, err := ar.data.Cache.GetString(ctx, constant.UserTokenCacheKey+accessToken)
	if err != nil {
		return nil, errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	userInfo = &entity.UserCacheInfo{}
	err = json.Unmarshal([]byte(userInfoCache), userInfo)
	if err != nil {
		return nil, errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return userInfo, nil
}

// SetUserCacheInfo set user cache info
func (ar *authRepo) SetUserCacheInfo(ctx context.Context, accessToken string, userInfo *entity.UserCacheInfo) (err error) {
	userInfoCache, err := json.Marshal(userInfo)
	if err != nil {
		return err
	}
	err = ar.data.Cache.SetString(ctx, constant.UserTokenCacheKey+accessToken,
		string(userInfoCache), constant.UserTokenCacheTime)
	if err != nil {
		return errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	if err := ar.AddUserTokenMapping(ctx, userInfo.UserID, accessToken); err != nil {
		log.Error(err)
	}
	return nil
}

// RemoveUserCacheInfo remove user cache info
func (ar *authRepo) RemoveUserCacheInfo(ctx context.Context, accessToken string) (err error) {
	err = ar.data.Cache.Del(ctx, constant.UserTokenCacheKey+accessToken)
	if err != nil {
		return errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return nil
}

// SetUserStatus set user status
func (ar *authRepo) SetUserStatus(ctx context.Context, userID string, userInfo *entity.UserCacheInfo) (err error) {
	userInfoCache, err := json.Marshal(userInfo)
	if err != nil {
		return err
	}
	err = ar.data.Cache.SetString(ctx, constant.UserStatusChangedCacheKey+userID,
		string(userInfoCache), constant.UserStatusChangedCacheTime)
	if err != nil {
		return errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return nil
}

// GetUserStatus get user status
func (ar *authRepo) GetUserStatus(ctx context.Context, userID string) (userInfo *entity.UserCacheInfo, err error) {
	userInfoCache, err := ar.data.Cache.GetString(ctx, constant.UserStatusChangedCacheKey+userID)
	if err != nil {
		return nil, errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	userInfo = &entity.UserCacheInfo{}
	err = json.Unmarshal([]byte(userInfoCache), userInfo)
	if err != nil {
		return nil, errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return userInfo, nil
}

// RemoveUserStatus remove user status
func (ar *authRepo) RemoveUserStatus(ctx context.Context, userID string) (err error) {
	err = ar.data.Cache.Del(ctx, constant.UserStatusChangedCacheKey+userID)
	if err != nil {
		return errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return nil
}

// GetAdminUserCacheInfo get admin user cache info
func (ar *authRepo) GetAdminUserCacheInfo(ctx context.Context, accessToken string) (userInfo *entity.UserCacheInfo, err error) {
	userInfoCache, err := ar.data.Cache.GetString(ctx, constant.AdminTokenCacheKey+accessToken)
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
		return
	}
	userInfo = &entity.UserCacheInfo{}
	err = json.Unmarshal([]byte(userInfoCache), userInfo)
	if err != nil {
		return nil, err
	}
	return userInfo, nil
}

// SetAdminUserCacheInfo set admin user cache info
func (ar *authRepo) SetAdminUserCacheInfo(ctx context.Context, accessToken string, userInfo *entity.UserCacheInfo) (err error) {
	userInfoCache, err := json.Marshal(userInfo)
	if err != nil {
		return err
	}

	err = ar.data.Cache.SetString(ctx, constant.AdminTokenCacheKey+accessToken, string(userInfoCache),
		constant.AdminTokenCacheTime)
	if err != nil {
		return errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return nil
}

// RemoveAdminUserCacheInfo remove admin user cache info
func (ar *authRepo) RemoveAdminUserCacheInfo(ctx context.Context, accessToken string) (err error) {
	err = ar.data.Cache.Del(ctx, constant.AdminTokenCacheKey+accessToken)
	if err != nil {
		return errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return nil
}

// AddUserTokenMapping add user token mapping
func (ar *authRepo) AddUserTokenMapping(ctx context.Context, userID, accessToken string) (err error) {
	key := constant.UserTokenMappingCacheKey + userID
	resp, _ := ar.data.Cache.GetString(ctx, key)
	mapping := make(map[string]bool, 0)
	if len(resp) > 0 {
		_ = json.Unmarshal([]byte(resp), &mapping)
	}
	mapping[accessToken] = true
	content, _ := json.Marshal(mapping)
	return ar.data.Cache.SetString(ctx, key, string(content), constant.UserTokenCacheTime)
}

// RemoveUserTokens Log out all users under this user id
func (ar *authRepo) RemoveUserTokens(ctx context.Context, userID string, remainToken string) {
	key := constant.UserTokenMappingCacheKey + userID
	resp, _ := ar.data.Cache.GetString(ctx, key)
	mapping := make(map[string]bool, 0)
	if len(resp) > 0 {
		_ = json.Unmarshal([]byte(resp), &mapping)
		log.Debugf("find %d user tokens by user id %s", len(mapping), userID)
	}

	for token := range mapping {
		if token == remainToken {
			continue
		}
		if err := ar.RemoveUserCacheInfo(ctx, token); err != nil {
			log.Error(err)
		} else {
			log.Debugf("del user %s token success")
		}
	}
	if err := ar.RemoveUserStatus(ctx, userID); err != nil {
		log.Error(err)
	}
	if err := ar.data.Cache.Del(ctx, key); err != nil {
		log.Error(err)
	}
}

// NewAuthRepo new repository
func NewAuthRepo(data *data.Data) auth.AuthRepo {
	return &authRepo{
		data: data,
	}
}
