package follow

import (
	"context"

	"github.com/segmentfault/answer/internal/entity"
	"github.com/segmentfault/answer/internal/schema"
	"github.com/segmentfault/answer/internal/service/activity_common"
	tagcommon "github.com/segmentfault/answer/internal/service/tag_common"
)

type FollowRepo interface {
	Follow(ctx context.Context, objectId, userId string) error
	FollowCancel(ctx context.Context, objectId, userId string) error
}

type FollowService struct {
	tagRepo          tagcommon.TagRepo
	followRepo       FollowRepo
	followCommonRepo activity_common.FollowRepo
}

func NewFollowService(
	followRepo FollowRepo,
	followCommonRepo activity_common.FollowRepo,
	tagRepo tagcommon.TagRepo,
) *FollowService {
	return &FollowService{
		followRepo:       followRepo,
		followCommonRepo: followCommonRepo,
		tagRepo:          tagRepo,
	}
}

// Follow or cancel follow object
func (fs *FollowService) Follow(ctx context.Context, dto *schema.FollowDTO) (resp schema.FollowResp, err error) {
	if dto.IsCancel {
		err = fs.followRepo.FollowCancel(ctx, dto.ObjectID, dto.UserID)
	} else {
		err = fs.followRepo.Follow(ctx, dto.ObjectID, dto.UserID)
	}
	if err != nil {
		return resp, err
	}
	follows, err := fs.followCommonRepo.GetFollowAmount(ctx, dto.ObjectID)
	if err != nil {
		return resp, err
	}

	resp.Follows = follows
	resp.IsFollowed = !dto.IsCancel
	return resp, nil
}

// UpdateFollowTags update user follow tags
func (fs *FollowService) UpdateFollowTags(ctx context.Context, req *schema.UpdateFollowTagsReq) (err error) {
	objIDs, err := fs.followCommonRepo.GetFollowIDs(ctx, req.UserID, entity.Tag{}.TableName())
	if err != nil {
		return
	}
	oldFollowTagList, err := fs.tagRepo.GetTagListByIDs(ctx, objIDs)
	if err != nil {
		return err
	}
	oldTagMapping := make(map[string]bool)
	for _, tag := range oldFollowTagList {
		oldTagMapping[tag.SlugName] = true
	}

	newTagMapping := make(map[string]bool)
	for _, tag := range req.SlugNameList {
		newTagMapping[tag] = true
	}

	// cancel follow
	for _, tag := range oldFollowTagList {
		if !newTagMapping[tag.SlugName] {
			err := fs.followRepo.FollowCancel(ctx, tag.ID, req.UserID)
			if err != nil {
				return err
			}
		}
	}

	// new follow
	for _, tagSlugName := range req.SlugNameList {
		if !oldTagMapping[tagSlugName] {
			tagInfo, exist, err := fs.tagRepo.GetTagBySlugName(ctx, tagSlugName)
			if err != nil {
				return err
			}
			if !exist {
				continue
			}
			err = fs.followRepo.Follow(ctx, tagInfo.ID, req.UserID)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
