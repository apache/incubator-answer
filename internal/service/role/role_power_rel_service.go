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
	rolePowerRelRepo   RolePowerRelRepo
	userRoleRelService *UserRoleRelService
}

// NewRolePowerRelService new role power rel service
func NewRolePowerRelService(rolePowerRelRepo RolePowerRelRepo,
	userRoleRelService *UserRoleRelService) *RolePowerRelService {
	return &RolePowerRelService{
		rolePowerRelRepo:   rolePowerRelRepo,
		userRoleRelService: userRoleRelService,
	}
}

// GetRolePowerList get role power list
func (rs *RolePowerRelService) GetRolePowerList(ctx context.Context, roleID int) (powers []string, err error) {
	return rs.rolePowerRelRepo.GetRolePowerTypeList(ctx, roleID)
}

// GetUserPowerList get  list all
func (rs *RolePowerRelService) GetUserPowerList(ctx context.Context, userID string) (powers []string, err error) {
	roleID, err := rs.userRoleRelService.GetUserRole(ctx, userID)
	if err != nil {
		return nil, err
	}
	return rs.rolePowerRelRepo.GetRolePowerTypeList(ctx, roleID)
}
