package collectioncommon

import (
	"context"

	"github.com/segmentfault/answer/internal/entity"
)

// CollectionRepo collection repository
type CollectionRepo interface {
	AddCollection(ctx context.Context, collection *entity.Collection) (err error)
	RemoveCollection(ctx context.Context, id string) (err error)
	UpdateCollection(ctx context.Context, collection *entity.Collection, cols []string) (err error)
	GetCollection(ctx context.Context, id int) (collection *entity.Collection, exist bool, err error)
	GetCollectionList(ctx context.Context, collection *entity.Collection) (collectionList []*entity.Collection, err error)
	GetOneByObjectIDAndUser(ctx context.Context, userId string, objectId string) (collection *entity.Collection, exist bool, err error)
	SearchByObjectIDsAndUser(ctx context.Context, userId string, objectIds []string) (collectionList []*entity.Collection, err error)
	CountByObjectID(ctx context.Context, objectId string) (total int64, err error)
	GetCollectionPage(ctx context.Context, page, pageSize int, collection *entity.Collection) (collectionList []*entity.Collection, total int64, err error)
	SearchObjectCollected(ctx context.Context, userId string, objectIds []string) (collectedMap map[string]bool, err error)
	SearchList(ctx context.Context, search *entity.CollectionSearch) ([]*entity.Collection, int64, error)
}

// CollectionService user service
type CollectionCommon struct {
	collectionRepo CollectionRepo
}

func NewCollectionCommon(collectionRepo CollectionRepo) *CollectionCommon {
	return &CollectionCommon{
		collectionRepo: collectionRepo,
	}
}

// SearchObjectCollected search object is collected
func (ccs *CollectionCommon) SearchObjectCollected(ctx context.Context, userId string, objectIds []string) (collectedMap map[string]bool, err error) {
	return ccs.collectionRepo.SearchObjectCollected(ctx, userId, objectIds)
}

func (ccs *CollectionCommon) SearchList(ctx context.Context, search *entity.CollectionSearch) ([]*entity.Collection, int64, error) {
	return ccs.collectionRepo.SearchList(ctx, search)
}
