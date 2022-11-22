package entity

import (
	"time"
)

const (
	QuestionStatusAvailable = 1
	QuestionStatusClosed    = 2
	QuestionStatusDeleted   = 10
)

var CmsQuestionSearchStatus = map[string]int{
	"available": QuestionStatusAvailable,
	"closed":    QuestionStatusClosed,
	"deleted":   QuestionStatusDeleted,
}

var CmsQuestionSearchStatusIntToString = map[int]string{
	QuestionStatusAvailable: "available",
	QuestionStatusClosed:    "closed",
	QuestionStatusDeleted:   "deleted",
}

type QuestionTag struct {
	Question `xorm:"extends"`
	TagRel   `xorm:"extends"`
}

// Question question
type Question struct {
	ID               string    `xorm:"not null pk BIGINT(20) id"`
	CreatedAt        time.Time `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP created_at"`
	UpdatedAt        time.Time `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP updated_at"`
	UserID           string    `xorm:"not null default 0 BIGINT(20) INDEX user_id"`
	Title            string    `xorm:"not null default '' VARCHAR(150) title"`
	OriginalText     string    `xorm:"not null MEDIUMTEXT original_text"`
	ParsedText       string    `xorm:"not null MEDIUMTEXT parsed_text"`
	Status           int       `xorm:"not null default 1 INT(11) status"`
	ViewCount        int       `xorm:"not null default 0 INT(11) view_count"`
	UniqueViewCount  int       `xorm:"not null default 0 INT(11) unique_view_count"`
	VoteCount        int       `xorm:"not null default 0 INT(11) vote_count"`
	AnswerCount      int       `xorm:"not null default 0 INT(11) answer_count"`
	CollectionCount  int       `xorm:"not null default 0 INT(11) collection_count"`
	FollowCount      int       `xorm:"not null default 0 INT(11) follow_count"`
	AcceptedAnswerID string    `xorm:"not null default 0 BIGINT(20) accepted_answer_id"`
	LastAnswerID     string    `xorm:"not null default 0 BIGINT(20) last_answer_id"`
	PostUpdateTime   time.Time `xorm:"default CURRENT_TIMESTAMP TIMESTAMP post_update_time"`
	RevisionID       string    `xorm:"not null default 0 BIGINT(20) revision_id"`
}

// TableName question table name
func (Question) TableName() string {
	return "question"
}

// QuestionWithTagsRevision question
type QuestionWithTagsRevision struct {
	Question
	Tags []*TagSimpleInfoForRevision `json:"tags"`
}

// TagSimpleInfoForRevision tag simple info for revision
type TagSimpleInfoForRevision struct {
	ID          string `xorm:"not null pk comment('tag_id') BIGINT(20) id"`
	MainTagID   int64  `xorm:"not null default 0 BIGINT(20) main_tag_id"`
	SlugName    string `xorm:"not null default '' unique VARCHAR(35) slug_name"`
	DisplayName string `xorm:"not null default '' VARCHAR(35) display_name"`
	Recommend   bool   `xorm:"not null default false BOOL recommend"`
	Reserved    bool   `xorm:"not null default false BOOL reserved"`
	RevisionID  string `xorm:"not null default 0 BIGINT(20) revision_id"`
}
