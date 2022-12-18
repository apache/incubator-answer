package entity

import "time"

// Power power
type Power struct {
	ID          int       `xorm:"not null pk autoincr INT(11) id"`
	CreatedAt   time.Time `xorm:"created TIMESTAMP created_at"`
	UpdatedAt   time.Time `xorm:"updated TIMESTAMP updated_at"`
	Name        string    `xorm:"not null default '' VARCHAR(50) name"`
	PowerType   string    `xorm:"not null default '' VARCHAR(100) power_type"`
	Description string    `xorm:"not null default '' VARCHAR(200) description"`
}

// TableName power table name
func (Power) TableName() string {
	return "power"
}
