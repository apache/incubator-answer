package data

// Database database config
type Database struct {
	Driver          string `json:"driver" mapstructure:"driver" yaml:"driver"`
	Connection      string `json:"connection" mapstructure:"connection" yaml:"connection"`
	ConnMaxLifeTime int    `json:"conn_max_life_time" mapstructure:"conn_max_life_time" yaml:"conn_max_life_time,omitempty"`
	MaxOpenConn     int    `json:"max_open_conn" mapstructure:"max_open_conn" yaml:"max_open_conn,omitempty"`
	MaxIdleConn     int    `json:"max_idle_conn" mapstructure:"max_idle_conn" yaml:"max_idle_conn,omitempty"`
}

// CacheConf cache
type CacheConf struct {
	FilePath string `json:"file_path" mapstructure:"file_path" yaml:"file_path"`
}
