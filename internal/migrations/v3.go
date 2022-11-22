package migrations

import (
	"time"

	"xorm.io/xorm"
)

func addActivityTimeline(x *xorm.Engine) error {
	type Reversion struct {
		ReviewUserID int64 `xorm:"not null default 0 BIGINT(20) review_user_id"`
	}
	type Activity struct {
		CancelledAt      time.Time `xorm:"TIMESTAMP cancelled_at"`
		RevisionID       int64     `xorm:"not null default 0 BIGINT(20) revision_id"`
		OriginalObjectID string    `xorm:"not null default 0 BIGINT(20) original_object_id"`
	}
	return x.Sync(new(Activity), new(Reversion))
}
