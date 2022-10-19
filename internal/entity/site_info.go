package entity

import "time"

// SiteInfo site information setting
type SiteInfo struct {
	ID        string    `xorm:"not null pk autoincr INT(11) id"`
	CreatedAt time.Time `xorm:"created TIMESTAMP created_at"`
	UpdatedAt time.Time `xorm:"updated TIMESTAMP updated_at"`
	Type      string    `xorm:"not null VARCHAR(64) type"`
	Content   string    `xorm:"not null MEDIUMTEXT content"`
	Status    int       `xorm:"not null default 1 INT(11) status"`
}

// TableName table name
func (*SiteInfo) TableName() string {
	return "site_info"
}
