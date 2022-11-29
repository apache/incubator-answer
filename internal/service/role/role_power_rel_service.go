package role

import (
	"context"
)

// RolePowerRelRepo rolePowerRel repository
type RolePowerRelRepo interface {
	GetRolePowerTypeList(ctx context.Context, roleID int) (powers []string, err error)
}

// RolePowerRelService user service
type RolePowerRelService struct {
	rolePowerRelRepo RolePowerRelRepo
}

// NewRolePowerRelService new role power rel service
func NewRolePowerRelService(rolePowerRelRepo RolePowerRelRepo) *RolePowerRelService {
	return &RolePowerRelService{
		rolePowerRelRepo: rolePowerRelRepo,
	}
}

// GetRolePowerList get role power list
func (rs *RolePowerRelService) GetRolePowerList(ctx context.Context, roleID int) (powers []string, err error) {
	return rs.rolePowerRelRepo.GetRolePowerTypeList(ctx, roleID)
}
