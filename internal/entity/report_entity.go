package entity

import "time"

const (
	ReportStatusPending   = 1
	ReportStatusCompleted = 2
	ReportStatusDeleted   = 10
)

var (
	ReportStatus = map[string]int{
		"pending":   ReportStatusPending,
		"completed": ReportStatusCompleted,
		"deleted":   ReportStatusDeleted,
	}
)

// Report report
type Report struct {
	ID             string    `xorm:"not null pk autoincr comment('id') BIGINT(20) id"`
	CreatedAt      time.Time `xorm:"created comment('create time') TIMESTAMP created_at"`
	UpdatedAt      time.Time `xorm:"updated comment('update time') TIMESTAMP updated_at"`
	UserID         string    `xorm:"not null comment('reporter user id') BIGINT(20) user_id"`
	ObjectID       string    `xorm:"not null comment('object id') BIGINT(20) object_id"`
	ReportedUserID string    `xorm:"not null default 0 comment('reported user id') BIGINT(20) reported_user_id"`
	ObjectType     int       `xorm:"not null default 0 comment('revision type') INT(11) object_type"`
	ReportType     int       `xorm:"not null default 0 comment('report type') INT(11) report_type"`
	Content        string    `xorm:"not null comment('report content') TEXT content"`
	FlaggedType    int       `xorm:"not null default 0 comment('flagged type') INT(11) flagged_type"`
	FlaggedContent string    `xorm:"comment('flagged content') TEXT flagged_content"`
	Status         int       `xorm:"not null default 1 comment('status(normal: 1; pending:2; delete: 10)') INT(11) status"`
}

// TableName report table name
func (Report) TableName() string {
	return "report"
}
