package schema

import (
	"github.com/answerdev/answer/internal/base/validator"
	"github.com/answerdev/answer/pkg/converter"
)

// RemoveAnswerReq delete answer request
type RemoveAnswerReq struct {
	// answer id
	ID string `validate:"required" json:"id"`
	// user id
	UserID string `json:"-"`
	// whether user can delete it
	CanDelete   bool   `json:"-"`
	CaptchaID   string `json:"captcha_id"` // captcha_id
	CaptchaCode string `json:"captcha_code"`
}

const (
	AnswerAcceptedFailed = 1
	AnswerAcceptedEnable = 2
)

type AnswerAddReq struct {
	QuestionID  string `json:"question_id"`
	Content     string `validate:"required,notblank,gte=6,lte=65535" json:"content"`
	HTML        string `json:"-"`
	UserID      string `json:"-"`
	CanEdit     bool   `json:"-"`
	CanDelete   bool   `json:"-"`
	CaptchaID   string `json:"captcha_id"` // captcha_id
	CaptchaCode string `json:"captcha_code"`
}

func (req *AnswerAddReq) Check() (errFields []*validator.FormErrorField, err error) {
	req.HTML = converter.Markdown2HTML(req.Content)
	return nil, nil
}

type AnswerUpdateReq struct {
	ID           string `json:"id"`
	QuestionID   string `json:"question_id"`
	Title        string `json:"title"`
	Content      string `validate:"required,notblank,gte=6,lte=65535" json:"content"`
	EditSummary  string `validate:"omitempty" json:"edit_summary"`
	HTML         string `json:"-"`
	UserID       string `json:"-"`
	NoNeedReview bool   `json:"-"`
	// whether user can edit it
	CanEdit     bool   `json:"-"`
	CaptchaID   string `json:"captcha_id"` // captcha_id
	CaptchaCode string `json:"captcha_code"`
}

func (req *AnswerUpdateReq) Check() (errFields []*validator.FormErrorField, err error) {
	req.HTML = converter.Markdown2HTML(req.Content)
	return nil, nil
}

// AnswerUpdateResp answer update resp
type AnswerUpdateResp struct {
	WaitForReview bool `json:"wait_for_review"`
}

type AnswerListReq struct {
	QuestionID string `json:"question_id" form:"question_id"` // question_id
	Order      string `json:"order" form:"order"`             // 1 Default 2 time
	Page       int    `json:"page" form:"page"`               // Query number of pages
	PageSize   int    `json:"page_size" form:"page_size"`     // Search page size
	UserID     string `json:"-" `
	IsAdmin    bool   `json:"-"`
	// whether user can edit it
	CanEdit bool `json:"-"`
	// whether user can delete it
	CanDelete bool `json:"-"`
}

type AnswerInfo struct {
	ID             string         `json:"id" xorm:"id"`                   // id
	QuestionID     string         `json:"question_id" xorm:"question_id"` // question_id
	Content        string         `json:"content" xorm:"content"`         // content
	HTML           string         `json:"html" xorm:"html"`               // html
	CreateTime     int64          `json:"create_time" xorm:"created"`     // create_time
	UpdateTime     int64          `json:"update_time" xorm:"updated"`     // update_time
	Accepted       int            `json:"accepted"`                       // 1 Failed 2 accepted
	UserID         string         `json:"-" `
	UpdateUserID   string         `json:"-" `
	UserInfo       *UserBasicInfo `json:"user_info,omitempty"`
	UpdateUserInfo *UserBasicInfo `json:"update_user_info,omitempty"`
	Collected      bool           `json:"collected"`
	VoteStatus     string         `json:"vote_status"`
	VoteCount      int            `json:"vote_count"`
	QuestionInfo   *QuestionInfo  `json:"question_info,omitempty"`
	Status         int            `json:"status"`

	// MemberActions
	MemberActions []*PermissionMemberAction `json:"member_actions"`
}

type AdminAnswerInfo struct {
	ID           string         `json:"id"`
	QuestionID   string         `json:"question_id"`
	Description  string         `json:"description"`
	CreateTime   int64          `json:"create_time"`
	UpdateTime   int64          `json:"update_time"`
	Accepted     int            `json:"accepted"`
	UserID       string         `json:"-" `
	UpdateUserID string         `json:"-" `
	UserInfo     *UserBasicInfo `json:"user_info"`
	VoteCount    int            `json:"vote_count"`
	QuestionInfo struct {
		Title string `json:"title"`
	} `json:"question_info"`
}

type AnswerAcceptedReq struct {
	QuestionID string `json:"question_id"`
	AnswerID   string `json:"answer_id"`
	UserID     string `json:"-" `
}

type AdminSetAnswerStatusRequest struct {
	StatusStr string `json:"status"`
	AnswerID  string `json:"answer_id"`
	UserID    string `json:"-" `
}
