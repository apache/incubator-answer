package role

import (
	"context"

	"github.com/answerdev/answer/internal/entity"
)

// PowerRepo power repository
type PowerRepo interface {
	GetPowerList(ctx context.Context, power *entity.Power) (powers []*entity.Power, err error)
}
