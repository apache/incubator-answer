package config

import (
	"github.com/segmentfault/answer/internal/base/data"
	"github.com/segmentfault/answer/internal/schema"
	"github.com/segmentfault/answer/internal/service/config"
	"github.com/segmentfault/pacman/log"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	dataSource *data.Data
	repo       config.ConfigRepo
)

func init() {
	s := "root:123456@tcp(10.0.10.200:3306)/answer_new?charset=utf8&interpolateParams=true&timeout=3s&readTimeout=3s&writeTimeout=3s"
	cache, _, _ := data.NewCache(&data.CacheConf{})
	dataSource, _, _ = data.NewData(data.NewDB(true, &data.Database{
		Connection: s,
	}), cache)
	repo = NewConfigRepo(dataSource)
}

func TestConfigRepo_GetConfigById(t *testing.T) {
	var (
		id    = 58
		value = schema.ReasonItem{}
		err   error
	)
	err = repo.GetConfigById(id, &value)
	assert.NoError(t, err)
	log.Info(value)
}
