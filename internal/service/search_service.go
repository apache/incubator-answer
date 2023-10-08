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
func (ss *SearchService) Search(ctx context.Context, dto *schema.SearchDTO) (resp *schema.SearchResp, err error) {
	if dto.Page < 1 {
		dto.Page = 1
	}
	if len(dto.Query) == 0 {
		return &schema.SearchResp{
			Total:         0,
			SearchResults: make([]*schema.SearchResult, 0),
		}, nil
	}

	// search type
	cond := ss.searchParser.ParseStructure(ctx, dto)

	// check search plugin
	var finder plugin.Search
	_ = plugin.CallSearch(func(search plugin.Search) error {
		finder = search
		return nil
	})

	resp = &schema.SearchResp{}
	// search plugin is not found, call system search
	if finder == nil {
		if cond.SearchAll() {
			resp.SearchResults, resp.Total, err =
				ss.searchRepo.SearchContents(ctx, cond.Words, cond.Tags, cond.UserID, cond.VoteAmount, dto.Page, dto.Size, dto.Order)
		} else if cond.SearchQuestion() {
			resp.SearchResults, resp.Total, err =
				ss.searchRepo.SearchQuestions(ctx, cond.Words, cond.Tags, cond.NotAccepted, cond.Views, cond.AnswerAmount, dto.Page, dto.Size, dto.Order)
		} else if cond.SearchAnswer() {
			resp.SearchResults, resp.Total, err =
				ss.searchRepo.SearchAnswers(ctx, cond.Words, cond.Tags, cond.Accepted, cond.QuestionID, dto.Page, dto.Size, dto.Order)
		}
		return
	}
	return ss.searchByPlugin(ctx, finder, cond, dto)
}

func (ss *SearchService) searchByPlugin(ctx context.Context, finder plugin.Search, cond *schema.SearchCondition, dto *schema.SearchDTO) (resp *schema.SearchResp, err error) {
	var res []plugin.SearchResult
	resp = &schema.SearchResp{}
	if cond.SearchAll() {
		res, resp.Total, err = finder.SearchContents(ctx, cond.Convert2PluginSearchCond(dto.Page, dto.Size, dto.Order))
	} else if cond.SearchQuestion() {
		res, resp.Total, err = finder.SearchQuestions(ctx, cond.Convert2PluginSearchCond(dto.Page, dto.Size, dto.Order))
	} else if cond.SearchAnswer() {
		res, resp.Total, err = finder.SearchAnswers(ctx, cond.Convert2PluginSearchCond(dto.Page, dto.Size, dto.Order))
	}

	resp.SearchResults, err = ss.searchRepo.ParseSearchPluginResult(ctx, res)
	return resp, err
}
