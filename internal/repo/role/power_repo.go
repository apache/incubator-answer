package role

import (
	"context"

	"github.com/answerdev/answer/internal/base/data"
	"github.com/answerdev/answer/internal/base/reason"
	"github.com/answerdev/answer/internal/entity"
	"github.com/answerdev/answer/internal/service/role"
	"github.com/segmentfault/pacman/errors"
)

// powerRepo power repository
type powerRepo struct {
	data *data.Data
}

// NewPowerRepo new repository
func NewPowerRepo(data *data.Data) role.PowerRepo {
	return &powerRepo{
		data: data,
	}
}

// GetPowerList get  list all
func (pr *powerRepo) GetPowerList(ctx context.Context, power *entity.Power) (powerList []*entity.Power, err error) {
	powerList = make([]*entity.Power, 0)
	err = pr.data.DB.Context(ctx).Find(powerList, power)
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}
