package schema

import (
	"encoding/json"

	"github.com/answerdev/answer/internal/base/constant"
	"github.com/answerdev/answer/internal/base/reason"
	"github.com/answerdev/answer/internal/base/validator"
	"github.com/answerdev/answer/internal/entity"
	"github.com/answerdev/answer/pkg/checker"
	"github.com/answerdev/answer/pkg/converter"
	"github.com/answerdev/answer/pkg/gravatar"
	"github.com/jinzhu/copier"
	"github.com/segmentfault/pacman/errors"
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
	BioHTML string `json:"bio_html"`
	// website
	Website string `json:"website"`
	// location
	Location string `json:"location"`
	// ip info
	IPInfo string `json:"ip_info"`
	// language
	Language string `json:"language"`
	// access token
	AccessToken string `json:"access_token"`
	// role id
	RoleID int `json:"role_id"`
	// user status
	Status string `json:"status"`
}

func (r *GetUserResp) GetFromUserEntity(userInfo *entity.User) {
	_ = copier.Copy(r, userInfo)
	r.Avatar = FormatAvatarInfo(userInfo.Avatar, userInfo.EMail)
	r.CreatedAt = userInfo.CreatedAt.Unix()
	r.LastLoginDate = userInfo.LastLoginDate.Unix()
	statusShow, ok := UserStatusShow[userInfo.Status]
	if ok {
		r.Status = statusShow
	}
}

type GetUserToSetShowResp struct {
	*GetUserResp
	Avatar *AvatarInfo `json:"avatar"`
}

func (r *GetUserToSetShowResp) GetFromUserEntity(userInfo *entity.User) {
	_ = copier.Copy(r, userInfo)
	r.CreatedAt = userInfo.CreatedAt.Unix()
	r.LastLoginDate = userInfo.LastLoginDate.Unix()
	statusShow, ok := UserStatusShow[userInfo.Status]
	if ok {
		r.Status = statusShow
	}
	avatarInfo := &AvatarInfo{}
	_ = json.Unmarshal([]byte(userInfo.Avatar), avatarInfo)
	if constant.DefaultAvatar == "gravatar" && avatarInfo.Type == "" {
		avatarInfo.Type = "gravatar"
		avatarInfo.Gravatar = gravatar.GetAvatarURL(userInfo.EMail)
	}
	// if json.Unmarshal Error avatarInfo.Type is Empty
	r.Avatar = avatarInfo
}

func FormatAvatarInfo(avatarJson, email string) (res string) {
	defer func() {
		if constant.DefaultAvatar == "gravatar" && len(res) == 0 {
			res = gravatar.GetAvatarURL(email)
		}
	}()

	if avatarJson == "" {
		return ""
	}
	avatarInfo := &AvatarInfo{}
	err := json.Unmarshal([]byte(avatarJson), avatarInfo)
	if err != nil {
		return ""
	}
	switch avatarInfo.Type {
	case "gravatar":
		return avatarInfo.Gravatar
	case "custom":
		return avatarInfo.Custom
	default:
		return ""
	}
}

// GetUserStatusResp get user status info
type GetUserStatusResp struct {
	// user status
	Status string `json:"status"`
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
	BioHTML string `json:"bio_html"`
	// website
	Website string `json:"website"`
	// location
	Location  string `json:"location"`
	Status    string `json:"status"`
	StatusMsg string `json:"status_msg,omitempty"`
}

func (r *GetOtherUserInfoByUsernameResp) GetFromUserEntity(userInfo *entity.User) {
	_ = copier.Copy(r, userInfo)
	Avatar := FormatAvatarInfo(userInfo.Avatar, userInfo.EMail)
	r.Avatar = Avatar

	r.CreatedAt = userInfo.CreatedAt.Unix()
	r.LastLoginDate = userInfo.LastLoginDate.Unix()
	statusShow, ok := UserStatusShow[userInfo.Status]
	if ok {
		r.Status = statusShow
	}
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
}

