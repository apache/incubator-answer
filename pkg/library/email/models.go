package email

type RegisterTemplateData struct {
	SiteName    string
	RegisterUrl string
}

type PassResetTemplateData struct {
	SiteName     string
	PassResetUrl string
}
