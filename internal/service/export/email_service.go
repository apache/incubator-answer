package export

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"net/smtp"

	"github.com/jordan-wright/email"
	"github.com/segmentfault/answer/internal/service/config"
	"github.com/segmentfault/pacman/log"
	"golang.org/x/net/context"
)

// EmailService kit service
type EmailService struct {
	configRepo config.ConfigRepo
	emailRepo  EmailRepo
}

// EmailRepo email repository
type EmailRepo interface {
	SetCode(ctx context.Context, code, content string) error
	VerifyCode(ctx context.Context, code string) (content string, err error)
}

// NewEmailService email service
func NewEmailService(configRepo config.ConfigRepo, emailRepo EmailRepo) *EmailService {
	return &EmailService{
		configRepo: configRepo,
		emailRepo:  emailRepo,
	}
}

// EmailConfig email config
type EmailConfig struct {
	EmailWebName        string `json:"email_web_name"`
	EmailFrom           string `json:"email_from"`
	EmailFromPass       string `json:"email_from_pass"`
	EmailFromHostname   string `json:"email_from_hostname"`
	EmailFromSMTP       string `json:"email_from_smtp"`
	EmailFromName       string `json:"email_from_name"`
	EmailRegisterTitle  string `json:"email_register_title"`
	EmailRegisterBody   string `json:"email_register_body"`
	EmailPassResetTitle string `json:"email_pass_reset_title"`
	EmailPassResetBody  string `json:"email_pass_reset_body"`
	EmailChangeTitle    string `json:"email_change_title"`
	EmailChangeBody     string `json:"email_change_body"`
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

// Send email send
func (es *EmailService) Send(ctx context.Context, emailAddr, title, body, code, content string) {
	emailClient := email.NewEmail()

	ec, err := es.getEmailConfig()
	if err != nil {
		log.Error(err)
		return
	}

	emailClient.From = fmt.Sprintf("%s <%s>", ec.EmailFromName, ec.EmailFrom)
	emailClient.To = []string{emailAddr}
	emailClient.Subject = title
	emailClient.HTML = []byte(body)
	err = emailClient.Send(ec.EmailFromSMTP, smtp.PlainAuth("", ec.EmailFrom, ec.EmailFromPass, ec.EmailFromHostname))
	if err != nil {
		log.Error(err)
	}

	err = es.emailRepo.SetCode(ctx, code, content)
	if err != nil {
		log.Error(err)
	}
	return
}

// VerifyUrlExpired email send
func (es *EmailService) VerifyUrlExpired(ctx context.Context, code string) (content string) {
	content, err := es.emailRepo.VerifyCode(ctx, code)
	if err != nil {
		log.Error(err)
	}
	return content
}

func (es *EmailService) RegisterTemplate(registerUrl string) (title, body string, err error) {
	ec, err := es.getEmailConfig()
	if err != nil {
		return
	}

	templateData := RegisterTemplateData{ec.EmailWebName, registerUrl}
	tmpl, err := template.New("register_title").Parse(ec.EmailRegisterTitle)
	if err != nil {
		return "", "", err
	}
	titleBuf := &bytes.Buffer{}
	bodyBuf := &bytes.Buffer{}
	err = tmpl.Execute(titleBuf, templateData)
	if err != nil {
		return "", "", err
	}

	tmpl, err = template.New("register_body").Parse(ec.EmailRegisterBody)
	err = tmpl.Execute(bodyBuf, templateData)
	if err != nil {
		return "", "", err
	}

	return titleBuf.String(), bodyBuf.String(), nil
}

func (es *EmailService) PassResetTemplate(passResetUrl string) (title, body string, err error) {
	ec, err := es.getEmailConfig()
	if err != nil {
		return
	}

	templateData := PassResetTemplateData{ec.EmailWebName, passResetUrl}
	tmpl, err := template.New("pass_reset_title").Parse(ec.EmailPassResetTitle)
	if err != nil {
		return "", "", err
	}
	titleBuf := &bytes.Buffer{}
	bodyBuf := &bytes.Buffer{}
	err = tmpl.Execute(titleBuf, templateData)
	if err != nil {
		return "", "", err
	}

	tmpl, err = template.New("pass_reset_body").Parse(ec.EmailPassResetBody)
	err = tmpl.Execute(bodyBuf, templateData)
	if err != nil {
		return "", "", err
	}
	return titleBuf.String(), bodyBuf.String(), nil
}

func (es *EmailService) ChangeEmailTemplate(changeEmailUrl string) (title, body string, err error) {
	ec, err := es.getEmailConfig()
	if err != nil {
		return
	}

	templateData := ChangeEmailTemplateData{
		SiteName:       ec.EmailWebName,
		ChangeEmailUrl: changeEmailUrl,
	}
	tmpl, err := template.New("email_change_title").Parse(ec.EmailChangeTitle)
	if err != nil {
		return "", "", err
	}
	titleBuf := &bytes.Buffer{}
	bodyBuf := &bytes.Buffer{}
	err = tmpl.Execute(titleBuf, templateData)
	if err != nil {
		return "", "", err
	}

	tmpl, err = template.New("email_change_body").Parse(ec.EmailChangeBody)
	err = tmpl.Execute(bodyBuf, templateData)
	if err != nil {
		return "", "", err
	}
	return titleBuf.String(), bodyBuf.String(), nil
}

func (es *EmailService) getEmailConfig() (ec *EmailConfig, err error) {
	emailConf, err := es.configRepo.GetString("email.config")
	if err != nil {
		return nil, err
	}
	ec = &EmailConfig{}
	err = json.Unmarshal([]byte(emailConf), ec)
	if err != nil {
		return nil, err
	}
	return ec, nil
}
