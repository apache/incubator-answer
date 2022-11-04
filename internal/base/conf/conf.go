package conf

import (
	"path/filepath"

	"github.com/answerdev/answer/internal/base/data"
	"github.com/answerdev/answer/internal/base/server"
	"github.com/answerdev/answer/internal/base/translator"
	"github.com/answerdev/answer/internal/cli"
	"github.com/answerdev/answer/internal/router"
	"github.com/answerdev/answer/internal/service/service_config"
	"github.com/answerdev/answer/pkg/writer"
	"github.com/segmentfault/pacman/contrib/conf/viper"
	"sigs.k8s.io/yaml"
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

// ReadConfig read config
func ReadConfig(configFilePath string) (c *AllConfig, err error) {
	if len(configFilePath) == 0 {
		configFilePath = filepath.Join(cli.ConfigFilePath, cli.DefaultConfigFileName)
	}
	c = &AllConfig{}
	config, err := viper.NewWithPath(configFilePath)
	if err != nil {
		return nil, err
	}
	if err = config.Parse(&c); err != nil {
		return nil, err
	}
	return c, nil
}

// RewriteConfig rewrite config file path
func RewriteConfig(configFilePath string, allConfig *AllConfig) error {
	content, err := yaml.Marshal(allConfig)
	if err != nil {
		return err
	}
	return writer.ReplaceFile(configFilePath, string(content))
}
