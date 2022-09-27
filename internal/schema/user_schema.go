package schema

import (
	"encoding/json"

	"github.com/davecgh/go-spew/spew"
	"github.com/jinzhu/copier"
	"github.com/segmentfault/answer/internal/base/validator"
	"github.com/segmentfault/answer/internal/entity"
	"github.com/segmentfault/answer/pkg/checker"
)

// UserVerifyEmailReq user verify email request
type UserVerifyEmailReq struct {
	// code
	Code string `validate:"required,gt=0,lte=500" form:"code"`
	// content
	Content string `json:"-"`
}

// GetUserResp get user response
type GetUserResp struct {
	// user id
	ID string `json:"id"`
	// create time
	CreatedAt int64 `json:"created_at"`
	// last login date
	LastLoginDate int64 `json:"last_login_date"`
	// username
	Username string `json:"username"`
	// email
	EMail string `json:"e_mail"`
	// mail status(1 pass 2 to be verified)
	MailStatus int `json:"mail_status"`
	// notice status(1 on 2off)
	NoticeStatus int `json:"notice_status"`
	// follow count
	FollowCount int `json:"follow_count"`
	// answer count
	AnswerCount int `json:"answer_count"`
	// question count
	QuestionCount int `json:"question_count"`
	// rank
	Rank int `json:"rank"`
	// authority group
	AuthorityGroup int `json:"authority_group"`
	// display name
	DisplayName string `json:"display_name"`
	// avatar
	Avatar string `json:"avatar"`
	// mobile
	Mobile string `json:"mobile"`
	// bio markdown
	Bio string `json:"bio"`
	// bio html
	BioHtml string `json:"bio_html"`
	// website
	Website string `json:"website"`
	// location
	Location string `json:"location"`
	// ip info
	IPInfo string `json:"ip_info"`
	// access token
	AccessToken string `json:"access_token"`
	// is admin
	IsAdmin bool `json:"is_admin"`
	// user status
	Status string `json:"status"`
}

func (r *GetUserResp) GetFromUserEntity(userInfo *entity.User) {
	_ = copier.Copy(r, userInfo)
	r.CreatedAt = userInfo.CreatedAt.Unix()
	r.LastLoginDate = userInfo.LastLoginDate.Unix()
	statusShow, ok := UserStatusShow[userInfo.Status]
	if ok {
		r.Status = statusShow
	}

}

// GetOtherUserInfoByUsernameResp get user response
type GetOtherUserInfoByUsernameResp struct {
	// user id
	ID string `json:"id"`
	// create time
	CreatedAt int64 `json:"created_at"`
	// last login date
	LastLoginDate int64 `json:"last_login_date"`
	// username
	Username string `json:"username"`
	// email
	// follow count
	FollowCount int `json:"follow_count"`
	// answer count
	AnswerCount int `json:"answer_count"`
	// question count
	QuestionCount int `json:"question_count"`
	// rank
	Rank int `json:"rank"`
	// display name
	DisplayName string `json:"display_name"`
	// avatar
	Avatar string `json:"avatar"`
	// mobile
	Mobile string `json:"mobile"`
	// bio markdown
	Bio string `json:"bio"`
	// bio html
	BioHtml string `json:"bio_html"`
	// website
	Website string `json:"website"`
	// location
	Location string `json:"location"`
	// ip info
	IPInfo string `json:"ip_info"`
	// is admin
	IsAdmin   bool   `json:"is_admin"`
	Status    string `json:"status"`
	StatusMsg string `json:"status_msg,omitempty"`
}

func (r *GetOtherUserInfoByUsernameResp) GetFromUserEntity(userInfo *entity.User) {
	_ = copier.Copy(r, userInfo)
	r.CreatedAt = userInfo.CreatedAt.Unix()
	r.LastLoginDate = userInfo.LastLoginDate.Unix()
	statusShow, ok := UserStatusShow[userInfo.Status]
	if ok {
		r.Status = statusShow
	}
	spew.Dump(userInfo)
	if userInfo.MailStatus == entity.EmailStatusToBeVerified {
		statusMsgShow, ok := UserStatusShowMsg[11]
		if ok {
			r.StatusMsg = statusMsgShow
		}
	} else {
		statusMsgShow, ok := UserStatusShowMsg[userInfo.Status]
		if ok {
			r.StatusMsg = statusMsgShow
		}
	}

	spew.Dump(r)

}

const (
	Mail_State_Pass   = 1
	Mail_State_Verifi = 2

	Notice_Status_On  = 1
	Notice_Status_Off = 2

	//ActionRecord ReportType
	ActionRecord_Type_Login     = "login"
	ActionRecord_Type_Email     = "e_mail"
	ActionRecord_Type_Find_Pass = "find_pass"
)

var UserStatusShow = map[int]string{
	1:  "normal",
	9:  "forbidden",
	10: "delete",
}
var UserStatusShowMsg = map[int]string{
	1:  "",
	9:  "<strong>This user was suspended forever.</strong> This user doesn’t meet a community guideline.",
	10: "This user was deleted.",
	11: "This user is inactive.",
}

