package entity

import "time"

// UserExternalLogin user external login
type UserExternalLogin struct {
	ID         int64     `xorm:"not null pk autoincr BIGINT(20) id"`
	CreatedAt  time.Time `xorm:"created TIMESTAMP created_at"`
	UpdatedAt  time.Time `xorm:"updated TIMESTAMP updated_at"`
	UserID     string    `xorm:"not null default 0 BIGINT(20) user_id"`
	Provider   string    `xorm:"not null default '' VARCHAR(100) provider"`
	ExternalID string    `xorm:"not null default '' VARCHAR(128) external_id"`
	MetaInfo   string    `xorm:"TEXT meta_info"`
}

// TableName  table name
func (UserExternalLogin) TableName() string {
	return "user_external_login"
}
