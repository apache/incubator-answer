package search

import (
	"context"
	"regexp"
	"strings"

	"github.com/segmentfault/answer/internal/entity"
	"github.com/segmentfault/answer/internal/schema"
	"github.com/segmentfault/answer/internal/service/search_common"
	usercommon "github.com/segmentfault/answer/internal/service/user_common"
)

type AuthorSearch struct {
	repo     search_common.SearchRepo
	userRepo usercommon.UserRepo
	exp      string
	w        string
	page     int
	size     int
}

func NewAuthorSearch(repo search_common.SearchRepo, userRepo usercommon.UserRepo) *AuthorSearch {
	return &AuthorSearch{
		repo:     repo,
		userRepo: userRepo,
	}
}

// Parse
// example: "user:12345" -> {exp="" w="12345"}
func (s *AuthorSearch) Parse(dto *schema.SearchDTO) (ok bool) {
	var (
		exp,
		q,
		w,
		p,
		me,
		name string
		user *entity.User
		has  bool
		err  error
	)
	exp = ""
	q = dto.Query
	w = q
	p = `(?m)^user:([a-z0-9._-]+)`
	me = "user:me"

	re := regexp.MustCompile(p)
	res := re.FindStringSubmatch(q)
	if len(res) == 2 {
		name = res[1]
		user, has, err = s.userRepo.GetByUsername(nil, name)

		if err == nil && has {
			exp = user.ID
			trimLen := len(res[0])
			w = q[trimLen:]
			ok = true
		}
	} else if strings.Index(q, me) == 0 {
		exp = dto.UserID
		w = strings.TrimPrefix(q, me)
		ok = true
	}

	w = strings.TrimSpace(w)
	s.exp = exp
	s.w = w
	s.page = dto.Page
	s.size = dto.Size
	return
}

func (s *AuthorSearch) Search(ctx context.Context) (resp []schema.SearchResp, total int64, err error) {
	var (
		words []string
	)

	if len(s.exp) == 0 {
		return
	}

	words = strings.Split(s.w, " ")
	if len(words) > 3 {
		words = words[:4]
	}

	resp, total, err = s.repo.SearchContents(ctx, words, "", s.exp, -1, s.page, s.size)

	return
}
