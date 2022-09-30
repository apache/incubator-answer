package schema

type SearchDTO struct {
	// Query the query string
	Query string
	// UserID current login user ID
	UserID string
	Page   int
	Size   int
}

type SearchObject struct {
	ID              string `json:"id"`
	Title           string `json:"title"`
	Excerpt         string `json:"excerpt"`
	CreatedAtParsed int64  `json:"created_at"`
	VoteCount       int    `json:"vote_count"`
	Accepted        bool   `json:"accepted"`
	AnswerCount     int    `json:"answer_count"`
	// user info
	UserInfo *UserBasicInfo `json:"user_info"`
	// tags
	Tags []TagResp `json:"tags"`
	// Status
	StatusStr string `json:"status"`
}

type TagResp struct {
	SlugName    string `json:"display_name"`
	DisplayName string `json:"slug_name"`
	// if main tag slug name is not empty, this tag is synonymous with the main tag
	MainTagSlugName string `json:"main_tag_slug_name"`
}

type SearchResp struct {
	// object_type
	ObjectType string `json:"object_type"`
	// this object
	Object SearchObject `json:"object"`
}

type SearchListResp struct {
	Total int64 `json:"count"`
	// search response
	SearchResp []SearchResp `json:"list"`
	// extra fields
	Extra interface{} `json:"extra"`
}
