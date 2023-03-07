package captcha

import (
	"context"
	"fmt"
	"time"

	"github.com/answerdev/answer/internal/base/data"
	"github.com/answerdev/answer/internal/base/reason"
	"github.com/answerdev/answer/internal/service/action"
	"github.com/segmentfault/pacman/errors"
	"github.com/segmentfault/pacman/log"
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
	cacheKey := fmt.Sprintf("ActionRecord:%s@", ip)
	err = cr.data.Cache.SetInt64(ctx, cacheKey, int64(amount), 6*time.Minute)
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}

func (cr *captchaRepo) GetActionType(ctx context.Context, ip, actionType string) (amount int, err error) {
	cacheKey := fmt.Sprintf("ActionRecord:%s@", ip)
	res, err := cr.data.Cache.GetInt64(ctx, cacheKey)
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return int(res), nil
}

func (cr *captchaRepo) DelActionType(ctx context.Context, ip, actionType string) (err error) {
	cacheKey := fmt.Sprintf("ActionRecord:%s@", ip)
	err = cr.data.Cache.Del(ctx, cacheKey)
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}

// SetCaptcha set captcha to cache
func (cr *captchaRepo) SetCaptcha(ctx context.Context, key, captcha string) (err error) {
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
		log.Debug(err)
	}
	return captcha, nil
}

func (cr *captchaRepo) DelCaptcha(ctx context.Context, key string) (err error) {
	err = cr.data.Cache.Del(ctx, key)
	if err != nil {
		log.Debug(err)
	}
	return nil
}
