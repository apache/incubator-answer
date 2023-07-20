package service

import (
	"context"
	"github.com/answerdev/answer/internal/schema"
	"github.com/answerdev/answer/internal/service/search_common"
	"github.com/answerdev/answer/internal/service/search_parser"
	"github.com/answerdev/answer/plugin"
)

type SearchService struct {
	searchParser *search_parser.SearchParser
	searchRepo   search_common.SearchRepo
}

func NewSearchService(
	searchParser *search_parser.SearchParser,
	searchRepo search_common.SearchRepo,
) *SearchService {
	return &SearchService{
		searchParser: searchParser,
		searchRepo:   searchRepo,
	}
}

// Search search contents
func (ss *SearchService) Search(ctx context.Context, dto *schema.SearchDTO) (resp []schema.SearchResp, total int64, err error) {
	if dto.Page < 1 {
		dto.Page = 1
	}

	// search type
	cond := ss.searchParser.ParseStructure(ctx, dto)

	// check search plugin
	var s plugin.Search
	_ = plugin.CallSearch(func(search plugin.Search) error {
		s = search
		return nil
	})

	// search plugin is not found, call system search
	if s == nil {
		if cond.SearchAll() {
			resp, total, err = ss.searchRepo.SearchContents(ctx, cond.Words, cond.Tags, cond.UserID, cond.VoteAmount, dto.Page, dto.Size, dto.Order)
		} else if cond.SearchQuestion() {
			resp, total, err = ss.searchRepo.SearchQuestions(ctx, cond.Words, cond.Tags, cond.NotAccepted, cond.Views, cond.AnswerAmount, dto.Page, dto.Size, dto.Order)
		} else if cond.SearchAnswer() {
			resp, total, err = ss.searchRepo.SearchAnswers(ctx, cond.Words, cond.Tags, cond.Accepted, cond.QuestionID, dto.Page, dto.Size, dto.Order)
		}
		return
	}
	return ss.searchByPlugin(ctx, s, cond, dto)
}

func (ss *SearchService) searchByPlugin(ctx context.Context, finder plugin.Search, cond *schema.SearchCondition, dto *schema.SearchDTO) (resp []schema.SearchResp, total int64, err error) {
	var res []plugin.SearchResult
	if cond.SearchAll() {
		res, total, err = finder.SearchContents(ctx, cond.Words, cond.Tags, cond.UserID, cond.VoteAmount, dto.Page, dto.Size, dto.Order)
	} else if cond.SearchQuestion() {
		res, total, err = finder.SearchQuestions(ctx, cond.Words, cond.Tags, cond.NotAccepted, cond.Views, cond.AnswerAmount, dto.Page, dto.Size, dto.Order)
	} else if cond.SearchAnswer() {
		res, total, err = finder.SearchAnswers(ctx, cond.Words, cond.Tags, cond.Accepted, cond.QuestionID, dto.Page, dto.Size, dto.Order)
	}

	resp, err = ss.searchRepo.ParseSearchPluginResult(ctx, res)
	if err != nil {
		return nil, 0, err
	}
	return
}
