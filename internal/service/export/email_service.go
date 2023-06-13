package export

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"mime"
	"os"
	"time"

	"github.com/answerdev/answer/internal/base/constant"
	"github.com/answerdev/answer/internal/base/handler"
	"github.com/answerdev/answer/internal/base/reason"
	"github.com/answerdev/answer/internal/base/translator"
	"github.com/answerdev/answer/internal/entity"
	"github.com/answerdev/answer/internal/schema"
	"github.com/answerdev/answer/internal/service/config"
	"github.com/answerdev/answer/internal/service/siteinfo_common"
	"github.com/segmentfault/pacman/errors"
	"github.com/segmentfault/pacman/log"
	"golang.org/x/net/context"
	"gopkg.in/gomail.v2"
)

// EmailService kit service
type EmailService struct {
	configService *config.ConfigService
	emailRepo     EmailRepo
	siteInfoRepo  siteinfo_common.SiteInfoRepo
}

// EmailRepo email repository
type EmailRepo interface {
	SetCode(ctx context.Context, code, content string, duration time.Duration) error
	VerifyCode(ctx context.Context, code string) (content string, err error)
}

// NewEmailService email service
func NewEmailService(configService *config.ConfigService, emailRepo EmailRepo, siteInfoRepo siteinfo_common.SiteInfoRepo) *EmailService {
	return &EmailService{
		configService: configService,
		emailRepo:     emailRepo,
		siteInfoRepo:  siteInfoRepo,
	}
}

// EmailConfig email config
type EmailConfig struct {
	FromEmail          string `json:"from_email"`
	FromName           string `json:"from_name"`
	SMTPHost           string `json:"smtp_host"`
	SMTPPort           int    `json:"smtp_port"`
	Encryption         string `json:"encryption"` // "" SSL
	SMTPUsername       string `json:"smtp_username"`
	SMTPPassword       string `json:"smtp_password"`
	SMTPAuthentication bool   `json:"smtp_authentication"`
}

func (e *EmailConfig) IsSSL() bool {
	return e.Encryption == "SSL"
}

type RegisterTemplateData struct {
	SiteName    string
	RegisterUrl string
}

type PassResetTemplateData struct {
	SiteName     string
	PassResetUrl string
}

type ChangeEmailTemplateData struct {
	SiteName       string
	ChangeEmailUrl string
}

type TestTemplateData struct {
	SiteName string
}

// SendAndSaveCode send email and save code
func (es *EmailService) SendAndSaveCode(ctx context.Context, toEmailAddr, subject, body, code, codeContent string) {
	es.Send(ctx, toEmailAddr, subject, body)
	err := es.emailRepo.SetCode(ctx, code, codeContent, 10*time.Minute)
	if err != nil {
		log.Error(err)
	}
}

// SendAndSaveCodeWithTime send email and save code
func (es *EmailService) SendAndSaveCodeWithTime(
	ctx context.Context, toEmailAddr, subject, body, code, codeContent string, duration time.Duration) {
	es.Send(ctx, toEmailAddr, subject, body)
	err := es.emailRepo.SetCode(ctx, code, codeContent, duration)
	if err != nil {
		log.Error(err)
	}
}

// Send email send
func (es *EmailService) Send(ctx context.Context, toEmailAddr, subject, body string) {
	log.Infof("try to send email to %s", toEmailAddr)
	ec, err := es.GetEmailConfig(ctx)
	if err != nil {
		log.Errorf("get email config failed: %s", err)
		return
	}

	m := gomail.NewMessage()
	fromName := mime.QEncoding.Encode("utf-8", ec.FromName)
	m.SetHeader("From", fmt.Sprintf("%s <%s>", fromName, ec.FromEmail))
	m.SetHeader("To", toEmailAddr)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gomail.NewDialer(ec.SMTPHost, ec.SMTPPort, ec.SMTPUsername, ec.SMTPPassword)
	if ec.IsSSL() {
		d.SSL = true
	}
	if len(os.Getenv("SKIP_SMTP_TLS_VERIFY")) > 0 {
		d.TLSConfig = &tls.Config{ServerName: d.Host, InsecureSkipVerify: true}
	}
	if err := d.DialAndSend(m); err != nil {
		log.Errorf("send email to %s failed: %s", toEmailAddr, err)
	} else {
		log.Infof("send email to %s success", toEmailAddr)
	}
}

