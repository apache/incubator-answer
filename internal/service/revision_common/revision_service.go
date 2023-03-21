package revision_common

import (
	"context"

	"github.com/answerdev/answer/internal/base/reason"
	"github.com/answerdev/answer/internal/service/revision"
	usercommon "github.com/answerdev/answer/internal/service/user_common"
	"github.com/answerdev/answer/pkg/uid"
	"github.com/segmentfault/pacman/errors"
	"github.com/segmentfault/pacman/log"

	"github.com/answerdev/answer/internal/entity"
	"github.com/answerdev/answer/internal/schema"
	"github.com/jinzhu/copier"
)

// RevisionService user service
type RevisionService struct {
	revisionRepo revision.RevisionRepo
	userRepo     usercommon.UserRepo
}

func NewRevisionService(revisionRepo revision.RevisionRepo,
	userRepo usercommon.UserRepo,
) *RevisionService {
	return &RevisionService{
		revisionRepo: revisionRepo,
		userRepo:     userRepo,
	}
}

func (rs *RevisionService) GetUnreviewedRevisionCount(ctx context.Context, req *schema.RevisionSearch) (count int64, err error) {
	if len(req.GetCanReviewObjectTypes()) == 0 {
		return 0, nil
	}
	_, count, err = rs.revisionRepo.GetUnreviewedRevisionPage(ctx, req.Page, 1, req.GetCanReviewObjectTypes())
	return count, err
}

// AddRevision add revision
//
// autoUpdateRevisionID bool : if autoUpdateRevisionID is true , the object.revision_id will be updated,
// if not need auto update object.revision_id, it must be false.
// example: user can edit the object, but need audit, the revision_id will be updated when admin approved
func (rs *RevisionService) AddRevision(ctx context.Context, req *schema.AddRevisionDTO, autoUpdateRevisionID bool) (
	revisionID string, err error) {
	req.ObjectID = uid.DeShortID(req.ObjectID)
	rev := &entity.Revision{}
	_ = copier.Copy(rev, req)
	err = rs.revisionRepo.AddRevision(ctx, rev, autoUpdateRevisionID)
	if err != nil {
		return "", err
	}
	return rev.ID, nil
}

// GetRevision get revision
func (rs *RevisionService) GetRevision(ctx context.Context, revisionID string) (
	revision *entity.Revision, err error) {
	revisionInfo, exist, err := rs.revisionRepo.GetRevisionByID(ctx, revisionID)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if !exist {
		return nil, errors.BadRequest(reason.ObjectNotFound)
	}
	return revisionInfo, nil
}

// ExistUnreviewedByObjectID
func (rs *RevisionService) ExistUnreviewedByObjectID(ctx context.Context, objectID string) (revision *entity.Revision, exist bool, err error) {
	objectID = uid.DeShortID(objectID)
	revision, exist, err = rs.revisionRepo.ExistUnreviewedByObjectID(ctx, objectID)
	return revision, exist, err
}
