package entity

import "time"

const (
	AnswerSearchOrderByDefault = "default"
	AnswerSearchOrderByTime    = "updated"
	AnswerSearchOrderByVote    = "vote"

	AnswerStatusAvailable = 1
	AnswerStatusDeleted   = 10
)

var AdminAnswerSearchStatus = map[string]int{
	"available": AnswerStatusAvailable,
	"deleted":   AnswerStatusDeleted,
}

// Answer answer
type Answer struct {
	ID             string    `xorm:"not null pk autoincr BIGINT(20) id"`
	CreatedAt      time.Time `xorm:"created not null default CURRENT_TIMESTAMP TIMESTAMP created_at"`
	UpdatedAt      time.Time `xorm:"updated_at TIMESTAMP"`
	QuestionID     string    `xorm:"not null default 0 BIGINT(20) question_id"`
	UserID         string    `xorm:"not null default 0 BIGINT(20) INDEX user_id"`
	LastEditUserID string    `xorm:"not null default 0 BIGINT(20) last_edit_user_id"`
	OriginalText   string    `xorm:"not null MEDIUMTEXT original_text"`
	ParsedText     string    `xorm:"not null MEDIUMTEXT parsed_text"`
	Status         int       `xorm:"not null default 1 INT(11) status"`
	Adopted        int       `xorm:"not null default 1 INT(11) adopted"`
	CommentCount   int       `xorm:"not null default 0 INT(11) comment_count"`
	VoteCount      int       `xorm:"not null default 0 INT(11) vote_count"`
	RevisionID     string    `xorm:"not null default 0 BIGINT(20) revision_id"`
}

type AnswerSearch struct {
	Answer
	Order    string `json:"order_by" `                  // default or updated
	Page     int    `json:"page" form:"page"`           // Query number of pages
	PageSize int    `json:"page_size" form:"page_size"` // Search page size
}

type AdminAnswerSearch struct {
	Page       int    `json:"page" form:"page"`           // Query number of pages
	PageSize   int    `json:"page_size" form:"page_size"` // Search page size
	Status     int    `json:"-" form:"-"`
	StatusStr  string `json:"status" form:"status"`                                             // Status 1 Available 2 closed 10 Deleted
	Query      string `validate:"omitempty,gt=0,lte=100" json:"query" form:"query" `            //Query string
	QuestionID string `validate:"omitempty,gt=0,lte=24" json:"question_id" form:"question_id" ` //Query string
}

// TableName answer table name
func (Answer) TableName() string {
	return "answer"
}
