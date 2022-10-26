package unique

import (
	"context"
)

// UniqueIDRepo unique id repository
type UniqueIDRepo interface {
	GenUniqueIDStr(ctx context.Context, key string) (uniqueID string, err error)
}
