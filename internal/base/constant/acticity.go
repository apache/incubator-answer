package constant

type ActivityTypeKey string

const (
	ActEdited    = "edited"
	ActClosed    = "closed"
	ActVotedDown = "voted_down"
	ActVotedUp   = "voted_up"
	ActVoteDown  = "vote_down"
	ActVoteUp    = "vote_up"
	ActUpVote    = "upvote"
	ActDownVote  = "downvote"
	ActFollow    = "follow"
	ActAccepted  = "accepted"
	ActAccept    = "accept"
	ActPin       = "pin"
	ActUnPin     = "unpin"
	ActShow      = "show"
	ActHide      = "hide"
)

const (
	ActQuestionAsked     ActivityTypeKey = "question.asked"
	ActQuestionClosed    ActivityTypeKey = "question.closed"
	ActQuestionReopened  ActivityTypeKey = "question.reopened"
	ActQuestionAnswered  ActivityTypeKey = "question.answered"
	ActQuestionCommented ActivityTypeKey = "question.commented"
	ActQuestionAccept    ActivityTypeKey = "question.accept"
	ActQuestionUpvote    ActivityTypeKey = "question.upvote"
	ActQuestionDownVote  ActivityTypeKey = "question.downvote"
	ActQuestionEdited    ActivityTypeKey = "question.edited"
	ActQuestionRollback  ActivityTypeKey = "question.rollback"
	ActQuestionDeleted   ActivityTypeKey = "question.deleted"
	ActQuestionUndeleted ActivityTypeKey = "question.undeleted"
	ActQuestionPin       ActivityTypeKey = "question.pin"
	ActQuestionUnPin     ActivityTypeKey = "question.unpin"
	ActQuestionHide      ActivityTypeKey = "question.hide"
	ActQuestionShow      ActivityTypeKey = "question.show"
)

const (
	ActAnswerAnswered  ActivityTypeKey = "answer.answered"
	ActAnswerCommented ActivityTypeKey = "answer.commented"
	ActAnswerAccept    ActivityTypeKey = "answer.accept"
	ActAnswerUpvote    ActivityTypeKey = "answer.upvote"
	ActAnswerDownVote  ActivityTypeKey = "answer.downvote"
	ActAnswerEdited    ActivityTypeKey = "answer.edited"
	ActAnswerRollback  ActivityTypeKey = "answer.rollback"
	ActAnswerDeleted   ActivityTypeKey = "answer.deleted"
	ActAnswerUndeleted ActivityTypeKey = "answer.undeleted"
)

const (
	ActTagCreated   ActivityTypeKey = "tag.created"
	ActTagEdited    ActivityTypeKey = "tag.edited"
	ActTagRollback  ActivityTypeKey = "tag.rollback"
	ActTagDeleted   ActivityTypeKey = "tag.deleted"
	ActTagUndeleted ActivityTypeKey = "tag.undeleted"
)
