package conf

import (
	"github.com/segmentfault/answer/internal/base/data"
	"github.com/segmentfault/answer/internal/base/server"
	"github.com/segmentfault/answer/internal/base/translator"
	"github.com/segmentfault/answer/internal/router"
	"github.com/segmentfault/answer/internal/service/service_config"
)

// AllConfig all config
type AllConfig struct {
	Debug         bool                          `json:"debug" mapstructure:"debug"`
	Data          *Data                         `json:"data" mapstructure:"data"`
	Server        *Server                       `json:"server" mapstructure:"server"`
	I18n          *translator.I18n              `json:"i18n" mapstructure:"i18n"`
	Swaggerui     *router.SwaggerConfig         `json:"swaggerui" mapstructure:"swaggerui"`
	ServiceConfig *service_config.ServiceConfig `json:"service_config" mapstructure:"service_config"`
}

// Server server config
type Server struct {
	HTTP *server.HTTP `json:"http" mapstructure:"http"`
}

// Data data config
type Data struct {
	Database *data.Database  `json:"database" mapstructure:"database"`
	Cache    *data.CacheConf `json:"cache" mapstructure:"cache"`
}

// ------------------ remove

// log .
type log struct {
	Dir        string `json:"dir"`
	Name       string `json:"name"`
	Access     bool   `json:"access"`
	Level      string `json:"level"`
	MaxSize    int    `json:"max_size"`
	MaxBackups int    `json:"max_backups"`
	MaxAge     int    `json:"max_age"`
}

// Local .
type Local struct {
	Address string `json:"address"`
	Debug   bool   `json:"debug"`
	log     log    `json:"log"`
}

// // SwaggerConfig .
// type SwaggerConfig struct {
// 	Show     bool   `json:"show"`
// 	Protocol string `json:"protocol"`
// 	Host     string `json:"host"`
// 	Address  string `json:"address"`
// }

// Answer .
type Answer struct {
	MaxIdle    int    `json:"max_idle"`
	MaxOpen    int    `json:"max_open"`
	IsDebug    bool   `json:"is_debug"`
	Datasource string `json:"datasource"`
}

// Mysql .
type Mysql struct {
	Answer Answer `json:"answer"`
}
