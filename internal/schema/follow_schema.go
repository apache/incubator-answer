package schema

// FollowReq follow object request
type FollowReq struct {
	// object id
	ObjectID string `validate:"required" form:"object_id" json:"object_id"`
	// is cancel
	IsCancel bool `validate:"omitempty" form:"is_cancel" json:"is_cancel"`
}

// FollowResp response object's follows and current user follow status
type FollowResp struct {
	// the followers of object
	Follows int `json:"follows"`
	// if user is followed object will be true,otherwise false
	IsFollowed bool `json:"is_followed"`
}

type FollowDTO struct {
	// object TagID
	ObjectID string
	// is cancel
	IsCancel bool
	// user TagID
	UserID string
}

// UpdateFollowTagsReq update user follow tags
type UpdateFollowTagsReq struct {
	// tag slug name list
	SlugNameList []string `json:"slug_name_list"`
	// user id
	UserID string `json:"-"`
}
