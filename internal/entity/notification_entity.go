package entity

import "time"

// Notification notification
type Notification struct {
	ID        string    `xorm:"not null pk autoincr comment('notification id') BIGINT(20) id"`
	CreatedAt time.Time `xorm:"created comment('create time') TIMESTAMP created_at"`
	UpdatedAt time.Time `xorm:"comment('update time') TIMESTAMP updated_at"`
	UserID    string    `xorm:"not null default 0 comment('user id') BIGINT(20) INDEX user_id"`
	ObjectID  string    `xorm:"not null default 0 comment('object id') INDEX BIGINT(20) object_id"`
	Content   string    `xorm:"not null comment('notification content') TEXT content"`
	Type      int       `xorm:"not null default 0 comment('notification type(1:inbox; 2:achievement)') INT(11) type"`
	IsRead    int       `xorm:"not null default 1 comment('read status(unread: 1; read 2)') INT(11) is_read"`
	Status    int       `xorm:"not null default 1 comment('notification status(normal: 1; delete 2)') INT(11) status"`
}

// TableName notification table name
func (Notification) TableName() string {
	return "notification"
}
