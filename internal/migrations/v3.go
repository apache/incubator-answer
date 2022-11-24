package migrations

import (
	"time"

	"github.com/answerdev/answer/internal/entity"
	"xorm.io/xorm"
)

func addActivityTimeline(x *xorm.Engine) error {
	defaultConfigTable := []*entity.Config{
		{ID: 87, Key: "question.asked", Value: `0`},
		{ID: 88, Key: "question.closed", Value: `0`},
		{ID: 89, Key: "question.reopened", Value: `0`},
		{ID: 90, Key: "question.answered", Value: `0`},
		{ID: 91, Key: "question.commented", Value: `0`},
		{ID: 92, Key: "question.accept", Value: `0`},
		{ID: 93, Key: "question.edited", Value: `0`},
		{ID: 94, Key: "question.rollback", Value: `0`},
		{ID: 95, Key: "question.deleted", Value: `0`},
		{ID: 96, Key: "question.undeleted", Value: `0`},
		{ID: 97, Key: "answer.answered", Value: `0`},
		{ID: 98, Key: "answer.commented", Value: `0`},
		{ID: 99, Key: "answer.edited", Value: `0`},
		{ID: 100, Key: "answer.rollback", Value: `0`},
		{ID: 101, Key: "answer.undeleted", Value: `0`},
		{ID: 102, Key: "tag.created", Value: `0`},
		{ID: 103, Key: "tag.edited", Value: `0`},
		{ID: 104, Key: "tag.rollback", Value: `0`},
		{ID: 105, Key: "tag.deleted", Value: `0`},
		{ID: 106, Key: "tag.undeleted", Value: `0`},
	}
	for _, c := range defaultConfigTable {
		exist, err := x.Get(&entity.Config{ID: c.ID, Key: c.Key})
		if err != nil {
			return err
		}
		if exist {
			continue
		}
		if _, err := x.Insert(&entity.Config{ID: c.ID, Key: c.Key, Value: c.Value}); err != nil {
			return err
		}
	}

	type Revision struct {
		ReviewUserID int64 `xorm:"not null default 0 BIGINT(20) review_user_id"`
	}
	type Activity struct {
		CancelledAt      time.Time `xorm:"TIMESTAMP cancelled_at"`
		RevisionID       int64     `xorm:"not null default 0 BIGINT(20) revision_id"`
		OriginalObjectID string    `xorm:"not null default 0 BIGINT(20) original_object_id"`
	}
	return x.Sync(new(Activity), new(Revision))
}
