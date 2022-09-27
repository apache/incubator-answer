package entity

// Uniqid uniqid
type Uniqid struct {
	ID         int64 `xorm:"not null pk autoincr comment('uniqid_id') BIGINT(20) id"`
	UniqidType int   `xorm:"not null default 0 comment('uniqid_type') INT(11) uniqid_type"`
}

// TableName uniqid table name
func (Uniqid) TableName() string {
	return "uniqid"
}
