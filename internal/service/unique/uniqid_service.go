package unique

import (
	"context"
)

// UniqueIDRepo unique id repository
type UniqueIDRepo interface {
	GenUniqueID(ctx context.Context, key string) (uniqueID int64, err error)
	GenUniqueIDStr(ctx context.Context, key string) (uniqueID string, err error)
}
