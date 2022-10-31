package search

import (
	"context"
	"strings"

	"github.com/answerdev/answer/internal/schema"
	"github.com/answerdev/answer/internal/service/search_common"
)

type NotAcceptedQuestion struct {
	repo  search_common.SearchRepo
	w     string
	page  int
	size  int
	order string
}

func NewNotAcceptedQuestion(repo search_common.SearchRepo) *NotAcceptedQuestion {
	return &NotAcceptedQuestion{
		repo: repo,
	}
}

func (s *NotAcceptedQuestion) Parse(dto *schema.SearchDTO) (ok bool) {
	var (
		q,
		w,
		p string
	)

	q = dto.Query
	w = dto.Query
	p = `hasaccepted:no`

	if strings.Index(q, p) == 0 {
		ok = true
		w = strings.TrimPrefix(q, p)
	}

	s.w = strings.TrimSpace(w)
	s.page = dto.Page
	s.size = dto.Size
	s.order = dto.Order
	return
}
func (s *NotAcceptedQuestion) Search(ctx context.Context) (resp []schema.SearchResp, total int64, err error) {
	var (
		words []string
	)

	words = strings.Split(s.w, " ")
	if len(words) > 3 {
		words = words[:4]
	}

	return s.repo.SearchQuestions(ctx, words, true, -1, s.page, s.size, s.order)
}
