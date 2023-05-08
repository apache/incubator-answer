package migrations

import (
	"fmt"

	"github.com/answerdev/answer/internal/entity"
	"github.com/segmentfault/pacman/log"
	"xorm.io/xorm"
)

func updateQuestionPostTime(x *xorm.Engine) error {
	questionList := make([]entity.Question, 0)
	err := x.Find(&questionList, &entity.Question{})
	if err != nil {
		return fmt.Errorf("get questions failed: %w", err)
	}
	for _, item := range questionList {
		if item.PostUpdateTime.IsZero() {
			if !item.UpdatedAt.IsZero() {
				item.PostUpdateTime = item.UpdatedAt
			} else if !item.CreatedAt.IsZero() {
				item.PostUpdateTime = item.CreatedAt
			}
			if _, err = x.Update(item, &entity.Question{ID: item.ID}); err != nil {
				log.Errorf("update %+v config failed: %s", item, err)
				return fmt.Errorf("update question failed: %w", err)
			}
		}

	}

	return nil
}
