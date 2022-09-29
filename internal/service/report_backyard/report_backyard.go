package report_backyard

import (
	"context"
	"strings"

	"github.com/segmentfault/answer/internal/service/config"

	"github.com/jinzhu/copier"
	"github.com/segmentfault/answer/internal/base/pager"
	"github.com/segmentfault/answer/internal/base/reason"
	"github.com/segmentfault/answer/internal/entity"
	"github.com/segmentfault/answer/internal/repo/common"
	"github.com/segmentfault/answer/internal/schema"
	answercommon "github.com/segmentfault/answer/internal/service/answer_common"
	"github.com/segmentfault/answer/internal/service/comment_common"
	questioncommon "github.com/segmentfault/answer/internal/service/question_common"
	"github.com/segmentfault/answer/internal/service/report_common"
	"github.com/segmentfault/answer/internal/service/report_handle_backyard"
	usercommon "github.com/segmentfault/answer/internal/service/user_common"
	"github.com/segmentfault/pacman/errors"
)

// ReportBackyardService user service
type ReportBackyardService struct {
	reportRepo        report_common.ReportRepo
	commonUser        *usercommon.UserCommon
	commonRepo        *common.CommonRepo
	answerRepo        answercommon.AnswerRepo
	questionRepo      questioncommon.QuestionRepo
	commentCommonRepo comment_common.CommentCommonRepo
	reportHandle      *report_handle_backyard.ReportHandle
	configRepo        config.ConfigRepo
}

// NewReportBackyardService new report service
func NewReportBackyardService(
	reportRepo report_common.ReportRepo,
	commonUser *usercommon.UserCommon,
	commonRepo *common.CommonRepo,
	answerRepo answercommon.AnswerRepo,
	questionRepo questioncommon.QuestionRepo,
	commentCommonRepo comment_common.CommentCommonRepo,
	reportHandle *report_handle_backyard.ReportHandle,
	configRepo config.ConfigRepo) *ReportBackyardService {
	return &ReportBackyardService{
		reportRepo:        reportRepo,
		commonUser:        commonUser,
		commonRepo:        commonRepo,
		answerRepo:        answerRepo,
		questionRepo:      questionRepo,
		commentCommonRepo: commentCommonRepo,
		reportHandle:      reportHandle,
		configRepo:        configRepo,
	}
}

// ListReportPage list report pages
func (rs *ReportBackyardService) ListReportPage(ctx context.Context, dto schema.GetReportListPageDTO) (pageModel *pager.PageModel, err error) {
	var (
		resp  []*schema.GetReportListPageResp
		flags []entity.Report
		total int64

		flagedUserIds,
		userIds []string

		flagedUsers,
		users map[string]*schema.UserBasicInfo
	)

	pageModel = &pager.PageModel{}

	flags, total, err = rs.reportRepo.GetReportListPage(ctx, dto)
	if err != nil {
		return
	}

	_ = copier.Copy(&resp, flags)
	for _, r := range resp {
		flagedUserIds = append(flagedUserIds, r.ReportedUserID)
		userIds = append(userIds, r.UserID)
		r.Format()
	}

	// flaged users
	flagedUsers, err = rs.commonUser.BatchUserBasicInfoByID(ctx, flagedUserIds)

	// flag users
	users, err = rs.commonUser.BatchUserBasicInfoByID(ctx, userIds)
	for _, r := range resp {
		r.ReportedUser = flagedUsers[r.ReportedUserID]
		r.ReportUser = users[r.UserID]
	}

	rs.parseObject(ctx, &resp)
	return pager.NewPageModel(total, resp), nil
}

// HandleReported handle the reported object
func (rs *ReportBackyardService) HandleReported(ctx context.Context, req schema.ReportHandleReq) (err error) {
	var (
		reported   = entity.Report{}
		handleData = entity.Report{
			FlagedContent: req.FlagedContent,
			FlagedType:    req.FlagedType,
			Status:        entity.ReportStatusCompleted,
		}
		exist = false
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

func (rs *ReportBackyardService) parseObject(ctx context.Context, resp *[]*schema.GetReportListPageResp) {
	var (
		res = *resp
	)

	for i, r := range res {
		var (
			objIds map[string]string
			exists,
			ok bool
			err error
			questionId,
			answerId,
			commentId string
			question *entity.Question
			answer   *entity.Answer
			cmt      *entity.Comment
		)

		objIds, err = rs.commonRepo.GetObjectIDMap(r.ObjectID)
		if err != nil {
			continue
		}

		questionId, ok = objIds["question"]
		if !ok {
			continue
		}

		question, exists, err = rs.questionRepo.GetQuestion(ctx, questionId)
		if err != nil || !exists {
			continue
		}

		answerId, ok = objIds["answer"]
		if ok {
			answer, _, err = rs.answerRepo.GetAnswer(ctx, answerId)
		}

		commentId, ok = objIds["comment"]
		if ok {
			cmt, _, err = rs.commentCommonRepo.GetComment(ctx, commentId)
		}

		switch r.OType {
		case "question":
			r.QuestionID = questionId
			r.Title = question.Title
			r.Excerpt = rs.cutOutTagParsedText(question.OriginalText)

		case "answer":
			r.QuestionID = questionId
			r.AnswerID = answerId
			r.Title = question.Title
			r.Excerpt = rs.cutOutTagParsedText(answer.OriginalText)

		case "comment":
			r.QuestionID = questionId
			r.AnswerID = answerId
			r.CommentID = commentId
			r.Title = question.Title
			r.Excerpt = rs.cutOutTagParsedText(cmt.OriginalText)
		}

		// parse reason
		if r.ReportType > 0 {
			r.Reason = &schema.ReasonItem{
				ReasonType: r.ReportType,
			}
			err = rs.configRepo.GetConfigById(r.ReportType, r.Reason)
		}
		if r.FlagedType > 0 {
			r.FlagedReason = &schema.ReasonItem{
				ReasonType: r.FlagedType,
			}
			_ = rs.configRepo.GetConfigById(r.FlagedType, r.FlagedReason)
		}

		res[i] = r
	}
	resp = &res
}

func (rs *ReportBackyardService) cutOutTagParsedText(parsedText string) string {
	parsedText = strings.TrimSpace(parsedText)
	idx := strings.Index(parsedText, "\n")
	if idx >= 0 {
		parsedText = parsedText[0:idx]
	}
	return parsedText
}
