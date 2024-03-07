package schema

import (
	"github.com/apache/incubator-answer/internal/base/validator"
	"github.com/apache/incubator-answer/pkg/uid"
)

// UpdateReviewReq update review request
type UpdateReviewReq struct {
	ReviewID int    `validate:"required" json:"review_id"`
	Status   string `validate:"required,oneof=approve reject" json:"status"`
	UserID   string `json:"-"`
}

func (r *UpdateReviewReq) IsApprove() bool {
	return r.Status == "approve"
}

func (r *UpdateReviewReq) IsReject() bool {
	return r.Status == "reject"
}

// GetUnreviewedPostPageReq get review page request
type GetUnreviewedPostPageReq struct {
	ObjectID        string            `validate:"omitempty" form:"object_id"`
	Page            int               `validate:"omitempty" form:"page"`
	ReviewerMapping map[string]string `json:"-"`
	UserID          string            `json:"-"`
}

func (r *GetUnreviewedPostPageReq) Check() (errField []*validator.FormErrorField, err error) {
	if len(r.ObjectID) > 0 {
		r.Page = 1
		r.ObjectID = uid.DeShortID(r.ObjectID)
	}
	return
}

// GetUnreviewedPostPageResp get review page response
type GetUnreviewedPostPageResp struct {
	ReviewID             int           `json:"review_id"`
	CreatedAt            int64         `json:"created_at"`
	ObjectID             string        `json:"object_id"`
	ObjectType           string        `json:"object_type" enums:"question,answer,comment"`
	Title                string        `json:"title"`
	ParsedText           string        `json:"parsed_text"`
	Tags                 []*TagResp    `json:"tags"`
	AuthorUserID         string        `json:"author_user_id"`
	UserInfo             UserBasicInfo `json:"user_info"`
	SubmitAt             int64         `json:"submit_at"`
	SubmitterDisplayName string        `json:"submitter_display_name"`
}
