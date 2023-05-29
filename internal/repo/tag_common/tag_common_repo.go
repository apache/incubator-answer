package tag_common

import (
	"context"
	"fmt"

	"github.com/answerdev/answer/internal/base/data"
	"github.com/answerdev/answer/internal/base/pager"
	"github.com/answerdev/answer/internal/base/reason"
	"github.com/answerdev/answer/internal/entity"
	tagcommon "github.com/answerdev/answer/internal/service/tag_common"
	"github.com/answerdev/answer/internal/service/unique"
	"github.com/segmentfault/pacman/errors"
	"xorm.io/builder"
)

// tagCommonRepo tag repository
type tagCommonRepo struct {
	data         *data.Data
	uniqueIDRepo unique.UniqueIDRepo
}

// NewTagCommonRepo new repository
func NewTagCommonRepo(
	data *data.Data,
	uniqueIDRepo unique.UniqueIDRepo,
) tagcommon.TagCommonRepo {
	return &tagCommonRepo{
		data:         data,
		uniqueIDRepo: uniqueIDRepo,
	}
}

// GetTagListByIDs get tag list all
func (tr *tagCommonRepo) GetTagListByIDs(ctx context.Context, ids []string) (tagList []*entity.Tag, err error) {
	tagList = make([]*entity.Tag, 0)
	session := tr.data.DB.Context(ctx).In("id", ids)
	session.Where(builder.Eq{"status": entity.TagStatusAvailable})
	err = session.OrderBy("recommend desc,reserved desc,id desc").Find(&tagList)
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}

// GetTagBySlugName get tag by slug name
func (tr *tagCommonRepo) GetTagBySlugName(ctx context.Context, slugName string) (tagInfo *entity.Tag, exist bool, err error) {
	tagInfo = &entity.Tag{}
	session := tr.data.DB.Context(ctx).Where("LOWER(slug_name) = ?", slugName)
	session.Where(builder.Eq{"status": entity.TagStatusAvailable})
	exist, err = session.Get(tagInfo)
	if err != nil {
		return nil, false, errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}

// GetTagListByName get tag list all like name
func (tr *tagCommonRepo) GetTagListByName(ctx context.Context, name string, hasReserved bool) (tagList []*entity.Tag, err error) {
	tagList = make([]*entity.Tag, 0)
	cond := &entity.Tag{}
	session := tr.data.DB.Context(ctx).Where("")
	if name != "" {
		session.Where("slug_name LIKE LOWER(?) or display_name LIKE ?", name+"%", name+"%")
	} else {
		session.UseBool("recommend")
		cond.Recommend = true
	}
	session.Where(builder.Eq{"status": entity.TagStatusAvailable})
	session.Asc("slug_name")
	err = session.OrderBy("recommend desc,reserved desc,id desc").Find(&tagList, cond)
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}

func (tr *tagCommonRepo) GetRecommendTagList(ctx context.Context) (tagList []*entity.Tag, err error) {
	tagList = make([]*entity.Tag, 0)
	cond := &entity.Tag{}
	session := tr.data.DB.Context(ctx).Where("")
	cond.Recommend = true
	// session.Where(builder.Eq{"status": entity.TagStatusAvailable})
	session.Asc("slug_name")
	session.UseBool("recommend")
	err = session.Find(&tagList, cond)
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}

func (tr *tagCommonRepo) GetReservedTagList(ctx context.Context) (tagList []*entity.Tag, err error) {
	tagList = make([]*entity.Tag, 0)
	cond := &entity.Tag{}
	session := tr.data.DB.Context(ctx).Where("")
	cond.Reserved = true
	// session.Where(builder.Eq{"status": entity.TagStatusAvailable})
	session.Asc("slug_name")
	session.UseBool("reserved")
	err = session.Find(&tagList, cond)
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}

// GetTagListByNames get tag list all like name
func (tr *tagCommonRepo) GetTagListByNames(ctx context.Context, names []string) (tagList []*entity.Tag, err error) {

	tagList = make([]*entity.Tag, 0)
	session := tr.data.DB.Context(ctx).In("slug_name", names).UseBool("recommend", "reserved")
	// session.Where(builder.Eq{"status": entity.TagStatusAvailable})
	err = session.OrderBy("recommend desc,reserved desc,id desc").Find(&tagList)
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}

// GetTagByID get tag one
func (tr *tagCommonRepo) GetTagByID(ctx context.Context, tagID string, includeDeleted bool) (
	tag *entity.Tag, exist bool, err error,
) {
	tag = &entity.Tag{}
	session := tr.data.DB.Context(ctx).Where(builder.Eq{"id": tagID})
	if !includeDeleted {
		session.Where(builder.Eq{"status": entity.TagStatusAvailable})
	}
	exist, err = session.Get(tag)
	if err != nil {
		return nil, false, errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}

// GetTagPage get tag page
func (tr *tagCommonRepo) GetTagPage(ctx context.Context, page, pageSize int, tag *entity.Tag, queryCond string) (
	tagList []*entity.Tag, total int64, err error,
) {
	tagList = make([]*entity.Tag, 0)
	session := tr.data.DB.Context(ctx)

	if len(tag.SlugName) > 0 {
		session.Where(builder.Or(builder.Like{"slug_name", fmt.Sprintf("LOWER(%s)", tag.SlugName)}, builder.Like{"display_name", tag.SlugName}))
		tag.SlugName = ""
	}
	session.Where(builder.Eq{"status": entity.TagStatusAvailable})
	session.Where("main_tag_id = 0") // if this tag is synonym, exclude it

	switch queryCond {
	case "popular":
		session.Desc("question_count")
	case "name":
		session.Asc("slug_name")
	case "newest":
		session.Desc("created_at")
	}

	total, err = pager.Help(page, pageSize, &tagList, tag, session)
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}

// AddTagList add tag
func (tr *tagCommonRepo) AddTagList(ctx context.Context, tagList []*entity.Tag) (err error) {
	for _, item := range tagList {
		item.ID, err = tr.uniqueIDRepo.GenUniqueIDStr(ctx, item.TableName())
		if err != nil {
			return err
		}
		item.RevisionID = "0"
	}
	_, err = tr.data.DB.Context(ctx).Insert(tagList)
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}

// UpdateTagQuestionCount update tag question count
func (tr *tagCommonRepo) UpdateTagQuestionCount(ctx context.Context, tagID string, questionCount int) (err error) {
	cond := &entity.Tag{QuestionCount: questionCount}
	_, err = tr.data.DB.Context(ctx).Where(builder.Eq{"id": tagID}).MustCols("question_count").Update(cond)
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}

func (tr *tagCommonRepo) UpdateTagsAttribute(ctx context.Context, tags []string, attribute string, value bool) (err error) {
	bean := &entity.Tag{}
	switch attribute {
	case "recommend":
		bean.Recommend = value
	case "reserved":
		bean.Reserved = value
	default:
		return
	}
	session := tr.data.DB.Context(ctx).In("slug_name", tags).Cols(attribute).UseBool(attribute)
	_, err = session.Update(bean)
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}
