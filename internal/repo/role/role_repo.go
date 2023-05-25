package role

import (
	"context"

	"github.com/answerdev/answer/internal/base/data"
	"github.com/answerdev/answer/internal/base/reason"
	"github.com/answerdev/answer/internal/entity"
	service "github.com/answerdev/answer/internal/service/role"
	"github.com/segmentfault/pacman/errors"
)

// roleRepo role repository
type roleRepo struct {
	data *data.Data
}

// NewRoleRepo new repository
func NewRoleRepo(data *data.Data) service.RoleRepo {
	return &roleRepo{
		data: data,
	}
}

// GetRoleAllList get role list all
func (rr *roleRepo) GetRoleAllList(ctx context.Context) (roleList []*entity.Role, err error) {
	roleList = make([]*entity.Role, 0)
	err = rr.data.DB.Context(ctx).Find(&roleList)
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}

// GetRoleAllMapping get role all mapping
func (rr *roleRepo) GetRoleAllMapping(ctx context.Context) (roleMapping map[int]*entity.Role, err error) {
	roleList, err := rr.GetRoleAllList(ctx)
	if err != nil {
		return nil, err
	}
	roleMapping = make(map[int]*entity.Role, 0)
	for _, role := range roleList {
		roleMapping[role.ID] = role
	}
	return roleMapping, nil
}
