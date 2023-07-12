package repo_test

import (
	"testing"

	"github.com/answerdev/answer/internal/schema"
	"github.com/stretchr/testify/assert"
)

func Test_configRepo_Get(t *testing.T) {
	configRepo := config_common.NewConfigRepo(testDataSource)
	_, err := configRepo.Get("email.config")
	assert.NoError(t, err)
}

func Test_configRepo_GetArrayString(t *testing.T) {
	configRepo := config_common.NewConfigRepo(testDataSource)
	got, err := configRepo.GetArrayString("daily_rank_limit.exclude")
	assert.NoError(t, err)
	assert.Equal(t, 1, len(got))
	assert.Equal(t, "answer.accepted", got[0])
}

func Test_configRepo_GetConfigById(t *testing.T) {
	configRepo := config_common.NewConfigRepo(testDataSource)

	closeInfo := &schema.GetReportTypeResp{}
	err := configRepo.GetJsonConfigByIDAndSetToObject(74, closeInfo)

	assert.NoError(t, err)
	assert.Equal(t, "needs close", closeInfo.Name)
}

func Test_configRepo_GetConfigType(t *testing.T) {
	configRepo := config_common.NewConfigRepo(testDataSource)
	configType, err := configRepo.GetConfigType("answer.accepted")
	assert.NoError(t, err)
	assert.Equal(t, 1, configType)
}

func Test_configRepo_GetInt(t *testing.T) {
	configRepo := config_common.NewConfigRepo(testDataSource)
	got, err := configRepo.GetInt("answer.accepted")
	assert.NoError(t, err)
	assert.Equal(t, 15, got)
}

func Test_configRepo_GetString(t *testing.T) {
	configRepo := config_common.NewConfigRepo(testDataSource)
	_, err := configRepo.GetString("email.config")
	assert.NoError(t, err)
}

func Test_configRepo_SetConfig(t *testing.T) {
	configRepo := config_common.NewConfigRepo(testDataSource)
	got, err := configRepo.GetString("email.config")
	assert.NoError(t, err)

	err = configRepo.SetConfig("email.config", got)
	assert.NoError(t, err)
}
