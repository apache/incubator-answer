package activity_type

const (
	QuestionVoteUp    = "question.vote_up"
	QuestionVoteDown  = "question.vote_down"
	AnswerVoteUp      = "answer.vote_up"
	AnswerVoteDown    = "answer.vote_down"
	CommentVoteUp     = "comment.vote_up"
	CommentVoteDown   = "comment.vote_down"
	AnswerAccepted    = "answer.accepted"
	AnswerAccept      = "answer.accept"
	QuestionVotedUp   = "question.voted_up"
	QuestionVotedDown = "question.voted_down"
	AnswerVotedUp     = "answer.voted_up"
	AnswerVotedDown   = "answer.voted_down"
)

var (
	ActivityTypeList = []string{
		QuestionVoteUp,
		QuestionVoteDown,
		AnswerVoteUp,
		AnswerVoteDown,
		CommentVoteUp,
		CommentVoteDown,
		AnswerAccepted,
		AnswerAccept,
		QuestionVotedUp,
		QuestionVotedDown,
		AnswerVotedUp,
		AnswerVotedDown,
	}
	activityTypeFlagMapping = map[string]string{
		QuestionVoteUp:    "upvote",
		QuestionVoteDown:  "downvote",
		AnswerVoteUp:      "upvote",
		AnswerVoteDown:    "downvote",
		CommentVoteUp:     "upvote",
		CommentVoteDown:   "downvote",
		AnswerAccepted:    "accepted",
		AnswerAccept:      "accept",
		QuestionVotedUp:   "upvoted",
		QuestionVotedDown: "downvoted",
		AnswerVotedUp:     "upvoted",
		AnswerVotedDown:   "downvoted",
	}
)

func Format(activityTypeID int) string {
	return ""
	//activityTypeStr := config_common.ID2KeyMapping[activityTypeID]
	//activityTypeFlag := activityTypeFlagMapping[activityTypeStr]
	//if len(activityTypeFlag) == 0 {
	//	return "edit" // to edit
	//}
	//return activityTypeFlag // todo i18n support
}
