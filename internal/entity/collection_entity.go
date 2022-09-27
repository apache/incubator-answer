package entity

import "time"

const ()

// Collection collection
type Collection struct {
	ID                    string    `xorm:"not null pk default 0 comment('collection id') BIGINT(20) id"`
	UserID                string    `xorm:"not null default 0 comment('user id') BIGINT(20) user_id"`
	ObjectID              string    `xorm:"not null default 0 comment('object id') BIGINT(20) object_id"`
	UserCollectionGroupID string    `xorm:"not null default 0 comment('user collection group id') BIGINT(20) user_collection_group_id"`
	CreatedAt             time.Time `xorm:"created not null default CURRENT_TIMESTAMP TIMESTAMP created_at"`
	UpdatedAt             time.Time `xorm:"updated not null default CURRENT_TIMESTAMP TIMESTAMP updated_at"`
}

type CollectionSearch struct {
	Collection
	Page     int `json:"page" form:"page"`           //Query number of pages
	PageSize int `json:"page_size" form:"page_size"` //Search page size
}

// TableName collection table name
func (Collection) TableName() string {
	return "collection"
}
