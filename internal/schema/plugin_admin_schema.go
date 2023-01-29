package schema

import "github.com/answerdev/answer/plugin"

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
	Name         string               `json:"name"`
	SlugName     string               `json:"slug_name"`
	Description  string               `json:"description"`
	Version      string               `json:"version"`
	ConfigFields []plugin.ConfigField `json:"config_fields"`
}

type UpdatePluginConfigReq struct {
	PluginSlugName string         `validate:"required,gt=1,lte=100" json:"plugin_slug_name"`
	ConfigFields   map[string]any `json:"config_fields"`
}
