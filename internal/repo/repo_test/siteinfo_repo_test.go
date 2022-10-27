package repo_test

import (
	"context"
	"testing"

	"github.com/answerdev/answer/internal/entity"
	"github.com/answerdev/answer/internal/repo/site_info"
	"github.com/stretchr/testify/assert"
)

func Test_siteInfoRepo_SaveByType(t *testing.T) {
	siteInfoRepo := site_info.NewSiteInfo(testDataSource)

	data := &entity.SiteInfo{Content: "site_info", Type: "test"}

	err := siteInfoRepo.SaveByType(context.TODO(), data.Type, data)
	assert.NoError(t, err)

	got, exist, err := siteInfoRepo.GetByType(context.TODO(), data.Type)
	assert.NoError(t, err)
	assert.True(t, exist)
	assert.Equal(t, data.Content, got.Content)

	data.Content = "new site_info"
	err = siteInfoRepo.SaveByType(context.TODO(), data.Type, data)
	assert.NoError(t, err)

	got, exist, err = siteInfoRepo.GetByType(context.TODO(), data.Type)
	assert.NoError(t, err)
	assert.True(t, exist)
	assert.Equal(t, data.Content, got.Content)
}
