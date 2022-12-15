package cron

import (
	"context"
	"fmt"

	"github.com/answerdev/answer/internal/service"
	"github.com/answerdev/answer/internal/service/siteinfo_common"
	"github.com/robfig/cron/v3"
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
	fmt.Println("start cron")
	s.questionService.SitemapCron(context.Background())
	c := cron.New()
	c.AddFunc("0 */1 * * *", func() {
		ctx := context.Background()
		fmt.Println("sitemap cron execution")
		s.questionService.SitemapCron(ctx)
	})
	c.Start()
}
