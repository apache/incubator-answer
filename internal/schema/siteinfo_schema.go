package schema

import (
	"context"
	"fmt"
	"net/mail"
	"net/url"

	"github.com/answerdev/answer/internal/base/handler"
	"github.com/answerdev/answer/internal/base/reason"
	"github.com/answerdev/answer/internal/base/translator"
	"github.com/answerdev/answer/internal/base/validator"
	"github.com/segmentfault/pacman/errors"
)

const PermaLinkQuestionIDAndTitle = 1
const PermaLinkQuestionID = 2

// SiteGeneralReq site general request
type SiteGeneralReq struct {
	Name             string `validate:"required,sanitizer,gt=1,lte=128" form:"name" json:"name"`
	ShortDescription string `validate:"omitempty,sanitizer,gt=3,lte=255" form:"short_description" json:"short_description"`
	Description      string `validate:"omitempty,sanitizer,gt=3,lte=2000" form:"description" json:"description"`
	SiteUrl          string `validate:"required,sanitizer,gt=1,lte=512,url" form:"site_url" json:"site_url"`
	ContactEmail     string `validate:"required,sanitizer,gt=1,lte=512,email" form:"contact_email" json:"contact_email"`
}

type SiteSeoReq struct {
	PermaLink int    `validate:"required,lte=3,gte=0" form:"permalink" json:"permalink"`
	Robots    string `validate:"required" form:"robots" json:"robots"`
}

func (r *SiteGeneralReq) FormatSiteUrl() {
	parsedUrl, err := url.Parse(r.SiteUrl)
	if err != nil {
		return
	}
	r.SiteUrl = fmt.Sprintf("%s://%s", parsedUrl.Scheme, parsedUrl.Host)
}

// SiteInterfaceReq site interface request
type SiteInterfaceReq struct {
	Language string `validate:"required,gt=1,lte=128" form:"language" json:"language"`
	TimeZone string `validate:"required,gt=1,lte=128" form:"time_zone" json:"time_zone"`
}

// SiteBrandingReq site branding request
type SiteBrandingReq struct {
	Logo       string `validate:"omitempty,gt=0,lte=512" form:"logo" json:"logo"`
	MobileLogo string `validate:"omitempty,gt=0,lte=512" form:"mobile_logo" json:"mobile_logo"`
	SquareIcon string `validate:"omitempty,gt=0,lte=512" form:"square_icon" json:"square_icon"`
	Favicon    string `validate:"omitempty,gt=0,lte=512" form:"favicon" json:"favicon"`
}

// SiteWriteReq site write request
type SiteWriteReq struct {
	RequiredTag   bool     `validate:"omitempty" form:"required_tag" json:"required_tag"`
	RecommendTags []string `validate:"omitempty" form:"recommend_tags" json:"recommend_tags"`
	ReservedTags  []string `validate:"omitempty" form:"reserved_tags" json:"reserved_tags"`
	UserID        string   `json:"-"`
}

// SiteLegalReq site branding request
type SiteLegalReq struct {
	TermsOfServiceOriginalText string `json:"terms_of_service_original_text"`
	TermsOfServiceParsedText   string `json:"terms_of_service_parsed_text"`
	PrivacyPolicyOriginalText  string `json:"privacy_policy_original_text"`
	PrivacyPolicyParsedText    string `json:"privacy_policy_parsed_text"`
}

// GetSiteLegalInfoReq site site legal request
type GetSiteLegalInfoReq struct {
	InfoType string `validate:"required,oneof=tos privacy" form:"info_type"`
}

func (r *GetSiteLegalInfoReq) IsTOS() bool {
	return r.InfoType == "tos"
}

func (r *GetSiteLegalInfoReq) IsPrivacy() bool {
	return r.InfoType == "privacy"
}

// GetSiteLegalInfoResp get site legal info response
type GetSiteLegalInfoResp struct {
	TermsOfServiceOriginalText string `json:"terms_of_service_original_text,omitempty"`
	TermsOfServiceParsedText   string `json:"terms_of_service_parsed_text,omitempty"`
	PrivacyPolicyOriginalText  string `json:"privacy_policy_original_text,omitempty"`
	PrivacyPolicyParsedText    string `json:"privacy_policy_parsed_text,omitempty"`
}

// SiteLoginReq site login request
type SiteLoginReq struct {
	AllowNewRegistrations bool `json:"allow_new_registrations"`
	LoginRequired         bool `json:"login_required"`
}

// SiteCustomCssHTMLReq site custom css html
type SiteCustomCssHTMLReq struct {
	CustomHead   string `validate:"omitempty,gt=0,lte=65536" json:"custom_head"`
	CustomCss    string `validate:"omitempty,gt=0,lte=65536" json:"custom_css"`
	CustomHeader string `validate:"omitempty,gt=0,lte=65536" json:"custom_header"`
	CustomFooter string `validate:"omitempty,gt=0,lte=65536" json:"custom_footer"`
}

