package search

import (
	"context"
	"github.com/segmentfault/answer/internal/schema"
	"github.com/segmentfault/answer/internal/service/search_common"
	"strings"
)

type AnswerSearch struct {
	repo search_common.SearchRepo
	w    string
	page int
	size int
}

func NewAnswerSearch(repo search_common.SearchRepo) *AnswerSearch {
	return &AnswerSearch{
		repo: repo,
	}
}

func (s *AnswerSearch) Parse(dto *schema.SearchDTO) (ok bool) {
	var (
		q,
		w,
		p string
	)

	q = dto.Query
	w = dto.Query
	p = `is:answer`

	if strings.Index(q, p) == 0 {
		ok = true
		w = strings.TrimPrefix(q, p)
	}

	s.w = strings.TrimSpace(w)
	s.page = dto.Page
	s.size = dto.Size
	return
}
func (s *AnswerSearch) Search(ctx context.Context) (resp []schema.SearchResp, total int64, err error) {

	words := strings.Split(s.w, " ")
	if len(words) > 3 {
		words = words[:4]
	}

	return s.repo.SearchAnswers(ctx, words, false, "", s.page, s.size)
}
