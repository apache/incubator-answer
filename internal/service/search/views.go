package search

import (
	"context"
	"github.com/segmentfault/answer/internal/schema"
	"github.com/segmentfault/answer/internal/service/search_common"
	"regexp"
	"strings"
)

type ViewsSearch struct {
	repo search_common.SearchRepo
	exp  string
	q    string
}

func NewViewsSearch(repo search_common.SearchRepo) *ViewsSearch {
	return &ViewsSearch{
		repo: repo,
	}
}

func (s *ViewsSearch) Parse(dto *schema.SearchDTO) (ok bool) {
	exp := ""
	w := dto.Query
	q := w
	p := `(?m)^views:([0-9]+)`

	re := regexp.MustCompile(p)
	res := re.FindStringSubmatch(q)
	if len(res) == 2 {
		exp = res[1]
		trimLen := len(res[0])
		q = w[trimLen:]
		ok = true
	}

	q = strings.TrimSpace(q)
	s.exp = exp
	s.q = q
	return
}
func (s *ViewsSearch) Search(ctx context.Context) (resp []schema.SearchResp, total int64, err error) {
	return
}
