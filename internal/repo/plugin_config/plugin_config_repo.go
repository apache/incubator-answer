package plugin_config

import (
	"context"

	"github.com/answerdev/answer/internal/base/data"
	"github.com/answerdev/answer/internal/base/reason"
	"github.com/answerdev/answer/internal/entity"
	"github.com/answerdev/answer/internal/service/plugin_common"
	"github.com/segmentfault/pacman/errors"
)

type pluginConfigRepo struct {
	data *data.Data
}

// NewPluginConfigRepo new repository
func NewPluginConfigRepo(data *data.Data) plugin_common.PluginConfigRepo {
	return &pluginConfigRepo{
		data: data,
	}
}

func (ur *pluginConfigRepo) SavePluginConfig(ctx context.Context, pluginSlugName, configValue string) (err error) {
	old := &entity.PluginConfig{PluginSlugName: pluginSlugName}
	exist, err := ur.data.DB.Context(ctx).Get(old)
	if err != nil {
		return errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	if exist {
		old.Value = configValue
		_, err = ur.data.DB.Context(ctx).ID(old.ID).Update(old)
	} else {
		_, err = ur.data.DB.Context(ctx).InsertOne(&entity.PluginConfig{PluginSlugName: pluginSlugName, Value: configValue})
	}
	if err != nil {
		return errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return nil
}

func (ur *pluginConfigRepo) GetPluginConfigAll(ctx context.Context) (pluginConfigs []*entity.PluginConfig, err error) {
	pluginConfigs = make([]*entity.PluginConfig, 0)
	err = ur.data.DB.Context(ctx).Find(&pluginConfigs)
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return pluginConfigs, err
}
