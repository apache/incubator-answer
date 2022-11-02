package schema

import "time"

const (
	CGDefault = 1
	CGDIY     = 2
)

// CollectionSwitchReq switch collection request
type CollectionSwitchReq struct {
	// object TagID
	ObjectID string `validate:"required" json:"object_id"`
	// user collection group TagID
	GroupID string `validate:"required" json:"group_id"`
}

// CollectionSwitchDTO  collection data transfer object
type CollectionSwitchDTO struct {
	ObjectID string
	GroupID  string
	UserID   string
}

// CollectionSwitchResp switch collection response
type CollectionSwitchResp struct {
	ObjectID              string `json:"object_id"`
	Switch                bool   `json:"switch"`
	ObjectCollectionCount string `json:"object_collection_count"`
}

// AddCollectionGroupReq add collection group request
type AddCollectionGroupReq struct {
	//
	UserID int64 `validate:"required" comment:"" json:"user_id"`
	// the collection group name
	Name string `validate:"required,gt=0,lte=50" comment:"the collection group name" json:"name"`
	// mark this group is default, default 1
	DefaultGroup int `validate:"required" comment:"mark this group is default, default 1" json:"default_group"`
	//
	CreateTime time.Time `validate:"required" comment:"" json:"create_time"`
	//
	UpdateTime time.Time `validate:"required" comment:"" json:"update_time"`
}

// UpdateCollectionGroupReq update collection group request
type UpdateCollectionGroupReq struct {
	//
	ID int64 `validate:"required" comment:"" json:"id"`
	//
	UserID int64 `validate:"omitempty" comment:"" json:"user_id"`
	// the collection group name
	Name string `validate:"omitempty,gt=0,lte=50" comment:"the collection group name" json:"name"`
	// mark this group is default, default 1
	DefaultGroup int `validate:"omitempty" comment:"mark this group is default, default 1" json:"default_group"`
	//
	CreateTime time.Time `validate:"omitempty" comment:"" json:"create_time"`
	//
	UpdateTime time.Time `validate:"omitempty" comment:"" json:"update_time"`
}

// GetCollectionGroupResp get collection group response
type GetCollectionGroupResp struct {
	//
	ID int64 `json:"id"`
	//
	UserID int64 `json:"user_id"`
	// the collection group name
	Name string `json:"name"`
	// mark this group is default, default 1
	DefaultGroup int `json:"default_group"`
	//
	CreateTime time.Time `json:"create_time"`
	//
	UpdateTime time.Time `json:"update_time"`
}
