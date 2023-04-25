package migrations

import (
	"fmt"

	"github.com/answerdev/answer/internal/entity"
	"github.com/segmentfault/pacman/log"
	"xorm.io/xorm"
)

func updateRolePinAndHideFeatures(x *xorm.Engine) error {

	defaultConfigTable := []*entity.Config{
		{ID: 119, Key: "question.pin", Value: `0`},
		{ID: 120, Key: "question.unpin", Value: `0`},
		{ID: 121, Key: "question.show", Value: `0`},
		{ID: 122, Key: "question.hide", Value: `0`},
		{ID: 123, Key: "rank.question.pin", Value: `-1`},
		{ID: 124, Key: "rank.question.unpin", Value: `-1`},
		{ID: 125, Key: "rank.question.show", Value: `-1`},
		{ID: 126, Key: "rank.question.hide", Value: `-1`},
	}
	for _, c := range defaultConfigTable {
		exist, err := x.Get(&entity.Config{ID: c.ID})
		if err != nil {
			return fmt.Errorf("get config failed: %w", err)
		}
		if exist {
			if _, err = x.Update(c, &entity.Config{ID: c.ID}); err != nil {
				log.Errorf("update %+v config failed: %s", c, err)
				return fmt.Errorf("update config failed: %w", err)
			}
			continue
		}
		if _, err = x.Insert(&entity.Config{ID: c.ID, Key: c.Key, Value: c.Value}); err != nil {
			log.Errorf("insert %+v config failed: %s", c, err)
			return fmt.Errorf("add config failed: %w", err)
		}
	}

	return nil
}
