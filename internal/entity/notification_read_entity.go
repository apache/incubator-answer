package entity

import "time"

// NotificationRead notification read record
type NotificationRead struct {
	ID        int       `xorm:"not null pk autoincr comment('id') INT(11) id"`
	CreatedAt time.Time `xorm:"created comment('create time') TIMESTAMP created_at"`
	UpdatedAt time.Time `xorm:"updated comment('update time') TIMESTAMP updated_at"`
	UserID    int64     `xorm:"not null default 0 comment('user id') BIGINT(20) user_id"`
	MessageID int64     `xorm:"not null default 0 comment('message id') BIGINT(20) message_id"`
	IsRead    int       `xorm:"not null default 1 comment('read status(unread: 1; read 2)') INT(11) is_read"`
}

// TableName notification read record table name
func (NotificationRead) TableName() string {
	return "notification_read"
}
