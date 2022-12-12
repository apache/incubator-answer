package cron

import (
	"context"

	"github.com/answerdev/answer/internal/service"
	"github.com/answerdev/answer/internal/service/siteinfo_common"
)

// ScheduledTaskManager scheduled task manager
type ScheduledTaskManager struct {
	siteInfoService *siteinfo_common.SiteInfoCommonService
	questionService *service.QuestionService
}

// NewScheduledTaskManager new scheduled task manager
func NewScheduledTaskManager(
	siteInfoService *siteinfo_common.SiteInfoCommonService,
	questionService *service.QuestionService,
) *ScheduledTaskManager {
	manager := &ScheduledTaskManager{
		siteInfoService: siteInfoService,
		questionService: questionService,
	}
	return manager
}

func (s *ScheduledTaskManager) Run() {
	ctx := context.Background()
	s.questionService.SitemapCron(ctx)
}
