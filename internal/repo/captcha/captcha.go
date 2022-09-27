package captcha

import (
	"context"
	"fmt"
	"time"

	"github.com/segmentfault/answer/internal/base/data"
	"github.com/segmentfault/answer/internal/base/reason"
	"github.com/segmentfault/answer/internal/service/action"
	"github.com/segmentfault/pacman/errors"
)

// captchaRepo captcha repository
type captchaRepo struct {
	data *data.Data
}

// NewCaptchaRepo new repository
func NewCaptchaRepo(data *data.Data) action.CaptchaRepo {
	return &captchaRepo{
		data: data,
	}
}

func (cr *captchaRepo) SetActionType(ctx context.Context, ip, actionType string, amount int) (err error) {
	cacheKey := fmt.Sprintf("ActionRecord:%s@%s", ip, actionType)
	err = cr.data.Cache.SetInt64(ctx, cacheKey, int64(amount), 6*time.Minute)
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}

func (cr *captchaRepo) GetActionType(ctx context.Context, ip, actionType string) (amount int, err error) {
	cacheKey := fmt.Sprintf("ActionRecord:%s@%s", ip, actionType)
	res, err := cr.data.Cache.GetInt64(ctx, cacheKey)
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	// TODO: cache reflect should return empty when key not found
	return int(res), nil
}

func (cr *captchaRepo) DelActionType(ctx context.Context, ip, actionType string) (err error) {
	cacheKey := fmt.Sprintf("ActionRecord:%s@%s", ip, actionType)
	err = cr.data.Cache.Del(ctx, cacheKey)
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}

// SetCaptcha set captcha to cache
func (cr *captchaRepo) SetCaptcha(ctx context.Context, key, captcha string) (err error) {
	// TODO make cache time to config
	err = cr.data.Cache.SetString(ctx, key, captcha, 6*time.Minute)
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}

// GetCaptcha get captcha from cache
func (cr *captchaRepo) GetCaptcha(ctx context.Context, key string) (captcha string, err error) {
	captcha, err = cr.data.Cache.GetString(ctx, key)
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	// TODO: cache reflect should return empty when key not found
	return captcha, nil
}
