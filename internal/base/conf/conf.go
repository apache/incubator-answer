package conf

import (
	"github.com/answerdev/answer/internal/base/data"
	"github.com/answerdev/answer/internal/base/server"
	"github.com/answerdev/answer/internal/base/translator"
	"github.com/answerdev/answer/internal/router"
	"github.com/answerdev/answer/internal/service/service_config"
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
