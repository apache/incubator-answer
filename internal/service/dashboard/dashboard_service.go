package dashboard

import (
	"context"

	"github.com/answerdev/answer/internal/schema"
	"github.com/answerdev/answer/internal/service/activity_common"
	answercommon "github.com/answerdev/answer/internal/service/answer_common"
	"github.com/answerdev/answer/internal/service/comment_common"
	"github.com/answerdev/answer/internal/service/config"
	questioncommon "github.com/answerdev/answer/internal/service/question_common"
	"github.com/answerdev/answer/internal/service/report_common"
	usercommon "github.com/answerdev/answer/internal/service/user_common"
)

type DashboardService struct {
	questionRepo questioncommon.QuestionRepo
	answerRepo   answercommon.AnswerRepo
	commentRepo  comment_common.CommentCommonRepo
	voteRepo     activity_common.VoteRepo
	userRepo     usercommon.UserRepo
	reportRepo   report_common.ReportRepo
	configRepo   config.ConfigRepo
}

func NewDashboardService(
	questionRepo questioncommon.QuestionRepo,
	answerRepo answercommon.AnswerRepo,
	commentRepo comment_common.CommentCommonRepo,
	voteRepo activity_common.VoteRepo,
	userRepo usercommon.UserRepo,
	reportRepo report_common.ReportRepo,
	configRepo config.ConfigRepo,

) *DashboardService {
	return &DashboardService{
		questionRepo: questionRepo,
		answerRepo:   answerRepo,
		commentRepo:  commentRepo,
		voteRepo:     voteRepo,
		userRepo:     userRepo,
		reportRepo:   reportRepo,
		configRepo:   configRepo,
	}
}

// Statistical
func (ds *DashboardService) Statistical(ctx context.Context) (*schema.DashboardInfo, error) {
	dashboardInfo := &schema.DashboardInfo{}
	questionCount, err := ds.questionRepo.GetQuestionCount(ctx)
	if err != nil {
		return dashboardInfo, err
	}
	answerCount, err := ds.answerRepo.GetAnswerCount(ctx)
	if err != nil {
		return dashboardInfo, err
	}
	commentCount, err := ds.commentRepo.GetCommentCount(ctx)
	if err != nil {
		return dashboardInfo, err
	}

	typeKeys := []string{
		"question.vote_up",
		"question.vote_down",
		"answer.vote_up",
		"answer.vote_down",
	}
	var activityTypes []int

	for _, typeKey := range typeKeys {
		var t int
		t, err = ds.configRepo.GetConfigType(typeKey)
		if err != nil {
			continue
		}
		activityTypes = append(activityTypes, t)
	}

	voteCount, err := ds.voteRepo.GetVoteCount(ctx, activityTypes)
	if err != nil {
		return dashboardInfo, err
	}
	userCount, err := ds.userRepo.GetUserCount(ctx)
	if err != nil {
		return dashboardInfo, err
	}

	reportCount, err := ds.reportRepo.GetReportCount(ctx)
	if err != nil {
		return dashboardInfo, err
	}
	dashboardInfo.QuestionCount = questionCount
	dashboardInfo.AnswerCount = answerCount
	dashboardInfo.CommentCount = commentCount
	dashboardInfo.VoteCount = voteCount
	dashboardInfo.UserCount = userCount
	dashboardInfo.ReportCount = reportCount

	dashboardInfo.UploadingFiles = "Allowed"
	dashboardInfo.SMTP = "Enabled"
	dashboardInfo.OccupyingStorageSpace = "1MB"
	dashboardInfo.AppStartTime = "102"
	return dashboardInfo, nil
}
