package data

// Database database config
type Database struct {
	Connection      string `json:"connection" mapstructure:"connection"`
	ConnMaxLifeTime int    `json:"conn_max_life_time" mapstructure:"conn_max_life_time"`
	MaxOpenConn     int    `json:"max_open_conn" mapstructure:"max_open_conn"`
	MaxIdleConn     int    `json:"max_idle_conn" mapstructure:"max_idle_conn"`
}

// CacheConf cache
type CacheConf struct {
	FilePath string `json:"file_path" mapstructure:"file_path"`
}
