package plugin_common

import (
	"context"
	"encoding/json"

	"github.com/answerdev/answer/internal/base/constant"
	"github.com/answerdev/answer/internal/base/reason"
	"github.com/answerdev/answer/internal/entity"
	"github.com/answerdev/answer/internal/plugin"
	"github.com/answerdev/answer/internal/schema"
	"github.com/answerdev/answer/internal/service/config"
	"github.com/segmentfault/pacman/errors"
	"github.com/segmentfault/pacman/log"
)

type PluginConfigRepo interface {
	SavePluginConfig(ctx context.Context, pluginSlugName, configValue string) (err error)
	GetPluginConfigAll(ctx context.Context) (pluginConfigs []*entity.PluginConfig, err error)
}

// PluginCommonService user service
type PluginCommonService struct {
	configRepo       config.ConfigRepo
	PluginConfigRepo PluginConfigRepo
}

// NewPluginCommonService new report service
func NewPluginCommonService(
	PluginConfigRepo PluginConfigRepo,
	configRepo config.ConfigRepo) *PluginCommonService {

	// init plugin status
	pluginStatus, err := configRepo.GetString(constant.PluginStatus)
	if err != nil {
		log.Error(err)
	} else {
		if err := plugin.StatusManager.UnmarshalJSON([]byte(pluginStatus)); err != nil {
			log.Error(err)
		}
	}

	// init plugin config
	pluginConfigs, err := PluginConfigRepo.GetPluginConfigAll(context.Background())
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
		configRepo:       configRepo,
		PluginConfigRepo: PluginConfigRepo,
	}
}

// UpdatePluginStatus update plugin status
func (ps *PluginCommonService) UpdatePluginStatus(ctx context.Context) (err error) {
	content, err := plugin.StatusManager.MarshalJSON()
	if err != nil {
		return errors.InternalServer(reason.UnknownError).WithError(err)
	}
	return ps.configRepo.SetConfig(constant.PluginStatus, string(content))
}

// UpdatePluginConfig update plugin config
func (ps *PluginCommonService) UpdatePluginConfig(ctx context.Context, req *schema.UpdatePluginConfigReq) (err error) {
	configValue, _ := json.Marshal(req.ConfigFields)
	return ps.PluginConfigRepo.SavePluginConfig(ctx, req.PluginSlugName, string(configValue))
}
