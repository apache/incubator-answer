package entity

import "time"

const (
	TagStatusAvailable = 1
	TagStatusDeleted   = 10
)

// Tag tag
type Tag struct {
	ID              string    `xorm:"not null comment('tag_id') BIGINT(20) id"`
	CreatedAt       time.Time `xorm:"created comment('create time') TIMESTAMP created_at"`
	UpdatedAt       time.Time `xorm:"updated comment('update time') TIMESTAMP updated_at"`
	MainTagID       int64     `xorm:"not null default 0 comment('main tag id') BIGINT(20) main_tag_id"`
	MainTagSlugName string    `xorm:"default '' comment('main_tag_slug_name') VARCHAR(50) main_tag_slug_name"`
	SlugName        string    `xorm:"default '' comment('slug_name') unique VARCHAR(50) slug_name"`
	DisplayName     string    `xorm:"not null default '' comment('display_name') VARCHAR(50) display_name"`
	OriginalText    string    `xorm:"not null comment('original comment content') MEDIUMTEXT original_text"`
	ParsedText      string    `xorm:"not null comment('parsed comment content') MEDIUMTEXT parsed_text"`
	FollowCount     int       `xorm:"not null default 0 comment('associated follow count') INT(11) follow_count"`
	QuestionCount   int       `xorm:"not null default 0 comment('associated question count') INT(11) question_count"`
	Status          int       `xorm:"not null default 1 comment('tag status(available: 1; deleted: 10)') INT(11) status"`
	RevisionID      string    `xorm:"not null default 0 BIGINT(20) revision_id"`
}

// TableName tag table name
func (Tag) TableName() string {
	return "tag"
}
