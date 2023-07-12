package tag

import (
	"context"

	"github.com/answerdev/answer/internal/base/data"
	"github.com/answerdev/answer/internal/base/reason"
	"github.com/answerdev/answer/internal/entity"
	"github.com/answerdev/answer/internal/service/tag_common"
	"github.com/answerdev/answer/internal/service/unique"
	"github.com/answerdev/answer/pkg/converter"
	"github.com/segmentfault/pacman/errors"
	"xorm.io/builder"
)

// tagRepo tag repository
type tagRepo struct {
	data         *data.Data
	uniqueIDRepo unique.UniqueIDRepo
}

// NewTagRepo new repository
func NewTagRepo(
	data *data.Data,
	uniqueIDRepo unique.UniqueIDRepo,
) tag_common.TagRepo {
	return &tagRepo{
		data:         data,
		uniqueIDRepo: uniqueIDRepo,
	}
}

// RemoveTag delete tag
func (tr *tagRepo) RemoveTag(ctx context.Context, tagID string) (err error) {
	session := tr.data.DB.Context(ctx).Where(builder.Eq{"id": tagID})
	_, err = session.Update(&entity.Tag{Status: entity.TagStatusDeleted})
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}

// UpdateTag update tag
func (tr *tagRepo) UpdateTag(ctx context.Context, tag *entity.Tag) (err error) {
	_, err = tr.data.DB.Context(ctx).Where(builder.Eq{"id": tag.ID}).Update(tag)
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}

// UpdateTagSynonym update synonym tag
func (tr *tagRepo) UpdateTagSynonym(ctx context.Context, tagSlugNameList []string, mainTagID int64,
	mainTagSlugName string,
) (err error) {
	bean := &entity.Tag{MainTagID: mainTagID, MainTagSlugName: mainTagSlugName}
	session := tr.data.DB.Context(ctx).In("slug_name", tagSlugNameList).MustCols("main_tag_id", "main_tag_slug_name")
	_, err = session.Update(bean)
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}

func (tr *tagRepo) GetTagSynonymCount(ctx context.Context, tagID string) (count int64, err error) {
	count, err = tr.data.DB.Context(ctx).Count(&entity.Tag{MainTagID: converter.StringToInt64(tagID), Status: entity.TagStatusAvailable})
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}

// GetTagList get tag list all
func (tr *tagRepo) GetTagList(ctx context.Context, tag *entity.Tag) (tagList []*entity.Tag, err error) {
	tagList = make([]*entity.Tag, 0)
	session := tr.data.DB.Context(ctx).Where(builder.Eq{"status": entity.TagStatusAvailable})
	err = session.Find(&tagList, tag)
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}
