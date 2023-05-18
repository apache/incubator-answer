package report

import (
	"github.com/answerdev/answer/internal/entity"
	"github.com/answerdev/answer/internal/schema"
	"github.com/answerdev/answer/internal/service/object_info"
	"github.com/answerdev/answer/internal/service/report_common"
	"github.com/answerdev/answer/pkg/obj"
	"golang.org/x/net/context"
)

// ReportService user service
type ReportService struct {
	reportRepo        report_common.ReportRepo
	objectInfoService *object_info.ObjService
}

// NewReportService new report service
func NewReportService(reportRepo report_common.ReportRepo,
	objectInfoService *object_info.ObjService,
) *ReportService {
	return &ReportService{
		reportRepo:        reportRepo,
		objectInfoService: objectInfoService,
	}
}

// AddReport add report
func (rs *ReportService) AddReport(ctx context.Context, req *schema.AddReportReq) (err error) {
	objectTypeNumber, err := obj.GetObjectTypeNumberByObjectID(req.ObjectID)
	if err != nil {
		return err
	}

	// TODO this reported user id should be get by revision
	objInfo, err := rs.objectInfoService.GetInfo(ctx, req.ObjectID)
	if err != nil {
		return err
	}

	report := &entity.Report{
		UserID:         req.UserID,
		ReportedUserID: objInfo.ObjectCreatorUserID,
		ObjectID:       req.ObjectID,
		ObjectType:     objectTypeNumber,
		ReportType:     req.ReportType,
		Content:        req.Content,
		Status:         entity.ReportStatusPending,
	}
	return rs.reportRepo.AddReport(ctx, report)
}
