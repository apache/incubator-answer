package reason_common

import (
	"context"

	"github.com/answerdev/answer/internal/schema"
)

type ReasonRepo interface {
	ListReasons(ctx context.Context, objectType, action string) (resp []*schema.ReasonItem, err error)
}
