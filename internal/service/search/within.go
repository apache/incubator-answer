package search

import (
	"context"

	"github.com/segmentfault/answer/internal/schema"
	"github.com/segmentfault/answer/internal/service/search_common"
)

type WithinSearch struct {
	repo  search_common.SearchRepo
	w     string
	page  int
	size  int
	order string
}

func NewWithinSearch(repo search_common.SearchRepo) *WithinSearch {
	return &WithinSearch{
		repo: repo,
	}
}

func (s *WithinSearch) Parse(dto *schema.SearchDTO) (ok bool) {
	var (
		q      string
		w      []rune
		hasEnd bool
	)

	q = dto.Query

	if q[0:1] == `"` {
		for _, v := range []rune(q) {
			if len(w) == 0 && string(v) == `"` {
				continue
			} else if string(v) == `"` {
				hasEnd = true
				break
			} else {
				w = append(w, v)
			}
		}
	}

	if hasEnd {
		ok = true
	}

	s.w = string(w)
	s.page = dto.Page
	s.size = dto.Size
	s.order = dto.Order
	return
}

func (s *WithinSearch) Search(ctx context.Context) (resp []schema.SearchResp, total int64, err error) {
	return s.repo.SearchContents(ctx, []string{s.w}, "", "", -1, s.page, s.size, s.order)
}