const (
	MailStatePass   = 1
	MailStateVerifi = 2

	NoticeStatusOn  = 1
	NoticeStatusOff = 2

	ActionRecordTypeLogin    = "login"
	ActionRecordTypeEmail    = "e_mail"
	ActionRecordTypeFindPass = "find_pass"
)

var UserStatusShow = map[int]string{
	1:  "normal",
	9:  "forbidden",
	10: "deleted",
}

var UserStatusShowMsg = map[int]string{
	1:  "",
	9:  "<strong>This user was suspended forever.</strong> This user doesn’t meet a community guideline.",
	10: "This user was deleted.",
	11: "This user is inactive.",
}

// EmailLogin
type UserEmailLogin struct {
	Email       string `validate:"required,email,gt=0,lte=500" json:"e_mail"` // e_mail
	Pass        string `validate:"required,gte=8,lte=32" json:"pass"`         // password
	CaptchaID   string `json:"captcha_id"`                                    // captcha_id
	CaptchaCode string `json:"captcha_code"`                                  // captcha_code
}

// UserRegisterReq user register request
type UserRegisterReq struct {
	// name
	Name string `validate:"required,gt=3,lte=30" json:"name"`
	// email
	Email string `validate:"required,email,gt=0,lte=500" json:"e_mail" `
	// password
	Pass        string `validate:"required,gte=8,lte=32" json:"pass"`
	IP          string `json:"-" `
	CaptchaID   string `json:"captcha_id"`   // captcha_id
	CaptchaCode string `json:"captcha_code"` // captcha_code
}

func (u *UserRegisterReq) Check() (errFields []*validator.FormErrorField, err error) {
	// TODO i18n
	err = checker.CheckPassword(8, 32, 0, u.Pass)
	if err != nil {
		errField := &validator.FormErrorField{
			ErrorField: "pass",
			ErrorMsg:   err.Error(),
		}
		errFields = append(errFields, errField)
		return errFields, err
	}
	return nil, nil
}

// UserModifyPassWordRequest
type UserModifyPassWordRequest struct {
	UserID  string `json:"-" `        // user_id
	OldPass string `json:"old_pass" ` // old password
	Pass    string `json:"pass" `     // password
}

func (u *UserModifyPassWordRequest) Check() (errFields []*validator.FormErrorField, err error) {
	// TODO i18n
	err = checker.CheckPassword(8, 32, 0, u.Pass)
	if err != nil {
		errField := &validator.FormErrorField{
			ErrorField: "pass",
			ErrorMsg:   err.Error(),
		}
		errFields = append(errFields, errField)
		return errFields, err
	}
	return nil, nil
}

type UpdateInfoRequest struct {
	// display_name
	DisplayName string `validate:"required,gt=0,lte=30" json:"display_name"`
	// username
	Username string `validate:"omitempty,gt=3,lte=30" json:"username"`
	// avatar
	Avatar AvatarInfo `json:"avatar"`
	// bio
	Bio string `validate:"omitempty,gt=0,lte=4096" json:"bio"`
	// bio
	BioHTML string `json:"-"`
	// website
	Website string `validate:"omitempty,gt=0,lte=500" json:"website"`
	// location
	Location string `validate:"omitempty,gt=0,lte=100" json:"location"`
	// user id
	UserID string `json:"-" `
}

type AvatarInfo struct {
	Type     string `validate:"omitempty,gt=0,lte=100"  json:"type"`
	Gravatar string `validate:"omitempty,gt=0,lte=200"  json:"gravatar"`
	Custom   string `validate:"omitempty,gt=0,lte=200"  json:"custom"`
}

func (req *UpdateInfoRequest) Check() (errFields []*validator.FormErrorField, err error) {
	if len(req.Username) > 0 {
		if checker.IsInvalidUsername(req.Username) {
			errField := &validator.FormErrorField{
				ErrorField: "username",
				ErrorMsg:   reason.UsernameInvalid,
			}
			errFields = append(errFields, errField)
			return errFields, errors.BadRequest(reason.UsernameInvalid)
		}
	}
	req.BioHTML = converter.Markdown2BasicHTML(req.Bio)
	return nil, nil
}

