package export

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"

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
	SetCode(ctx context.Context, code, content string) error
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

	RegisterTitle  string `json:"register_title"`
	RegisterBody   string `json:"register_body"`
	PassResetTitle string `json:"pass_reset_title"`
	PassResetBody  string `json:"pass_reset_body"`
	ChangeTitle    string `json:"change_title"`
	ChangeBody     string `json:"change_body"`
	TestTitle      string `json:"test_title"`
	TestBody       string `json:"test_body"`
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

// Send email send
func (es *EmailService) Send(ctx context.Context, toEmailAddr, subject, body, code, codeContent string) {
	log.Infof("try to send email to %s", toEmailAddr)
	ec, err := es.GetEmailConfig()
	if err != nil {
		log.Errorf("get email config failed: %s", err)
		return
	}

	m := gomail.NewMessage()
	m.SetHeader("From", ec.FromEmail)
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

	if len(code) > 0 {
		err = es.emailRepo.SetCode(ctx, code, codeContent)
		if err != nil {
			log.Error(err)
		}
	}
}

// VerifyUrlExpired email send
func (es *EmailService) VerifyUrlExpired(ctx context.Context, code string) (content string) {
	content, err := es.emailRepo.VerifyCode(ctx, code)
	if err != nil {
		log.Error(err)
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
	ec, err := es.GetEmailConfig()
	if err != nil {
		return
	}
	siteinfo, err := es.GetSiteGeneral(ctx)
	if err != nil {
		return
	}

	templateData := RegisterTemplateData{
		SiteName: siteinfo.Name, RegisterUrl: registerUrl,
	}
	tmpl, err := template.New("register_title").Parse(ec.RegisterTitle)
	if err != nil {
		return "", "", err
	}
	titleBuf := &bytes.Buffer{}
	bodyBuf := &bytes.Buffer{}
	err = tmpl.Execute(titleBuf, templateData)
	if err != nil {
		return "", "", err
	}

	tmpl, err = template.New("register_body").Parse(ec.RegisterBody)
	if err != nil {
		return "", "", err
	}
	err = tmpl.Execute(bodyBuf, templateData)
	if err != nil {
		return "", "", err
	}

	return titleBuf.String(), bodyBuf.String(), nil
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

func (es *EmailService) TestTemplate(ctx context.Context) (title, body string, err error) {
	ec, err := es.GetEmailConfig()
	if err != nil {
		return
	}

	siteinfo, err := es.GetSiteGeneral(ctx)
	if err != nil {
		return
	}

	templateData := TestTemplateData{
		SiteName: siteinfo.Name,
	}

	titleBuf := &bytes.Buffer{}
	bodyBuf := &bytes.Buffer{}

	tmpl, err := template.New("test_title").Parse(ec.TestTitle)
	if err != nil {
		return "", "", fmt.Errorf("email test title template parse error: %s", err)
	}
	err = tmpl.Execute(titleBuf, templateData)
	if err != nil {
		return "", "", fmt.Errorf("email test body template parse error: %s", err)
	}
	tmpl, err = template.New("test_body").Parse(ec.TestBody)
	err = tmpl.Execute(bodyBuf, templateData)
	if err != nil {
		return "", "", err
	}
	return titleBuf.String(), bodyBuf.String(), nil
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
