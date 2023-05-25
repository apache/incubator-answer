package config

import "context"

// ConfigRepo config repository
type ConfigRepo interface {
	Get(key string) (interface{}, error)
	GetString(key string) (string, error)
	GetInt(key string) (int, error)
	GetArrayString(key string) ([]string, error)
	GetConfigType(key string) (int, error)
	GetJsonConfigByIDAndSetToObject(id int, value any) (err error)
	SetConfig(ctx context.Context, key, value string) (err error)
}

// ConfigService user service
type ConfigService struct {
	configRepo ConfigRepo
}

func NewConfigService(configRepo ConfigRepo) *ConfigService {
	return &ConfigService{
		configRepo: configRepo,
	}
}