// UpdateUserInterfaceRequest update user interface request
type UpdateUserInterfaceRequest struct {
	// language
	Language string `validate:"required,gt=1,lte=100" json:"language"`
	// user id
	UserId string `json:"-" `
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

func (u *UserRePassWordRequest) Check() (errFields []*validator.FormErrorField, err error) {
	// TODO i18n
	err = checker.CheckPassword(8, 32, 0, u.Pass)
	if err != nil {
		errField := &validator.FormErrorField{
			ErrorField: "pass",
			ErrorMsg:   err.Error(),
		}
		errFields = append(errFields, errField)
		return errFields, err
	}
	return nil, nil
}

type UserNoticeSetRequest struct {
	NoticeSwitch bool   `json:"notice_switch"`
	UserID       string `json:"-"`
}

type UserNoticeSetResp struct {
	NoticeSwitch bool `json:"notice_switch"`
}

type ActionRecordReq struct {
	// action
	Action string `validate:"required,oneof=login e_mail find_pass" form:"action"`
	IP     string `json:"-"`
}

type ActionRecordResp struct {
	CaptchaID  string `json:"captcha_id"`
	CaptchaImg string `json:"captcha_img"`
	Verify     bool   `json:"verify"`
}

type UserBasicInfo struct {
	ID          string `json:"id"`           // user_id
	Username    string `json:"username" `    // name
	Rank        int    `json:"rank" `        // rank
	DisplayName string `json:"display_name"` // display_name
	Avatar      string `json:"avatar" `      // avatar
	Website     string `json:"website" `     // website
	Location    string `json:"location" `    // location
	IPInfo      string `json:"ip_info"`      // ip info
	Status      string `json:"status"`       // status
}

type GetOtherUserInfoByUsernameReq struct {
	Username string `validate:"required,gt=0,lte=500" form:"username"`
}

type GetOtherUserInfoResp struct {
	Info *GetOtherUserInfoByUsernameResp `json:"info"`
}

type UserChangeEmailSendCodeReq struct {
	UserVerifyEmailSendReq
	Email  string `validate:"required,email,gt=0,lte=500" json:"e_mail"`
	Pass   string `validate:"omitempty,gte=8,lte=32" json:"pass"`
	UserID string `json:"-"`
}

type UserChangeEmailVerifyReq struct {
	Code    string `validate:"required,gt=0,lte=500" json:"code"`
	Content string `json:"-"`
}

type UserVerifyEmailSendReq struct {
	CaptchaID   string `validate:"omitempty,gt=0,lte=500" json:"captcha_id"`
	CaptchaCode string `validate:"omitempty,gt=0,lte=500" json:"captcha_code"`
}

// UserRankingResp user ranking response
type UserRankingResp struct {
	UsersWithTheMostReputation []*UserRankingSimpleInfo `json:"users_with_the_most_reputation"`
	UsersWithTheMostVote       []*UserRankingSimpleInfo `json:"users_with_the_most_vote"`
	Staffs                     []*UserRankingSimpleInfo `json:"staffs"`
}

// UserRankingSimpleInfo user ranking simple info
type UserRankingSimpleInfo struct {
	// username
	Username string `json:"username"`
	// rank
	Rank int `json:"rank"`
	// vote
	VoteCount int `json:"vote_count"`
	// display name
	DisplayName string `json:"display_name"`
	// avatar
	Avatar string `json:"avatar"`
}

// UserUnsubscribeEmailNotificationReq user unsubscribe email notification request
type UserUnsubscribeEmailNotificationReq struct {
	Code    string `validate:"required,gt=0,lte=500" json:"code"`
	Content string `json:"-"`
}
