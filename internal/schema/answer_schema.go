package schema

import (
	"time"
)

// RemoveAnswerReq delete answer request
type RemoveAnswerReq struct {
	// answer id
	ID string `validate:"required" comment:"answer id" json:"id"`
	// user id
	UserID string `json:"-"`
}

// GetAnswerListReq get answer list all request
type GetAnswerListReq struct {
	// question id
	QuestionID int64 `validate:"omitempty" comment:"question id" form:"question_id"`
	// answer user id
	UserID int64 `validate:"omitempty" comment:"answer user id" form:"user_id"`
	// content markdown
	Content string `validate:"omitempty" comment:"content markdown" form:"content"`
	// content html
	Html string `validate:"omitempty" comment:"content html" form:"html"`
	//  answer status(available: 1; deleted: 10)
	Status int `validate:"omitempty" comment:" answer status(available: 1; deleted: 10)" form:"status"`
	// adopted (1 failed 2 adopted)
	Adopted int `validate:"omitempty" comment:"adopted (1 failed 2 adopted)" form:"adopted"`
	// comment count
	CommentCount int `validate:"omitempty" comment:"comment count" form:"comment_count"`
	// vote count
	VoteCount int `validate:"omitempty" comment:"vote count" form:"vote_count"`
	//
	CreateTime time.Time `validate:"omitempty" comment:"" form:"create_time"`
	//
	UpdateTime time.Time `validate:"omitempty" comment:"" form:"update_time"`
}

// GetAnswerWithPageReq get answer list page request
type GetAnswerWithPageReq struct {
	// page
	Page int `validate:"omitempty,min=1" form:"page"`
	// page size
	PageSize int `validate:"omitempty,min=1" form:"page_size"`
	// question id
	QuestionID int64 `validate:"omitempty" comment:"question id" form:"question_id"`
	// answer user id
	UserID int64 `validate:"omitempty" comment:"answer user id" form:"user_id"`
	// content markdown
	Content string `validate:"omitempty" comment:"content markdown" form:"content"`
	// content html
	Html string `validate:"omitempty" comment:"content html" form:"html"`
	//  answer status(available: 1; deleted: 10)
	Status int `validate:"omitempty" comment:" answer status(available: 1; deleted: 10)" form:"status"`
	// adopted (1 failed 2 adopted)
	Adopted int `validate:"omitempty" comment:"adopted (1 failed 2 adopted)" form:"adopted"`
	// comment count
	CommentCount int `validate:"omitempty" comment:"comment count" form:"comment_count"`
	// vote count
	VoteCount int `validate:"omitempty" comment:"vote count" form:"vote_count"`
	//
	CreateTime time.Time `validate:"omitempty" comment:"" form:"create_time"`
	//
	UpdateTime time.Time `validate:"omitempty" comment:"" form:"update_time"`
}

// GetAnswerResp get answer response
type GetAnswerResp struct {
	// answer id
	ID int64 `json:"id"`
	// question id
	QuestionID int64 `json:"question_id"`
	// answer user id
	UserID int64 `json:"user_id"`
	// content markdown
	Content string `json:"content"`
	// content html
	Html string `json:"html"`
	//  answer status(available: 1; deleted: 10)
	Status int `json:"status"`
	// adopted (1 failed 2 adopted)
	Adopted int `json:"adopted"`
	// comment count
	CommentCount int `json:"comment_count"`
	// vote count
	VoteCount int `json:"vote_count"`
	//
	CreateTime time.Time `json:"create_time"`
	//
	UpdateTime time.Time `json:"update_time"`
}

const (
	Answer_Adopted_Failed = 1
	Answer_Adopted_Enable = 2
)

type AnswerAddReq struct {
	QuestionId string `json:"question_id" ` // question_id
	Content    string `json:"content" `     // content
	Html       string `json:"html" `        // 解析后的html
	UserID     string `json:"-" `           // user_id
}

type AnswerUpdateReq struct {
	ID          string `json:"id"`                                // id
	QuestionId  string `json:"question_id" `                      // question_id
	UserID      string `json:"-" `                                // user_id
	Title       string `json:"title" `                            // title
	Content     string `json:"content"`                           // content
	Html        string `json:"html" `                             // 解析后的html
	EditSummary string `validate:"omitempty" json:"edit_summary"` //edit_summary
}

type AnswerList struct {
	QuestionId  string `json:"question_id" `               // question_id
	Order       string `json:"order" `                     // 1 Default 2 time
	Page        int    `json:"page" form:"page"`           //Query number of pages
	PageSize    int    `json:"page_size" form:"page_size"` //Search page size
	LoginUserID string `json:"-" `
}

type AnswerInfo struct {
	ID             string         `json:"id" xorm:"id"`                   // id
	QuestionId     string         `json:"question_id" xorm:"question_id"` // question_id
	Content        string         `json:"content" xorm:"content"`         // content
	Html           string         `json:"html" xorm:"html"`               // html
	CreateTime     int64          `json:"create_time" xorm:"created"`     // create_time
	UpdateTime     int64          `json:"update_time" xorm:"updated"`     // update_time
	Adopted        int            `json:"adopted"`                        // 1 Failed 2 Adopted
	UserId         string         `json:"-" `
	UserInfo       *UserBasicInfo `json:"user_info,omitempty"`
	UpdateUserInfo *UserBasicInfo `json:"update_user_info,omitempty"`
	Collected      bool           `json:"collected"`
	VoteStatus     string         `json:"vote_status"`
	VoteCount      int            `json:"vote_count"`
	QuestionInfo   *QuestionInfo  `json:"question_info,omitempty"`

	// MemberActions
	MemberActions []*PermissionMemberAction `json:"member_actions"`
}

type AdminAnswerInfo struct {
	ID           string         `json:"id"`
	QuestionId   string         `json:"question_id"`
	Description  string         `json:"description"`
	CreateTime   int64          `json:"create_time"`
	UpdateTime   int64          `json:"update_time"`
	Adopted      int            `json:"adopted"`
	UserId       string         `json:"-" `
	UserInfo     *UserBasicInfo `json:"user_info"`
	VoteCount    int            `json:"vote_count"`
	QuestionInfo struct {
		Title string `json:"title"`
	} `json:"question_info"`
}

type AnswerAdoptedReq struct {
	QuestionID string `json:"question_id" ` // question_id
	AnswerID   string `json:"answer_id" `
	UserID     string `json:"-" `
}
