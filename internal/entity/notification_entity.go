package entity

import "time"

// Notification notification
type Notification struct {
	ID        string    `xorm:"not null pk autoincr BIGINT(20) id"`
	CreatedAt time.Time `xorm:"created TIMESTAMP created_at"`
	UpdatedAt time.Time `xorm:"TIMESTAMP updated_at"`
	UserID    string    `xorm:"not null default 0 BIGINT(20) INDEX user_id"`
	ObjectID  string    `xorm:"not null default 0 INDEX BIGINT(20) object_id"`
	Content   string    `xorm:"not null TEXT content"`
	Type      int       `xorm:"not null default 0 INT(11) type"`
	MsgType   int       `xorm:"not null default 0 INT(11) msg_type"`
	IsRead    int       `xorm:"not null default 1 INT(11) is_read"`
	Status    int       `xorm:"not null default 1 INT(11) status"`
}

// TableName notification table name
func (Notification) TableName() string {
	return "notification"
}
