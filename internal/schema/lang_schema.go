package schema

// GetLangOption get label option
type GetLangOption struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

var GetLangOptions = []*GetLangOption{
	{
		Label: "English(US)",
		Value: "en_US",
	},
	{
		Label: "中文(CN)",
		Value: "zh_CN",
	},
	{
		Label: "Tiếng Việt",
		Value: "vi_VN",
	},
}
