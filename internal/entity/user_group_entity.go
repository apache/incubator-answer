package entity

// UserGroup user group
type UserGroup struct {
	ID int64 `xorm:"not null pk autoincr comment('user group id') unique BIGINT(20) id"`
}

// TableName user group table name
func (UserGroup) TableName() string {
	return "user_group"
}
