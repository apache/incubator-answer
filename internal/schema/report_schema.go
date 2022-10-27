package schema

import (
	"time"

	"github.com/answerdev/answer/internal/base/constant"
)

// AddReportReq add report request
type AddReportReq struct {
	// object id
	ObjectID string `validate:"required,gt=0,lte=20" json:"object_id"`
	// report type
	ReportType int `validate:"required" json:"report_type"`
	// report content
	Content string `validate:"omitempty,gt=0,lte=500" json:"content"`
	// user id
	UserID string `json:"-"`
}

// GetReportListReq get report list all request
type GetReportListReq struct {
	// report source
	Source string `validate:"required,oneof=question answer comment" form:"source"`
}

// GetReportTypeResp get report response
type GetReportTypeResp struct {
	// report name
	Name string `json:"name"`
	// report description
	Description string `json:"description"`
	// report source
	Source string `json:"source"`
	// report type
	Type int `json:"type"`
	// is have content
	HaveContent bool `json:"have_content"`
	// content type
	ContentType string `json:"content_type"`
}

// ReportHandleReq request handle request
type ReportHandleReq struct {
	ID             string `validate:"required" comment:"report id" form:"id" json:"id"`
	FlaggedType    int    `validate:"required" comment:"flagged type" form:"flagged_type" json:"flagged_type"`
	FlaggedContent string `validate:"omitempty" comment:"flagged content" form:"flagged_content" json:"flagged_content"`
}

// GetReportListPageDTO report list data transfer object
type GetReportListPageDTO struct {
	ObjectType string
	Status     string
	Page       int
	PageSize   int
}

// GetReportListPageResp get report list
type GetReportListPageResp struct {
	ID           string         `json:"id"`
	ReportedUser *UserBasicInfo `json:"reported_user"`
	ReportUser   *UserBasicInfo `json:"report_user"`

	Content        string `json:"content"`
	FlaggedContent string `json:"flagged_content"`
	OType          string `json:"object_type"`

	ObjectID   string `json:"-"`
	QuestionID string `json:"question_id"`
	AnswerID   string `json:"answer_id"`
	CommentID  string `json:"comment_id"`

	Title   string `json:"title"`
	Excerpt string `json:"excerpt"`

	// create time
	CreatedAt       time.Time `json:"-"`
	CreatedAtParsed int64     `json:"created_at"`

	UpdatedAt       time.Time `json:"_"`
	UpdatedAtParsed int64     `json:"updated_at"`

	Reason        *ReasonItem `json:"reason"`
	FlaggedReason *ReasonItem `json:"flagged_reason"`

	UserID         string `json:"-"`
	ReportedUserID string `json:"-"`
	Status         int    `json:"-"`
	ObjectType     int    `json:"-"`
	ReportType     int    `json:"-"`
	FlaggedType    int    `json:"-"`
}

// Format format result
func (r *GetReportListPageResp) Format() {
	r.OType = constant.ObjectTypeNumberMapping[r.ObjectType]

	r.CreatedAtParsed = r.CreatedAt.Unix()
	r.UpdatedAtParsed = r.UpdatedAt.Unix()
}
