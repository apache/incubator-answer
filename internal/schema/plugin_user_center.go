package schema

type UserCenterAgentResp struct {
	Enabled   bool       `json:"enabled"`
	AgentInfo *AgentInfo `json:"agent_info"`
}

type AgentInfo struct {
	Name                      string           `json:"name"`
	DisplayName               string           `json:"display_name"`
	Icon                      string           `json:"icon"`
	Url                       string           `json:"url"`
	LoginRedirectURL          string           `json:"login_redirect_url"`
	SignUpRedirectURL         string           `json:"sign_up_redirect_url"`
	ControlCenterItems        []*ControlCenter `json:"control_center"`
	EnabledOriginalUserSystem bool             `json:"enabled_original_user_system"`
}

type ControlCenter struct {
	Name  string `json:"name"`
	Label string `json:"label"`
	Url   string `json:"url"`
}

type UserCenterPersonalBranding struct {
	Enabled          bool                `json:"enabled"`
	PersonalBranding []*PersonalBranding `json:"personal_branding"`
}

type PersonalBranding struct {
	Icon  string `json:"icon"`
	Name  string `json:"name"`
	Label string `json:"label"`
	Url   string `json:"url"`
}
