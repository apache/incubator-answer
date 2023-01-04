package plugin

type Connector interface {
	Base

	// ConnectorLogo presents the logo binary data of the connector
	ConnectorLogo() []byte

	// ConnectorLogoContentType presents the content type of the logo
	// e.g. image/png, image/jpeg, image/gif
	ConnectorLogoContentType() string

	// ConnectorName presents the name of the connector
	// e.g. Facebook, Twitter, Instagram
	ConnectorName() string

	// ConnectorSlugName presents the slug name of the connector
	// Please use lowercase and hyphen as the separator
	// e.g. facebook, twitter, instagram
	ConnectorSlugName() string

	// ConnectorSender presents the sender of the connector
	// It handles the start endpoint of the connector
	ConnectorSender(ctx *GinContext)
	ConnectorReceiver(ctx *GinContext)

	//ConnectorLoginURL(redirectURL string) (loginURL string)
	//ConnectorLoginUserInfo(code string) (userInfo *UserExternalLogin, err error)
}

type UserExternalLogin struct {
	Provider    string
	ExternalID  string
	Email       string
	Name        string
	FirstName   string
	LastName    string
	NickName    string
	Description string
	AvatarUrl   string
}

var (
	// CallConnector is a function that calls all registered connectors
	CallConnector,
	registerConnector = MakePlugin[Connector]()
)
