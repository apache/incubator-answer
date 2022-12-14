package siteinfo_common

import (
	"testing"

	"github.com/answerdev/answer/internal/base/constant"
	"github.com/answerdev/answer/internal/entity"
	"github.com/answerdev/answer/internal/service/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var (
	mockSiteInfoRepo *mock.MockSiteInfoRepo
)

func mockInit(ctl *gomock.Controller) {
	mockSiteInfoRepo = mock.NewMockSiteInfoRepo(ctl)
	mockSiteInfoRepo.EXPECT().GetByType(gomock.Any(), constant.SiteTypeGeneral).
		Return(&entity.SiteInfo{Content: `{"name":"name"}`}, true, nil)
}

func TestSiteInfoCommonService_GetSiteGeneral(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mockInit(ctl)
	siteInfoCommonService := NewSiteInfoCommonService(mockSiteInfoRepo)
	resp, err := siteInfoCommonService.GetSiteGeneral(nil)
	assert.NoError(t, err)
	assert.Equal(t, resp.Name, "name")
}
