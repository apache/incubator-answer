package search

import (
	"context"
	"strings"

	"github.com/answerdev/answer/internal/schema"
	"github.com/answerdev/answer/internal/service/search_common"
)

type ObjectSearch struct {
	repo  search_common.SearchRepo
	w     string
	page  int
	size  int
	order string
}

func NewObjectSearch(repo search_common.SearchRepo) *ObjectSearch {
	return &ObjectSearch{
		repo: repo,
	}
}

func (s *ObjectSearch) Parse(dto *schema.SearchDTO) (ok bool) {
	var (
		w string
	)
	w = strings.TrimSpace(dto.Query)
	if len(w) > 0 {
		ok = true
	}

	s.w = w
	s.page = dto.Page
	s.size = dto.Size
	s.order = dto.Order
	return
}
func (s *ObjectSearch) Search(ctx context.Context) (resp []schema.SearchResp, total int64, err error) {

	words := strings.Split(s.w, " ")
	if len(words) > 3 {
		words = words[:4]
	}
	return s.repo.SearchContents(ctx, words, "", "", -1, s.page, s.size, s.order)
}
