package controller

import (
	"github.com/answerdev/answer/internal/service/activity"
	"github.com/answerdev/answer/internal/service/activity_common"
)

type ActivityController struct {
	activityCommonService *activity_common.ActivityCommon
	activityService       *activity.ActivityService
}

// NewActivityController new activity controller.
func NewActivityController(
	activityCommonService *activity_common.ActivityCommon,
	activityService *activity.ActivityService) *ActivityController {
	return &ActivityController{activityCommonService: activityCommonService, activityService: activityService}
}
