package schema

type VoteReq struct {
	ObjectID string `validate:"required"form:"object_id" json:"object_id"`  //	 id
	IsCancel bool   `validate:"omitempty"form:"is_cancel" json:"is_cancel"` // is cancel
}

type VoteDTO struct {
	// object TagID
	ObjectID string
	// is cancel
	IsCancel bool
	// user TagID
	UserID string
}

type VoteResp struct {
	UpVotes    int    `json:"up_votes"`
	DownVotes  int    `json:"down_votes"`
	Votes      int    `json:"votes"`
	VoteStatus string `json:"vote_status"`
}

type GetVoteWithPageReq struct {
	// page
	Page int `validate:"omitempty,min=1" form:"page"`
	// page size
	PageSize int `validate:"omitempty,min=1" form:"page_size"`
	// user id
	UserID string `validate:"required" form:"user_id"`
}

type VoteQuestion struct {
	// object ID
	ID string `json:"id"`
	// title
	Title string `json:"title"`
}

type VoteAnswer struct {
	// object ID
	ID string `json:"id"`
	// question ID
	QuestionID string `json:"question_id"`
	// title
	Title string `json:"title"`
}

type GetVoteWithPageResp struct {
	// create time
	CreatedAt int64 `json:"created_at"`
	// object id
	ObjectID string `json:"object_id"`
	// question id
	QuestionID string `json:"question_id"`
	// answer id
	AnswerID string `json:"answer_id"`
	// object type
	ObjectType string `json:"object_type" enums:"question,answer,tag,comment"`
	// title
	Title string `json:"title"`
	// content
	Content string `json:"content"`
	// vote type
	VoteType string `json:"vote_type"`
}