// EmailLogin
type UserEmailLogin struct {
	Email       string `json:"e_mail" `       // e_mail
	Pass        string `json:"pass" `         // password
	CaptchaID   string `json:"captcha_id" `   // captcha_id
	CaptchaCode string `json:"captcha_code" ` // captcha_code
}

// Register
type UserRegister struct {
	// name
	Name string `validate:"required,gt=5,lte=50" json:"name"`
	// email
	Email string `validate:"required,email,gt=0,lte=500" json:"e_mail" `
	// password
	Pass string `validate:"required,gte=8,lte=32" json:"pass"`
	IP   string `json:"-" `
}

func (u *UserRegister) Check() (errField *validator.ErrorField, err error) {
	// TODO i18n
	err = checker.PassWordCheck(8, 32, 0, u.Pass)
	if err != nil {
		return &validator.ErrorField{
			Key:   "pass",
			Value: err.Error(),
		}, err
	}
	return nil, nil
}

// UserModifyPassWordRequest
type UserModifyPassWordRequest struct {
	UserId  string `json:"-" `        // user_id
	OldPass string `json:"old_pass" ` // old password
	Pass    string `json:"pass" `     // password
}

func (u *UserModifyPassWordRequest) Check() (errField *validator.ErrorField, err error) {
	// TODO i18n
	err = checker.PassWordCheck(8, 32, 0, u.Pass)
	if err != nil {
		return &validator.ErrorField{
			Key:   "pass",
			Value: err.Error(),
		}, err
	}
	return nil, nil
}

type UpdateInfoRequest struct {
	UserId      string `json:"-" `            // user_id
	UserName    string `json:"username"`      // name
	DisplayName string `json:"display_name" ` // display_name
	Avatar      string `json:"avatar" `       // avatar
	Bio         string `json:"bio"`
	BioHtml     string `json:"bio_html"`
	Website     string `json:"website" ` // website
	Location    string `json:"location"` // location
}

type UserRetrievePassWordRequest struct {
	Email       string `validate:"required,email,gt=0,lte=500" json:"e_mail" ` // e_mail
	CaptchaID   string `json:"captcha_id" `                                    // captcha_id
	CaptchaCode string `json:"captcha_code" `                                  // captcha_code
}

type UserRePassWordRequest struct {
	Code    string `validate:"required,gt=0,lte=100" json:"code" ` // code
	Pass    string `validate:"required,gt=0,lte=32" json:"pass" `  // Password
	Content string `json:"-"`
}

func (u *UserRePassWordRequest) Check() (errField *validator.ErrorField, err error) {
	// TODO i18n
	err = checker.PassWordCheck(8, 32, 0, u.Pass)
	if err != nil {
		return &validator.ErrorField{
			Key:   "pass",
			Value: err.Error(),
		}, err
	}
	return nil, nil
}

type UserNoticeSetRequest struct {
	UserId       string `json:"-" ` // user_id
	NoticeSwitch bool   `json:"notice_switch" `
}

type UserNoticeSetResp struct {
	NoticeSwitch bool `json:"notice_switch"`
}

type ActionRecordReq struct {
	// action
	Action string `validate:"required,oneof=login e_mail find_pass" form:"action"`
	Ip     string `json:"-"`
}

type ActionRecordResp struct {
	CaptchaID  string `json:"captcha_id"`
	CaptchaImg string `json:"captcha_img"`
	Verify     bool   `json:"verify"`
}

type UserBasicInfo struct {
	UserId      string `json:"-" `           // user_id
	UserName    string `json:"username" `    // name
	Rank        int    `json:"rank" `        // rank
	DisplayName string `json:"display_name"` // display_name
	Avatar      string `json:"avatar" `      // avatar
	Website     string `json:"website" `     // website
	Location    string `json:"location" `    // location
	IpInfo      string `json:"ip_info"`      // ip info
	Status      int    `json:"status"`       // status
}

type GetOtherUserInfoByUsernameReq struct {
	Username string `validate:"required,gt=0,lte=500" form:"username"`
}

type GetOtherUserInfoResp struct {
	Info *GetOtherUserInfoByUsernameResp `json:"info"`
	Has  bool                            `json:"has"`
}

type UserChangeEmailSendCodeReq struct {
	Email  string `validate:"required,email,gt=0,lte=500" json:"e_mail"`
	UserID string `json:"-"`
}

type EmailCodeContent struct {
	Email  string `json:"e_mail"`
	UserID string `json:"user_id"`
}

func (r *EmailCodeContent) ToJSONString() string {
	codeBytes, _ := json.Marshal(r)
	return string(codeBytes)
}

func (r *EmailCodeContent) FromJSONString(data string) error {
	return json.Unmarshal([]byte(data), &r)
}

type UserChangeEmailVerifyReq struct {
	Code    string `validate:"required,gt=0,lte=500" json:"code"`
	Content string `json:"-"`
}

type UserVerifyEmailSendReq struct {
	CaptchaID   string `validate:"omitempty,gt=0,lte=500" json:"captcha_id"`
	CaptchaCode string `validate:"omitempty,gt=0,lte=500" json:"captcha_code"`
}
