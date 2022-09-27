package service

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/segmentfault/answer/internal/base/reason"
	"github.com/segmentfault/answer/internal/entity"
	"github.com/segmentfault/answer/internal/schema"
	"github.com/segmentfault/pacman/errors"
)

// CollectionGroupRepo collectionGroup repository
type CollectionGroupRepo interface {
	AddCollectionGroup(ctx context.Context, collectionGroup *entity.CollectionGroup) (err error)
	AddCollectionDefaultGroup(ctx context.Context, userID string) (collectionGroup *entity.CollectionGroup, err error)
	UpdateCollectionGroup(ctx context.Context, collectionGroup *entity.CollectionGroup, cols []string) (err error)
	GetCollectionGroup(ctx context.Context, id string) (collectionGroup *entity.CollectionGroup, exist bool, err error)
	GetCollectionGroupPage(ctx context.Context, page, pageSize int, collectionGroup *entity.CollectionGroup) (collectionGroupList []*entity.CollectionGroup, total int64, err error)
	GetDefaultID(ctx context.Context, userId string) (collectionGroup *entity.CollectionGroup, has bool, err error)
}

// CollectionGroupService user service
type CollectionGroupService struct {
	collectionGroupRepo CollectionGroupRepo
}

func NewCollectionGroupService(collectionGroupRepo CollectionGroupRepo) *CollectionGroupService {
	return &CollectionGroupService{
		collectionGroupRepo: collectionGroupRepo,
	}
}

// AddCollectionGroup add collection group
func (cs *CollectionGroupService) AddCollectionGroup(ctx context.Context, req *schema.AddCollectionGroupReq) (err error) {
	collectionGroup := &entity.CollectionGroup{}
	_ = copier.Copy(collectionGroup, req)
	return cs.collectionGroupRepo.AddCollectionGroup(ctx, collectionGroup)
}

// UpdateCollectionGroup update collection group
func (cs *CollectionGroupService) UpdateCollectionGroup(ctx context.Context, req *schema.UpdateCollectionGroupReq, cols []string) (err error) {
	collectionGroup := &entity.CollectionGroup{}
	_ = copier.Copy(collectionGroup, req)
	return cs.collectionGroupRepo.UpdateCollectionGroup(ctx, collectionGroup, cols)
}

// GetCollectionGroup get collection group one
func (cs *CollectionGroupService) GetCollectionGroup(ctx context.Context, id string) (resp *schema.GetCollectionGroupResp, err error) {
	collectionGroup, exist, err := cs.collectionGroupRepo.GetCollectionGroup(ctx, id)
	if err != nil {
		return
	}
	if !exist {
		return nil, errors.BadRequest(reason.UnknownError)
	}

	resp = &schema.GetCollectionGroupResp{}
	_ = copier.Copy(resp, collectionGroup)
	return resp, nil
}
