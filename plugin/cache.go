package plugin

import (
	"context"
	"time"
)

type Cache interface {
	Base

	GetString(ctx context.Context, key string) (data string, exist bool, err error)
	SetString(ctx context.Context, key, value string, ttl time.Duration) (err error)
	GetInt64(ctx context.Context, key string) (data int64, exist bool, err error)
	SetInt64(ctx context.Context, key string, value int64, ttl time.Duration) (err error)
	Increase(ctx context.Context, key string, value int64) (data int64, err error)
	Decrease(ctx context.Context, key string, value int64) (data int64, err error)
	Del(ctx context.Context, key string) (err error)
	Flush(ctx context.Context) (err error)
}

var (
	// CallCache is a function that calls all registered cache
	CallCache,
	registerCache = MakePlugin[Cache](false)
)
