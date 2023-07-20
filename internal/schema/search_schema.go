package schema

import "github.com/answerdev/answer/internal/base/constant"

type SearchDTO struct {
	UserID string // UserID current login user ID
	Query  string `validate:"required,gte=1,lte=60" json:"q" form:"q"`                   // Query the query string
	Page   int    `validate:"omitempty,min=1" form:"page,default=1" json:"page"`         //Query number of pages
	Size   int    `validate:"omitempty,min=1,max=50" form:"size,default=30" json:"size"` //Search page size
	Order  string `validate:"required,oneof=newest active score relevance" form:"order,default=relevance" json:"order" enums:"newest,active,score,relevance"`
}

type SearchCondition struct {
	// search target type: all/question/answer
	TargetType string
	// search query user id
	UserID string
	// vote amount
	VoteAmount int
	// only show not accepted answer's question
	NotAccepted bool
	// view amount
	Views int
	// answer count
	AnswerAmount int
	// only show accepted answer
	Accepted bool
	// only show this question's answer
	QuestionID string
	// search query tags
	Tags []string
	// search query keywords
	Words []string
}

// SearchAll check if search all
func (s *SearchCondition) SearchAll() bool {
	return len(s.TargetType) == 0
}

// SearchQuestion check if search only need question
func (s *SearchCondition) SearchQuestion() bool {
	return s.TargetType == constant.QuestionObjectType
}

// SearchAnswer check if search only need answer
func (s *SearchCondition) SearchAnswer() bool {
	return s.TargetType == constant.AnswerObjectType
}

type SearchObject struct {
	ID              string `json:"id"`
	QuestionID      string `json:"question_id"`
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
	ID          string `json:"-"`
	SlugName    string `json:"slug_name"`
	DisplayName string `json:"display_name"`
	// if main tag slug name is not empty, this tag is synonymous with the main tag
	MainTagSlugName string `json:"main_tag_slug_name"`
	Recommend       bool   `json:"recommend"`
	Reserved        bool   `json:"reserved"`
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
