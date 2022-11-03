package entity

import "time"

// Revision revision
type Revision struct {
	ID         string    `xorm:"not null pk autoincr BIGINT(20) id"`
	CreatedAt  time.Time `xorm:"created TIMESTAMP created_at"`
	UpdatedAt  time.Time `xorm:"updated TIMESTAMP updated_at"`
	UserID     string    `xorm:"not null default 0 BIGINT(20) user_id"`
	ObjectType int       `xorm:"not null default 0 ) INT(11) object_type"`
	ObjectID   string    `xorm:"not null default 0 BIGINT(20) INDEX object_id"`
	Title      string    `xorm:"not null default '' VARCHAR(255) title"`
	Content    string    `xorm:"not null TEXT content"`
	Log        string    `xorm:"VARCHAR(255) log"`
	// Status todo: this field is not used, will be removed in the future
	Status int `xorm:"not null default 1 INT(11) status"`
}

// TableName revision table name
func (Revision) TableName() string {
	return "revision"
}
