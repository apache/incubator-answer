package plugin

type Notification interface {
	Base

	// Notify sends a notification to the user
	Notify(msg *NotificationMessage)
}

type NotificationMessage struct {
	Type                      NotificationType          `json:"notification_type"`
	ReceiverUserID            string                    `json:"receiver_user_id"`
	ReceiverLang              string                    `json:"receiver_lang"`
	ExternalID                string                    `json:"external_id"`
	NewAnswerNoticeData       NewAnswerNoticeData       `json:"new_answer_template_raw_data,omitempty"`
	NewInviteAnswerNoticeData NewInviteAnswerNoticeData `json:"new_invite_answer_template_raw_data,omitempty"`
	NewCommentNoticeData      NewCommentNoticeData      `json:"new_comment_template_raw_data,omitempty"`
	NewQuestionNoticeData     NewQuestionNoticeData     `json:"new_question_template_raw_data,omitempty"`
}

type NotificationType string

const (
	NewAnswer              NotificationType = "new_answer"
	NewInviteAnswer        NotificationType = "new_invite_answer"
	NewComment             NotificationType = "new_comment"
	NewQuestion            NotificationType = "new_question"
	NewQuestionFollowedTag NotificationType = "new_question_followed_tag"
)

type NewAnswerNoticeData struct {
	AnswerUserDisplayName string `json:"answer_user_display_name"`
	QuestionTitle         string `json:"question_title"`
	QuestionID            string `json:"question_id"`
	AnswerID              string `json:"answer_id"`
	AnswerSummary         string `json:"answer_summary"`
}

type NewInviteAnswerNoticeData struct {
	InviterDisplayName string `json:"inviter_display_name"`
	QuestionTitle      string `json:"question_title"`
	QuestionID         string `json:"question_id"`
}

type NewCommentNoticeData struct {
	CommentUserDisplayName string `json:"comment_user_display_name"`
	QuestionTitle          string `json:"question_title"`
	QuestionID             string `json:"question_id"`
	AnswerID               string `json:"answer_id"`
	CommentID              string `json:"comment_id"`
	CommentSummary         string `json:"comment_summary"`
}

type NewQuestionNoticeData struct {
	QuestionTitle string `json:"question_title"`
	QuestionUrl   string `json:"question_url"`
	Tags          string `json:"tags"`
}

var (
	// CallNotification is a function that calls all registered notification plugins
	CallNotification,
	registerNotification = MakePlugin[Notification](false)
)
