package search_common

import (
	"context"
	"github.com/answerdev/answer/internal/schema"
)

type SearchRepo interface {
	SearchContents(ctx context.Context, words []string, tagIDs []string, userID string, votes, page, size int, order string) (resp []schema.SearchResp, total int64, err error)
	SearchQuestions(ctx context.Context, words []string, notAccepted bool, views, answers int, page, size int, order string) (resp []schema.SearchResp, total int64, err error)
	SearchAnswers(ctx context.Context, words []string, tagIDs []string, accepted bool, questionID string, page, size int, order string) (resp []schema.SearchResp, total int64, err error)
}
