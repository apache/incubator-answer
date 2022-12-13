package cron

import "github.com/answerdev/answer/internal/service/siteinfo_common"

// ScheduledTaskManager scheduled task manager
type ScheduledTaskManager struct {
	siteInfoService *siteinfo_common.SiteInfoCommonService
}

// NewScheduledTaskManager new scheduled task manager
func NewScheduledTaskManager(siteInfoService *siteinfo_common.SiteInfoCommonService) *ScheduledTaskManager {
	manager := &ScheduledTaskManager{siteInfoService: siteInfoService}
	return manager
}

func (s *ScheduledTaskManager) Run() {

}
