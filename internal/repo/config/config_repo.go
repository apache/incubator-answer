package config

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/answerdev/answer/internal/service/config"
	"github.com/answerdev/answer/pkg/converter"

	"github.com/answerdev/answer/internal/base/data"
	"github.com/answerdev/answer/internal/base/reason"
	"github.com/answerdev/answer/internal/entity"
	"github.com/segmentfault/pacman/errors"
)

var (
	Key2ValueMapping = make(map[string]interface{})
	Key2IDMapping    = make(map[string]int)
	ID2KeyMapping    = make(map[int]string)
)

// configRepo config repository
type configRepo struct {
	data *data.Data
	mu   sync.Mutex
}

// NewConfigRepo new repository
func NewConfigRepo(data *data.Data) config.ConfigRepo {
	repo := &configRepo{
		data: data,
	}
	repo.init()
	return repo
}

// init initializes the Key2ValueMapping map data structures
func (cr *configRepo) init() {
	cr.mu.Lock()
	defer cr.mu.Unlock()
	rows := &[]entity.Config{}
	err := cr.data.DB.Context(context.TODO()).Find(rows)
	if err == nil {
		for _, row := range *rows {
			Key2ValueMapping[row.Key] = row.Value
			Key2IDMapping[row.Key] = row.ID
			ID2KeyMapping[row.ID] = row.Key
		}
	}
}

// Get Base method for getting the config value
// Key string
func (cr *configRepo) Get(key string) (interface{}, error) {
	value, ok := Key2ValueMapping[key]
	if ok {
		return value, nil
	} else {
		return value, errors.InternalServer(reason.DatabaseError).WithMsg(fmt.Sprintf("no such config key: %v", key))
	}
}

// GetString method for getting the config value to string
// key string
func (cr *configRepo) GetString(key string) (string, error) {
	value, err := cr.Get(key)
	if err != nil {
		return "", err
	}
	str, ok := value.(string)
	if !ok {
		return "", errors.InternalServer(reason.DatabaseError).WithMsg(fmt.Sprintf("config value is wrong type: %v", key))
	}
	return str, nil
}

// GetInt method for getting the config value to int64
// key string
func (cr *configRepo) GetInt(key string) (int, error) {
	value, err := cr.GetString(key)
	if err != nil {
		return 0, err
	}
	return converter.StringToInt(value), nil
}

// GetArrayString method for getting the config value to string array
func (cr *configRepo) GetArrayString(key string) ([]string, error) {
	arr := &[]string{}
	value, err := cr.GetString(key)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(value), arr)
	return *arr, err
}

// GetConfigType method for getting the config type
func (cr *configRepo) GetConfigType(key string) (int, error) {
	value, ok := Key2IDMapping[key]
	if !ok {
		return 0, errors.InternalServer(reason.DatabaseError).WithMsg(fmt.Sprintf("no such config type: %v", key))
	}
	return value, nil
}

// GetJsonConfigByIDAndSetToObject get config key from config id
func (cr *configRepo) GetJsonConfigByIDAndSetToObject(id int, object any) (err error) {
	key, ok := ID2KeyMapping[id]
	if !ok {
		return errors.InternalServer(reason.DatabaseError).WithMsg(fmt.Sprintf("no such config id: %v", id))
	}

	conf, err := cr.Get(key)
	if err != nil {
		return errors.InternalServer(reason.DatabaseError).WithError(err)
	}
	str, ok := conf.(string)
	if !ok {
		return errors.InternalServer(reason.DatabaseError).WithMsg(fmt.Sprintf("no such config id: %v", id))
	}
	err = json.Unmarshal([]byte(str), object)
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithMsg(fmt.Sprintf("no such config id: %v", id))
	}
	return
}

// SetConfig set config
func (cr *configRepo) SetConfig(ctx context.Context, key, value string) (err error) {
	id := Key2IDMapping[key]
	_, err = cr.data.DB.Context(ctx).ID(id).Update(&entity.Config{Value: value})
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	} else {
		Key2ValueMapping[key] = value
	}
	return
}
