package role

import (
	"context"

	"github.com/answerdev/answer/internal/base/data"
	"github.com/answerdev/answer/internal/base/reason"
	"github.com/answerdev/answer/internal/entity"
	"github.com/answerdev/answer/internal/service/role"
	"github.com/segmentfault/pacman/errors"
	"xorm.io/builder"
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
		session = session.Context(ctx)
		item := &entity.UserRoleRel{UserID: userID}
		exist, err := session.Get(item)
		if err != nil {
			return nil, err
		}
		if exist {
			item.RoleID = roleID
			_, err = session.ID(item.ID).Update(item)
		} else {
			_, err = session.Insert(&entity.UserRoleRel{UserID: userID, RoleID: roleID})
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
	err = ur.data.DB.Context(ctx).In("user_id", userIDs).Find(&userRoleRelList)
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}

// GetUserRoleRelListByRoleID get user role all by role id
func (ur *userRoleRelRepo) GetUserRoleRelListByRoleID(ctx context.Context, roleIDs []int) (
	userRoleRelList []*entity.UserRoleRel, err error) {
	userRoleRelList = make([]*entity.UserRoleRel, 0)
	err = ur.data.DB.Context(ctx).In("role_id", roleIDs).Find(&userRoleRelList)
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}

// GetUserRoleRel get user role
func (ur *userRoleRelRepo) GetUserRoleRel(ctx context.Context, userID string) (
	rolePowerRel *entity.UserRoleRel, exist bool, err error) {
	rolePowerRel = &entity.UserRoleRel{}
	exist, err = ur.data.DB.Context(ctx).Where(builder.Eq{"user_id": userID}).Get(rolePowerRel)
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}
