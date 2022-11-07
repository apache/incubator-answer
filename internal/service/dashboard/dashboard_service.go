package dashboard

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/answerdev/answer/internal/base/constant"
	"github.com/answerdev/answer/internal/base/reason"
	"github.com/answerdev/answer/internal/schema"
	"github.com/answerdev/answer/internal/service/activity_common"
	answercommon "github.com/answerdev/answer/internal/service/answer_common"
	"github.com/answerdev/answer/internal/service/comment_common"
	"github.com/answerdev/answer/internal/service/config"
	"github.com/answerdev/answer/internal/service/export"
	questioncommon "github.com/answerdev/answer/internal/service/question_common"
	"github.com/answerdev/answer/internal/service/report_common"
	"github.com/answerdev/answer/internal/service/siteinfo_common"
	usercommon "github.com/answerdev/answer/internal/service/user_common"
	"github.com/segmentfault/pacman/errors"
	"github.com/segmentfault/pacman/log"
)

type DashboardService struct {
	questionRepo    questioncommon.QuestionRepo
	answerRepo      answercommon.AnswerRepo
	commentRepo     comment_common.CommentCommonRepo
	voteRepo        activity_common.VoteRepo
	userRepo        usercommon.UserRepo
	reportRepo      report_common.ReportRepo
	configRepo      config.ConfigRepo
	siteInfoService *siteinfo_common.SiteInfoCommonService
}

func NewDashboardService(
	questionRepo questioncommon.QuestionRepo,
	answerRepo answercommon.AnswerRepo,
	commentRepo comment_common.CommentCommonRepo,
	voteRepo activity_common.VoteRepo,
	userRepo usercommon.UserRepo,
	reportRepo report_common.ReportRepo,
	configRepo config.ConfigRepo,
	siteInfoService *siteinfo_common.SiteInfoCommonService,
) *DashboardService {
	return &DashboardService{
		questionRepo:    questionRepo,
		answerRepo:      answerRepo,
		commentRepo:     commentRepo,
		voteRepo:        voteRepo,
		userRepo:        userRepo,
		reportRepo:      reportRepo,
		configRepo:      configRepo,
		siteInfoService: siteInfoService,
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

	siteInfoInterface, err := ds.siteInfoService.GetSiteInterface(ctx)
	if err != nil {
		return dashboardInfo, err
	}

	dashboardInfo.QuestionCount = questionCount
	dashboardInfo.AnswerCount = answerCount
	dashboardInfo.CommentCount = commentCount
	dashboardInfo.VoteCount = voteCount
	dashboardInfo.UserCount = userCount
	dashboardInfo.ReportCount = reportCount

	dashboardInfo.UploadingFiles = true
	emailconfig, err := ds.GetEmailConfig()
	if err != nil {
		return dashboardInfo, err
	}
	if emailconfig.SMTPHost != "" {
		dashboardInfo.SMTP = true
	}
	dashboardInfo.HTTPS = true
	dashboardInfo.OccupyingStorageSpace = "1MB"
	dashboardInfo.AppStartTime = "102"
	dashboardInfo.TimeZone = siteInfoInterface.TimeZone
	dashboardInfo.VersionInfo.Version = constant.Version
	dashboardInfo.VersionInfo.RemoteVersion = ds.RemoteVersion(ctx)
	return dashboardInfo, nil
}

func (ds *DashboardService) RemoteVersion(ctx context.Context) string {
	url := "https://answer.dev/getlatest"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Answer/"+constant.Version)
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		log.Error("http.Client error", err)
		return ""
	}
	defer resp.Body.Close()

	respByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("http.Client error", err)
		return ""
	}
	remoteVersion := &schema.RemoteVersion{}
	err = json.Unmarshal(respByte, remoteVersion)
	if err != nil {
		log.Error("json.Unmarshal error", err)
		return ""
	}
	return remoteVersion.Release.Version
}

func (ds *DashboardService) GetEmailConfig() (ec *export.EmailConfig, err error) {
	emailConf, err := ds.configRepo.GetString("email.config")
	if err != nil {
		return nil, err
	}
	ec = &export.EmailConfig{}
	err = json.Unmarshal([]byte(emailConf), ec)
	if err != nil {
		return nil, errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return ec, nil
}
