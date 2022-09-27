package entity

import "time"

const (
	ActivityAvailable = 0
	ActivityCancelled = 1
)

// Activity activity
type Activity struct {
	ID            string    `xorm:"not null pk autoincr comment('Activity TagID autoincrement') BIGINT(20) id"`
	CreatedAt     time.Time `xorm:"created comment('create time') TIMESTAMP created_at"`
	UpdatedAt     time.Time `xorm:"updated comment('update time') TIMESTAMP updated_at"`
	UserID        string    `xorm:"not null comment('the user ID that generated the activity or affected by the activity') index BIGINT(20) user_id"`
	TriggerUserID int64     `xorm:"not null default 0 comment('the trigger user TagID that generated the activity or affected by the activity') index BIGINT(20) trigger_user_id"`
	ObjectID      string    `xorm:"not null default 0 comment('the object TagID that affected by the activity') index BIGINT(20) object_id"`
	ActivityType  int       `xorm:"not null comment('activity type, correspond to config id') INT(11) activity_type"`
	Cancelled     int       `xorm:"not null default 0 comment('mark this activity if cancelled or not,default 0(not cancelled)') TINYINT(4) cancelled"`
	Rank          int       `xorm:"not null default 0 comment('rank of current operating user affected') INT(11) rank"`
	HasRank       int       `xorm:"not null default 0 comment('this activity has rank or not') TINYINT(4) has_rank"`
}

type ActivityRunkSum struct {
	Rank int `xorm:"not null default 0 comment('rank of current operating user affected') INT(11) rank"`
}

// TableName activity table name
func (Activity) TableName() string {
	return "activity"
}
