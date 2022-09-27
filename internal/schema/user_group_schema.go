package schema

// AddUserGroupReq add user group request
type AddUserGroupReq struct {
}

// RemoveUserGroupReq delete user group request
type RemoveUserGroupReq struct {
	// user group id
	ID int64 `validate:"required" comment:"user group id" json:"id"`
}

// UpdateUserGroupReq update user group request
type UpdateUserGroupReq struct {
	// user group id
	ID int64 `validate:"required" comment:"user group id" json:"id"`
}

// GetUserGroupListReq get user group list all request
type GetUserGroupListReq struct {
}

// GetUserGroupWithPageReq get user group list page request
type GetUserGroupWithPageReq struct {
	// page
	Page int `validate:"omitempty,min=1" form:"page"`
	// page size
	PageSize int `validate:"omitempty,min=1" form:"page_size"`
}

// GetUserGroupResp get user group response
type GetUserGroupResp struct {
	// user group id
	ID int64 `json:"id"`
}
