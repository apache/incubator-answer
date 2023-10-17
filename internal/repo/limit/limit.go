package limit

import (
	"context"
	"fmt"
	"github.com/answerdev/answer/internal/base/constant"
	"github.com/answerdev/answer/internal/base/data"
	"github.com/answerdev/answer/internal/base/reason"
	"github.com/segmentfault/pacman/errors"
	"time"
)

// LimitRepo auth repository
type LimitRepo struct {
	data *data.Data
}

// NewRateLimitRepo new repository
func NewRateLimitRepo(data *data.Data) *LimitRepo {
	return &LimitRepo{
		data: data,
	}
}

// CheckAndRecord check
func (lr *LimitRepo) CheckAndRecord(ctx context.Context, key string) (limit bool, err error) {
	_, exist, err := lr.data.Cache.GetString(ctx, constant.RateLimitCacheKeyPrefix+key)
	if err != nil {
		return false, errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	if exist {
		return true, nil
	}
	err = lr.data.Cache.SetString(ctx, constant.RateLimitCacheKeyPrefix+key,
		fmt.Sprintf("%d", time.Now().Unix()), constant.RateLimitCacheTime)
	if err != nil {
		return false, errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return false, nil
}

// ClearRecord clear
func (lr *LimitRepo) ClearRecord(ctx context.Context, key string) error {
	return lr.data.Cache.Del(ctx, constant.RateLimitCacheKeyPrefix+key)
}
