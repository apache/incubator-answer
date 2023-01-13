package schema

type ConnectorInfoResp struct {
	Name string `json:"name"`
	Icon string `json:"icon"`
	Link string `json:"link"`
}

type ConnectorUserInfoResp struct {
	Name       string `json:"name"`
	Icon       string `json:"icon"`
	Link       string `json:"link"`
	Binding    bool   `json:"binding"`
	ExternalID string `json:"external_id"`
}
