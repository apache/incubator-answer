package plugin

type Connector interface {
	Base

	// ConnectorLogoSVG presents the logo in svg format
	ConnectorLogoSVG() string

	// ConnectorName presents the name of the connector
	// e.g. Facebook, Twitter, Instagram
	ConnectorName() Translator

	// ConnectorSlugName presents the slug name of the connector
	// Please use lowercase and hyphen as the separator
	// e.g. facebook, twitter, instagram
	ConnectorSlugName() string

	// ConnectorSender presents the sender of the connector
	// It handles the start endpoint of the connector
	// receiverURL is the whole URL of the receiver
	ConnectorSender(ctx *GinContext, receiverURL string) (redirectURL string)

	// ConnectorReceiver presents the receiver of the connector
	// It handles the callback endpoint of the connector, and returns the
	ConnectorReceiver(ctx *GinContext) (userInfo ExternalLoginUserInfo, err error)
}

// ExternalLoginUserInfo external login user info
type ExternalLoginUserInfo struct {
	// required. The unique user ID provided by the third-party login
	ExternalID string
	// optional. This name is used preferentially during registration
	Name string
	// optional. If email exist will bind the existing user
	Email string
	// optional. The original user information provided by the third-party login platform
	MetaInfo string
}

var (
	// CallConnector is a function that calls all registered connectors
	CallConnector,
	registerConnector = MakePlugin[Connector](false)
)
