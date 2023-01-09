package user_external_login

import (
	"context"
	"encoding/json"

	"github.com/answerdev/answer/internal/base/constant"
	"github.com/answerdev/answer/internal/base/data"
	"github.com/answerdev/answer/internal/base/reason"
	"github.com/answerdev/answer/internal/entity"
	"github.com/answerdev/answer/internal/schema"
	"github.com/answerdev/answer/internal/service/user_external_login"
	"github.com/segmentfault/pacman/errors"
)

type userExternalLoginRepo struct {
	data *data.Data
}

// NewUserExternalLoginRepo new repository
func NewUserExternalLoginRepo(data *data.Data) user_external_login.UserExternalLoginRepo {
	return &userExternalLoginRepo{
		data: data,
	}
}

// AddUserExternalLogin add external login information
func (ur *userExternalLoginRepo) AddUserExternalLogin(ctx context.Context, user *entity.UserExternalLogin) (err error) {
	_, err = ur.data.DB.Insert(user)
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}

// UpdateInfo update user info
func (ur *userExternalLoginRepo) UpdateInfo(ctx context.Context, userInfo *entity.UserExternalLogin) (err error) {
	_, err = ur.data.DB.ID(userInfo.ID).Update(userInfo)
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}

// GetByExternalID get by external ID
func (ur *userExternalLoginRepo) GetByExternalID(ctx context.Context, externalID string) (
	userInfo *entity.UserExternalLogin, exist bool, err error) {
	userInfo = &entity.UserExternalLogin{}
	exist, err = ur.data.DB.Where("external_id = ?", externalID).Get(userInfo)
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}

// SetCacheUserExternalLoginInfo cache user info for external login
func (ur *userExternalLoginRepo) SetCacheUserExternalLoginInfo(
	ctx context.Context, key string, info *schema.ExternalLoginUserInfoCache) (err error) {
	cacheData, _ := json.Marshal(info)
	return ur.data.Cache.SetString(ctx, constant.ConnectorUserExternalInfoCacheKey+key,
		string(cacheData), constant.ConnectorUserExternalInfoCacheTime)
}

// GetCacheUserExternalLoginInfo cache user info for external login
func (ur *userExternalLoginRepo) GetCacheUserExternalLoginInfo(
	ctx context.Context, key string) (info *schema.ExternalLoginUserInfoCache, err error) {
	res, err := ur.data.Cache.GetString(ctx, constant.ConnectorUserExternalInfoCacheKey+key)
	if err != nil {
		return info, err
	}
	_ = json.Unmarshal([]byte(res), &info)
	return info, nil
}
