package role

import (
	"context"

	"github.com/answerdev/answer/internal/base/data"
	"github.com/answerdev/answer/internal/base/reason"
	"github.com/answerdev/answer/internal/entity"
	"github.com/answerdev/answer/internal/service/role"
	"github.com/segmentfault/pacman/errors"
	"xorm.io/xorm"
)

// userRoleRelRepo userRoleRel repository
type userRoleRelRepo struct {
	data *data.Data
}

// NewUserRoleRelRepo new repository
func NewUserRoleRelRepo(data *data.Data) role.UserRoleRelRepo {
	return &userRoleRelRepo{
		data: data,
	}
}

// SaveUserRoleRel save user role rel
func (ur *userRoleRelRepo) SaveUserRoleRel(ctx context.Context, userID string, roleID int) (err error) {
	_, err = ur.data.DB.Transaction(func(session *xorm.Session) (interface{}, error) {
		item := &entity.UserRoleRel{UserID: userID}
		exist, err := ur.data.DB.Get(item)
		if err != nil {
			return nil, err
		}
		if exist {
			item.RoleID = roleID
			_, err = ur.data.DB.Update(item)
		} else {
			_, err = ur.data.DB.Insert(&entity.UserRoleRel{UserID: userID, RoleID: roleID})
		}
		if err != nil {
			return nil, err
		}
		return nil, nil
	})
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}

// GetUserRoleRelList get user role all
func (ur *userRoleRelRepo) GetUserRoleRelList(ctx context.Context, userIDs []string) (
	userRoleRelList []*entity.UserRoleRel, err error) {
	userRoleRelList = make([]*entity.UserRoleRel, 0)
	err = ur.data.DB.In("user_id", userIDs).Find(&userRoleRelList)
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}

// GetUserRoleRel get user role
func (ur *userRoleRelRepo) GetUserRoleRel(ctx context.Context, userID string) (
	rolePowerRel *entity.RolePowerRel, exist bool, err error) {
	rolePowerRel = &entity.RolePowerRel{}
	exist, err = ur.data.DB.Where("user_id", userID).Get(rolePowerRel)
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}
