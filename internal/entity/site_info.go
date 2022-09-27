package entity

import "time"

// SiteInfo site information setting
type SiteInfo struct {
	ID        string    `xorm:"not null pk autoincr comment('id') INT(11) id"`
	CreatedAt time.Time `xorm:"created comment('create time') TIMESTAMP created_at"`
	UpdatedAt time.Time `xorm:"updated comment('update time') TIMESTAMP updated_at"`
	Type      string    `xorm:"not null comment('content') VARCHAR(64) type"`
	Content   string    `xorm:"not null comment('content') MEDIUMTEXT content"`
	Status    int       `xorm:"not null default 1 comment('site info status(available: 1; deleted: 10)') INT(11) status"`
}

// TableName table name
func (*SiteInfo) TableName() string {
	return "site_info"
}
