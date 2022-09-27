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
	ID             string        `xorm:"not null pk autoincr comment('comment id') BIGINT(20) id"`
	CreatedAt      time.Time     `xorm:"created comment('create time') TIMESTAMP created_at"`
	UpdatedAt      time.Time     `xorm:"updated comment('update time') TIMESTAMP updated_at"`
	UserID         string        `xorm:"not null default 0 comment('user id') BIGINT(20) user_id"`
	ReplyUserID    sql.NullInt64 `xorm:"comment('reply user id') BIGINT(20) reply_user_id"`
	ReplyCommentID sql.NullInt64 `xorm:"comment('reply comment id') BIGINT(20) reply_comment_id"`
	ObjectID       string        `xorm:"not null default 0 comment('object id') BIGINT(20) object_id"`
	QuestionID     string        `xorm:"not null default 0 comment('question id') BIGINT(20) question_id"`
	VoteCount      int           `xorm:"not null default 0 comment('user vote amount') INT(11) vote_count"`
	Status         int           `xorm:"not null default 0 comment('comment status(available: 1; deleted: 10)') TINYINT(4) status"`
	OriginalText   string        `xorm:"not null comment('original comment content') MEDIUMTEXT original_text"`
	ParsedText     string        `xorm:"not null comment('parsed comment content') MEDIUMTEXT parsed_text"`
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
