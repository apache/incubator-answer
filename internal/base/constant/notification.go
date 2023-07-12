package constant

const (
	// NotificationUpdateQuestion update question
	NotificationUpdateQuestion = "notification.action.update_question"
	// NotificationAnswerTheQuestion answer the question
	NotificationAnswerTheQuestion = "notification.action.answer_the_question"
	// NotificationUpVotedTheQuestion up voted the question
	NotificationUpVotedTheQuestion = "notification.action.up_voted_question"
	// NotificationDownVotedTheQuestion down voted the question
	NotificationDownVotedTheQuestion = "notification.action.down_voted_question"
	// NotificationUpdateAnswer update answer
	NotificationUpdateAnswer = "notification.action.update_answer"
	// NotificationAcceptAnswer accept answer
	NotificationAcceptAnswer = "notification.action.accept_answer"
	// NotificationUpVotedTheAnswer up voted the answer
	NotificationUpVotedTheAnswer = "notification.action.up_voted_answer"
	// NotificationDownVotedTheAnswer down voted the answer
	NotificationDownVotedTheAnswer = "notification.action.down_voted_answer"
	// NotificationCommentQuestion comment question
	NotificationCommentQuestion = "notification.action.comment_question"
	// NotificationCommentAnswer comment answer
	NotificationCommentAnswer = "notification.action.comment_answer"
	// NotificationUpVotedTheComment up voted the comment
	NotificationUpVotedTheComment = "notification.action.up_voted_comment"
	// NotificationReplyToYou reply to you
	NotificationReplyToYou = "notification.action.reply_to_you"
	// NotificationMentionYou mention you
	NotificationMentionYou = "notification.action.mention_you"
	// NotificationYourQuestionIsClosed your question is closed
	NotificationYourQuestionIsClosed = "notification.action.your_question_is_closed"
	// NotificationYourQuestionWasDeleted your question was deleted
	NotificationYourQuestionWasDeleted = "notification.action.your_question_was_deleted"
	// NotificationYourAnswerWasDeleted your answer was deleted
	NotificationYourAnswerWasDeleted = "notification.action.your_answer_was_deleted"
	// NotificationYourCommentWasDeleted your comment was deleted
	NotificationYourCommentWasDeleted = "notification.action.your_comment_was_deleted"
	// NotificationInvitedYouToAnswer invited you to answer
	NotificationInvitedYouToAnswer = "notification.action.invited_you_to_answer"
)

var (
	NotificationMsgTypeMapping = map[string]int{
		NotificationUpdateQuestion:         1,
		NotificationAnswerTheQuestion:      1,
		NotificationUpVotedTheQuestion:     2,
		NotificationDownVotedTheQuestion:   2,
		NotificationUpdateAnswer:           1,
		NotificationAcceptAnswer:           1,
		NotificationUpVotedTheAnswer:       2,
		NotificationDownVotedTheAnswer:     2,
		NotificationCommentQuestion:        1,
		NotificationCommentAnswer:          1,
		NotificationUpVotedTheComment:      2,
		NotificationReplyToYou:             1,
		NotificationMentionYou:             1,
		NotificationYourQuestionIsClosed:   1,
		NotificationYourQuestionWasDeleted: 1,
		NotificationYourAnswerWasDeleted:   1,
		NotificationYourCommentWasDeleted:  1,
		NotificationInvitedYouToAnswer:     3,
	}
)
