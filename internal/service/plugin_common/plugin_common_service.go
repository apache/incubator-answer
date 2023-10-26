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

package plugin_common

import (
	"context"
	"encoding/json"
	"github.com/answerdev/answer/internal/base/data"
	"github.com/answerdev/answer/internal/repo/search_sync"

	"github.com/answerdev/answer/internal/base/constant"
	"github.com/answerdev/answer/internal/base/reason"
	"github.com/answerdev/answer/internal/entity"
	"github.com/answerdev/answer/internal/schema"
	"github.com/answerdev/answer/internal/service/config"
	"github.com/answerdev/answer/plugin"
	"github.com/segmentfault/pacman/errors"
	"github.com/segmentfault/pacman/log"
)

type PluginConfigRepo interface {
	SavePluginConfig(ctx context.Context, pluginSlugName, configValue string) (err error)
	GetPluginConfigAll(ctx context.Context) (pluginConfigs []*entity.PluginConfig, err error)
}

// PluginCommonService user service
type PluginCommonService struct {
	configService    *config.ConfigService
	pluginConfigRepo PluginConfigRepo
	data             *data.Data
}

// NewPluginCommonService new report service
func NewPluginCommonService(
	pluginConfigRepo PluginConfigRepo,
	configService *config.ConfigService,
	data *data.Data,
) *PluginCommonService {

	// init plugin status
	pluginStatus, err := configService.GetStringValue(context.TODO(), constant.PluginStatus)
	if err != nil {
		log.Error(err)
	} else {
		if err := plugin.StatusManager.UnmarshalJSON([]byte(pluginStatus)); err != nil {
			log.Error(err)
		}
	}

	// init plugin config
	pluginConfigs, err := pluginConfigRepo.GetPluginConfigAll(context.Background())
	if err != nil {
		log.Error(err)
	} else {
		for _, pluginConfig := range pluginConfigs {
			err := plugin.CallConfig(func(fn plugin.Config) error {
				if fn.Info().SlugName == pluginConfig.PluginSlugName {
					return fn.ConfigReceiver([]byte(pluginConfig.Value))
				}
				return nil
			})
			if err != nil {
				log.Errorf("parse plugin config failed: %s %v", pluginConfig.PluginSlugName, err)
			}
		}
	}

	return &PluginCommonService{
		configService:    configService,
		pluginConfigRepo: pluginConfigRepo,
		data:             data,
	}
}

// UpdatePluginStatus update plugin status
func (ps *PluginCommonService) UpdatePluginStatus(ctx context.Context) (err error) {
	content, err := plugin.StatusManager.MarshalJSON()
	if err != nil {
		return errors.InternalServer(reason.UnknownError).WithError(err)
	}
	return ps.configService.UpdateConfig(ctx, constant.PluginStatus, string(content))
}

// UpdatePluginConfig update plugin config
func (ps *PluginCommonService) UpdatePluginConfig(ctx context.Context, req *schema.UpdatePluginConfigReq) (err error) {
	configValue, _ := json.Marshal(req.ConfigFields)
	err = ps.pluginConfigRepo.SavePluginConfig(ctx, req.PluginSlugName, string(configValue))
	if err != nil {
		return err
	}

	_ = plugin.CallSearch(func(search plugin.Search) error {
		if search.Info().SlugName == req.PluginSlugName {
			search.RegisterSyncer(ctx, search_sync.NewPluginSyncer(ps.data))
		}
		return nil
	})
	return nil
}
