/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package conf

import (
	"bytes"
	"path/filepath"

	"github.com/apache/incubator-answer/internal/base/data"
	"github.com/apache/incubator-answer/internal/base/server"
	"github.com/apache/incubator-answer/internal/base/translator"
	"github.com/apache/incubator-answer/internal/cli"
	"github.com/apache/incubator-answer/internal/router"
	"github.com/apache/incubator-answer/internal/service/service_config"
	"github.com/apache/incubator-answer/pkg/writer"
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
	UI            *server.UI                    `json:"ui" mapstructure:"ui" yaml:"ui"`
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

// SetDefault set default config
func (c *AllConfig) SetDefault() {
	if c.UI == nil {
		c.UI = &server.UI{}
	}
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
	c.SetDefault()
	return c, nil
}

// RewriteConfig rewrite config file path
func RewriteConfig(configFilePath string, allConfig *AllConfig) error {
	buf := bytes.Buffer{}
	enc := yaml.NewEncoder(&buf)
	defer enc.Close()
	enc.SetIndent(2)
	if err := enc.Encode(allConfig); err != nil {
		return err
	}
	return writer.ReplaceFile(configFilePath, buf.String())
}
