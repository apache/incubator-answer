package schema

import (
	"github.com/answerdev/answer/internal/entity"
	"github.com/answerdev/answer/pkg/uid"
)

type ExternalNotificationMsg struct {
	ReceiverUserID string `json:"receiver_user_id"`
	ReceiverEmail  string `json:"receiver_email"`
	ReceiverLang   string `json:"receiver_lang"`

	NewAnswerTemplateRawData       *NewAnswerTemplateRawData       `json:"new_answer_template_raw_data,omitempty"`
	NewInviteAnswerTemplateRawData *NewInviteAnswerTemplateRawData `json:"new_invite_answer_template_raw_data,omitempty"`
	NewCommentTemplateRawData      *NewCommentTemplateRawData      `json:"new_comment_template_raw_data,omitempty"`
	NewQuestionTemplateRawData     *NewQuestionTemplateRawData     `json:"new_question_template_raw_data,omitempty"`
}

func CreateNewQuestionNotificationMsg(questionID, questionTitle string, tags []*entity.Tag) *ExternalNotificationMsg {
	questionID = uid.DeShortID(questionID)
	msg := &ExternalNotificationMsg{
		NewQuestionTemplateRawData: &NewQuestionTemplateRawData{
			QuestionID:    questionID,
			QuestionTitle: questionTitle,
		},
	}
	for _, tag := range tags {
		msg.NewQuestionTemplateRawData.Tags = append(msg.NewQuestionTemplateRawData.Tags, tag.SlugName)
		msg.NewQuestionTemplateRawData.TagIDs = append(msg.NewQuestionTemplateRawData.TagIDs, tag.ID)
	}
	return msg
}
