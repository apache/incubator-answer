package entity

import "time"

// UserRoleRel role
type UserRoleRel struct {
	ID        int       `xorm:"not null pk autoincr INT(11) id"`
	CreatedAt time.Time `xorm:"created TIMESTAMP created_at"`
	UpdatedAt time.Time `xorm:"updated TIMESTAMP updated_at"`
	UserID    string    `xorm:"not null default 0 BIGINT(20) user_id"`
	RoleID    int       `xorm:"not null default 0 INT(11) role_id"`
}

// TableName user role rel table name
func (UserRoleRel) TableName() string {
	return "user_role_rel"
}
