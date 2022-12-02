package migrations

import (
	"time"

	"github.com/answerdev/answer/internal/entity"
	"github.com/segmentfault/pacman/log"
	"xorm.io/xorm"
)

func addActivityTimeline(x *xorm.Engine) error {
	// only increasing field length to 128
	type Config struct {
		Key string `xorm:"unique VARCHAR(128) key"`
	}
	if err := x.Sync(new(Config)); err != nil {
		return err
	}
	defaultConfigTable := []*entity.Config{
		{ID: 36, Key: "rank.question.add", Value: `1`},
		{ID: 37, Key: "rank.question.edit", Value: `200`},
		{ID: 38, Key: "rank.question.delete", Value: `-1`},
		{ID: 39, Key: "rank.question.vote_up", Value: `15`},
		{ID: 40, Key: "rank.question.vote_down", Value: `125`},
		{ID: 41, Key: "rank.answer.add", Value: `1`},
		{ID: 42, Key: "rank.answer.edit", Value: `200`},
		{ID: 43, Key: "rank.answer.delete", Value: `-1`},
		{ID: 44, Key: "rank.answer.accept", Value: `1`},
		{ID: 45, Key: "rank.answer.vote_up", Value: `15`},
		{ID: 46, Key: "rank.answer.vote_down", Value: `125`},
		{ID: 47, Key: "rank.comment.add", Value: `1`},
		{ID: 48, Key: "rank.comment.edit", Value: `-1`},
		{ID: 49, Key: "rank.comment.delete", Value: `-1`},
		{ID: 50, Key: "rank.report.add", Value: `1`},
		{ID: 51, Key: "rank.tag.add", Value: `1`},
		{ID: 52, Key: "rank.tag.edit", Value: `100`},
		{ID: 53, Key: "rank.tag.delete", Value: `-1`},
		{ID: 54, Key: "rank.tag.synonym", Value: `20000`},
		{ID: 55, Key: "rank.link.url_limit", Value: `10`},
		{ID: 56, Key: "rank.vote.detail", Value: `0`},

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

		{ID: 107, Key: "rank.comment.vote_up", Value: `1`},
		{ID: 108, Key: "rank.comment.vote_down", Value: `1`},
		{ID: 109, Key: "rank.question.edit_without_review", Value: `2000`},
		{ID: 110, Key: "rank.answer.edit_without_review", Value: `2000`},
		{ID: 111, Key: "rank.tag.edit_without_review", Value: `20000`},
		{ID: 112, Key: "rank.answer.audit", Value: `2000`},
		{ID: 113, Key: "rank.question.audit", Value: `2000`},
		{ID: 114, Key: "rank.tag.audit", Value: `20000`},
	}
	for _, c := range defaultConfigTable {
		exist, err := x.Get(&entity.Config{ID: c.ID, Key: c.Key})
		if err != nil {
			return err
		}
		if exist {
			if _, err = x.Update(c, &entity.Config{ID: c.ID, Key: c.Key}); err != nil {
				log.Errorf("update %+v config failed: %s", c, err)
				return err
			}
			continue
		}
		if _, err = x.Insert(&entity.Config{ID: c.ID, Key: c.Key, Value: c.Value}); err != nil {
			log.Errorf("insert %+v config failed: %s", c, err)
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
	type Tag struct {
		UserID string `xorm:"not null default 0 BIGINT(20) user_id"`
	}
	type Question struct {
		UpdatedAt      time.Time `xorm:"updated_at TIMESTAMP"`
		LastEditUserID string    `xorm:"not null default 0 BIGINT(20) last_edit_user_id"`
		PostUpdateTime time.Time `xorm:"post_update_time TIMESTAMP"`
	}
	type Answer struct {
		UpdatedAt      time.Time `xorm:"updated_at TIMESTAMP"`
		LastEditUserID string    `xorm:"not null default 0 BIGINT(20) last_edit_user_id"`
	}
	return x.Sync(new(Activity), new(Revision), new(Tag), new(Question), new(Answer))
}
