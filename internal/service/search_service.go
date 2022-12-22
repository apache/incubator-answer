package service

import (
	"context"
	"github.com/answerdev/answer/internal/schema"
	"github.com/answerdev/answer/internal/service/search_common"
	"github.com/answerdev/answer/internal/service/search_parser"
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
func (ss *SearchService) Search(ctx context.Context, dto *schema.SearchDTO) (resp []schema.SearchResp, total int64, extra interface{}, err error) {
	extra = nil
	if dto.Page < 1 {
		dto.Page = 1
	}

	// search type
	searchType,
		// search all
		userID,
		votes,
		// search questions
		notAccepted,
		_,
		views,
		answers,
		// search answers
		accepted,
		questionID,
		_,
		// common fields
		tags,
		words := ss.searchParser.ParseStructure(dto)

	switch searchType {
	case "all":
		resp, total, err = ss.searchRepo.SearchContents(ctx, words, tags, userID, votes, dto.Page, dto.Size, dto.Order)
		if err != nil {
			return nil, 0, nil, err
		}
	case "question":
		resp, total, err = ss.searchRepo.SearchQuestions(ctx, words, tags, notAccepted, views, answers, dto.Page, dto.Size, dto.Order)
	case "answer":
		resp, total, err = ss.searchRepo.SearchAnswers(ctx, words, tags, accepted, questionID, dto.Page, dto.Size, dto.Order)
	}
	return
}
