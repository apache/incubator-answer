package search

import (
	"context"
	"github.com/segmentfault/answer/internal/entity"
	"github.com/segmentfault/answer/internal/schema"
	"github.com/segmentfault/answer/internal/service/activity_common"
	"github.com/segmentfault/answer/internal/service/search_common"
	tagcommon "github.com/segmentfault/answer/internal/service/tag_common"
	"regexp"
	"strings"
)

type TagSearch struct {
	repo         search_common.SearchRepo
	tagRepo      tagcommon.TagRepo
	followCommon activity_common.FollowRepo
	page         int
	size         int
	exp          string
	w            string
	userID       string
	Extra        schema.GetTagPageResp
}

func NewTagSearch(repo search_common.SearchRepo, tagRepo tagcommon.TagRepo, followCommon activity_common.FollowRepo) *TagSearch {
	return &TagSearch{
		repo:         repo,
		tagRepo:      tagRepo,
		followCommon: followCommon,
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
	return ok
}

func (ts *TagSearch) Search(ctx context.Context) (resp []schema.SearchResp, total int64, err error) {
	var (
		words            []string
		tag              *entity.Tag
		exists, followed bool
	)
	tag, exists, err = ts.tagRepo.GetTagBySlugName(nil, ts.exp)
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
	}
	ts.Extra.GetExcerpt()

	if !exists {
		return
	}
	words = strings.Split(ts.w, " ")
	if len(words) > 3 {
		words = words[:4]
	}

	resp, total, err = ts.repo.SearchContents(ctx, words, tag.ID, "", -1, ts.page, ts.size)

	return
}
