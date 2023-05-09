package export

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"mime"
	"time"

	"github.com/answerdev/answer/internal/base/reason"
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
	configRepo   config.ConfigRepo
	emailRepo    EmailRepo
	siteInfoRepo siteinfo_common.SiteInfoRepo
}

// EmailRepo email repository
type EmailRepo interface {
	SetCode(ctx context.Context, code, content string, duration time.Duration) error
	VerifyCode(ctx context.Context, code string) (content string, err error)
}

// NewEmailService email service
func NewEmailService(configRepo config.ConfigRepo, emailRepo EmailRepo, siteInfoRepo siteinfo_common.SiteInfoRepo) *EmailService {
	return &EmailService{
		configRepo:   configRepo,
		emailRepo:    emailRepo,
		siteInfoRepo: siteInfoRepo,
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

	RegisterTitle   string `json:"register_title"`
	RegisterBody    string `json:"register_body"`
	PassResetTitle  string `json:"pass_reset_title"`
	PassResetBody   string `json:"pass_reset_body"`
	ChangeTitle     string `json:"change_title"`
	ChangeBody      string `json:"change_body"`
	TestTitle       string `json:"test_title"`
	TestBody        string `json:"test_body"`
	NewAnswerTitle  string `json:"new_answer_title"`
	NewAnswerBody   string `json:"new_answer_body"`
	NewCommentTitle string `json:"new_comment_title"`
	NewCommentBody  string `json:"new_comment_body"`
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
	ec, err := es.GetEmailConfig()
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
	emailConfig, err := es.GetEmailConfig()
	if err != nil {
		return
	}

	siteInfo, err := es.GetSiteGeneral(ctx)
	if err != nil {
		return
	}
	templateData := RegisterTemplateData{
		SiteName:    siteInfo.Name,
		RegisterUrl: registerUrl,
	}

	title, err = es.parseTemplateData(emailConfig.RegisterTitle, templateData)
	if err != nil {
		return "", "", fmt.Errorf("email template parse error: %s", err)
	}

	body, err = es.parseTemplateData(emailConfig.RegisterBody, templateData)
	if err != nil {
		return "", "", fmt.Errorf("email template parse error: %s", err)
	}
	return title, body, nil
}

func (es *EmailService) PassResetTemplate(ctx context.Context, passResetUrl string) (title, body string, err error) {
	ec, err := es.GetEmailConfig()
	if err != nil {
		return
	}

	siteinfo, err := es.GetSiteGeneral(ctx)
	if err != nil {
		return
	}

	templateData := PassResetTemplateData{SiteName: siteinfo.Name, PassResetUrl: passResetUrl}
	tmpl, err := template.New("pass_reset_title").Parse(ec.PassResetTitle)
	if err != nil {
		return "", "", err
	}
	titleBuf := &bytes.Buffer{}
	bodyBuf := &bytes.Buffer{}
	err = tmpl.Execute(titleBuf, templateData)
	if err != nil {
		return "", "", err
	}

	tmpl, err = template.New("pass_reset_body").Parse(ec.PassResetBody)
	if err != nil {
		return "", "", err
	}
	err = tmpl.Execute(bodyBuf, templateData)
	if err != nil {
		return "", "", err
	}
	return titleBuf.String(), bodyBuf.String(), nil
}

func (es *EmailService) ChangeEmailTemplate(ctx context.Context, changeEmailUrl string) (title, body string, err error) {
	ec, err := es.GetEmailConfig()
	if err != nil {
		return
	}

	siteinfo, err := es.GetSiteGeneral(ctx)
	if err != nil {
		return
	}
	templateData := ChangeEmailTemplateData{
		SiteName:       siteinfo.Name,
		ChangeEmailUrl: changeEmailUrl,
	}
	tmpl, err := template.New("email_change_title").Parse(ec.ChangeTitle)
	if err != nil {
		return "", "", err
	}
	titleBuf := &bytes.Buffer{}
	bodyBuf := &bytes.Buffer{}
	err = tmpl.Execute(titleBuf, templateData)
	if err != nil {
		return "", "", err
	}

	tmpl, err = template.New("email_change_body").Parse(ec.ChangeBody)
	if err != nil {
		return "", "", err
	}
	err = tmpl.Execute(bodyBuf, templateData)
	if err != nil {
		return "", "", err
	}
	return titleBuf.String(), bodyBuf.String(), nil
}

// TestTemplate send test email template parse
func (es *EmailService) TestTemplate(ctx context.Context) (title, body string, err error) {
	emailConfig, err := es.GetEmailConfig()
	if err != nil {
		return
	}

	siteInfo, err := es.GetSiteGeneral(ctx)
	if err != nil {
		return
	}
	templateData := TestTemplateData{
		SiteName: siteInfo.Name,
	}

	title, err = es.parseTemplateData(emailConfig.TestTitle, templateData)
	if err != nil {
		return "", "", fmt.Errorf("email template parse error: %s", err)
	}

	body, err = es.parseTemplateData(emailConfig.TestBody, templateData)
	if err != nil {
		return "", "", fmt.Errorf("email template parse error: %s", err)
	}
	return title, body, nil
}

// NewAnswerTemplate new answer template
func (es *EmailService) NewAnswerTemplate(ctx context.Context, raw *schema.NewAnswerTemplateRawData) (
	title, body string, err error) {
	emailConfig, err := es.GetEmailConfig()
	if err != nil {
		return
	}

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
	templateData.SiteName = siteInfo.Name

	title, err = es.parseTemplateData(emailConfig.NewAnswerTitle, templateData)
	if err != nil {
		return "", "", fmt.Errorf("email template parse error: %s", err)
	}

	body, err = es.parseTemplateData(emailConfig.NewAnswerBody, templateData)
	if err != nil {
		return "", "", fmt.Errorf("email template parse error: %s", err)
	}
	return title, body, nil
}

// NewCommentTemplate new comment template
func (es *EmailService) NewCommentTemplate(ctx context.Context, raw *schema.NewCommentTemplateRawData) (
	title, body string, err error) {
	emailConfig, err := es.GetEmailConfig()
	if err != nil {
		return
	}

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
	templateData.SiteName = siteInfo.Name

	title, err = es.parseTemplateData(emailConfig.NewCommentTitle, templateData)
	if err != nil {
		return "", "", fmt.Errorf("email template parse error: %s", err)
	}

	body, err = es.parseTemplateData(emailConfig.NewCommentBody, templateData)
	if err != nil {
		return "", "", fmt.Errorf("email template parse error: %s", err)
	}
	return title, body, nil
}

func (es *EmailService) parseTemplateData(templateContent string, templateData interface{}) (parsedData string, err error) {
	parsedDataBuf := &bytes.Buffer{}
	tmpl, err := template.New("").Parse(templateContent)
	if err != nil {
		return "", err
	}
	err = tmpl.Execute(parsedDataBuf, templateData)
	if err != nil {
		return "", err
	}
	return parsedDataBuf.String(), nil
}

func (es *EmailService) GetEmailConfig() (ec *EmailConfig, err error) {
	emailConf, err := es.configRepo.GetString("email.config")
	if err != nil {
		return nil, err
	}
	ec = &EmailConfig{}
	err = json.Unmarshal([]byte(emailConf), ec)
	if err != nil {
		return nil, errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return ec, nil
}

// SetEmailConfig set email config
func (es *EmailService) SetEmailConfig(ec *EmailConfig) (err error) {
	data, _ := json.Marshal(ec)
	return es.configRepo.SetConfig("email.config", string(data))
}
