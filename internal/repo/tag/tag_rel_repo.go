package tag

import (
	"context"

	"github.com/answerdev/answer/internal/base/data"
	"github.com/answerdev/answer/internal/base/reason"
	"github.com/answerdev/answer/internal/entity"
	tagcommon "github.com/answerdev/answer/internal/service/tag_common"
	"github.com/segmentfault/pacman/errors"
)

// tagListRepo tagList repository
type tagListRepo struct {
	data *data.Data
}

// NewTagListRepo new repository
func NewTagListRepo(data *data.Data) tagcommon.TagRelRepo {
	return &tagListRepo{
		data: data,
	}
}

// AddTagRelList add tag list
func (tr *tagListRepo) AddTagRelList(ctx context.Context, tagList []*entity.TagRel) (err error) {
	_, err = tr.data.DB.Insert(tagList)
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}

// RemoveTagRelListByObjectID delete tag list
func (tr *tagListRepo) RemoveTagRelListByObjectID(ctx context.Context, objectID string) (err error) {
	_, err = tr.data.DB.Where("object_id = ?", objectID).Update(&entity.TagRel{Status: entity.TagRelStatusDeleted})
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}

// RemoveTagRelListByIDs delete tag list
func (tr *tagListRepo) RemoveTagRelListByIDs(ctx context.Context, ids []int64) (err error) {
	_, err = tr.data.DB.In("id", ids).Update(&entity.TagRel{Status: entity.TagRelStatusDeleted})
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}

// GetObjectTagRelWithoutStatus get object tag relation no matter status
func (tr *tagListRepo) GetObjectTagRelWithoutStatus(ctx context.Context, objectID, tagID string) (
	tagRel *entity.TagRel, exist bool, err error,
) {
	tagRel = &entity.TagRel{}
	session := tr.data.DB.Where("object_id = ?", objectID).And("tag_id = ?", tagID)
	exist, err = session.Get(tagRel)
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}

// EnableTagRelByIDs update tag status to available
func (tr *tagListRepo) EnableTagRelByIDs(ctx context.Context, ids []int64) (err error) {
	_, err = tr.data.DB.In("id", ids).Update(&entity.TagRel{Status: entity.TagRelStatusAvailable})
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}

// GetObjectTagRelList get object tag relation list all
func (tr *tagListRepo) GetObjectTagRelList(ctx context.Context, objectID string) (tagListList []*entity.TagRel, err error) {
	tagListList = make([]*entity.TagRel, 0)
	session := tr.data.DB.Where("object_id = ?", objectID)
	session.Where("status = ?", entity.TagRelStatusAvailable)
	err = session.Find(&tagListList)
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}

// BatchGetObjectTagRelList get object tag relation list all
func (tr *tagListRepo) BatchGetObjectTagRelList(ctx context.Context, objectIds []string) (tagListList []*entity.TagRel, err error) {
	tagListList = make([]*entity.TagRel, 0)
	session := tr.data.DB.In("object_id", objectIds)
	session.Where("status = ?", entity.TagRelStatusAvailable)
	err = session.Find(&tagListList)
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}

// CountTagRelByTagID count tag relation
func (tr *tagListRepo) CountTagRelByTagID(ctx context.Context, tagID string) (count int64, err error) {
	count, err = tr.data.DB.Count(&entity.TagRel{TagID: tagID, Status: entity.AnswerStatusAvailable})
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}
