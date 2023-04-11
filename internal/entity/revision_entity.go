package entity

import (
	"time"
)

const (
	// RevisioNnormalStatus this revision is Nnormal
	RevisioNnormalStatus = 0
	// RevisionUnreviewedStatus this revision is unreviewed
	RevisionUnreviewedStatus = 1
	// RevisionReviewPassStatus this revision is reviewed and approved by operator
	RevisionReviewPassStatus = 2
	// RevisionReviewRejectStatus this revision is reviewed and rejected by operator
	RevisionReviewRejectStatus = 3
)

// Revision revision
type Revision struct {
	ID           string    `xorm:"not null pk autoincr BIGINT(20) id"`
	CreatedAt    time.Time `xorm:"created TIMESTAMP created_at"`
	UpdatedAt    time.Time `xorm:"updated TIMESTAMP updated_at"`
	UserID       string    `xorm:"not null default 0 BIGINT(20) user_id"`
	ObjectType   int       `xorm:"not null default 0 INT(11) object_type"`
	ObjectID     string    `xorm:"not null default 0 BIGINT(20) INDEX object_id"`
	Title        string    `xorm:"not null default '' VARCHAR(255) title"`
	Content      string    `xorm:"not null TEXT content"`
	Log          string    `xorm:"VARCHAR(255) log"`
	Status       int       `xorm:"not null default 1 INT(11) status"`
	ReviewUserID int64     `xorm:"not null default 0 BIGINT(20) review_user_id"`
}

// TableName revision table name
func (Revision) TableName() string {
	return "revision"
}
