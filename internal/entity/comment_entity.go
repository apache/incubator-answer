package entity

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/segmentfault/answer/pkg/converter"
)

const (
	CommentStatusAvailable = 1
	CommentStatusDeleted   = 10
)

// Comment comment
type Comment struct {
	ID             string        `xorm:"not null pk autoincr BIGINT(20) id"`
	CreatedAt      time.Time     `xorm:"created TIMESTAMP created_at"`
	UpdatedAt      time.Time     `xorm:"updated TIMESTAMP updated_at"`
	UserID         string        `xorm:"not null default 0 BIGINT(20) user_id"`
	ReplyUserID    sql.NullInt64 `xorm:"BIGINT(20) reply_user_id"`
	ReplyCommentID sql.NullInt64 `xorm:"BIGINT(20) reply_comment_id"`
	ObjectID       string        `xorm:"not null default 0 BIGINT(20) INDEX object_id"`
	QuestionID     string        `xorm:"not null default 0 BIGINT(20) question_id"`
	VoteCount      int           `xorm:"not null default 0 INT(11) vote_count"`
	Status         int           `xorm:"not null default 0 TINYINT(4) status"`
	OriginalText   string        `xorm:"not null MEDIUMTEXT original_text"`
	ParsedText     string        `xorm:"not null MEDIUMTEXT parsed_text"`
}

// TableName comment table name
func (c *Comment) TableName() string {
	return "comment"
}

// GetReplyUserID get reply user id
func (c *Comment) GetReplyUserID() string {
	if c.ReplyUserID.Valid {
		return fmt.Sprintf("%d", c.ReplyUserID.Int64)
	}
	return ""
}

// GetReplyCommentID get reply comment id
func (c *Comment) GetReplyCommentID() string {
	if c.ReplyCommentID.Valid {
		return fmt.Sprintf("%d", c.ReplyCommentID.Int64)
	}
	return ""
}

// SetReplyUserID set reply user id
func (c *Comment) SetReplyUserID(str string) {
	if len(str) > 0 {
		c.ReplyUserID = sql.NullInt64{Int64: converter.StringToInt64(str), Valid: true}
	} else {
		c.ReplyUserID = sql.NullInt64{Valid: false}
	}
}

// SetReplyCommentID set reply comment id
func (c *Comment) SetReplyCommentID(str string) {
	if len(str) > 0 {
		c.ReplyCommentID = sql.NullInt64{Int64: converter.StringToInt64(str), Valid: true}
	} else {
		c.ReplyCommentID = sql.NullInt64{Valid: false}
	}
}
