package report

import (
	"context"

	"github.com/segmentfault/answer/internal/base/constant"
	"github.com/segmentfault/answer/internal/base/pager"
	"github.com/segmentfault/answer/internal/schema"
	"github.com/segmentfault/answer/internal/service/report_common"

	"github.com/segmentfault/answer/internal/base/data"
	"github.com/segmentfault/answer/internal/base/reason"
	"github.com/segmentfault/answer/internal/entity"
	"github.com/segmentfault/answer/internal/service/unique"
	"github.com/segmentfault/pacman/errors"
)

// reportRepo report repository
type reportRepo struct {
	data         *data.Data
	uniqueIDRepo unique.UniqueIDRepo
}

// NewReportRepo new repository
func NewReportRepo(data *data.Data, uniqueIDRepo unique.UniqueIDRepo) report_common.ReportRepo {
	return &reportRepo{
		data:         data,
		uniqueIDRepo: uniqueIDRepo,
	}
}

// AddReport add report
func (rr *reportRepo) AddReport(ctx context.Context, report *entity.Report) (err error) {
	report.ID, err = rr.uniqueIDRepo.GenUniqueIDStr(ctx, report.TableName())
	if err != nil {
		return err
	}
	_, err = rr.data.DB.Insert(report)
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}

// GetReportListPage get report list page
func (rr *reportRepo) GetReportListPage(ctx context.Context, dto schema.GetReportListPageDTO) (reports []entity.Report, total int64, err error) {
	var (
		ok         bool
		status     int
		objectType int
		session    = rr.data.DB.NewSession()
		cond       = entity.Report{}
	)

	// parse status
	status, ok = entity.ReportStatus[dto.Status]
	if !ok {
		status = entity.ReportStatus["pending"]
	}
	cond.Status = status

	// parse object type
	objectType, ok = constant.ObjectTypeStrMapping[dto.ObjectType]
	if ok {
		cond.ObjectType = objectType
	}

	// order
	session.OrderBy("created_at desc")

	total, err = pager.Help(dto.Page, dto.PageSize, &reports, cond, session)
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}

// GetByID get report by ID
func (ar *reportRepo) GetByID(ctx context.Context, id string) (report entity.Report, exist bool, err error) {
	report = entity.Report{}
	exist, err = ar.data.DB.ID(id).Get(&report)
	return
}

// UpdateByID handle report by ID
func (ar *reportRepo) UpdateByID(
	ctx context.Context,
	id string,
	handleData entity.Report) (err error) {
	_, err = ar.data.DB.ID(id).Update(&handleData)
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}
