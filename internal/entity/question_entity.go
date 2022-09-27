package entity

import (
	"time"
)

const (
	QuestionStatusAvailable = 1
	QuestionStatusclosed    = 2
	QuestionStatusDeleted   = 10
)

var CmsQuestionSearchStatus = map[string]int{
	"available": QuestionStatusAvailable,
	"closed":    QuestionStatusclosed,
	"deleted":   QuestionStatusDeleted,
}

type QuestionTag struct {
	Question `xorm:"extends"`
	TagRel   `xorm:"extends"`
}

// Question question
type Question struct {
	ID               string    `xorm:"not null pk comment('question id') BIGINT(20) id"`
	UserID           string    `xorm:"not null default 0 comment('user id') BIGINT(20) user_id"`
	Title            string    `xorm:"not null default '' comment('question title') VARCHAR(255) title"`
	OriginalText     string    `xorm:"not null comment('original content') MEDIUMTEXT original_text"`
	ParsedText       string    `xorm:"not null comment('parsed content') MEDIUMTEXT parsed_text"`
	Status           int       `xorm:"not null default 1 comment(' question status(available: 1; deleted: 10)') INT(11) status"`
	ViewCount        int       `xorm:"not null default 0 comment('view count') INT(11) view_count"`
	UniqueViewCount  int       `xorm:"not null default 0 comment('unique view count') INT(11) unique_view_count"`
	VoteCount        int       `xorm:"not null default 0 comment('vote count') INT(11) vote_count"`
	AnswerCount      int       `xorm:"not null default 0 comment('answer count') INT(11) answer_count"`
	CollectionCount  int       `xorm:"not null default 0 comment('collection count') INT(11) collection_count"`
	FollowCount      int       `xorm:"not null default 0 comment('follow count') INT(11) follow_count"`
	AcceptedAnswerID string    `xorm:"not null default 0 comment('accepted answer id') BIGINT(20) accepted_answer_id"`
	LastAnswerID     string    `xorm:"not null default 0 comment('last answer id') BIGINT(20) last_answer_id"`
	CreatedAt        time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('create time') TIMESTAMP created_at"`
	UpdatedAt        time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('update time') TIMESTAMP updated_at"`
	PostUpdateTime   time.Time `xorm:"default CURRENT_TIMESTAMP comment('answer the last update time') TIMESTAMP post_update_time"`
	RevisionID       string    `xorm:"not null default 0 BIGINT(20) revision_id"`
}

// TableName question table name
func (Question) TableName() string {
	return "question"
}
