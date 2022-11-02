package comment_common

import (
	"context"

	"github.com/answerdev/answer/internal/base/reason"
	"github.com/answerdev/answer/internal/entity"
	"github.com/answerdev/answer/internal/schema"
	"github.com/segmentfault/pacman/errors"
)

// CommentCommonRepo comment repository
type CommentCommonRepo interface {
	GetComment(ctx context.Context, commentID string) (comment *entity.Comment, exist bool, err error)
	GetCommentCount(ctx context.Context) (count int64, err error)
}

// CommentCommonService user service
type CommentCommonService struct {
	commentRepo CommentCommonRepo
}

// NewCommentCommonService new comment service
func NewCommentCommonService(
	commentRepo CommentCommonRepo) *CommentCommonService {
	return &CommentCommonService{
		commentRepo: commentRepo,
	}
}

// GetComment get comment one
func (cs *CommentCommonService) GetComment(ctx context.Context, commentID string) (resp *schema.GetCommentResp, err error) {
	comment, exist, err := cs.commentRepo.GetComment(ctx, commentID)
	if err != nil {
		return
	}
	if !exist {
		return nil, errors.BadRequest(reason.UnknownError)
	}

	resp = &schema.GetCommentResp{}
	resp.SetFromComment(comment)
	return resp, nil
}