// VerifyUrlExpired email send
func (es *EmailService) VerifyUrlExpired(ctx context.Context, code string) (content string) {
	content, err := es.emailRepo.VerifyCode(ctx, code)
	if err != nil {
		log.Warn(err)
	}
	return content
}

func (es *EmailService) GetSiteGeneral(ctx context.Context) (resp schema.SiteGeneralResp, err error) {
	var (
		siteType = "general"
		siteInfo *entity.SiteInfo
		exist    bool
	)
	resp = schema.SiteGeneralResp{}

	siteInfo, exist, err = es.siteInfoRepo.GetByType(ctx, siteType)
	if !exist {
		return
	}

	_ = json.Unmarshal([]byte(siteInfo.Content), &resp)
	return
}

func (es *EmailService) RegisterTemplate(ctx context.Context, registerUrl string) (title, body string, err error) {
	siteInfo, err := es.GetSiteGeneral(ctx)
	if err != nil {
		return
	}
	templateData := RegisterTemplateData{
		SiteName:    siteInfo.Name,
		RegisterUrl: registerUrl,
	}

	lang := handler.GetLangByCtx(ctx)
	title = translator.TrWithData(lang, constant.EmailTplKeyRegisterTitle, templateData)
	body = translator.TrWithData(lang, constant.EmailTplKeyRegisterBody, templateData)
	return title, body, nil
}

func (es *EmailService) PassResetTemplate(ctx context.Context, passResetUrl string) (title, body string, err error) {
	siteInfo, err := es.GetSiteGeneral(ctx)
	if err != nil {
		return
	}

	templateData := PassResetTemplateData{SiteName: siteInfo.Name, PassResetUrl: passResetUrl}

	lang := handler.GetLangByCtx(ctx)
	title = translator.TrWithData(lang, constant.EmailTplKeyPassResetTitle, templateData)
	body = translator.TrWithData(lang, constant.EmailTplKeyPassResetBody, templateData)
	return title, body, nil
}

func (es *EmailService) ChangeEmailTemplate(ctx context.Context, changeEmailUrl string) (title, body string, err error) {
	siteInfo, err := es.GetSiteGeneral(ctx)
	if err != nil {
		return
	}
	templateData := ChangeEmailTemplateData{
		SiteName:       siteInfo.Name,
		ChangeEmailUrl: changeEmailUrl,
	}

	lang := handler.GetLangByCtx(ctx)
	title = translator.TrWithData(lang, constant.EmailTplKeyChangeEmailTitle, templateData)
	body = translator.TrWithData(lang, constant.EmailTplKeyChangeEmailBody, templateData)
	return title, body, nil
}

// TestTemplate send test email template parse
func (es *EmailService) TestTemplate(ctx context.Context) (title, body string, err error) {
	siteInfo, err := es.GetSiteGeneral(ctx)
	if err != nil {
		return
	}
	templateData := TestTemplateData{SiteName: siteInfo.Name}

	lang := handler.GetLangByCtx(ctx)
	title = translator.TrWithData(lang, constant.EmailTplKeyTestTitle, templateData)
	body = translator.TrWithData(lang, constant.EmailTplKeyTestBody, templateData)
	return title, body, nil
}

