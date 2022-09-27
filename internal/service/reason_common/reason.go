package reason_common

import (
	"context"
	"github.com/segmentfault/answer/internal/schema"
)

type ReasonRepo interface {
	ListReasons(ctx context.Context, req schema.ReasonReq) (resp []schema.ReasonItem, err error)
}
