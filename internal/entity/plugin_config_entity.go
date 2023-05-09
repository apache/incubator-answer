package entity

// PluginConfig plugin config
type PluginConfig struct {
	ID             int    `xorm:"not null pk autoincr INT(11) id"`
	PluginSlugName string `xorm:"unique VARCHAR(128) plugin_slug_name"`
	Value          string `xorm:"TEXT value"`
}

// TableName config table name
func (PluginConfig) TableName() string {
	return "plugin_config"
}
