package schema

// SiteGeneralReq site general request
type SiteGeneralReq struct {
	Name             string `validate:"required,gt=1,lte=128" comment:"site name" form:"name" json:"name"`
	ShortDescription string `validate:"required,gt=3,lte=255" comment:"short site description" form:"short_description" json:"short_description"`
	Description      string `validate:"required,gt=3,lte=2000" comment:"site description" form:"description" json:"description"`
}

// SiteInterfaceReq site interface request
type SiteInterfaceReq struct {
	Logo     string `validate:"omitempty,gt=0,lte=256" comment:"logo" form:"logo" json:"logo"`
	Theme    string `validate:"required,gt=1,lte=128" comment:"theme" form:"theme" json:"theme"`
	Language string `validate:"required,gt=1,lte=128" comment:"interface language" form:"language" json:"language"`
}

// SiteGeneralResp site general response
type SiteGeneralResp SiteGeneralReq

// SiteInterfaceResp site interface response
type SiteInterfaceResp SiteInterfaceReq

type SiteInfoResp struct {
	General *SiteGeneralResp   `json:"general"`
	Face    *SiteInterfaceResp `json:"interface"`
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
