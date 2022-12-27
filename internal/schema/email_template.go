package schema

import "encoding/json"

const (
	AccountActivationSourceType SourceType = "account-activation"
	PasswordResetSourceType     SourceType = "password-reset"
	ConfirmNewEmailSourceType   SourceType = "password-reset"
	UnsubscribeSourceType       SourceType = "unsubscribe"
)

type SourceType string

type EmailCodeContent struct {
	SourceType SourceType `json:"source_type"`
	Email      string     `json:"e_mail"`
	UserID     string     `json:"user_id"`
}

func (r *EmailCodeContent) ToJSONString() string {
	codeBytes, _ := json.Marshal(r)
	return string(codeBytes)
}

func (r *EmailCodeContent) FromJSONString(data string) error {
	return json.Unmarshal([]byte(data), &r)
}

type NewAnswerTemplateRawData struct {
	AnswerUserDisplayName string
	QuestionTitle         string
	QuestionID            string
	AnswerID              string
	AnswerSummary         string
	UnsubscribeCode       string
}

type NewAnswerTemplateData struct {
	SiteName       string
	DisplayName    string
	QuestionTitle  string
	AnswerUrl      string
	AnswerSummary  string
	UnsubscribeUrl string
}

type NewCommentTemplateRawData struct {
	CommentUserDisplayName string
	QuestionTitle          string
	QuestionID             string
	AnswerID               string
	CommentID              string
	CommentSummary         string
	UnsubscribeCode        string
}

type NewCommentTemplateData struct {
	SiteName       string
	DisplayName    string
	QuestionTitle  string
	CommentUrl     string
	CommentSummary string
	UnsubscribeUrl string
}
