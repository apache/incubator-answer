package search

import (
	"context"
	"strings"

	"github.com/segmentfault/answer/internal/schema"
	"github.com/segmentfault/answer/internal/service/search_common"
)

type AcceptedAnswerSearch struct {
	repo  search_common.SearchRepo
	w     string
	page  int
	size  int
	order string
}

func NewAcceptedAnswerSearch(repo search_common.SearchRepo) *AcceptedAnswerSearch {
	return &AcceptedAnswerSearch{
		repo: repo,
	}
}

func (s *AcceptedAnswerSearch) Parse(dto *schema.SearchDTO) (ok bool) {
	var (
		q,
		w,
		p string
	)

	q = dto.Query
	w = dto.Query
	p = `isaccepted:yes`

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
func (s *AcceptedAnswerSearch) Search(ctx context.Context) (resp []schema.SearchResp, total int64, err error) {

	words := strings.Split(s.w, " ")
	if len(words) > 3 {
		words = words[:4]
	}

	return s.repo.SearchAnswers(ctx, words, true, "", s.page, s.size, s.order)
}
