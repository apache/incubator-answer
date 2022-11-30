package search

import (
	"context"
	"regexp"
	"strings"

	"github.com/answerdev/answer/internal/entity"
	"github.com/answerdev/answer/internal/schema"
	"github.com/answerdev/answer/internal/service/activity_common"
	"github.com/answerdev/answer/internal/service/search_common"
	"github.com/answerdev/answer/internal/service/tag_common"
)

type TagSearch struct {
	repo             search_common.SearchRepo
	tagCommonService *tag_common.TagCommonService
	followCommon     activity_common.FollowRepo
	page             int
	size             int
	exp              string
	w                string
	userID           string
	Extra            schema.GetTagPageResp
	order            string
}

func NewTagSearch(repo search_common.SearchRepo,
	tagCommonService *tag_common.TagCommonService, followCommon activity_common.FollowRepo) *TagSearch {
	return &TagSearch{
		repo:             repo,
		tagCommonService: tagCommonService,
		followCommon:     followCommon,
	}
}

// Parse
// example: "[tag]hello" -> {exp="tag" w="hello"}
func (ts *TagSearch) Parse(dto *schema.SearchDTO) (ok bool) {
	exp := ""
	w := dto.Query
	q := w
	p := `(?m)^\[([a-zA-Z0-9-\+\.#]+)\]`

	re := regexp.MustCompile(p)
	res := re.FindStringSubmatch(q)
	if len(res) == 2 {
		exp = res[1]
		trimLen := len(res[0])
		w = q[trimLen:]
		ok = true
	}
	w = strings.TrimSpace(w)
	ts.exp = exp
	ts.w = w
	ts.page = dto.Page
	ts.size = dto.Size
	ts.userID = dto.UserID
	ts.order = dto.Order
	return ok
}

func (ts *TagSearch) Search(ctx context.Context) (resp []schema.SearchResp, total int64, err error) {
	var (
		words            []string
		tag              *entity.Tag
		exists, followed bool
	)
	tag, exists, err = ts.tagCommonService.GetTagBySlugName(ctx, ts.exp)
	if err != nil {
		return
	}

	if ts.userID != "" {
		followed, err = ts.followCommon.IsFollowed(ts.userID, tag.ID)
	}

	ts.Extra = schema.GetTagPageResp{
		TagID:         tag.ID,
		SlugName:      tag.SlugName,
		DisplayName:   tag.DisplayName,
		OriginalText:  tag.OriginalText,
		ParsedText:    tag.ParsedText,
		QuestionCount: tag.QuestionCount,
		IsFollower:    followed,
		Recommend:     tag.Recommend,
		Reserved:      tag.Reserved,
	}
	ts.Extra.GetExcerpt()

	if !exists {
		return
	}
	words = strings.Split(ts.w, " ")
	if len(words) > 3 {
		words = words[:4]
	}

	resp, total, err = ts.repo.SearchContents(ctx, words, tag.ID, "", -1, ts.page, ts.size, ts.order)

	return
}
