package schema

import "time"

// AddNotificationReadReq add notification read record request
type AddNotificationReadReq struct {
	// user id
	UserID int64 `validate:"required" comment:"user id" json:"user_id"`
	// message id
	MessageID int64 `validate:"required" comment:"message id" json:"message_id"`
	// read status(unread: 1; read 2)
	IsRead int `validate:"required" comment:"read status(unread: 1; read 2)" json:"is_read"`
}

// RemoveNotificationReadReq delete notification read record request
type RemoveNotificationReadReq struct {
	// id
	ID int `validate:"required" comment:"id" json:"id"`
}

// UpdateNotificationReadReq update notification read record request
type UpdateNotificationReadReq struct {
	// id
	ID int `validate:"required" comment:"id" json:"id"`
	// user id
	UserID int64 `validate:"omitempty" comment:"user id" json:"user_id"`
	// message id
	MessageID int64 `validate:"omitempty" comment:"message id" json:"message_id"`
	// read status(unread: 1; read 2)
	IsRead int `validate:"omitempty" comment:"read status(unread: 1; read 2)" json:"is_read"`
}

// GetNotificationReadListReq get notification read record list all request
type GetNotificationReadListReq struct {
	// user id
	UserID int64 `validate:"omitempty" comment:"user id" form:"user_id"`
	// message id
	MessageID int64 `validate:"omitempty" comment:"message id" form:"message_id"`
	// read status(unread: 1; read 2)
	IsRead int `validate:"omitempty" comment:"read status(unread: 1; read 2)" form:"is_read"`
}

// GetNotificationReadWithPageReq get notification read record list page request
type GetNotificationReadWithPageReq struct {
	// page
	Page int `validate:"omitempty,min=1" form:"page"`
	// page size
	PageSize int `validate:"omitempty,min=1" form:"page_size"`
	// user id
	UserID int64 `validate:"omitempty" comment:"user id" form:"user_id"`
	// message id
	MessageID int64 `validate:"omitempty" comment:"message id" form:"message_id"`
	// read status(unread: 1; read 2)
	IsRead int `validate:"omitempty" comment:"read status(unread: 1; read 2)" form:"is_read"`
}

// GetNotificationReadResp get notification read record response
type GetNotificationReadResp struct {
	// id
	ID int `json:"id"`
	// create time
	CreatedAt time.Time `json:"created_at"`
	// update time
	UpdatedAt time.Time `json:"updated_at"`
	// user id
	UserID int64 `json:"user_id"`
	// message id
	MessageID int64 `json:"message_id"`
	// read status(unread: 1; read 2)
	IsRead int `json:"is_read"`
}
