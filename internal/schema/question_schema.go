package schema

import (
	"time"

	"github.com/answerdev/answer/internal/base/validator"
	"github.com/answerdev/answer/pkg/converter"
)

const (
	SitemapMaxSize         = 50000
	SitemapCachekey        = "answer@sitemap"
	SitemapPageCachekey    = "answer@sitemap@page%d"
	QuestionOperationPin   = "pin"
	QuestionOperationUnPin = "unpin"
	QuestionOperationHide  = "hide"
	QuestionOperationShow  = "show"
)

// RemoveQuestionReq delete question request
type RemoveQuestionReq struct {
	// question id
	ID      string `validate:"required" json:"id"`
	UserID  string `json:"-" ` // user_id
	IsAdmin bool   `json:"-"`
}

type CloseQuestionReq struct {
	ID        string `validate:"required" json:"id"`
	CloseType int    `json:"close_type"` // close_type
	CloseMsg  string `json:"close_msg"`  // close_type
	UserID    string `json:"-"`          // user_id
}

type OperationQuestionReq struct {
	ID        string `validate:"required" json:"id"`
	Operation string `json:"operation"` // operation [pin unpin hide show]
	UserID    string `json:"-"`         // user_id
	CanPin    bool   `json:"-"`
	CanList   bool   `json:"-"`
}

type CloseQuestionMeta struct {
	CloseType int    `json:"close_type"`
	CloseMsg  string `json:"close_msg"`
}

// ReopenQuestionReq reopen question request
type ReopenQuestionReq struct {
	QuestionID string `json:"question_id"`
	UserID     string `json:"-"`
}

type QuestionAdd struct {
	// question title
	Title string `validate:"required,notblank,gte=6,lte=150" json:"title"`
	// content
	Content string `validate:"required,notblank,gte=6,lte=65535" json:"content"`
	// html
	HTML string `json:"-"`
	// tags
	Tags []*TagItem `validate:"required,dive" json:"tags"`
	// user id
	UserID string `json:"-"`
	QuestionPermission
}

func (req *QuestionAdd) Check() (errFields []*validator.FormErrorField, err error) {
	req.HTML = converter.Markdown2HTML(req.Content)
	for _, tag := range req.Tags {
		if len(tag.OriginalText) > 0 {
			tag.ParsedText = converter.Markdown2HTML(tag.OriginalText)
		}
	}
	return nil, nil
}

type QuestionAddByAnswer struct {
	// question title
	Title string `validate:"required,notblank,gte=6,lte=150" json:"title"`
	// content
	Content string `validate:"required,notblank,gte=6,lte=65535" json:"content"`
	// html
	HTML          string `json:"-"`
	AnswerContent string `validate:"required,notblank,gte=6,lte=65535" json:"answer_content"`
	AnswerHTML    string `json:"-"`
	// tags
	Tags []*TagItem `validate:"required,dive" json:"tags"`
	// user id
	UserID              string   `json:"-"`
	MentionUsernameList []string `validate:"omitempty" json:"mention_username_list"`
	QuestionPermission
}

func (req *QuestionAddByAnswer) Check() (errFields []*validator.FormErrorField, err error) {
	req.HTML = converter.Markdown2HTML(req.Content)
	req.AnswerHTML = converter.Markdown2HTML(req.AnswerContent)
	for _, tag := range req.Tags {
		if len(tag.OriginalText) > 0 {
			tag.ParsedText = converter.Markdown2HTML(tag.OriginalText)
		}
	}
	return nil, nil
}

