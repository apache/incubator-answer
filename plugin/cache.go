package plugin

import (
	"context"
	"time"
)

type Cache interface {
	Base

	GetString(ctx context.Context, key string) (string, error)
	SetString(ctx context.Context, key, value string, ttl time.Duration) error
	GetInt64(ctx context.Context, key string) (int64, error)
	SetInt64(ctx context.Context, key string, value int64, ttl time.Duration) error
	Del(ctx context.Context, key string) error
	Flush(ctx context.Context) error
}

var (
	// CallCache is a function that calls all registered cache
	CallCache,
	registerCache = MakePlugin[Cache](false)
)
