package service

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/segmentfault/answer/internal/base/pager"
	"github.com/segmentfault/answer/internal/base/reason"
	"github.com/segmentfault/answer/internal/entity"
	"github.com/segmentfault/answer/internal/schema"
	"github.com/segmentfault/pacman/errors"
)

// UserGroupRepo userGroup repository
type UserGroupRepo interface {
	AddUserGroup(ctx context.Context, userGroup *entity.UserGroup) (err error)
	RemoveUserGroup(ctx context.Context, id int) (err error)
	UpdateUserGroup(ctx context.Context, userGroup *entity.UserGroup) (err error)
	GetUserGroup(ctx context.Context, id int) (userGroup *entity.UserGroup, exist bool, err error)
	GetUserGroupList(ctx context.Context, userGroup *entity.UserGroup) (userGroups []*entity.UserGroup, err error)
	GetUserGroupPage(ctx context.Context, page, pageSize int, userGroup *entity.UserGroup) (userGroups []*entity.UserGroup, total int64, err error)
}

// UserGroupService user service
type UserGroupService struct {
	userGroupRepo UserGroupRepo
}

func NewUserGroupService(userGroupRepo UserGroupRepo) *UserGroupService {
	return &UserGroupService{
		userGroupRepo: userGroupRepo,
	}
}

// AddUserGroup add user group
func (us *UserGroupService) AddUserGroup(ctx context.Context, req *schema.AddUserGroupReq) (err error) {
	userGroup := &entity.UserGroup{}
	_ = copier.Copy(userGroup, req)
	return us.userGroupRepo.AddUserGroup(ctx, userGroup)
}

// RemoveUserGroup delete user group
func (us *UserGroupService) RemoveUserGroup(ctx context.Context, id int) (err error) {
	return us.userGroupRepo.RemoveUserGroup(ctx, id)
}

// UpdateUserGroup update user group
func (us *UserGroupService) UpdateUserGroup(ctx context.Context, req *schema.UpdateUserGroupReq) (err error) {
	userGroup := &entity.UserGroup{}
	_ = copier.Copy(userGroup, req)
	return us.userGroupRepo.UpdateUserGroup(ctx, userGroup)
}

// GetUserGroup get user group one
func (us *UserGroupService) GetUserGroup(ctx context.Context, id int) (resp *schema.GetUserGroupResp, err error) {
	userGroup, exist, err := us.userGroupRepo.GetUserGroup(ctx, id)
	if err != nil {
		return
	}
	if !exist {
		return nil, errors.BadRequest(reason.UnknownError)
	}

	resp = &schema.GetUserGroupResp{}
	_ = copier.Copy(resp, userGroup)
	return resp, nil
}

// GetUserGroupList get user group list all
func (us *UserGroupService) GetUserGroupList(ctx context.Context, req *schema.GetUserGroupListReq) (resp *[]schema.GetUserGroupResp, err error) {
	userGroup := &entity.UserGroup{}
	_ = copier.Copy(userGroup, req)

	userGroups, err := us.userGroupRepo.GetUserGroupList(ctx, userGroup)
	if err != nil {
		return
	}

	resp = &[]schema.GetUserGroupResp{}
	_ = copier.Copy(resp, userGroups)
	return
}

// GetUserGroupWithPage get user group list page
func (us *UserGroupService) GetUserGroupWithPage(ctx context.Context, req *schema.GetUserGroupWithPageReq) (pageModel *pager.PageModel, err error) {
	userGroup := &entity.UserGroup{}
	_ = copier.Copy(userGroup, req)

	page := req.Page
	pageSize := req.PageSize

	userGroups, total, err := us.userGroupRepo.GetUserGroupPage(ctx, page, pageSize, userGroup)
	if err != nil {
		return
	}

	resp := &[]schema.GetUserGroupResp{}
	_ = copier.Copy(resp, userGroups)

	return pager.NewPageModel(total, resp), nil
}
