package entity

import "time"

// Revision revision
type Revision struct {
	ID         string    `xorm:"not null pk autoincr comment('id') BIGINT(20) id"`
	CreatedAt  time.Time `xorm:"created comment('create time') TIMESTAMP created_at"`
	UpdatedAt  time.Time `xorm:"updated comment('update time') TIMESTAMP updated_at"`
	UserID     string    `xorm:"not null default 0 comment('user id') BIGINT(20) user_id"`
	ObjectType int       `xorm:"not null default 0 comment('revision type(question: 1; answer 2; tag 3)') INT(11) object_type"`
	ObjectID   string    `xorm:"not null default 0 comment('object id') BIGINT(20) object_id"`
	Title      string    `xorm:"not null default '' comment('title') VARCHAR(255) title"`
	Content    string    `xorm:"not null comment('content') TEXT content"`
	Log        string    `xorm:"comment('log') VARCHAR(255) log"`
	Status     int       `xorm:"not null default 1 comment('revision status(normal: 1; delete 2)') INT(11) status"`
}

// TableName revision table name
func (Revision) TableName() string {
	return "revision"
}
