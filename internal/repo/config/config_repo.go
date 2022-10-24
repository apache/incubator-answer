package config

import (
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
	err := cr.data.DB.Find(rows)
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
	if value != nil {
		return value.(string), err
	}
	return "", err
}

// GetInt method for getting the config value to int64
// key string
func (cr *configRepo) GetInt(key string) (int, error) {
	value, err := cr.GetString(key)
	if err != nil {
		return 0, err
	} else {
		return converter.StringToInt(value), nil
	}
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
	if ok {
		return value, nil
	} else {
		return 0, errors.InternalServer(reason.DatabaseError).WithMsg(fmt.Sprintf("no such config type: %v", key))
	}
}

// GetConfigById get config key from config id
func (cr *configRepo) GetConfigById(id int, value any) (err error) {
	var (
		ok   = true
		key  string
		conf interface{}
	)
	key, ok = ID2KeyMapping[id]
	if !ok {
		err = errors.InternalServer(reason.DatabaseError).WithMsg(fmt.Sprintf("no such config id: %v", id))
		return
	}

	conf, err = cr.Get(key)
	value = json.Unmarshal([]byte(conf.(string)), value)
	return
}

func (cr *configRepo) SetConfig(key, value string) (err error) {
	id := Key2IDMapping[key]
	_, err = cr.data.DB.ID(id).Update(&entity.Config{Value: value})
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	} else {
		Key2ValueMapping[key] = value
	}
	return
}