// SiteThemeReq site theme config
type SiteThemeReq struct {
	Theme       string                 `validate:"required,gt=0,lte=255" json:"theme"`
	ThemeConfig map[string]interface{} `validate:"omitempty" json:"theme_config"`
}

// SiteGeneralResp site general response
type SiteGeneralResp SiteGeneralReq

// SiteInterfaceResp site interface response
type SiteInterfaceResp SiteInterfaceReq

// SiteBrandingResp site branding response
type SiteBrandingResp SiteBrandingReq

// SiteLoginResp site login response
type SiteLoginResp SiteLoginReq

// SiteCustomCssHTMLResp site custom css html response
type SiteCustomCssHTMLResp SiteCustomCssHTMLReq

// SiteThemeResp site theme response
type SiteThemeResp struct {
	ThemeOptions []*ThemeOption         `json:"theme_options"`
	Theme        string                 `json:"theme"`
	ThemeConfig  map[string]interface{} `json:"theme_config"`
}

func (s *SiteThemeResp) TrTheme(ctx context.Context) {
	la := handler.GetLangByCtx(ctx)
	for _, option := range s.ThemeOptions {
		tr := translator.Tr(la, option.Value)
		// if tr is equal the option value means not found translation, so use the original label
		if tr != option.Value {
			option.Label = tr
		}
	}
}

// ThemeOption get label option
type ThemeOption struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

// SiteWriteResp site write response
type SiteWriteResp SiteWriteReq

// SiteLegalResp site write response
type SiteLegalResp SiteLegalReq

// SiteSeoResp site write response
type SiteSeoResp SiteSeoReq

// SiteInfoResp get site info response
type SiteInfoResp struct {
	General       *SiteGeneralResp       `json:"general"`
	Interface     *SiteInterfaceResp     `json:"interface"`
	Branding      *SiteBrandingResp      `json:"branding"`
	Login         *SiteLoginResp         `json:"login"`
	Theme         *SiteThemeResp         `json:"theme"`
	CustomCssHtml *SiteCustomCssHTMLResp `json:"custom_css_html"`
	SiteSeo       *SiteSeoReq            `json:"site_seo"`
}
type TemplateSiteInfoResp struct {
	General       *SiteGeneralResp       `json:"general"`
	Interface     *SiteInterfaceResp     `json:"interface"`
	Branding      *SiteBrandingResp      `json:"branding"`
	SiteSeo       *SiteSeoReq            `json:"site_seo"`
	CustomCssHtml *SiteCustomCssHTMLResp `json:"custom_css_html"`
	Title         string
	Year          string
	Canonical     string
	JsonLD        string
	Keywords      string
	Description   string
}

// UpdateSMTPConfigReq get smtp config request
type UpdateSMTPConfigReq struct {
	FromEmail          string `validate:"omitempty,gt=0,lte=256" json:"from_email"`
	FromName           string `validate:"omitempty,gt=0,lte=256" json:"from_name"`
	SMTPHost           string `validate:"omitempty,gt=0,lte=256" json:"smtp_host"`
	SMTPPort           int    `validate:"omitempty,min=1,max=65535" json:"smtp_port"`
	Encryption         string `validate:"omitempty,oneof=SSL" json:"encryption"` // "" SSL
	SMTPUsername       string `validate:"omitempty,gt=0,lte=256" json:"smtp_username"`
	SMTPPassword       string `validate:"omitempty,gt=0,lte=256" json:"smtp_password"`
	SMTPAuthentication bool   `validate:"omitempty" json:"smtp_authentication"`
	TestEmailRecipient string `validate:"omitempty,email" json:"test_email_recipient"`
}

func (r *UpdateSMTPConfigReq) Check() (errField []*validator.FormErrorField, err error) {
	_, err = mail.ParseAddress(r.FromName)
	if err == nil {
		return append(errField, &validator.FormErrorField{
			ErrorField: "from_name",
			ErrorMsg:   reason.SMTPConfigFromNameCannotBeEmail,
		}), errors.BadRequest(reason.SMTPConfigFromNameCannotBeEmail)
	}
	return nil, nil
}

// GetSMTPConfigResp get smtp config response
type GetSMTPConfigResp struct {
	FromEmail          string `json:"from_email"`
	FromName           string `json:"from_name"`
	SMTPHost           string `json:"smtp_host"`
	SMTPPort           int    `json:"smtp_port"`
	Encryption         string `json:"encryption"` // "" SSL
	SMTPUsername       string `json:"smtp_username"`
	SMTPPassword       string `json:"smtp_password"`
	SMTPAuthentication bool   `json:"smtp_authentication"`
}

// GetManifestJsonResp get manifest json response
type GetManifestJsonResp struct {
	ManifestVersion int               `json:"manifest_version"`
	Version         string            `json:"version"`
	ShortName       string            `json:"short_name"`
	Name            string            `json:"name"`
	Icons           map[string]string `json:"icons"`
	StartUrl        string            `json:"start_url"`
	Display         string            `json:"display"`
	ThemeColor      string            `json:"theme_color"`
	BackgroundColor string            `json:"background_color"`
}