type QuestionPermission struct {
	// whether user can add it
	CanAdd bool `json:"-"`
	// whether user can edit it
	CanEdit bool `json:"-"`
	// whether user can delete it
	CanDelete bool `json:"-"`
	// whether user can close it
	CanClose bool `json:"-"`
	// whether user can reopen it
	CanReopen bool `json:"-"`
	// whether user can pin it
	CanPin   bool `json:"-"`
	CanUnPin bool `json:"-"`
	// whether user can hide it
	CanHide bool `json:"-"`
	CanShow bool `json:"-"`
	// whether user can use reserved it
	CanUseReservedTag bool `json:"-"`
	// whether user can invite other user to answer this question
	CanInviteOtherToAnswer bool `json:"-"`
	CanAddTag              bool `json:"-"`
}

type CheckCanQuestionUpdate struct {
	// question id
	ID string `validate:"required" form:"id"`
	// user id
	UserID  string `json:"-"`
	IsAdmin bool   `json:"-"`
}

type QuestionUpdate struct {
	// question id
	ID string `validate:"required" json:"id"`
	// question title
	Title string `validate:"required,notblank,gte=6,lte=150" json:"title"`
	// content
	Content string `validate:"required,notblank,gte=6,lte=65535" json:"content"`
	// html
	HTML       string   `json:"-"`
	InviteUser []string `validate:"omitempty"  json:"invite_user"`
	// tags
	Tags []*TagItem `validate:"required,dive" json:"tags"`
	// edit summary
	EditSummary string `validate:"omitempty" json:"edit_summary"`
	// user id
	UserID       string `json:"-"`
	NoNeedReview bool   `json:"-"`
	QuestionPermission
}

type QuestionUpdateInviteUser struct {
	ID         string   `validate:"required" json:"id"`
	InviteUser []string `validate:"omitempty"  json:"invite_user"`
	UserID     string   `json:"-"`
	QuestionPermission
}

func (req *QuestionUpdate) Check() (errFields []*validator.FormErrorField, err error) {
	req.HTML = converter.Markdown2HTML(req.Content)
	return nil, nil
}

type QuestionBaseInfo struct {
	ID              string `json:"id" `
	Title           string `json:"title" xorm:"title"`                       // title
	ViewCount       int    `json:"view_count" xorm:"view_count"`             // view count
	AnswerCount     int    `json:"answer_count" xorm:"answer_count"`         // answer count
	CollectionCount int    `json:"collection_count" xorm:"collection_count"` // collection count
	FollowCount     int    `json:"follow_count" xorm:"follow_count"`         // follow count
	Status          string `json:"status"`
	AcceptedAnswer  bool   `json:"accepted_answer"`
}

type QuestionInfo struct {
	ID                   string         `json:"id" `
	Title                string         `json:"title" xorm:"title"`                         // title
	UrlTitle             string         `json:"url_title" xorm:"url_title"`                 // title
	Content              string         `json:"content" xorm:"content"`                     // content
	HTML                 string         `json:"html" xorm:"html"`                           // html
	Description          string         `json:"description"`                                //description
	Tags                 []*TagResp     `json:"tags" `                                      // tags
	ViewCount            int            `json:"view_count" xorm:"view_count"`               // view_count
	UniqueViewCount      int            `json:"unique_view_count" xorm:"unique_view_count"` // unique_view_count
	VoteCount            int            `json:"vote_count" xorm:"vote_count"`               // vote_count
	AnswerCount          int            `json:"answer_count" xorm:"answer_count"`           // answer count
	CollectionCount      int            `json:"collection_count" xorm:"collection_count"`   // collection count
	FollowCount          int            `json:"follow_count" xorm:"follow_count"`           // follow count
	AcceptedAnswerID     string         `json:"accepted_answer_id" `                        // accepted_answer_id
	LastAnswerID         string         `json:"last_answer_id" `                            // last_answer_id
	CreateTime           int64          `json:"create_time" `                               // create_time
	UpdateTime           int64          `json:"-"`                                          // update_time
	PostUpdateTime       int64          `json:"update_time"`
	QuestionUpdateTime   int64          `json:"edit_time"`
	Pin                  int            `json:"pin"`  // 1: unpin, 2: pin
	Show                 int            `json:"show"` // 0: show, 1: hide
	Status               int            `json:"status"`
	Operation            *Operation     `json:"operation,omitempty"`
	UserID               string         `json:"-" `
	LastEditUserID       string         `json:"-" `
	LastAnsweredUserID   string         `json:"-" `
	UserInfo             *UserBasicInfo `json:"user_info"`
	UpdateUserInfo       *UserBasicInfo `json:"update_user_info,omitempty"`
	LastAnsweredUserInfo *UserBasicInfo `json:"last_answered_user_info,omitempty"`
	Answered             bool           `json:"answered"`
	Collected            bool           `json:"collected"`
	VoteStatus           string         `json:"vote_status"`
	IsFollowed           bool           `json:"is_followed"`

	// MemberActions
	MemberActions  []*PermissionMemberAction `json:"member_actions"`
	ExtendsActions []*PermissionMemberAction `json:"extends_actions"`
}

