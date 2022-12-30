package schema

type ConnectorInfoResp struct {
	Name string `json:"name"`
	Icon []byte `json:"icon"`
	Link string `json:"link"`
}
