package plugin

type Reviewer interface {
	Base
	Review(content *ReviewContent) (result *ReviewResult)
}

// ReviewContent is a struct that contains the content of a review
type ReviewContent struct {
	// The type of the content, e.g. question, answer
	ObjectType string
	Title      string
	Content    string
	Tags       []string
	// The author of the content
	Author ReviewContentAuthor
}

type ReviewContentAuthor struct {
	// The user's reputation
	Rank int
	// The amount of questions and answers that the user has approved
	ApprovedQuestionAmount int64
	ApprovedAnswerAmount   int64
	// 1:User 2:Admin 3:Moderator
	Role int
}

// ReviewResult is a struct that contains the result of a review
type ReviewResult struct {
	// If the review is approved
	Approved bool
	// The reason for the result
	Reason string
}

var (
	// CallReviewer is a function that calls all registered parsers
	CallReviewer,
	registerReviewer = MakePlugin[Reviewer](false)
)
