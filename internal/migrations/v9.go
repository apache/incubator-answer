package migrations

import (
	"context"
	"fmt"

	"github.com/answerdev/answer/internal/entity"
	"github.com/segmentfault/pacman/log"
	"xorm.io/xorm"
)

func updateAcceptAnswerRank(ctx context.Context, x *xorm.Engine) error {
	c := &entity.Config{ID: 44, Key: "rank.answer.accept", Value: `-1`}
	if _, err := x.Context(ctx).Update(c, &entity.Config{ID: 44, Key: "rank.answer.accept"}); err != nil {
		log.Errorf("update %+v config failed: %s", c, err)
		return fmt.Errorf("update config failed: %w", err)
	}
	return nil
}
