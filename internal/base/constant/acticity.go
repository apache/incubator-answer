package constant

// question activity

type ActivityTypeKey string

const (
	ActQuestionAsked     ActivityTypeKey = "question.asked"
	ActQuestionClosed    ActivityTypeKey = "question.closed"
	ActQuestionReopened  ActivityTypeKey = "question.reopened"
	ActQuestionAnswered  ActivityTypeKey = "question.answered"
	ActQuestionCommented ActivityTypeKey = "question.commented"
	ActQuestionAccept    ActivityTypeKey = "question.accept"
	ActQuestionUpvote    ActivityTypeKey = "question.upvote"
	ActQuestionDownvote  ActivityTypeKey = "question.downvote"
	ActQuestionEdit      ActivityTypeKey = "question.edit"
	ActQuestionRollback  ActivityTypeKey = "question.rollback"
	ActQuestionDeleted   ActivityTypeKey = "question.deleted"
	ActQuestionUndeleted ActivityTypeKey = "question.undeleted"
)

// answer activity

const (
	ActAnswerAnswered  ActivityTypeKey = "answer.answered"
	ActAnswerCommented ActivityTypeKey = "answer.commented"
	ActAnswerAccept    ActivityTypeKey = "answer.accept"
	ActAnswerUpvote    ActivityTypeKey = "answer.upvote"
	ActAnswerDownvote  ActivityTypeKey = "answer.downvote"
	ActAnswerEdit      ActivityTypeKey = "answer.edit"
	ActAnswerRollback  ActivityTypeKey = "answer.rollback"
	ActAnswerDeleted   ActivityTypeKey = "answer.deleted"
	ActAnswerUndeleted ActivityTypeKey = "answer.undeleted"
)

// tag activity

const (
	ActTagCreated   ActivityTypeKey = "tag.created"
	ActTagEdit      ActivityTypeKey = "tag.edit"
	ActTagRollback  ActivityTypeKey = "tag.rollback"
	ActTagDeleted   ActivityTypeKey = "tag.deleted"
	ActTagUndeleted ActivityTypeKey = "tag.undeleted"
)
