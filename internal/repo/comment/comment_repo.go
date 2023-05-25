package comment

import (
	"context"

	"github.com/answerdev/answer/internal/base/data"
	"github.com/answerdev/answer/internal/base/pager"
	"github.com/answerdev/answer/internal/base/reason"
	"github.com/answerdev/answer/internal/entity"
	"github.com/answerdev/answer/internal/service/comment"
	"github.com/answerdev/answer/internal/service/comment_common"
	"github.com/answerdev/answer/internal/service/unique"
	"github.com/segmentfault/pacman/errors"
)

// commentRepo comment repository
type commentRepo struct {
	data         *data.Data
	uniqueIDRepo unique.UniqueIDRepo
}

// NewCommentRepo new repository
func NewCommentRepo(data *data.Data, uniqueIDRepo unique.UniqueIDRepo) comment.CommentRepo {
	return &commentRepo{
		data:         data,
		uniqueIDRepo: uniqueIDRepo,
	}
}

// NewCommentCommonRepo new repository
func NewCommentCommonRepo(data *data.Data, uniqueIDRepo unique.UniqueIDRepo) comment_common.CommentCommonRepo {
	return &commentRepo{
		data:         data,
		uniqueIDRepo: uniqueIDRepo,
	}
}

// AddComment add comment
func (cr *commentRepo) AddComment(ctx context.Context, comment *entity.Comment) (err error) {
	comment.ID, err = cr.uniqueIDRepo.GenUniqueIDStr(ctx, comment.TableName())
	if err != nil {
		return err
	}
	_, err = cr.data.DB.Context(ctx).Insert(comment)
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}

// RemoveComment delete comment
func (cr *commentRepo) RemoveComment(ctx context.Context, commentID string) (err error) {
	session := cr.data.DB.Context(ctx).ID(commentID)
	_, err = session.Update(&entity.Comment{Status: entity.CommentStatusDeleted})
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}

// UpdateComment update comment
func (cr *commentRepo) UpdateComment(ctx context.Context, comment *entity.Comment) (err error) {
	_, err = cr.data.DB.Context(ctx).ID(comment.ID).Where("user_id = ?", comment.UserID).Update(comment)
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}

// GetComment get comment one
func (cr *commentRepo) GetComment(ctx context.Context, commentID string) (
	comment *entity.Comment, exist bool, err error,
) {
	comment = &entity.Comment{}
	exist, err = cr.data.DB.Context(ctx).ID(commentID).Get(comment)
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}

func (cr *commentRepo) GetCommentCount(ctx context.Context) (count int64, err error) {
	list := make([]*entity.Comment, 0)
	count, err = cr.data.DB.Context(ctx).Where("status = ?", entity.CommentStatusAvailable).FindAndCount(&list)
	if err != nil {
		return count, errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}

// GetCommentPage get comment page
func (cr *commentRepo) GetCommentPage(ctx context.Context, commentQuery *comment.CommentQuery) (
	commentList []*entity.Comment, total int64, err error,
) {
	commentList = make([]*entity.Comment, 0)

	session := cr.data.DB.Context(ctx)
	session.OrderBy(commentQuery.GetOrderBy())
	session.Where("status = ?", entity.CommentStatusAvailable)

	cond := &entity.Comment{ObjectID: commentQuery.ObjectID, UserID: commentQuery.UserID}
	total, err = pager.Help(commentQuery.Page, commentQuery.PageSize, &commentList, cond, session)
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}
