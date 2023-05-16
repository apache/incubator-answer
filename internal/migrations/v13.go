package migrations

import (
	"fmt"

	"github.com/answerdev/answer/internal/entity"
	"github.com/segmentfault/pacman/log"
	"xorm.io/xorm"
)

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

	//tag question count

	//user question count

	//user answer count

	return nil
}
