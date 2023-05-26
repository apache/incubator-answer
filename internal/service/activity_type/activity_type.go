package activity_type

const (
	QuestionVoteUp    = "question.vote_up"
	QuestionVoteDown  = "question.vote_down"
	QuestionVotedUp   = "question.voted_up"
	QuestionVotedDown = "question.voted_down"
	AnswerVoteUp      = "answer.vote_up"
	AnswerVoteDown    = "answer.vote_down"
	AnswerVotedUp     = "answer.voted_up"
	AnswerVotedDown   = "answer.voted_down"
	AnswerAccepted    = "answer.accepted"
	AnswerAccept      = "answer.accept"
	CommentVoteUp     = "comment.vote_up"
)

var (
	ActivityTypeList = []string{
		QuestionVoteUp,
		QuestionVoteDown,
		QuestionVotedUp,
		QuestionVotedDown,
		AnswerVoteUp,
		AnswerVoteDown,
		AnswerVotedUp,
		AnswerVotedDown,
		AnswerAccepted,
		AnswerAccept,
		CommentVoteUp,
	}
	VoteActivityTypeList = []string{
		QuestionVoteUp,
		QuestionVoteDown,
		QuestionVotedUp,
		QuestionVotedDown,
		AnswerVoteUp,
		AnswerVoteDown,
		AnswerVotedUp,
		AnswerVotedDown,
		CommentVoteUp,
	}
	activityTypeFlagMapping = map[string]string{
		QuestionVoteUp:    "upvote",
		QuestionVoteDown:  "downvote",
		QuestionVotedUp:   "upvoted",
		QuestionVotedDown: "downvoted",
		AnswerVoteUp:      "upvote",
		AnswerVoteDown:    "downvote",
		AnswerVotedUp:     "upvoted",
		AnswerVotedDown:   "downvoted",
		AnswerAccepted:    "accepted",
		AnswerAccept:      "accept",
		CommentVoteUp:     "upvote",
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