// NewAnswerTemplate new answer template
func (es *EmailService) NewAnswerTemplate(ctx context.Context, raw *schema.NewAnswerTemplateRawData) (
	title, body string, err error) {
	siteInfo, err := es.GetSiteGeneral(ctx)
	if err != nil {
		return
	}
	templateData := &schema.NewAnswerTemplateData{
		SiteName:       siteInfo.Name,
		DisplayName:    raw.AnswerUserDisplayName,
		QuestionTitle:  raw.QuestionTitle,
		AnswerUrl:      fmt.Sprintf("%s/questions/%s/%s", siteInfo.SiteUrl, raw.QuestionID, raw.AnswerID),
		AnswerSummary:  raw.AnswerSummary,
		UnsubscribeUrl: fmt.Sprintf("%s/users/unsubscribe?code=%s", siteInfo.SiteUrl, raw.UnsubscribeCode),
	}

	lang := handler.GetLangByCtx(ctx)
	title = translator.TrWithData(lang, constant.EmailTplKeyNewAnswerTitle, templateData)
	body = translator.TrWithData(lang, constant.EmailTplKeyNewAnswerBody, templateData)
	return title, body, nil
}

// NewInviteAnswerTemplate new invite answer template
func (es *EmailService) NewInviteAnswerTemplate(ctx context.Context, raw *schema.NewInviteAnswerTemplateRawData) (
	title, body string, err error) {
	siteInfo, err := es.GetSiteGeneral(ctx)
	if err != nil {
		return
	}
	templateData := &schema.NewInviteAnswerTemplateData{
		SiteName:       siteInfo.Name,
		DisplayName:    raw.InviterDisplayName,
		QuestionTitle:  raw.QuestionTitle,
		InviteUrl:      fmt.Sprintf("%s/questions/%s", siteInfo.SiteUrl, raw.QuestionID),
		UnsubscribeUrl: fmt.Sprintf("%s/users/unsubscribe?code=%s", siteInfo.SiteUrl, raw.UnsubscribeCode),
	}

	lang := handler.GetLangByCtx(ctx)
	title = translator.TrWithData(lang, constant.EmailTplKeyInvitedAnswerTitle, templateData)
	body = translator.TrWithData(lang, constant.EmailTplKeyInvitedAnswerBody, templateData)
	return title, body, nil
}

// NewCommentTemplate new comment template
func (es *EmailService) NewCommentTemplate(ctx context.Context, raw *schema.NewCommentTemplateRawData) (
	title, body string, err error) {
	siteInfo, err := es.GetSiteGeneral(ctx)
	if err != nil {
		return
	}
	templateData := &schema.NewCommentTemplateData{
		SiteName:       siteInfo.Name,
		DisplayName:    raw.CommentUserDisplayName,
		QuestionTitle:  raw.QuestionTitle,
		CommentSummary: raw.CommentSummary,
		UnsubscribeUrl: fmt.Sprintf("%s/users/unsubscribe?code=%s", siteInfo.SiteUrl, raw.UnsubscribeCode),
	}
	if len(raw.AnswerID) > 0 {
		templateData.CommentUrl = fmt.Sprintf("%s/questions/%s/%s?commentId=%s", siteInfo.SiteUrl, raw.QuestionID,
			raw.AnswerID, raw.CommentID)
	} else {
		templateData.CommentUrl = fmt.Sprintf("%s/questions/%s?commentId=%s", siteInfo.SiteUrl,
			raw.QuestionID, raw.CommentID)
	}

	lang := handler.GetLangByCtx(ctx)
	title = translator.TrWithData(lang, constant.EmailTplKeyNewCommentTitle, templateData)
	body = translator.TrWithData(lang, constant.EmailTplKeyNewCommentBody, templateData)
	return title, body, nil
}

func (es *EmailService) GetEmailConfig(ctx context.Context) (ec *EmailConfig, err error) {
	emailConf, err := es.configService.GetStringValue(ctx, "email.config")
	if err != nil {
		return nil, err
	}
	ec = &EmailConfig{}
	err = json.Unmarshal([]byte(emailConf), ec)
	if err != nil {
		log.Errorf("old email config format is invalid, you need to update smtp config: %v", err)
		return nil, errors.BadRequest(reason.SiteInfoConfigNotFound)
	}
	return ec, nil
}

// SetEmailConfig set email config
func (es *EmailService) SetEmailConfig(ctx context.Context, ec *EmailConfig) (err error) {
	data, _ := json.Marshal(ec)
	return es.configService.UpdateConfig(ctx, "email.config", string(data))
}
