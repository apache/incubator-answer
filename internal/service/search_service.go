package service

import (
	"context"

	"github.com/segmentfault/answer/internal/schema"
	"github.com/segmentfault/answer/internal/service/activity_common"
	"github.com/segmentfault/answer/internal/service/search"
	"github.com/segmentfault/answer/internal/service/search_common"
	tagcommon "github.com/segmentfault/answer/internal/service/tag_common"
	usercommon "github.com/segmentfault/answer/internal/service/user_common"
)

type Search interface {
	Parse(dto *schema.SearchDTO) (ok bool)
	Search(ctx context.Context) (resp []schema.SearchResp, total int64, err error)
}

type SearchService struct {
	searchRepo           search_common.SearchRepo
	tagSearch            *search.TagSearch
	withinSearch         *search.WithinSearch
	authorSearch         *search.AuthorSearch
	scoreSearch          *search.ScoreSearch
	answersSearch        *search.AnswersSearch
	notAcceptedQuestion  *search.NotAcceptedQuestion
	acceptedAnswerSearch *search.AcceptedAnswerSearch
	inQuestionSearch     *search.InQuestionSearch
	questionSearch       *search.QuestionSearch
	answerSearch         *search.AnswerSearch
	viewsSearch          *search.ViewsSearch
	objectSearch         *search.ObjectSearch
}

func NewSearchService(
	searchRepo search_common.SearchRepo,
	tagRepo tagcommon.TagRepo,
	userRepo usercommon.UserRepo,
	followCommon activity_common.FollowRepo,
) *SearchService {
	return &SearchService{
		searchRepo:           searchRepo,
		tagSearch:            search.NewTagSearch(searchRepo, tagRepo, followCommon),
		withinSearch:         search.NewWithinSearch(searchRepo),
		authorSearch:         search.NewAuthorSearch(searchRepo, userRepo),
		scoreSearch:          search.NewScoreSearch(searchRepo),
		answersSearch:        search.NewAnswersSearch(searchRepo),
		acceptedAnswerSearch: search.NewAcceptedAnswerSearch(searchRepo),
		notAcceptedQuestion:  search.NewNotAcceptedQuestion(searchRepo),
		inQuestionSearch:     search.NewInQuestionSearch(searchRepo),
		questionSearch:       search.NewQuestionSearch(searchRepo),
		answerSearch:         search.NewAnswerSearch(searchRepo),
		viewsSearch:          search.NewViewsSearch(searchRepo),
		objectSearch:         search.NewObjectSearch(searchRepo),
	}
}

func (ss *SearchService) Search(ctx context.Context, dto *schema.SearchDTO) (resp []schema.SearchResp, total int64, extra interface{}, err error) {
	extra = nil
	switch {
	case ss.tagSearch.Parse(dto):
		resp, total, err = ss.tagSearch.Search(ctx)
		extra = ss.tagSearch.Extra
	case ss.withinSearch.Parse(dto):
		resp, total, err = ss.withinSearch.Search(ctx)
	case ss.authorSearch.Parse(dto):
		resp, total, err = ss.authorSearch.Search(ctx)
	case ss.scoreSearch.Parse(dto):
		resp, total, err = ss.scoreSearch.Search(ctx)
	case ss.answersSearch.Parse(dto):
		resp, total, err = ss.answersSearch.Search(ctx)
	case ss.acceptedAnswerSearch.Parse(dto):
		resp, total, err = ss.acceptedAnswerSearch.Search(ctx)
	case ss.notAcceptedQuestion.Parse(dto):
		resp, total, err = ss.notAcceptedQuestion.Search(ctx)
	case ss.inQuestionSearch.Parse(dto):
		resp, total, err = ss.inQuestionSearch.Search(ctx)
	case ss.questionSearch.Parse(dto):
		resp, total, err = ss.questionSearch.Search(ctx)
	case ss.answerSearch.Parse(dto):
		resp, total, err = ss.answerSearch.Search(ctx)
	case ss.viewsSearch.Parse(dto):
		resp, total, err = ss.viewsSearch.Search(ctx)
	default:
		ss.objectSearch.Parse(dto)
		resp, total, err = ss.objectSearch.Search(ctx)
	}
	return resp, total, extra, nil
}
