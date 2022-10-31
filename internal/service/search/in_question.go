package search

import (
	"context"
	"github.com/answerdev/answer/internal/schema"
	"github.com/answerdev/answer/internal/service/search_common"
	"regexp"
	"strings"
)

type InQuestionSearch struct {
	repo  search_common.SearchRepo
	w     string
	exp   string
	page  int
	size  int
	order string
}

func NewInQuestionSearch(repo search_common.SearchRepo) *InQuestionSearch {
	return &InQuestionSearch{
		repo: repo,
	}
}

func (s *InQuestionSearch) Parse(dto *schema.SearchDTO) (ok bool) {
	var (
		w,
		q,
		p,
		exp string
	)

	q = dto.Query
	w = dto.Query
	p = `(?m)^inquestion:([0-9]+)`

	re := regexp.MustCompile(p)
	res := re.FindStringSubmatch(q)
	if len(res) == 2 {
		exp = res[1]
		trimLen := len(res[0])
		w = q[trimLen:]
		ok = true
	}

	s.exp = exp
	s.w = strings.TrimSpace(w)
	s.page = dto.Page
	s.size = dto.Size
	s.order = dto.Order
	return
}
func (s *InQuestionSearch) Search(ctx context.Context) (resp []schema.SearchResp, total int64, err error) {
	var (
		words []string
	)

	words = strings.Split(s.w, " ")
	if len(words) > 3 {
		words = words[:4]
	}

	return s.repo.SearchAnswers(ctx, words, false, s.exp, s.page, s.size, s.order)
}
