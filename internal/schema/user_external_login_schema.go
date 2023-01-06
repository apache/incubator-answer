package schema

// UserExternalLoginResp user external login resp
type UserExternalLoginResp struct {
	BindingKey  string `json:"binding_key"`
	AccessToken string `json:"access_token"`
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
	EmailExistAndMustBeConfirmed bool `json:"email_exist_and_must_be_confirmed"`
}

// ExternalLoginBindingUserReq external login binding user request
type ExternalLoginBindingUserReq struct {
	Code string `validate:"required,gt=0,lte=500" json:"code"`
}

// ExternalLoginBindingUserResp external login binding user response
type ExternalLoginBindingUserResp struct {
	AccessToken string `json:"access_token"`
}
