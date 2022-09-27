package entity

import "time"

const (
	TagRelStatusAvailable = 1
	TagRelStatusDeleted   = 10
)

// TagRel tag relation
type TagRel struct {
	ID        int64     `xorm:"not null pk autoincr comment('tag_list_id') BIGINT(20) id"`
	CreatedAt time.Time `xorm:"created comment('create time') TIMESTAMP created_at"`
	UpdatedAt time.Time `xorm:"updated comment('update time') TIMESTAMP updated_at"`
	ObjectID  string    `xorm:"not null comment('object_id') index BIGINT(20) object_id"`
	TagID     string    `xorm:"not null comment('tag_id') index BIGINT(20) tag_id"`
	Status    int       `xorm:"not null default 1 comment('tag_list_status(available: 1; deleted: 10)') INT(11) status"`
}

// TableName tag list table name
func (TagRel) TableName() string {
	return "tag_rel"
}
