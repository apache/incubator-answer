package entity

import "time"

const (
	TagRelStatusAvailable = 1
	TagRelStatusHide      = 2
	TagRelStatusDeleted   = 10
)

// TagRel tag relation
type TagRel struct {
	ID        int64     `xorm:"not null pk autoincr BIGINT(20) id"`
	CreatedAt time.Time `xorm:"created TIMESTAMP created_at"`
	UpdatedAt time.Time `xorm:"updated TIMESTAMP updated_at"`
	ObjectID  string    `xorm:"not null INDEX UNIQUE(s) BIGINT(20) object_id"`
	TagID     string    `xorm:"not null INDEX UNIQUE(s) BIGINT(20) tag_id"`
	Status    int       `xorm:"not null default 1 INT(11) status"`
}

// TableName tag list table name
func (TagRel) TableName() string {
	return "tag_rel"
}
