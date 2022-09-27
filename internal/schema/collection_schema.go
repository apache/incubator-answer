package schema

import "time"

// AddCollectionReq add collection request
type AddCollectionReq struct {
	// user id
	UserID int64 `validate:"required" comment:"user id" json:"user_id"`
	// object id
	ObjectID int64 `validate:"required" comment:"object id" json:"object_id"`
	// user collection group id
	UserCollectionGroupID int64 `validate:"required" comment:"user collection group id" json:"user_collection_group_id"`
	//
	CreateTime time.Time `validate:"required" comment:"" json:"create_time"`
	//
	UpdateTime time.Time `validate:"required" comment:"" json:"update_time"`
}

// RemoveCollectionReq delete collection request
type RemoveCollectionReq struct {
	// collection id
	ID int64 `validate:"required" comment:"collection id" json:"id"`
}

// UpdateCollectionReq update collection request
type UpdateCollectionReq struct {
	// collection id
	ID int64 `validate:"required" comment:"collection id" json:"id"`
	// user id
	UserID int64 `validate:"omitempty" comment:"user id" json:"user_id"`
	// object id
	ObjectID int64 `validate:"omitempty" comment:"object id" json:"object_id"`
	// user collection group id
	UserCollectionGroupID int64 `validate:"omitempty" comment:"user collection group id" json:"user_collection_group_id"`
	//
	CreateTime time.Time `validate:"omitempty" comment:"" json:"create_time"`
	//
	UpdateTime time.Time `validate:"omitempty" comment:"" json:"update_time"`
}

// GetCollectionListReq get collection list all request
type GetCollectionListReq struct {
	// user id
	UserID int64 `validate:"omitempty" comment:"user id" form:"user_id"`
	// object id
	ObjectID int64 `validate:"omitempty" comment:"object id" form:"object_id"`
	// user collection group id
	UserCollectionGroupID int64 `validate:"omitempty" comment:"user collection group id" form:"user_collection_group_id"`
	//
	CreateTime time.Time `validate:"omitempty" comment:"" form:"create_time"`
	//
	UpdateTime time.Time `validate:"omitempty" comment:"" form:"update_time"`
}

// GetCollectionWithPageReq get collection list page request
type GetCollectionWithPageReq struct {
	// page
	Page int `validate:"omitempty,min=1" form:"page"`
	// page size
	PageSize int `validate:"omitempty,min=1" form:"page_size"`
	// user id
	UserID int64 `validate:"omitempty" comment:"user id" form:"user_id"`
	// object id
	ObjectID int64 `validate:"omitempty" comment:"object id" form:"object_id"`
	// user collection group id
	UserCollectionGroupID int64 `validate:"omitempty" comment:"user collection group id" form:"user_collection_group_id"`
	//
	CreateTime time.Time `validate:"omitempty" comment:"" form:"create_time"`
	//
	UpdateTime time.Time `validate:"omitempty" comment:"" form:"update_time"`
}

// GetCollectionResp get collection response
type GetCollectionResp struct {
	// collection id
	ID int64 `json:"id"`
	// user id
	UserID int64 `json:"user_id"`
	// object id
	ObjectID int64 `json:"object_id"`
	// user collection group id
	UserCollectionGroupID int64 `json:"user_collection_group_id"`
	//
	CreateTime time.Time `json:"create_time"`
	//
	UpdateTime time.Time `json:"update_time"`
}
