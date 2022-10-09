package search

import (
	"context"
	"github.com/segmentfault/answer/internal/schema"
	"github.com/segmentfault/answer/internal/service/search_common"
	"strings"
)

type QuestionSearch struct {
	repo  search_common.SearchRepo
	w     string
	page  int
	size  int
	order string
}

func NewQuestionSearch(repo search_common.SearchRepo) *QuestionSearch {
	return &QuestionSearch{
		repo: repo,
	}
}

func (s *QuestionSearch) Parse(dto *schema.SearchDTO) (ok bool) {
	var (
		q,
		w,
		p string
	)

	q = dto.Query
	w = dto.Query
	p = `is:question`

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

func (s *QuestionSearch) Search(ctx context.Context) (resp []schema.SearchResp, total int64, err error) {

	words := strings.Split(s.w, " ")
	if len(words) > 3 {
		words = words[:4]
	}

	return s.repo.SearchQuestions(ctx, words, false, -1, s.page, s.size, s.order)
}
