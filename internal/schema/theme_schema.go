package schema

// GetThemeOption get label option
type GetThemeOption struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

var GetThemeOptions = []*GetThemeOption{
	{
		Label: "Default",
		Value: "default",
	},
}
