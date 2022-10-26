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
)

// authRepo activity repository
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

// GetBackyardUserCacheInfo get backyard user cache info
func (ar *authRepo) GetBackyardUserCacheInfo(ctx context.Context, accessToken string) (userInfo *entity.UserCacheInfo, err error) {
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

// SetBackyardUserCacheInfo set backyard user cache info
func (ar *authRepo) SetBackyardUserCacheInfo(ctx context.Context, accessToken string, userInfo *entity.UserCacheInfo) (err error) {
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

// RemoveBackyardUserCacheInfo remove backyard user cache info
func (ar *authRepo) RemoveBackyardUserCacheInfo(ctx context.Context, accessToken string) (err error) {
	err = ar.data.Cache.Del(ctx, constant.AdminTokenCacheKey+accessToken)
	if err != nil {
		return errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return nil
}

// NewAuthRepo new repository
func NewAuthRepo(data *data.Data) auth.AuthRepo {
	return &authRepo{
		data: data,
	}
}
