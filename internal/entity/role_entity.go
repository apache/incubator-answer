package entity

import "time"

// Role role
type Role struct {
	ID          int       `xorm:"not null pk autoincr INT(11) id"`
	CreatedAt   time.Time `xorm:"created TIMESTAMP created_at"`
	UpdatedAt   time.Time `xorm:"updated TIMESTAMP updated_at"`
	Name        string    `xorm:"not null default '' VARCHAR(50) name"`
	Description string    `xorm:"not null default '' VARCHAR(200) description"`
}

// TableName user table name
func (Role) TableName() string {
	return "role"
}
