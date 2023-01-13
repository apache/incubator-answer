package schema

const (
	PluginStatusActive   PluginStatus = "active"
	PluginStatusInactive PluginStatus = "inactive"
)

type PluginStatus string

type GetPluginListReq struct {
	Status     PluginStatus `form:"status"`
	HaveConfig bool         `form:"have_config"`
}

type GetPluginListResp struct {
	Name        string `json:"name"`
	SlugName    string `json:"slug_name"`
	Description string `json:"description"`
	Version     string `json:"version"`
	Enabled     bool   `json:"enabled"`
	HaveConfig  bool   `json:"have_config"`
}

type UpdatePluginStatusReq struct {
	PluginSlugName string `validate:"required,gt=1,lte=100" json:"plugin_slug_name"`
	Enabled        bool   `json:"enabled"`
}

type GetPluginConfigReq struct {
	PluginSlugName string `validate:"required,gt=1,lte=100" form:"plugin_slug_name"`
}

type GetPluginConfigResp struct {
	//ConfigFields []plugin.ConfigField `json:"config_fields"`
	Name         string         `json:"name"`
	SlugName     string         `json:"slug_name"`
	Description  string         `json:"description"`
	Version      string         `json:"version"`
	ConfigFields []*ConfigField `json:"config_fields"`
}

const (
	TEXT     ConfigFieldType = "text"
	Select   ConfigFieldType = "select"
	Checkbox ConfigFieldType = "checkbox"
)

type ConfigFieldType string

type ConfigField struct {
	Name        string          `json:"name"`
	Type        ConfigFieldType `json:"type"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
	Required    bool            `json:"required"`
	Value       string          `json:"value"`
	UIOptions   UIOptions       `json:"ui_options"`
	Options     []Option        `json:"options"`
}

type UIOptions struct {
	Placeholder string `json:"placeholder"`
	Type        string `json:"type"`
}

type Option struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

type UpdatePluginConfigReq struct {
	PluginSlugName string         `validate:"required,gt=1,lte=100" json:"plugin_slug_name"`
	ConfigFields   map[string]any `json:"config_fields"`
}
