package plugin

import (
	"context"
)

type SearchResult struct {
	// ID content ID
	ID string
	// Type content type, example: "answer", "question"
	Type string
}

type SearchContent struct {
	ObjectID    string   `json:"objectID"`
	Title       string   `json:"title"`
	Type        string   `json:"type"`
	Content     string   `json:"content"`
	Answers     int64    `json:"answers"`
	Status      int64    `json:"status"`
	Tags        []string `json:"tags"`
	QuestionID  string   `json:"questionID"`
	UserID      string   `json:"userID"`
	Views       int64    `json:"views"`
	Created     int64    `json:"created"`
	Active      int64    `json:"active"`
	Score       int64    `json:"score"`
	HasAccepted bool     `json:"hasAccepted"`
}

type SearchBasicCond struct {
	// From zero-based page number
	Page int
	// Page size
	PageSize int

	// The keywords for search.
	Words []string
	// TagIDs is a list of tag IDs.
	TagIDs []string
	// The object's owner user ID.
	UserID string
	// The order of the search result.
	Order string

	// Weathers the question is accepted or not. Only support search question.
	QuestionAccepted AcceptedCond
	// Weathers the answer is accepted or not. Only support search answer.
	AnswerAccepted AcceptedCond

	// Only support search answer.
	QuestionID string

	// greater than or equal to the number of votes.
	VoteAmount int
	// greater than or equal to the number of views.
	ViewAmount int
	// greater than or equal to the number of answers. Only support search question.
	AnswerAmount int
}

type AcceptedCond int

const (
	AcceptedCondAll AcceptedCond = iota
	AcceptedCondTrue
	AcceptedCondFalse
)

type Search interface {
	Base
	SearchContents(ctx context.Context, cond *SearchBasicCond) (res []SearchResult, total int64, err error)
	SearchQuestions(ctx context.Context, cond *SearchBasicCond) (res []SearchResult, total int64, err error)
	SearchAnswers(ctx context.Context, cond *SearchBasicCond) (res []SearchResult, total int64, err error)
	UpdateContent(ctx context.Context, contentID string, content *SearchContent) error
	DeleteContent(ctx context.Context, contentID string) error
}

var (
	// CallUserCenter is a function that calls all registered parsers
	CallSearch,
	registerSearch = MakePlugin[Search](false)
)
