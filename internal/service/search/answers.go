package search

import (
	"context"
	"regexp"
	"strings"

	"github.com/answerdev/answer/internal/schema"
	"github.com/answerdev/answer/internal/service/search_common"
	"github.com/answerdev/answer/pkg/converter"
)

type AnswersSearch struct {
	repo  search_common.SearchRepo
	exp   int
	w     string
	page  int
	size  int
	order string
}

func NewAnswersSearch(repo search_common.SearchRepo) *AnswersSearch {
	return &AnswersSearch{
		repo: repo,
	}
}

func (s *AnswersSearch) Parse(dto *schema.SearchDTO) (ok bool) {
	var (
		q,
		w,
		p,
		exp string
	)

	q = dto.Query
	w = dto.Query
	p = `(?m)^answers:([0-9]+)`

	re := regexp.MustCompile(p)
	res := re.FindStringSubmatch(q)
	if len(res) == 2 {
		exp = res[1]
		trimLen := len(res[0])
		w = q[trimLen:]
		ok = true
	}

	s.exp = converter.StringToInt(exp)
	s.w = strings.TrimSpace(w)
	s.page = dto.Page
	s.size = dto.Size
	s.order = dto.Order
	return
}

func (s *AnswersSearch) Search(ctx context.Context) (resp []schema.SearchResp, total int64, err error) {

	words := strings.Split(s.w, " ")
	if len(words) > 3 {
		words = words[:4]
	}

	return s.repo.SearchQuestions(ctx, words, false, s.exp, s.page, s.size, s.order)
}
