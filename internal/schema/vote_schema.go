package schema

type VoteReq struct {
	ObjectID string `validate:"required" form:"object_id" json:"object_id"`  //	 id
	IsCancel bool   `validate:"omitempty" form:"is_cancel" json:"is_cancel"` // is cancel
	UserID   string `json:"-"`
}

type VoteResp struct {
	UpVotes    int64  `json:"up_votes"`
	DownVotes  int64  `json:"down_votes"`
	Votes      int64  `json:"votes"`
	VoteStatus string `json:"vote_status"`
}

// VoteOperationInfo vote operation info
type VoteOperationInfo struct {
	// operation object id
	ObjectID string
	// question answer comment
	ObjectType string
	// object owner user id
	ObjectCreatorUserID string
	// operation user id
	OperatingUserID string
	// vote up
	VoteUp bool
	// vote down
	VoteDown bool
	// vote activity info
	Activities []*VoteActivity
}

// VoteActivity vote activity
type VoteActivity struct {
	ActivityType   int
	ActivityUserID string
	TriggerUserID  string
	Rank           int
}

func (v *VoteActivity) HasRank() int {
	if v.Rank != 0 {
		return 1
	}
	return 0
}

type GetVoteWithPageReq struct {
	// page
	Page int `validate:"omitempty,min=1" form:"page"`
	// page size
	PageSize int `validate:"omitempty,min=1" form:"page_size"`
	// user id
	UserID string `json:"-"`
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
	// url title
	UrlTitle string `json:"url_title"`
	// content
	Content string `json:"content"`
	// vote type
	VoteType string `json:"vote_type"`
}
