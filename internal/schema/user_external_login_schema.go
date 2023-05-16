package schema

// UserExternalLoginResp user external login resp
type UserExternalLoginResp struct {
	BindingKey  string `json:"binding_key"`
	AccessToken string `json:"access_token"`
	// ErrMsg error message, if not empty, means login failed and this message should be displayed.
	ErrMsg   string `json:"-"`
	ErrTitle string `json:"-"`
}

// ExternalLoginBindingUserSendEmailReq external login binding user request
type ExternalLoginBindingUserSendEmailReq struct {
	BindingKey string `validate:"required,gt=1,lte=100" json:"binding_key"`
	Email      string `validate:"required,gt=1,lte=512,email" json:"email"`
	// If must is true, whatever email if exists, try to bind user.
	// If must is false, when email exist, will only be prompted with a warning.
	Must bool `json:"must"`
}

// ExternalLoginBindingUserSendEmailResp external login binding user response
type ExternalLoginBindingUserSendEmailResp struct {
	EmailExistAndMustBeConfirmed bool   `json:"email_exist_and_must_be_confirmed"`
	AccessToken                  string `json:"access_token"`
}

// ExternalLoginBindingUserReq external login binding user request
type ExternalLoginBindingUserReq struct {
	Code    string `validate:"required,gt=0,lte=500" json:"code"`
	Content string `json:"-"`
}

// ExternalLoginBindingUserResp external login binding user response
type ExternalLoginBindingUserResp struct {
	AccessToken string `json:"access_token"`
}

// ExternalLoginUserInfoCache external login user info
type ExternalLoginUserInfoCache struct {
	// Third party identification
	// e.g. facebook, twitter, instagram
	Provider string
	// required. The unique user ID provided by the third-party login
	ExternalID string
	// optional. This name is used preferentially during registration
	DisplayName string
	// optional. This username is used preferentially during registration
	Username string
	// optional. If email exist will bind the existing user
	Email string
	// optional. The avatar URL provided by the third-party login platform
	Avatar string
	// optional. The original user information provided by the third-party login platform
	MetaInfo string
	// optional. The bio provided by the third-party login platform
	Bio string
}

// ExternalLoginUnbindingReq external login unbinding user
type ExternalLoginUnbindingReq struct {
	ExternalID string `validate:"required,gt=0,lte=128" json:"external_id"`
	UserID     string `json:"-"`
}

// UserCenterUserSettingsResp user center user info response
type UserCenterUserSettingsResp struct {
	ProfileSettingAgent UserSettingAgent `json:"profile_setting_agent"`
	AccountSettingAgent UserSettingAgent `json:"account_setting_agent"`
}

type UserCenterAdminFunctionAgentResp struct {
	AllowCreateUser         bool `json:"allow_create_user"`
	AllowUpdateUserStatus   bool `json:"allow_update_user_status"`
	AllowUpdateUserPassword bool `json:"allow_update_user_password"`
	AllowUpdateUserRole     bool `json:"allow_update_user_role"`
}

type UserSettingAgent struct {
	Enabled     bool   `json:"enabled"`
	RedirectURL string `json:"redirect_url"`
}
