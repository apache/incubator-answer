package schema

type ReasonItem struct {
	ReasonType  int    `json:"reason_type"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ContentType string `json:"content_type"`
}

type ReasonReq struct {
	// ObjectType
	ObjectType string `validate:"required" form:"object_type" json:"object_type"`
	// Action
	Action string `validate:"required" form:"action" json:"action"`
}
