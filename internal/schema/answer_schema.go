package schema

// RemoveAnswerReq delete answer request
type RemoveAnswerReq struct {
	// answer id
	ID string `validate:"required" json:"id"`
	// user id
	UserID  string `json:"-"`
	IsAdmin bool   `json:"-"`
}

const (
	AnswerAdoptedFailed = 1
	AnswerAdoptedEnable = 2
)

type AnswerAddReq struct {
	QuestionID string `json:"question_id" ` // question_id
	Content    string `json:"content" `     // content
	HTML       string `json:"html" `        // html
	UserID     string `json:"-" `           // user_id
}

type AnswerUpdateReq struct {
	ID           string `json:"id"`                                // id
	QuestionID   string `json:"question_id" `                      // question_id
	UserID       string `json:"-" `                                // user_id
	Title        string `json:"title" `                            // title
	Content      string `json:"content"`                           // content
	HTML         string `json:"html" `                             // html
	EditSummary  string `validate:"omitempty" json:"edit_summary"` // edit_summary
	NoNeedReview bool   `json:"-"`
	// whether user can edit it
	CanEdit bool `json:"-"`
}

// AnswerUpdateResp answer update resp
type AnswerUpdateResp struct {
	AnswerInfo    *AnswerInfo   `json:"info"`
	QuestionInfo  *QuestionInfo `json:"question"`
	WaitForReview bool          `json:"wait_for_review"`
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
	Adopted        int            `json:"adopted"`                        // 1 Failed 2 Adopted
	UserID         string         `json:"-" `
	UpdateUserID   string         `json:"-" `
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
	QuestionID   string         `json:"question_id"`
	Description  string         `json:"description"`
	CreateTime   int64          `json:"create_time"`
	UpdateTime   int64          `json:"update_time"`
	Adopted      int            `json:"adopted"`
	UserID       string         `json:"-" `
	UpdateUserID string         `json:"-" `
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
