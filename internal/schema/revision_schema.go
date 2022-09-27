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
}

// GetRevisionListReq get revision list all request
type GetRevisionListReq struct {
	// object id
	ObjectID string `validate:"required" comment:"object_id" form:"object_id"`
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
