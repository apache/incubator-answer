package entity

const (
	CaptchaActionEmail            = "email"
	CaptchaActionPassword         = "password"
	CaptchaActionEditUserinfo     = "edit_userinfo"
	CaptchaActionQuestion         = "question"
	CaptchaActionAnswer           = "answer"
	CaptchaActionComment          = "comment"
	CaptchaActionEdit             = "edit"
	CaptchaActionInvitationAnswer = "invitation_answer"
	CaptchaActionSearch           = "search"
	CaptchaActionReport           = "report"
	CaptchaActionDelete           = "delete"
	CaptchaActionVote             = "vote"
)

type ActionRecordInfo struct {
	LastTime int64  `json:"last_time"`
	Num      int    `json:"num"`
	Config   string `json:"config"`
}
