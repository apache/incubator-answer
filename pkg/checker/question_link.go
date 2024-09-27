package checker

import (
	"github.com/apache/incubator-answer/internal/base/constant"
	"github.com/apache/incubator-answer/pkg/obj"
)

const (
	QuestionLinkTypeURL = 1
	QuestionLinkTypeID  = 2
)

type QuestionLink struct {
	LinkType   int
	QuestionID string
	AnswerID   string
}

func GetQuestionLink(content string) []QuestionLink {
	uniqueIDs := make(map[string]struct{})
	var questionLinks []QuestionLink

	// use two pointer to find the link
	left, right := 0, 0
	for right < len(content) {
		// find "/questions/" or "#"
		if right+11 < len(content) && content[right:right+11] == "/questions/" {
			left = right
			right += 11
			processURL(content, &left, &right, uniqueIDs, &questionLinks)
		} else if content[right] == '#' {
			left = right + 1
			right = left
			processID(content, &left, &right, uniqueIDs, &questionLinks)
		} else {
			right++
		}
	}

	return questionLinks
}

func processURL(content string, left, right *int, uniqueIDs map[string]struct{}, questionLinks *[]QuestionLink) {
	for *right < len(content) && isDigit(content[*right]) {
		*right++
	}
	questionID := content[*left+len("/questions/") : *right]

	answerID := ""
	if *right < len(content) && content[*right] == '/' {
		*left = *right + 1
		*right = *left
		for *right < len(content) && isDigit(content[*right]) {
			*right++
		}
		answerID = content[*left:*right]
	}

	addUniqueID(questionID, answerID, QuestionLinkTypeURL, uniqueIDs, questionLinks)
}

func processID(content string, left, right *int, uniqueIDs map[string]struct{}, questionLinks *[]QuestionLink) {
	for *right < len(content) && isDigit(content[*right]) {
		*right++
	}
	id := content[*left:*right]
	addUniqueID(id, "", QuestionLinkTypeID, uniqueIDs, questionLinks)
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func addUniqueID(questionID, answerID string, linkType int, uniqueIDs map[string]struct{}, questionLinks *[]QuestionLink) {
	if questionID == "" && answerID == "" {
		return
	}

	isAdd := false
	if answerID != "" {
		if objectType, err := obj.GetObjectTypeStrByObjectID(answerID); err == nil && objectType == constant.AnswerObjectType {
			if _, ok := uniqueIDs[answerID]; !ok {
				uniqueIDs[answerID] = struct{}{}
				isAdd = true
			}
		}
	}

	if objectType, err := obj.GetObjectTypeStrByObjectID(questionID); err == nil {
		if _, ok := uniqueIDs[questionID]; !ok {
			uniqueIDs[questionID] = struct{}{}
			isAdd = true
			if objectType == constant.AnswerObjectType {
				answerID = questionID
				questionID = ""
			}
		}
	}

	if isAdd {
		*questionLinks = append(*questionLinks, QuestionLink{
			LinkType:   linkType,
			QuestionID: questionID,
			AnswerID:   answerID,
		})
	}
}
