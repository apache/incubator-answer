package entity

// Uniqid uniqid
type Uniqid struct {
	ID         int64 `xorm:"not null pk autoincr BIGINT(20) id"`
	UniqidType int   `xorm:"not null default 0 INT(11) uniqid_type"`
}

// TableName uniqid table name
func (Uniqid) TableName() string {
	return "uniqid"
}
