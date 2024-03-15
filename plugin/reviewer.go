package plugin

type Reviewer interface {
	Base
	Review(content *ReviewContent) (result *ReviewResult)
}

// ReviewContent is a struct that contains the content of a review
type ReviewContent struct {
	// The type of the content, e.g. question, answer
	ObjectType string
	// The title of the content, only available for the question
	Title string
	// The content of the review, always available
	Content string
	// The tags of the content, only available for the question
	Tags []string
	// The author of the content
	Author ReviewContentAuthor
	// Review Language, the site language. e.g. en_US
	// The plugin may reply the review result according to the language
	Language string
}

type ReviewContentAuthor struct {
	// The user's reputation
	Rank int
	// The amount of questions that has approved
	ApprovedQuestionAmount int64
	// The amount of answers that has approved
	ApprovedAnswerAmount int64
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
