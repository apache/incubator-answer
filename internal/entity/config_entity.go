package entity

// Config config
type Config struct {
	ID    int    `xorm:"not null pk autoincr comment('config id') INT(11) id"`
	Key   string `xorm:"comment('the config key') unique VARCHAR(32) key"`
	Value string `xorm:"comment('the config value, custom data structures and types') TEXT value"`
}

// TableName config table name
func (Config) TableName() string {
	return "config"
}
