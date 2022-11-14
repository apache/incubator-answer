package entity

import "time"

const (
	TagStatusAvailable = 1
	TagStatusDeleted   = 10
)

// Tag tag
type Tag struct {
	ID              string    `xorm:"not null pk comment('tag_id') BIGINT(20) id"`
	CreatedAt       time.Time `xorm:"created TIMESTAMP created_at"`
	UpdatedAt       time.Time `xorm:"updated TIMESTAMP updated_at"`
	MainTagID       int64     `xorm:"not null default 0 BIGINT(20) main_tag_id"`
	MainTagSlugName string    `xorm:"not null default '' VARCHAR(35) main_tag_slug_name"`
	SlugName        string    `xorm:"not null default '' unique VARCHAR(35) slug_name"`
	DisplayName     string    `xorm:"not null default '' VARCHAR(35) display_name"`
	OriginalText    string    `xorm:"not null MEDIUMTEXT original_text"`
	ParsedText      string    `xorm:"not null MEDIUMTEXT parsed_text"`
	FollowCount     int       `xorm:"not null default 0 INT(11) follow_count"`
	QuestionCount   int       `xorm:"not null default 0 INT(11) question_count"`
	Status          int       `xorm:"not null default 1 INT(11) status"`
	Recommend       bool      `xorm:"not null default false BOOL recommend"`
	RevisionID      string    `xorm:"not null default 0 BIGINT(20) revision_id"`
}

// TableName tag table name
func (Tag) TableName() string {
	return "tag"
}
