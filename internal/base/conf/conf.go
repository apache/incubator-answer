package conf

import (
	"bytes"
	"path/filepath"

	"github.com/answerdev/answer/internal/base/data"
	"github.com/answerdev/answer/internal/base/server"
	"github.com/answerdev/answer/internal/base/translator"
	"github.com/answerdev/answer/internal/cli"
	"github.com/answerdev/answer/internal/router"
	"github.com/answerdev/answer/internal/service/service_config"
	"github.com/answerdev/answer/pkg/writer"
	"github.com/segmentfault/pacman/contrib/conf/viper"
	"gopkg.in/yaml.v3"
)

// AllConfig all config
type AllConfig struct {
	Debug         bool                          `json:"debug" mapstructure:"debug" yaml:"debug"`
	Server        *Server                       `json:"server" mapstructure:"server" yaml:"server"`
	Data          *Data                         `json:"data" mapstructure:"data" yaml:"data"`
	I18n          *translator.I18n              `json:"i18n" mapstructure:"i18n" yaml:"i18n"`
	ServiceConfig *service_config.ServiceConfig `json:"service_config" mapstructure:"service_config" yaml:"service_config"`
	Swaggerui     *router.SwaggerConfig         `json:"swaggerui" mapstructure:"swaggerui" yaml:"swaggerui"`
}

type PathIgnore struct {
	Users []string `yaml:"users"`
}

// Server server config
type Server struct {
	HTTP *server.HTTP `json:"http" mapstructure:"http" yaml:"http"`
}

// Data data config
type Data struct {
	Database *data.Database  `json:"database" mapstructure:"database" yaml:"database"`
	Cache    *data.CacheConf `json:"cache" mapstructure:"cache" yaml:"cache"`
}

// ReadConfig read config
func ReadConfig(configFilePath string) (c *AllConfig, err error) {
	if len(configFilePath) == 0 {
		configFilePath = filepath.Join(cli.ConfigFileDir, cli.DefaultConfigFileName)
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
	buf := bytes.Buffer{}
	enc := yaml.NewEncoder(&buf)
	enc.SetIndent(2)
	if err := enc.Encode(allConfig); err != nil {
		return err
	}
	return writer.ReplaceFile(configFilePath, buf.String())
}
