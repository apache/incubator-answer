package email

import (
	"bytes"
	"fmt"
	"net/smtp"
	"text/template"

	"github.com/jordan-wright/email"
	"github.com/segmentfault/pacman/log"
)

// EmailClient
type EmailClient struct {
	email  *email.Email
	config *Config
}

// Config .
type Config struct {
	WebName           string `json:"web_name"`
	WebHost           string `json:"web_host"`
	SecretKey         string `json:"secret_key"`
	UserSessionKey    string `json:"user_session_key"`
	EmailFrom         string `json:"email_from"`
	EmailFromPass     string `json:"email_from_pass"`
	EmailFromHostname string `json:"email_from_hostname"`
	EmailFromSMTP     string `json:"email_from_smtp"`
	EmailFromName     string `json:"email_from_name"`
	RegisterTitle     string `json:"register_title"`
	RegisterBody      string `json:"register_body"`
	PassResetTitle    string `json:"pass_reset_title"`
	PassResetBody     string `json:"pass_reset_body"`
}

// NewEmailClient
func NewEmailClient() *EmailClient {
	return &EmailClient{
		email: email.NewEmail(),
	}
}

func (s *EmailClient) Send(ToEmail, Title, Body string) {
	from := s.config.EmailFrom
	fromPass := s.config.EmailFromPass
	fromName := s.config.EmailFromName
	fromSmtp := s.config.EmailFromSMTP
	fromHostName := s.config.EmailFromHostname
	s.email.From = fmt.Sprintf("%s <%s>", fromName, from)
	s.email.To = []string{ToEmail}
	s.email.Subject = Title
	s.email.HTML = []byte(Body)
	err := s.email.Send(fromSmtp, smtp.PlainAuth("", from, fromPass, fromHostName))
	if err != nil {
		log.Error(err)
	}
}

func (s *EmailClient) RegisterTemplate(RegisterUrl string) (Title, Body string, err error) {
	webName := s.config.WebName
	templateData := RegisterTemplateData{webName, RegisterUrl}
	tmpl, err := template.New("register_title").Parse(s.config.RegisterTitle)
	if err != nil {
		return "", "", err
	}
	title := new(bytes.Buffer)
	body := new(bytes.Buffer)
	err = tmpl.Execute(title, templateData)
	if err != nil {
		return "", "", err
	}

	tmpl, err = template.New("register_body").Parse(s.config.RegisterBody)
	err = tmpl.Execute(body, templateData)
	if err != nil {
		return "", "", err
	}

	return title.String(), body.String(), nil
}

func (s *EmailClient) PassResetTemplate(PassResetUrl string) (Title, Body string, err error) {
	webName := s.config.WebName
	templateData := PassResetTemplateData{webName, PassResetUrl}
	tmpl, err := template.New("pass_reset_title").Parse(s.config.PassResetTitle)
	if err != nil {
		return "", "", err
	}
	title := new(bytes.Buffer)
	body := new(bytes.Buffer)
	err = tmpl.Execute(title, templateData)
	if err != nil {
		return "", "", err
	}

	tmpl, err = template.New("pass_reset_body").Parse(s.config.PassResetBody)
	err = tmpl.Execute(body, templateData)
	if err != nil {
		return "", "", err
	}

	return title.String(), body.String(), nil
}
