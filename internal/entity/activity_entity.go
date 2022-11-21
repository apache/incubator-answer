package entity

import "time"

const (
	ActivityAvailable = 0
	ActivityCancelled = 1
)

// Activity activity
type Activity struct {
	ID            string    `xorm:"not null pk autoincr BIGINT(20) id"`
	CreatedAt     time.Time `xorm:"created TIMESTAMP created_at"`
	UpdatedAt     time.Time `xorm:"updated TIMESTAMP updated_at"`
	CancelledAt   time.Time `xorm:"TIMESTAMP cancelled_at"`
	UserID        string    `xorm:"not null index BIGINT(20) user_id"`
	TriggerUserID int64     `xorm:"not null default 0 index BIGINT(20) trigger_user_id"`
	ObjectID      string    `xorm:"not null default 0 index BIGINT(20) object_id"`
	ActivityType  int       `xorm:"not null INT(11) activity_type"`
	Cancelled     int       `xorm:"not null default 0 TINYINT(4) cancelled"`
	Rank          int       `xorm:"not null default 0 INT(11) rank"`
	HasRank       int       `xorm:"not null default 0 TINYINT(4) has_rank"`
}

type ActivityRankSum struct {
	Rank int `xorm:"not null default 0 INT(11) rank"`
}

// TableName activity table name
func (Activity) TableName() string {
	return "activity"
}
