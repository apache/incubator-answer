package entity

import "time"

// CollectionGroup collection group
type CollectionGroup struct {
	ID           string    `xorm:"not null pk autoincr BIGINT(20) id"`
	CreatedAt    time.Time `xorm:"created not null default CURRENT_TIMESTAMP TIMESTAMP created_at"`
	UpdatedAt    time.Time `xorm:"updated not null default CURRENT_TIMESTAMP TIMESTAMP updated_at"`
	UserID       string    `xorm:"not null default 0 BIGINT(20) INDEX user_id"`
	Name         string    `xorm:"not null default '' VARCHAR(50) name"`
	DefaultGroup int       `xorm:"not null default 1 INT(11) default_group"`
}

// TableName collection group table name
func (CollectionGroup) TableName() string {
	return "collection_group"
}
