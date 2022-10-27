package meta

import (
	"context"

	"github.com/answerdev/answer/internal/base/data"
	"github.com/answerdev/answer/internal/base/pager"
	"github.com/answerdev/answer/internal/base/reason"
	"github.com/answerdev/answer/internal/entity"
	"github.com/answerdev/answer/internal/service/meta"
	"github.com/segmentfault/pacman/errors"
	"xorm.io/builder"
)

// metaRepo meta repository
type metaRepo struct {
	data *data.Data
}

// NewMetaRepo new repository
func NewMetaRepo(data *data.Data) meta.MetaRepo {
	return &metaRepo{
		data: data,
	}
}

// AddMeta add meta
func (mr *metaRepo) AddMeta(ctx context.Context, meta *entity.Meta) (err error) {
	_, err = mr.data.DB.Insert(meta)
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}

// RemoveMeta delete meta
func (mr *metaRepo) RemoveMeta(ctx context.Context, id int) (err error) {
	_, err = mr.data.DB.ID(id).Delete(&entity.Meta{})
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}

// UpdateMeta update meta
func (mr *metaRepo) UpdateMeta(ctx context.Context, meta *entity.Meta) (err error) {
	_, err = mr.data.DB.ID(meta.ID).Update(meta)
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}

// GetMetaByObjectIdAndKey get meta one
func (mr *metaRepo) GetMetaByObjectIdAndKey(ctx context.Context, objectID, key string) (
	meta *entity.Meta, exist bool, err error) {
	meta = &entity.Meta{}
	exist, err = mr.data.DB.Where(builder.Eq{"object_id": objectID}.And(builder.Eq{"`key`": key})).Desc("created_at").Get(meta)
	if err != nil {
		return nil, false, errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}

// GetMetaList get meta list all
func (mr *metaRepo) GetMetaList(ctx context.Context, meta *entity.Meta) (metaList []*entity.Meta, err error) {
	metaList = make([]*entity.Meta, 0)
	err = mr.data.DB.Find(metaList, meta)
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}

// GetMetaPage get meta page
func (mr *metaRepo) GetMetaPage(ctx context.Context, page, pageSize int, meta *entity.Meta) (metaList []*entity.Meta, total int64, err error) {
	metaList = make([]*entity.Meta, 0)
	total, err = pager.Help(page, pageSize, metaList, meta, mr.data.DB.NewSession())
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}
