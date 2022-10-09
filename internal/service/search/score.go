package search

import (
	"context"
	"regexp"
	"strings"

	"github.com/segmentfault/answer/internal/schema"
	"github.com/segmentfault/answer/internal/service/search_common"
	"github.com/segmentfault/answer/pkg/converter"
)

type ScoreSearch struct {
	repo  search_common.SearchRepo
	exp   int
	w     string
	page  int
	size  int
	order string
}

func NewScoreSearch(repo search_common.SearchRepo) *ScoreSearch {
	return &ScoreSearch{
		repo: repo,
	}
}

func (s *ScoreSearch) Parse(dto *schema.SearchDTO) (ok bool) {
	exp := ""
	q := dto.Query
	w := q
	p := `(?m)^score:([0-9]+)`

	re := regexp.MustCompile(p)
	res := re.FindStringSubmatch(w)
	if len(res) == 2 {
		exp = res[1]
		trimLen := len(res[0])
		w = q[trimLen:]
		ok = true
	}

	w = strings.TrimSpace(w)
	s.exp = converter.StringToInt(exp)
	s.w = w
	s.page = dto.Page
	s.size = dto.Size
	s.order = dto.Order
	return
}
func (s *ScoreSearch) Search(ctx context.Context) (resp []schema.SearchResp, total int64, err error) {
	var (
		words []string
	)

	words = strings.Split(s.w, " ")
	if len(words) > 3 {
		words = words[:4]
	}

	resp, total, err = s.repo.SearchContents(ctx, words, "", "", s.exp, s.page, s.size, s.order)
	return
}
