package entity

import "time"

// UserNotificationConfig user notification config
type UserNotificationConfig struct {
	ID        string    `xorm:"not null pk autoincr BIGINT(20) id"`
	CreatedAt time.Time `xorm:"created TIMESTAMP created_at"`
	UpdatedAt time.Time `xorm:"updated TIMESTAMP updated_at"`
	UserID    string    `xorm:"not null default 0 INDEX UNIQUE(uk_us) BIGINT(20) INDEX user_id"`
	Source    string    `xorm:"not null default '' INDEX UNIQUE(uk_us) VARCHAR(64) source"`
	Channels  string    `xorm:"not null TEXT channels"`
	Enabled   bool      `xorm:"not null default false BOOL enabled"`
}

// TableName notification table name
func (UserNotificationConfig) TableName() string {
	return "user_notification_config"
}
