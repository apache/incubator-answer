package schema

// RemoveQuestionReq delete question request
type RemoveQuestionReq struct {
	// question id
	ID      string `validate:"required" comment:"question id" json:"id"`
	UserID  string `json:"-" ` // user_id
	IsAdmin bool   `json:"-"`
}

type CloseQuestionReq struct {
	ID        string `validate:"required" comment:"question id" json:"id"`
	UserID    string `json:"-" `          // user_id
	CloseType int    `json:"close_type" ` // close_type
	CloseMsg  string `json:"close_msg" `  // close_type
	IsAdmin   bool   `json:"-"`
}

type CloseQuestionMeta struct {
	CloseType int    `json:"close_type"`
	CloseMsg  string `json:"close_msg"`
}

type QuestionAdd struct {
	// question title
	Title string `validate:"required,gte=6,lte=150" json:"title"`
	// content
	Content string `validate:"required,gte=6,lte=65535" json:"content"`
	// html
	HTML string `validate:"required,gte=6,lte=65535" json:"html"`
	// tags
	Tags []*TagItem `validate:"required,dive" json:"tags"`
	// user id
	UserID string `json:"-"`
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
	Title string `validate:"required,gte=6,lte=150" json:"title"`
	// content
	Content string `validate:"required,gte=6,lte=65535" json:"content"`
	// html
	HTML string `validate:"required,gte=6,lte=65535" json:"html"`
	// tags
	Tags []*TagItem `validate:"required,dive" json:"tags"`
	// edit summary
	EditSummary string `validate:"omitempty" json:"edit_summary"`
	// user id
	UserID       string `json:"-"`
	IsAdmin      bool   `json:"-"`
	NoNeedReview bool   `json:"-"`
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
	Content              string         `json:"content" xorm:"content"`                     // content
	HTML                 string         `json:"html" xorm:"html"`                           // html
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
	Status               int            `json:"status"`
	Operation            *Operation     `json:"operation,omitempty"`
	UserID               string         `json:"-" `
	UserInfo             *UserBasicInfo `json:"user_info"`
	UpdateUserInfo       *UserBasicInfo `json:"update_user_info,omitempty"`
	LastAnsweredUserInfo *UserBasicInfo `json:"last_answered_user_info,omitempty"`
	Answered             bool           `json:"answered"`
	Collected            bool           `json:"collected"`
	VoteStatus           string         `json:"vote_status"`
	IsFollowed           bool           `json:"is_followed"`

	// MemberActions
	MemberActions []*PermissionMemberAction `json:"member_actions"`
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

type Operation struct {
	OperationType        string `json:"operation_type"`
	OperationDescription string `json:"operation_description"`
	OperationMsg         string `json:"operation_msg"`
	OperationTime        int64  `json:"operation_time"`
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
	Adopted      int    `json:"adopted"`
	VoteCount    int    `json:"vote_count"`
	CreateTime   int    `json:"create_time"`
	UpdateTime   int    `json:"update_time"`
	QuestionInfo struct {
		Title string        `json:"title"`
		Tags  []interface{} `json:"tags"`
	} `json:"question_info"`
}

type UserQuestionInfo struct {
	ID               string        `json:"question_id"`
	Title            string        `json:"title"`
	VoteCount        int           `json:"vote_count"`
	Tags             []interface{} `json:"tags"`
	ViewCount        int           `json:"view_count"`
	AnswerCount      int           `json:"answer_count"`
	CollectionCount  int           `json:"collection_count"`
	CreateTime       int           `json:"create_time"`
	AcceptedAnswerID string        `json:"accepted_answer_id"`
	Status           string        `json:"status"`
}

type QuestionSearch struct {
	Page     int    `json:"page" form:"page"`           // Query number of pages
	PageSize int    `json:"page_size" form:"page_size"` // Search page size
	Order    string `json:"order" form:"order"`         // Search order by
	// Tags     []string `json:"tags" form:"tags"`           // Search tag
	Tag      string   `json:"tag" form:"tag"`           //Search tag
	TagIDs   []string `json:"-" form:"-"`               // Search tag
	UserName string   `json:"username" form:"username"` // Search username
	UserID   string   `json:"-" form:"-"`
}

type CmsQuestionSearch struct {
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
