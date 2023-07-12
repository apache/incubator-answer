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
	ActivityTypeFlagMapping = map[string]string{
		QuestionVoteUp:    "action_activity_type.upvote",
		QuestionVoteDown:  "action_activity_type.downvote",
		QuestionVotedUp:   "action_activity_type.upvoted",
		QuestionVotedDown: "action_activity_type.downvoted",
		AnswerVoteUp:      "action_activity_type.upvote",
		AnswerVoteDown:    "action_activity_type.downvote",
		AnswerVotedUp:     "action_activity_type.upvoted",
		AnswerVotedDown:   "action_activity_type.downvoted",
		AnswerAccepted:    "action_activity_type.accepted",
		AnswerAccept:      "action_activity_type.accept",
		CommentVoteUp:     "action_activity_type.upvote",
	}
)
