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
	// receiverURL is the whole URL of the receiver
	ConnectorSender(ctx *GinContext, receiverURL string) (redirectURL string)

	// ConnectorReceiver presents the receiver of the connector
	// It handles the callback endpoint of the connector, and returns the
	ConnectorReceiver(ctx *GinContext) (userInfo ExternalLoginUserInfo, err error)
}

type ExternalLoginUserInfo struct {
	ExternalID string
	Name       string
	Email      string
	MetaInfo   string
}

var (
	// CallConnector is a function that calls all registered connectors
	CallConnector,
	registerConnector = MakePlugin[Connector]()
)
