package entity

import "time"

// RolePowerRel role power rel
type RolePowerRel struct {
	ID        int       `xorm:"not null pk autoincr INT(11) id"`
	CreatedAt time.Time `xorm:"created TIMESTAMP created_at"`
	UpdatedAt time.Time `xorm:"updated TIMESTAMP updated_at"`
	RoleID    int       `xorm:"not null default 0 INT(11) role_id"`
	PowerType string    `xorm:"not null default '' VARCHAR(200) power_type"`
}

// TableName role power rel table name
func (RolePowerRel) TableName() string {
	return "role_power_rel"
}
