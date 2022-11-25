package schema

import (
	"time"
)

// AddRevisionDTO add revision request
type AddRevisionDTO struct {
	// user id
	UserID string
	// object id
	ObjectID string
	// title
	Title string
	// content
	Content string
	// log
	Log string
	// status
	Status int
}

// GetRevisionListReq get revision list all request
type GetRevisionListReq struct {
	// object id
	ObjectID string `validate:"required" comment:"object_id" form:"object_id"`
}

const RevisionAuditApprove = "approve"
const RevisionAuditReject = "reject"

type RevisionAuditReq struct {
	// object id
	ID                string `validate:"required" comment:"id" form:"id"`
	Operation         string `validate:"required" comment:"operation" form:"operation"` //approve or reject
	UserID            string `json:"-"`
	CanReviewQuestion bool   `json:"-"`
	CanReviewAnswer   bool   `json:"-"`
	CanReviewTag      bool   `json:"-"`
}

type RevisionSearch struct {
	Page              int    `json:"page" form:"page"` // Query number of pages
	CanReviewQuestion bool   `json:"-"`
	CanReviewAnswer   bool   `json:"-"`
	CanReviewTag      bool   `json:"-"`
	UserID            string `json:"-"`
}

type GetUnreviewedRevisionResp struct {
	Type           string                      `json:"type"`
	Info           *UnreviewedRevisionInfoInfo `json:"info"`
	UnreviewedInfo *GetRevisionResp            `json:"unreviewed_info"`
}

// GetRevisionResp get revision response
type GetRevisionResp struct {
	// id
	ID string `json:"id"`
	// user id
	UserID string `json:"use_id"`
	// object id
	ObjectID string `json:"object_id"`
	// object type
	ObjectType int `json:"-"`
	// title
	Title string `json:"title"`
	// content
	Content string `json:"-"`
	// content parsed
	ContentParsed interface{} `json:"content"`
	// revision status(normal: 1; delete 2)
	Status int `json:"status"`
	// create time
	CreatedAt       time.Time     `json:"-"`
	CreatedAtParsed int64         `json:"create_at"`
	UserInfo        UserBasicInfo `json:"user_info"`
	Log             string        `json:"reason"`
}
