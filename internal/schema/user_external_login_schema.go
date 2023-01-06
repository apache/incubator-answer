package schema

// UserExternalLoginResp user external login resp
type UserExternalLoginResp struct {
	ExternalID  string `json:"external_id"`
	AccessToken string `json:"access_token"`
}
