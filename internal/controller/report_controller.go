package controller

import (
	"github.com/answerdev/answer/internal/base/handler"
	"github.com/answerdev/answer/internal/base/middleware"
	"github.com/answerdev/answer/internal/base/reason"
	"github.com/answerdev/answer/internal/schema"
	"github.com/answerdev/answer/internal/service/permission"
	"github.com/answerdev/answer/internal/service/rank"
	"github.com/answerdev/answer/internal/service/report"
	"github.com/answerdev/answer/pkg/uid"
	"github.com/gin-gonic/gin"
	"github.com/segmentfault/pacman/errors"
)

// ReportController report controller
type ReportController struct {
	reportService *report.ReportService
	rankService   *rank.RankService
}

// NewReportController new controller
func NewReportController(reportService *report.ReportService, rankService *rank.RankService) *ReportController {
	return &ReportController{reportService: reportService, rankService: rankService}
}

// AddReport add report
// @Summary add report
// @Description add report <br> source (question, answer, comment, user)
// @Security ApiKeyAuth
// @Tags Report
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param data body schema.AddReportReq true "report"
// @Success 200 {object} handler.RespBody
// @Router /answer/api/v1/report [post]
func (rc *ReportController) AddReport(ctx *gin.Context) {
	req := &schema.AddReportReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}
	req.ObjectID = uid.DeShortID(req.ObjectID)
	req.UserID = middleware.GetLoginUserIDFromContext(ctx)
	can, err := rc.rankService.CheckOperationPermission(ctx, req.UserID, permission.ReportAdd, "")
	if err != nil {
		handler.HandleResponse(ctx, err, nil)
		return
	}
	if !can {
		handler.HandleResponse(ctx, errors.Forbidden(reason.RankFailToMeetTheCondition), nil)
		return
	}

	err = rc.reportService.AddReport(ctx, req)
	handler.HandleResponse(ctx, err, nil)
}
