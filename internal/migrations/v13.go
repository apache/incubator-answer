package migrations

import (
	"encoding/json"
	"fmt"

	"github.com/answerdev/answer/internal/base/constant"
	"github.com/answerdev/answer/internal/entity"
	"github.com/answerdev/answer/internal/schema"
	"github.com/segmentfault/pacman/log"
	"xorm.io/xorm"
)

func addGravatarBaseURL(x *xorm.Engine) error {
	usersSiteInfo := &entity.SiteInfo{
		Type: constant.SiteTypeUsers,
	}
	exist, err := x.Get(usersSiteInfo)
	if err != nil {
		return fmt.Errorf("get config failed: %w", err)
	}
	if exist {
		content := &schema.SiteUsersReq{}
		_ = json.Unmarshal([]byte(usersSiteInfo.Content), content)
		content.GravatarBaseURL = "https://www.gravatar.com/avatar/"
		data, _ := json.Marshal(content)
		usersSiteInfo.Content = string(data)

		_, err = x.ID(usersSiteInfo.ID).Cols("content").Update(usersSiteInfo)
		if err != nil {
			return fmt.Errorf("update site info failed: %w", err)
		}
	}

	//search all answers
	answers := make([]entity.Answer, 0)
	err = x.Find(&answers, &entity.Answer{Status: entity.AnswerStatusAvailable})
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
