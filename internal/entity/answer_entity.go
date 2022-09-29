package entity

import "time"

const (
	Answer_Search_OrderBy_Default = "default"
	Answer_Search_OrderBy_Time    = "updated"
	Answer_Search_OrderBy_Vote    = "vote"

	AnswerStatusAvailable = 1
	AnswerStatusDeleted   = 10
)

var CmsAnswerSearchStatus = map[string]int{
	"available": AnswerStatusAvailable,
	"deleted":   AnswerStatusDeleted,
}

// Answer answer
type Answer struct {
	ID           string    `xorm:"not null pk autoincr comment('answer id') BIGINT(20) id"`
	CreatedAt    time.Time `xorm:"created not null default CURRENT_TIMESTAMP TIMESTAMP created_at"`
	UpdatedAt    time.Time `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP updated_at"`
	QuestionID   string    `xorm:"not null default 0 comment('question id') BIGINT(20) question_id"`
	UserID       string    `xorm:"not null default 0 comment('answer user id') BIGINT(20) user_id"`
	OriginalText string    `xorm:"not null comment('original content') MEDIUMTEXT original_text"`
	ParsedText   string    `xorm:"not null comment('parsed content') MEDIUMTEXT parsed_text"`
	Status       int       `xorm:"not null default 1 comment(' answer status(available: 1; deleted: 10)') INT(11) status"`
	Adopted      int       `xorm:"not null default 1 comment('adopted (1 failed 2 adopted)') INT(11) adopted"`
	CommentCount int       `xorm:"not null default 0 comment('comment count') INT(11) comment_count"`
	VoteCount    int       `xorm:"not null default 0 comment('vote count') INT(11) vote_count"`
	RevisionID   string    `xorm:"not null default 0 BIGINT(20) revision_id"`
}

type AnswerSearch struct {
	Answer
	Order    string `json:"order_by" `                  // default or updated
	Page     int    `json:"page" form:"page"`           //Query number of pages
	PageSize int    `json:"page_size" form:"page_size"` //Search page size
}

type CmsAnswerSearch struct {
	Page      int    `json:"page" form:"page"`           //Query number of pages
	PageSize  int    `json:"page_size" form:"page_size"` //Search page size
	Status    int    `json:"-" form:"-"`
	StatusStr string `json:"status" form:"status"` //Status 1 Available 2 closed 10 Deleted
}

type AdminSetAnswerStatusRequest struct {
	StatusStr string `json:"status" form:"status"`
	AnswerID  string `json:"answer_id" form:"answer_id"`
}

// TableName answer table name
func (Answer) TableName() string {
	return "answer"
}
