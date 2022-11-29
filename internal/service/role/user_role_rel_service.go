package role

import (
	"context"

	"github.com/answerdev/answer/internal/entity"
)

// UserRoleRelRepo userRoleRel repository
type UserRoleRelRepo interface {
	SaveUserRoleRel(ctx context.Context, userID string, roleID int) (err error)
	GetUserRoleRelList(ctx context.Context, userIDs []string) (userRoleRelList []*entity.UserRoleRel, err error)
}

// UserRoleRelService user service
type UserRoleRelService struct {
	userRoleRelRepo UserRoleRelRepo
	roleService     *RoleService
}

// NewUserRoleRelService new user role rel service
func NewUserRoleRelService(userRoleRelRepo UserRoleRelRepo, roleService *RoleService) *UserRoleRelService {
	return &UserRoleRelService{
		userRoleRelRepo: userRoleRelRepo,
		roleService:     roleService,
	}
}

// SaveUserRole save user role
func (us *UserRoleRelService) SaveUserRole(ctx context.Context, userID string, roleID int) (err error) {
	return us.userRoleRelRepo.SaveUserRoleRel(ctx, userID, roleID)
}

// GetUserRoleMapping get user role mapping
func (us *UserRoleRelService) GetUserRoleMapping(ctx context.Context, userIDs []string) (
	userRoleMapping map[string]*entity.Role, err error) {
	userRoleMapping = make(map[string]*entity.Role, 0)
	roleMapping, err := us.roleService.GetRoleMapping(ctx)
	if err != nil {
		return userRoleMapping, err
	}
	if len(roleMapping) == 0 {
		return userRoleMapping, nil
	}

	relMapping, err := us.GetUserRoleRelMapping(ctx, userIDs)
	if err != nil {
		return userRoleMapping, err
	}

	// default role is user
	defaultRole := roleMapping[1]
	for _, userID := range userIDs {
		roleID, ok := relMapping[userID]
		if !ok {
			userRoleMapping[userID] = defaultRole
			continue
		}
		userRoleMapping[userID] = roleMapping[roleID]
		if userRoleMapping[userID] == nil {
			userRoleMapping[userID] = defaultRole
		}
	}
	return userRoleMapping, nil
}

// GetUserRoleRelMapping get user role rel mapping
func (us *UserRoleRelService) GetUserRoleRelMapping(ctx context.Context, userIDs []string) (
	userRoleRelMapping map[string]int, err error) {
	userRoleRelMapping = make(map[string]int, 0)

	relList, err := us.userRoleRelRepo.GetUserRoleRelList(ctx, userIDs)
	if err != nil {
		return userRoleRelMapping, err
	}

	for _, rel := range relList {
		userRoleRelMapping[rel.UserID] = rel.RoleID
	}
	return userRoleRelMapping, nil
}
