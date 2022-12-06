package entity

// Config config
type Config struct {
	ID    int    `xorm:"not null pk autoincr INT(11) id"`
	Key   string `xorm:"unique VARCHAR(128) key"`
	Value string `xorm:"TEXT value"`
}

// TableName config table name
func (Config) TableName() string {
	return "config"
}
