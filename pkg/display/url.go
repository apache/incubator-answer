package display

import (
	"github.com/answerdev/answer/internal/base/constant"
	"github.com/answerdev/answer/pkg/htmltext"
	"github.com/answerdev/answer/pkg/uid"
)

// QuestionURL get question url
func QuestionURL(permalink int, siteUrl, questionID, title string) string {
	u := siteUrl + "/questions"
	if permalink == constant.PermalinkQuestionIDAndTitle || permalink == constant.PermalinkQuestionID {
		questionID = uid.DeShortID(questionID)
	} else {
		questionID = uid.EnShortID(questionID)
	}
	u += "/" + questionID
	if permalink == constant.PermalinkQuestionIDAndTitle || permalink == constant.PermalinkQuestionIDAndTitleByShortID {
		u += "/" + htmltext.UrlTitle(title)
	}
	return u
}

// AnswerURL get answer url
func AnswerURL(permalink int, siteUrl, questionID, title, answerID string) string {
	if permalink == constant.PermalinkQuestionIDAndTitle ||
		permalink == constant.PermalinkQuestionID {
		answerID = uid.DeShortID(answerID)
	} else {
		answerID = uid.EnShortID(answerID)
	}
	return QuestionURL(permalink, siteUrl, questionID, title) + "/" + answerID
}

// CommentURL get comment url
func CommentURL(permalink int, siteUrl, questionID, title, answerID, commentID string) string {
	if len(answerID) > 0 {
		return AnswerURL(permalink, siteUrl, questionID, answerID, title) + "?commentId=" + commentID
	}
	return QuestionURL(permalink, siteUrl, questionID, title) + "?commentId=" + commentID
}
