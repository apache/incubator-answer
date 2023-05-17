package migrations

import (
	"fmt"

	"github.com/answerdev/answer/internal/entity"
	"github.com/segmentfault/pacman/log"
	"xorm.io/xorm"
)

func updateCount(x *xorm.Engine) error {
	// updateQuestionCount(x)
	// updateTagCount(x)
	// updateUserQuestionCount(x)
	updateUserAnswerCount(x)
	return nil
}

func updateQuestionCount(x *xorm.Engine) error {
	//question answer count
	answers := make([]entity.Answer, 0)
	err := x.Find(&answers, &entity.Answer{Status: entity.AnswerStatusAvailable})
	if err != nil {
		return fmt.Errorf("get answers failed: %w", err)
	}
	questionAnswerCount := make(map[string]int)
	for _, answer := range answers {
		_, ok := questionAnswerCount[answer.QuestionID]
		if !ok {
			questionAnswerCount[answer.QuestionID] = 1
		} else {
			questionAnswerCount[answer.QuestionID]++
		}
	}
	questionList := make([]entity.Question, 0)
	err = x.Find(&questionList, &entity.Question{})
	if err != nil {
		return fmt.Errorf("get questions failed: %w", err)
	}
	for _, item := range questionList {
		_, ok := questionAnswerCount[item.ID]
		if ok {
			item.AnswerCount = questionAnswerCount[item.ID]
			if _, err = x.Update(item, &entity.Question{ID: item.ID}); err != nil {
				log.Errorf("update %+v config failed: %s", item, err)
				return fmt.Errorf("update question failed: %w", err)
			}
		}
	}

	return nil
}

// updateTagCount update tag count
func updateTagCount(x *xorm.Engine) error {
	tagRelList := make([]entity.TagRel, 0)
	err := x.Find(&tagRelList, &entity.TagRel{})
	if err != nil {
		return fmt.Errorf("get tag rel failed: %w", err)
	}
	questionIDs := make([]string, 0)
	questionsAvailableMap := make(map[string]bool)
	questionsHideMap := make(map[string]bool)
	for _, item := range tagRelList {
		questionIDs = append(questionIDs, item.ObjectID)
		questionsAvailableMap[item.ObjectID] = false
		questionsHideMap[item.ObjectID] = false
	}
	questionList := make([]entity.Question, 0)
	err = x.In("id", questionIDs).In("question.status", []int{entity.QuestionStatusAvailable, entity.QuestionStatusClosed}).Find(&questionList, &entity.Question{})
	if err != nil {
		return fmt.Errorf("get questions failed: %w", err)
	}
	for _, question := range questionList {
		_, ok := questionsAvailableMap[question.ID]
		if ok {
			questionsAvailableMap[question.ID] = true
			if question.Show == entity.QuestionHide {
				questionsHideMap[question.ID] = true
			}
		}
	}

	for id, ok := range questionsHideMap {
		if ok {
			if _, err = x.Cols("status").Update(&entity.TagRel{Status: entity.TagRelStatusHide}, &entity.TagRel{ObjectID: id}); err != nil {
				log.Errorf("update %+v config failed: %s", id, err)
			}
		}
	}

	for id, ok := range questionsAvailableMap {
		if !ok {
			if _, err = x.Cols("status").Update(&entity.TagRel{Status: entity.TagRelStatusDeleted}, &entity.TagRel{ObjectID: id}); err != nil {
				log.Errorf("update %+v config failed: %s", id, err)
			}
		}
	}

	//select tag count
	newTagRelList := make([]entity.TagRel, 0)
	err = x.Find(&newTagRelList, &entity.TagRel{Status: entity.TagRelStatusAvailable})
	if err != nil {
		return fmt.Errorf("get tag rel failed: %w", err)
	}
	tagCountMap := make(map[string]int)
	for _, v := range newTagRelList {
		_, ok := tagCountMap[v.TagID]
		if !ok {
			tagCountMap[v.TagID] = 1
		} else {
			tagCountMap[v.TagID]++
		}
	}
	TagList := make([]entity.Tag, 0)
	err = x.Find(&TagList, &entity.Tag{})
	if err != nil {
		return fmt.Errorf("get tag  failed: %w", err)
	}
	for _, tag := range TagList {
		_, ok := tagCountMap[tag.ID]
		if ok {
			tag.QuestionCount = tagCountMap[tag.ID]
			if _, err = x.Update(tag, &entity.Tag{ID: tag.ID}); err != nil {
				log.Errorf("update %+v tag failed: %s", tag.ID, err)
				return fmt.Errorf("update tag failed: %w", err)
			}
		} else {
			tag.QuestionCount = 0
			if _, err = x.Update(tag, &entity.Tag{ID: tag.ID}); err != nil {
				log.Errorf("update %+v tag failed: %s", tag.ID, err)
				return fmt.Errorf("update tag failed: %w", err)
			}
		}
	}
	return nil
}

// updateUserQuestionCount update user question count
func updateUserQuestionCount(x *xorm.Engine) error {
	questionList := make([]entity.Question, 0)
	err := x.In("status", []int{entity.QuestionStatusAvailable, entity.QuestionStatusClosed}).Find(&questionList, &entity.Question{})
	if err != nil {
		return fmt.Errorf("get question  failed: %w", err)
	}
	userQuestionCountMap := make(map[string]int)
	for _, question := range questionList {
		_, ok := userQuestionCountMap[question.UserID]
		if !ok {
			userQuestionCountMap[question.UserID] = 1
		} else {
			userQuestionCountMap[question.UserID]++
		}
	}
	userList := make([]entity.User, 0)
	err = x.Find(&userList, &entity.User{})
	if err != nil {
		return fmt.Errorf("get user  failed: %w", err)
	}
	for _, user := range userList {
		_, ok := userQuestionCountMap[user.ID]
		if ok {
			user.QuestionCount = userQuestionCountMap[user.ID]
			if _, err = x.Cols("question_count").Update(user, &entity.User{ID: user.ID}); err != nil {
				log.Errorf("update %+v user failed: %s", user.ID, err)
				return fmt.Errorf("update user failed: %w", err)
			}
		} else {
			user.QuestionCount = 0
			if _, err = x.Cols("question_count").Update(user, &entity.User{ID: user.ID}); err != nil {
				log.Errorf("update %+v user failed: %s", user.ID, err)
				return fmt.Errorf("update user failed: %w", err)
			}
		}
	}
	return nil
}

// updateUserAnswerCount update user answer count
func updateUserAnswerCount(x *xorm.Engine) error {
	return nil
}
