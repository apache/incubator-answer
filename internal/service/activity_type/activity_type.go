package activity_type

import "github.com/answerdev/answer/internal/repo/config"

const (
	QuestionVoteUp   = "question.vote_up"
	QuestionVoteDown = "question.vote_down"
	AnswerVoteUp     = "answer.vote_up"
	AnswerVoteDown   = "answer.vote_down"
	CommentVoteUp    = "comment.vote_up"
	CommentVoteDown  = "comment.vote_down"
)

var (
	activityTypeFlagMapping = map[string]string{
		QuestionVoteUp:   "upvote",
		QuestionVoteDown: "downvote",
		AnswerVoteUp:     "upvote",
		AnswerVoteDown:   "downvote",
		CommentVoteUp:    "upvote",
		CommentVoteDown:  "downvote",
	}
)

func Format(activityTypeID int) string {
	activityTypeStr := config.ID2KeyMapping[activityTypeID]
	activityTypeFlag := activityTypeFlagMapping[activityTypeStr]
	if len(activityTypeFlag) == 0 {
		return "edit" // to edit
	}
	return activityTypeFlag // todo i18n support
}
