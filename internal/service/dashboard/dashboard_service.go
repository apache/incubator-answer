package dashboard

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/answerdev/answer/internal/base/constant"
	"github.com/answerdev/answer/internal/base/data"
	"github.com/answerdev/answer/internal/base/reason"
	"github.com/answerdev/answer/internal/schema"
	"github.com/answerdev/answer/internal/service/activity_common"
	answercommon "github.com/answerdev/answer/internal/service/answer_common"
	"github.com/answerdev/answer/internal/service/comment_common"
	"github.com/answerdev/answer/internal/service/config"
	"github.com/answerdev/answer/internal/service/export"
	questioncommon "github.com/answerdev/answer/internal/service/question_common"
	"github.com/answerdev/answer/internal/service/report_common"
	"github.com/answerdev/answer/internal/service/service_config"
	"github.com/answerdev/answer/internal/service/siteinfo_common"
	usercommon "github.com/answerdev/answer/internal/service/user_common"
	"github.com/answerdev/answer/pkg/dir"
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
	configService   *config.ConfigService
	siteInfoService *siteinfo_common.SiteInfoCommonService
	serviceConfig   *service_config.ServiceConfig

	data *data.Data
}

func NewDashboardService(
	questionRepo questioncommon.QuestionRepo,
	answerRepo answercommon.AnswerRepo,
	commentRepo comment_common.CommentCommonRepo,
	voteRepo activity_common.VoteRepo,
	userRepo usercommon.UserRepo,
	reportRepo report_common.ReportRepo,
	configService *config.ConfigService,
	siteInfoService *siteinfo_common.SiteInfoCommonService,
	serviceConfig *service_config.ServiceConfig,

	data *data.Data,
) *DashboardService {
	return &DashboardService{
		questionRepo:    questionRepo,
		answerRepo:      answerRepo,
		commentRepo:     commentRepo,
		voteRepo:        voteRepo,
		userRepo:        userRepo,
		reportRepo:      reportRepo,
		configService:   configService,
		siteInfoService: siteInfoService,
		serviceConfig:   serviceConfig,

		data: data,
	}
}

func (ds *DashboardService) StatisticalByCache(ctx context.Context) (*schema.DashboardInfo, error) {
	dashboardInfo := &schema.DashboardInfo{}
	infoStr, err := ds.data.Cache.GetString(ctx, schema.DashBoardCachekey)
	if err != nil {
		info, statisticalErr := ds.Statistical(ctx)
		if statisticalErr != nil {
			return nil, statisticalErr
		}
		if setCacheErr := ds.SetCache(ctx, info); setCacheErr != nil {
			log.Errorf("set dashboard statistical failed: %s", setCacheErr)
		}
		return info, nil
	}
	if err = json.Unmarshal([]byte(infoStr), dashboardInfo); err != nil {
		log.Errorf("parsing dashboard information failed: %s", err)
		return nil, errors.InternalServer(reason.UnknownError)
	}
	startTime := time.Now().Unix() - schema.AppStartTime.Unix()
	dashboardInfo.AppStartTime = fmt.Sprintf("%d", startTime)
	dashboardInfo.VersionInfo.Version = constant.Version
	dashboardInfo.VersionInfo.Revision = constant.Revision
	return dashboardInfo, nil
}

func (ds *DashboardService) SetCache(ctx context.Context, info *schema.DashboardInfo) error {
	infoStr, err := json.Marshal(info)
	if err != nil {
		return errors.InternalServer(reason.UnknownError).WithError(err).WithStack()
	}
	err = ds.data.Cache.SetString(ctx, schema.DashBoardCachekey, string(infoStr), schema.DashBoardCacheTime)
	if err != nil {
		return errors.InternalServer(reason.UnknownError).WithError(err).WithStack()
	}
	return nil
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
		cfg, err := ds.configService.GetConfigByKey(ctx, typeKey)
		if err != nil {
			continue
		}
		activityTypes = append(activityTypes, cfg.ID)
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
	emailconfig, err := ds.GetEmailConfig(ctx)
	if err != nil {
		return dashboardInfo, err
	}
	if emailconfig.SMTPHost != "" {
		dashboardInfo.SMTP = true
	}
	siteGeneral, err := ds.siteInfoService.GetSiteGeneral(ctx)
	if err != nil {
		return dashboardInfo, err
	}
	siteUrl, err := url.Parse(siteGeneral.SiteUrl)
	if err != nil {
		return dashboardInfo, err
	}
	if siteUrl.Scheme == "https" {
		dashboardInfo.HTTPS = true
	}

	dirSize, err := dir.DirSize(ds.serviceConfig.UploadPath)
	if err != nil {
		return dashboardInfo, err
	}
	size := dir.FormatFileSize(dirSize)
	dashboardInfo.OccupyingStorageSpace = size
	startTime := time.Now().Unix() - schema.AppStartTime.Unix()
	dashboardInfo.AppStartTime = fmt.Sprintf("%d", startTime)
	dashboardInfo.TimeZone = siteInfoInterface.TimeZone
	dashboardInfo.VersionInfo.Version = constant.Version
	dashboardInfo.VersionInfo.Revision = constant.Revision
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

	respByte, err := io.ReadAll(resp.Body)
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

func (ds *DashboardService) GetEmailConfig(ctx context.Context) (ec *export.EmailConfig, err error) {
	emailConf, err := ds.configService.GetStringValue(ctx, "email.config")
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
