package report_admin

import (
	"context"
	"encoding/json"

	"github.com/answerdev/answer/internal/base/handler"
	"github.com/answerdev/answer/internal/service/config"
	"github.com/answerdev/answer/internal/service/object_info"
	"github.com/answerdev/answer/pkg/htmltext"
	"github.com/segmentfault/pacman/log"

	"github.com/answerdev/answer/internal/base/pager"
	"github.com/answerdev/answer/internal/base/reason"
	"github.com/answerdev/answer/internal/entity"
	"github.com/answerdev/answer/internal/schema"
	answercommon "github.com/answerdev/answer/internal/service/answer_common"
	"github.com/answerdev/answer/internal/service/comment_common"
	questioncommon "github.com/answerdev/answer/internal/service/question_common"
	"github.com/answerdev/answer/internal/service/report_common"
	"github.com/answerdev/answer/internal/service/report_handle_admin"
	usercommon "github.com/answerdev/answer/internal/service/user_common"
	"github.com/jinzhu/copier"
	"github.com/segmentfault/pacman/errors"
)

// ReportAdminService user service
type ReportAdminService struct {
	reportRepo        report_common.ReportRepo
	commonUser        *usercommon.UserCommon
	answerRepo        answercommon.AnswerRepo
	questionRepo      questioncommon.QuestionRepo
	commentCommonRepo comment_common.CommentCommonRepo
	reportHandle      *report_handle_admin.ReportHandle
	configService     *config.ConfigService
	objectInfoService *object_info.ObjService
}

// NewReportAdminService new report service
func NewReportAdminService(
	reportRepo report_common.ReportRepo,
	commonUser *usercommon.UserCommon,
	answerRepo answercommon.AnswerRepo,
	questionRepo questioncommon.QuestionRepo,
	commentCommonRepo comment_common.CommentCommonRepo,
	reportHandle *report_handle_admin.ReportHandle,
	configService *config.ConfigService,
	objectInfoService *object_info.ObjService) *ReportAdminService {
	return &ReportAdminService{
		reportRepo:        reportRepo,
		commonUser:        commonUser,
		answerRepo:        answerRepo,
		questionRepo:      questionRepo,
		commentCommonRepo: commentCommonRepo,
		reportHandle:      reportHandle,
		configService:     configService,
		objectInfoService: objectInfoService,
	}
}

// ListReportPage list report pages
func (rs *ReportAdminService) ListReportPage(ctx context.Context, dto schema.GetReportListPageDTO) (pageModel *pager.PageModel, err error) {
	var (
		resp  []*schema.GetReportListPageResp
		flags []entity.Report
		total int64

		flaggedUserIds,
		userIds []string

		flaggedUsers,
		users map[string]*schema.UserBasicInfo
	)

	flags, total, err = rs.reportRepo.GetReportListPage(ctx, dto)
	if err != nil {
		return
	}

	_ = copier.Copy(&resp, flags)
	for _, r := range resp {
		flaggedUserIds = append(flaggedUserIds, r.ReportedUserID)
		userIds = append(userIds, r.UserID)
		r.Format()
	}

	// flagged users
	flaggedUsers, err = rs.commonUser.BatchUserBasicInfoByID(ctx, flaggedUserIds)
	if err != nil {
		return nil, err
	}

	// flag users
	users, err = rs.commonUser.BatchUserBasicInfoByID(ctx, userIds)
	if err != nil {
		return nil, err
	}
	for _, r := range resp {
		r.ReportedUser = flaggedUsers[r.ReportedUserID]
		r.ReportUser = users[r.UserID]
		rs.decorateReportResp(ctx, r)
	}
	return pager.NewPageModel(total, resp), nil
}

// HandleReported handle the reported object
func (rs *ReportAdminService) HandleReported(ctx context.Context, req schema.ReportHandleReq) (err error) {
	var (
		reported   *entity.Report
		handleData = entity.Report{
			FlaggedContent: req.FlaggedContent,
			FlaggedType:    req.FlaggedType,
			Status:         entity.ReportStatusCompleted,
		}
		exist bool
	)

	reported, exist, err = rs.reportRepo.GetByID(ctx, req.ID)
	if err != nil {
		err = errors.BadRequest(reason.ReportHandleFailed).WithError(err).WithStack()
		return
	}
	if !exist {
		err = errors.NotFound(reason.ReportNotFound)
		return
	}

	// check if handle or not
	if reported.Status != entity.ReportStatusPending {
		return
	}

	if err = rs.reportHandle.HandleObject(ctx, reported, req); err != nil {
		return
	}

	err = rs.reportRepo.UpdateByID(ctx, reported.ID, handleData)
	return
}

func (rs *ReportAdminService) decorateReportResp(ctx context.Context, resp *schema.GetReportListPageResp) {
	lang := handler.GetLangByCtx(ctx)
	objectInfo, err := rs.objectInfoService.GetInfo(ctx, resp.ObjectID)
	if err != nil {
		log.Error(err)
		return
	}

	resp.QuestionID = objectInfo.QuestionID
	resp.AnswerID = objectInfo.AnswerID
	resp.CommentID = objectInfo.CommentID
	resp.Title = objectInfo.Title
	resp.Excerpt = htmltext.FetchExcerpt(objectInfo.Content, "...", 240)

	if resp.ReportType > 0 {
		resp.Reason = &schema.ReasonItem{ReasonType: resp.ReportType}
		cf, err := rs.configService.GetConfigByID(ctx, resp.ReportType)
		if err != nil {
			log.Error(err)
		} else {
			_ = json.Unmarshal([]byte(cf.Value), resp.Reason)
			resp.Reason.Translate(cf.Key, lang)
		}
	}
	if resp.FlaggedType > 0 {
		resp.FlaggedReason = &schema.ReasonItem{ReasonType: resp.FlaggedType}
		cf, err := rs.configService.GetConfigByID(ctx, resp.FlaggedType)
		if err != nil {
			log.Error(err)
		} else {
			_ = json.Unmarshal([]byte(cf.Value), resp.Reason)
			resp.Reason.Translate(cf.Key, lang)
		}
	}
}
