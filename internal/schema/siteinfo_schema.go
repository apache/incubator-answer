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
