package schema

// UpdateUserStatusReq update user request
type UpdateUserStatusReq struct {
	// user id
	UserID string `validate:"required" json:"user_id"`
	// user status
	Status string `validate:"required,oneof=normal suspended deleted inactive" json:"status" enums:"normal,suspended,deleted,inactive"`
}

const (
	UserNormal    = "normal"
	UserSuspended = "suspended"
	UserDeleted   = "deleted"
	UserInactive  = "inactive"
)

func (r *UpdateUserStatusReq) IsNormal() bool    { return r.Status == UserNormal }
func (r *UpdateUserStatusReq) IsSuspended() bool { return r.Status == UserSuspended }
func (r *UpdateUserStatusReq) IsDeleted() bool   { return r.Status == UserDeleted }
func (r *UpdateUserStatusReq) IsInactive() bool  { return r.Status == UserInactive }

// GetUserPageReq get user list page request
type GetUserPageReq struct {
	// page
	Page int `validate:"omitempty,min=1" form:"page"`
	// page size
	PageSize int `validate:"omitempty,min=1" form:"page_size"`
	// email
	Query string `validate:"omitempty,gt=0,lte=100" form:"query"`
	// user status
	Status string `validate:"omitempty,oneof=suspended deleted inactive" form:"status"`
}

func (r *GetUserPageReq) IsSuspended() bool { return r.Status == UserSuspended }
func (r *GetUserPageReq) IsDeleted() bool   { return r.Status == UserDeleted }
func (r *GetUserPageReq) IsInactive() bool  { return r.Status == UserInactive }

// GetUserPageResp get user response
type GetUserPageResp struct {
	// user id
	UserID string `json:"user_id"`
	// create time
	CreatedAt int64 `json:"created_at"`
	// delete time
	DeletedAt int64 `json:"deleted_at"`
	// suspended time
	SuspendedAt int64 `json:"suspended_at"`
	// username
	Username string `json:"username"`
	// email
	EMail string `json:"e_mail"`
	// rank
	Rank int `json:"rank"`
	// user status(normal,suspended,deleted,inactive)
	Status string `json:"status"`
	// display name
	DisplayName string `json:"display_name"`
	// avatar
	Avatar string `json:"avatar"`
}

// GetUserInfoReq get user request
type GetUserInfoReq struct {
	UserID string `validate:"required" json:"user_id"`
}

// GetUserInfoResp get user response
type GetUserInfoResp struct {
}
