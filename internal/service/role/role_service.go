package role

import (
	"context"

	"github.com/answerdev/answer/internal/base/handler"
	"github.com/answerdev/answer/internal/base/translator"
	"github.com/answerdev/answer/internal/entity"
	"github.com/answerdev/answer/internal/schema"
	"github.com/jinzhu/copier"
)

const (
	// Since there is currently no need to edit roles to add roles and other operations,
	// the current role information is translated directly.
	// Later on, when the relevant ability is available, it can be adjusted by the user himself.

	RoleUserID      = 1
	RoleAdminID     = 2
	RoleModeratorID = 3

	roleUserName      = "User"
	roleAdminName     = "Admin"
	roleModeratorName = "Moderator"

	trRoleNameUser      = "role.name.user"
	trRoleNameAdmin     = "role.name.admin"
	trRoleNameModerator = "role.name.moderator"

	trRoleDescriptionUser      = "role.description.user"
	trRoleDescriptionAdmin     = "role.description.admin"
	trRoleDescriptionModerator = "role.description.moderator"
)

// RoleRepo role repository
type RoleRepo interface {
	GetRoleAllList(ctx context.Context) (roles []*entity.Role, err error)
	GetRoleAllMapping(ctx context.Context) (roleMapping map[int]*entity.Role, err error)
}

// RoleService user service
type RoleService struct {
	roleRepo RoleRepo
}

func NewRoleService(roleRepo RoleRepo) *RoleService {
	return &RoleService{
		roleRepo: roleRepo,
	}
}

// GetRoleList get role list all
func (rs *RoleService) GetRoleList(ctx context.Context) (resp []*schema.GetRoleResp, err error) {
	roles, err := rs.roleRepo.GetRoleAllList(ctx)
	if err != nil {
		return
	}

	for _, role := range roles {
		rs.translateRole(ctx, role)
	}

	resp = []*schema.GetRoleResp{}
	_ = copier.Copy(&resp, roles)
	return
}

func (rs *RoleService) GetRoleMapping(ctx context.Context) (roleMapping map[int]*entity.Role, err error) {
	return rs.roleRepo.GetRoleAllMapping(ctx)
}

func (rs *RoleService) translateRole(ctx context.Context, role *entity.Role) {
	switch role.Name {
	case roleUserName:
		role.Name = translator.GlobalTrans.Tr(handler.GetLangByCtx(ctx), trRoleNameUser)
		role.Description = translator.GlobalTrans.Tr(handler.GetLangByCtx(ctx), trRoleDescriptionUser)
	case roleAdminName:
		role.Name = translator.GlobalTrans.Tr(handler.GetLangByCtx(ctx), trRoleNameAdmin)
		role.Description = translator.GlobalTrans.Tr(handler.GetLangByCtx(ctx), trRoleDescriptionAdmin)
	case roleModeratorName:
		role.Name = translator.GlobalTrans.Tr(handler.GetLangByCtx(ctx), trRoleNameModerator)
		role.Description = translator.GlobalTrans.Tr(handler.GetLangByCtx(ctx), trRoleDescriptionModerator)
	}
	return
}
