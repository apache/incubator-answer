package plugin

type UserCenter interface {
	Base
	Description() UserCenterDesc
	ControlCenterItems() []ControlCenter
	LoginCallback(ctx *GinContext) (userInfo *UserCenterBasicUserInfo, err error)
	SignUpCallback(ctx *GinContext) (userInfo *UserCenterBasicUserInfo, err error)
	UserInfo(externalID string) (userInfo *UserCenterBasicUserInfo, err error)
	UserList(externalIDs []string) (userInfo []*UserCenterBasicUserInfo, err error)
	UserSettings(externalID string) (userSettings *SettingInfo, err error)
	PersonalBranding(externalID string) (branding []*PersonalBranding)
}

type UserCenterDesc struct {
	Name              string `json:"name"`
	Icon              string `json:"icon"`
	Url               string `json:"url"`
	LoginRedirectURL  string `json:"login_redirect_url"`
	SignUpRedirectURL string `json:"sign_up_redirect_url"`
	RankAgentEnabled  bool   `json:"rank_agent_enabled"`
}

type UserCenterBasicUserInfo struct {
	ExternalID  string `json:"external_id"`
	Username    string `json:"username"`
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
	Rank        int    `json:"rank"`
	Avatar      string `json:"avatar"`
	Mobile      string `json:"mobile"`
}

type ControlCenter struct {
	Name  string `json:"name"`
	Label string `json:"label"`
	Url   string `json:"url"`
}

type SettingInfo struct {
	ProfileSettingRedirectURL string `json:"profile_setting_redirect_url"`
	AccountSettingRedirectURL string `json:"account_setting_redirect_url"`
}

type PersonalBranding struct {
	Icon  string `json:"icon"`
	Name  string `json:"name"`
	Label string `json:"label"`
	Url   string `json:"url"`
}

var (
	// CallUserCenter is a function that calls all registered parsers
	CallUserCenter,
	registerUserCenter = MakePlugin[UserCenter](false)
)

func UserCenterEnabled() (enabled bool) {
	_ = CallUserCenter(func(fn UserCenter) error {
		enabled = true
		return nil
	})
	return
}

func RankAgentEnabled() (enabled bool) {
	_ = CallUserCenter(func(fn UserCenter) error {
		enabled = fn.Description().RankAgentEnabled
		return nil
	})
	return
}

func GetUserCenter() (uc UserCenter, ok bool) {
	_ = CallUserCenter(func(fn UserCenter) error {
		uc = fn
		ok = true
		return nil
	})
	return
}
