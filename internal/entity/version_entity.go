package entity

// Version version
type Version struct {
	ID            int   `xorm:"not null pk autoincr INT(11) id"`
	VersionNumber int64 `xorm:"not null default 0 INT(11) version_number"`
}

// TableName config table name
func (Version) TableName() string {
	return "version"
}
