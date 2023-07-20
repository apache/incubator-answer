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

type Search interface {
	Base
	SearchContents(ctx context.Context, words []string, tagIDs []string, userID string, votes int, page, size int, order string) (res []SearchResult, total int64, err error)
	SearchQuestions(ctx context.Context, words []string, tagIDs []string, notAccepted bool, views, answers int, page, size int, order string) (res []SearchResult, total int64, err error)
	SearchAnswers(ctx context.Context, words []string, tagIDs []string, accepted bool, questionID string, page, size int, order string) (res []SearchResult, total int64, err error)
	UpdateContent(ctx context.Context, contentID string, content *SearchContent) error
	DeleteContent(ctx context.Context, contentID string) error
}

var (
	// CallUserCenter is a function that calls all registered parsers
	CallSearch,
	registerSearch = MakePlugin[Search](false)
)
