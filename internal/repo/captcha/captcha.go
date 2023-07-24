package captcha

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/answerdev/answer/internal/base/data"
	"github.com/answerdev/answer/internal/base/reason"
	"github.com/answerdev/answer/internal/entity"
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

func (cr *captchaRepo) SetActionType(ctx context.Context, unit, actionType, config string, amount int) (err error) {
	now := time.Now()
	cacheKey := fmt.Sprintf("ActionRecord:%s@%s@%s", unit, actionType, now.Format("2006-1-02"))
	value := &entity.ActionRecordInfo{}
	value.LastTime = now.Unix()
	value.Num = amount
	valueStr, err := json.Marshal(value)
	if err != nil {
		return nil
	}
	err = cr.data.Cache.SetString(ctx, cacheKey, string(valueStr), 6*time.Minute)
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}

func (cr *captchaRepo) GetActionType(ctx context.Context, unit, actionType string) (actioninfo *entity.ActionRecordInfo, err error) {
	now := time.Now()
	cacheKey := fmt.Sprintf("ActionRecord:%s@%s@%s", unit, actionType, now.Format("2006-1-02"))
	actioninfo = &entity.ActionRecordInfo{}
	res, err := cr.data.Cache.GetString(ctx, cacheKey)
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	err = json.Unmarshal([]byte(res), actioninfo)
	if err != nil {
		return actioninfo, nil
	}
	return actioninfo, nil
}

func (cr *captchaRepo) DelActionType(ctx context.Context, unit, actionType string) (err error) {
	now := time.Now()
	cacheKey := fmt.Sprintf("ActionRecord:%s@%s@%s", unit, actionType, now.Format("2006-1-02"))
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
