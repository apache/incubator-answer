package config

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/answerdev/answer/internal/entity"
)

// ConfigRepo config repository
type ConfigRepo interface {
	GetConfigByID(ctx context.Context, id int) (c *entity.Config, err error)
	GetConfigByKey(ctx context.Context, key string) (c *entity.Config, err error)
	UpdateConfig(ctx context.Context, key, value string) (err error)
}

// ConfigService user service
type ConfigService struct {
	configRepo ConfigRepo
}

// NewConfigService new config service
func NewConfigService(configRepo ConfigRepo) *ConfigService {
	return &ConfigService{
		configRepo: configRepo,
	}
}

// GetIntValue get config int value
func (cs *ConfigService) GetIntValue(ctx context.Context, key string) (val int, err error) {
	cf, err := cs.configRepo.GetConfigByKey(ctx, key)
	if err != nil {
		return 0, err
	}
	return cf.GetIntValue(), nil
}

// GetStringValue get config string value
func (cs *ConfigService) GetStringValue(ctx context.Context, key string) (val string, err error) {
	cf, err := cs.configRepo.GetConfigByKey(ctx, key)
	if err != nil {
		return "", err
	}
	return cf.Value, nil
}

// GetArrayStringValue get config array string value
func (cs *ConfigService) GetArrayStringValue(ctx context.Context, key string) (val []string, err error) {
	cf, err := cs.configRepo.GetConfigByKey(ctx, key)
	if err != nil {
		return nil, err
	}
	return cf.GetArrayStringValue(), nil
}

func (cs *ConfigService) GetJsonConfigByIDAndSetToObject(ctx context.Context, id int, obj any) (err error) {
	cf, err := cs.configRepo.GetConfigByID(ctx, id)
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(cf.Value), obj)
	if err != nil {
		return fmt.Errorf("[%s] config value is not json format", cf.Key)
	}
	return nil
}

// GetConfigByID get config by id
func (cs *ConfigService) GetConfigByID(ctx context.Context, id int) (c *entity.Config, err error) {
	return cs.configRepo.GetConfigByID(ctx, id)
}

func (cs *ConfigService) GetConfigByKey(ctx context.Context, key string) (c *entity.Config, err error) {
	return cs.configRepo.GetConfigByKey(ctx, key)
}

// GetIDByKey get config id by key
func (cs *ConfigService) GetIDByKey(ctx context.Context, key string) (id int, err error) {
	cf, err := cs.configRepo.GetConfigByKey(ctx, key)
	if err != nil {
		return 0, err
	}
	return cf.ID, nil
}

func (cs *ConfigService) UpdateConfig(ctx context.Context, key, value string) (err error) {
	return cs.configRepo.UpdateConfig(ctx, key, value)
}