// UpdateQuestionResp update question resp
type UpdateQuestionResp struct {
	WaitForReview bool `json:"wait_for_review"`
}

type AdminQuestionInfo struct {
	ID               string         `json:"id"`
	Title            string         `json:"title"`
	VoteCount        int            `json:"vote_count"`
	AnswerCount      int            `json:"answer_count"`
	AcceptedAnswerID string         `json:"accepted_answer_id"`
	CreateTime       int64          `json:"create_time"`
	UpdateTime       int64          `json:"update_time"`
	EditTime         int64          `json:"edit_time"`
	UserID           string         `json:"-" `
	UserInfo         *UserBasicInfo `json:"user_info"`
}

type OperationLevel string

const (
	OperationLevelInfo    OperationLevel = "info"
	OperationLevelDanger  OperationLevel = "danger"
	OperationLevelWarning OperationLevel = "warning"
)

type Operation struct {
	Type        string         `json:"type"`
	Description string         `json:"description"`
	Msg         string         `json:"msg"`
	Time        int64          `json:"time"`
	Level       OperationLevel `json:"level"`
}

type GetCloseTypeResp struct {
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

type UserAnswerInfo struct {
	AnswerID     string `json:"answer_id"`
	QuestionID   string `json:"question_id"`
	Accepted     int    `json:"accepted"`
	VoteCount    int    `json:"vote_count"`
	CreateTime   int    `json:"create_time"`
	UpdateTime   int    `json:"update_time"`
	QuestionInfo struct {
		Title    string        `json:"title"`
		UrlTitle string        `json:"url_title"`
		Tags     []interface{} `json:"tags"`
	} `json:"question_info"`
}

type UserQuestionInfo struct {
	ID               string        `json:"question_id"`
	Title            string        `json:"title"`
	UrlTitle         string        `json:"url_title"`
	VoteCount        int           `json:"vote_count"`
	Tags             []interface{} `json:"tags"`
	ViewCount        int           `json:"view_count"`
	AnswerCount      int           `json:"answer_count"`
	CollectionCount  int           `json:"collection_count"`
	CreatedAt        int64         `json:"created_at"`
	AcceptedAnswerID string        `json:"accepted_answer_id"`
	Status           string        `json:"status"`
}

const (
	QuestionOrderCondNewest     = "newest"
	QuestionOrderCondActive     = "active"
	QuestionOrderCondFrequent   = "frequent"
	QuestionOrderCondScore      = "score"
	QuestionOrderCondUnanswered = "unanswered"
)

// QuestionPageReq query questions page
type QuestionPageReq struct {
	Page      int    `validate:"omitempty,min=1" form:"page"`
	PageSize  int    `validate:"omitempty,min=1" form:"page_size"`
	OrderCond string `validate:"omitempty,oneof=newest active frequent score unanswered" form:"order"`
	Tag       string `validate:"omitempty,gt=0,lte=100" form:"tag"`
	Username  string `validate:"omitempty,gt=0,lte=100" form:"username"`
	InDays    int    `validate:"omitempty,min=1" form:"in_days"`

	LoginUserID      string `json:"-"`
	UserIDBeSearched string `json:"-"`
	TagID            string `json:"-"`
}

const (
	QuestionPageRespOperationTypeAsked    = "asked"
	QuestionPageRespOperationTypeAnswered = "answered"
	QuestionPageRespOperationTypeModified = "modified"
)

type QuestionPageResp struct {
	ID          string     `json:"id" `
	CreatedAt   int64      `json:"created_at"`
	Title       string     `json:"title"`
	UrlTitle    string     `json:"url_title"`
	Description string     `json:"description"`
	Pin         int        `json:"pin"`  // 1: unpin, 2: pin
	Show        int        `json:"show"` // 0: show, 1: hide
	Status      int        `json:"status"`
	Tags        []*TagResp `json:"tags"`

	// question statistical information
	ViewCount       int `json:"view_count"`
	UniqueViewCount int `json:"unique_view_count"`
	VoteCount       int `json:"vote_count"`
	AnswerCount     int `json:"answer_count"`
	CollectionCount int `json:"collection_count"`
	FollowCount     int `json:"follow_count"`

	// answer information
	AcceptedAnswerID   string    `json:"accepted_answer_id"`
	LastAnswerID       string    `json:"last_answer_id"`
	LastAnsweredUserID string    `json:"-"`
	LastAnsweredAt     time.Time `json:"-"`

	// operator information
	OperatedAt    int64                     `json:"operated_at"`
	Operator      *QuestionPageRespOperator `json:"operator"`
	OperationType string                    `json:"operation_type"`
}

type QuestionPageRespOperator struct {
	ID          string `json:"id"`
	Username    string `json:"username"`
	Rank        int    `json:"rank"`
	DisplayName string `json:"display_name"`
}

type AdminQuestionSearch struct {
	Page      int    `json:"page" form:"page"`           // Query number of pages
	PageSize  int    `json:"page_size" form:"page_size"` // Search page size
	Status    int    `json:"-" form:"-"`
	StatusStr string `json:"status" form:"status"`                                  // Status 1 Available 2 closed 10 UserDeleted
	Query     string `validate:"omitempty,gt=0,lte=100" json:"query" form:"query" ` //Query string
}

type AdminSetQuestionStatusRequest struct {
	StatusStr  string `json:"status" form:"status"`
	QuestionID string `json:"question_id" form:"question_id"`
}

type SiteMapList struct {
	QuestionIDs []*SiteMapQuestionInfo `json:"question_ids"`
	MaxPageNum  []int                  `json:"max_page_num"`
}

type SiteMapPageList struct {
	PageData []*SiteMapQuestionInfo `json:"page_data"`
}

type SiteMapQuestionInfo struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	UpdateTime string `json:"time"`
}

type PersonalQuestionPageReq struct {
	Page        int    `validate:"omitempty,min=1" form:"page"`
	PageSize    int    `validate:"omitempty,min=1" form:"page_size"`
	OrderCond   string `validate:"omitempty,oneof=newest active frequent score unanswered" form:"order"`
	Username    string `validate:"omitempty,gt=0,lte=100" form:"username"`
	LoginUserID string `json:"-"`
}

type PersonalAnswerPageReq struct {
	Page        int    `validate:"omitempty,min=1" form:"page"`
	PageSize    int    `validate:"omitempty,min=1" form:"page_size"`
	OrderCond   string `validate:"omitempty,oneof=newest active frequent score unanswered" form:"order"`
	Username    string `validate:"omitempty,gt=0,lte=100" form:"username"`
	LoginUserID string `json:"-"`
}

type PersonalCollectionPageReq struct {
	Page     int    `validate:"omitempty,min=1" form:"page"`
	PageSize int    `validate:"omitempty,min=1" form:"page_size"`
	UserID   string `json:"-"`
}
