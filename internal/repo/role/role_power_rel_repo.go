package role

import (
	"context"

	"github.com/answerdev/answer/internal/base/data"
	"github.com/answerdev/answer/internal/base/reason"
	"github.com/answerdev/answer/internal/service/role"
	"github.com/segmentfault/pacman/errors"
	"xorm.io/builder"
)

// rolePowerRelRepo rolePowerRel repository
type rolePowerRelRepo struct {
	data *data.Data
}

// NewRolePowerRelRepo new repository
func NewRolePowerRelRepo(data *data.Data) role.RolePowerRelRepo {
	return &rolePowerRelRepo{
		data: data,
	}
}

// GetRolePowerTypeList get role power type list
func (rr *rolePowerRelRepo) GetRolePowerTypeList(ctx context.Context, roleID int) (powers []string, err error) {
	powers = make([]string, 0)
	err = rr.data.DB.Context(ctx).Table("role_power_rel").
		Cols("power_type").Where(builder.Eq{"role_id": roleID}).Find(&powers)
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}
