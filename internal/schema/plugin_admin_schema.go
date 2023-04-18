package schema

import (
	"github.com/answerdev/answer/plugin"
	"github.com/gin-gonic/gin"
)

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
	Link        string `json:"link"`
}

type UpdatePluginStatusReq struct {
	PluginSlugName string `validate:"required,gt=1,lte=100" json:"plugin_slug_name"`
	Enabled        bool   `json:"enabled"`
}

type GetPluginConfigReq struct {
	PluginSlugName string `validate:"required,gt=1,lte=100" form:"plugin_slug_name"`
}

type GetPluginConfigResp struct {
	Name         string        `json:"name"`
	SlugName     string        `json:"slug_name"`
	Description  string        `json:"description"`
	Version      string        `json:"version"`
	ConfigFields []ConfigField `json:"config_fields"`
}

func (g *GetPluginConfigResp) SetConfigFields(ctx *gin.Context, fields []plugin.ConfigField) {
	for _, field := range fields {
		configField := ConfigField{
			Name:        field.Name,
			Type:        string(field.Type),
			Title:       field.Title.Translate(ctx),
			Description: field.Description.Translate(ctx),
			Required:    field.Required,
			Value:       field.Value,
			UIOptions: ConfigFieldUIOptions{
				Rows:      field.UIOptions.Rows,
				InputType: string(field.UIOptions.InputType),
			},
		}
		configField.UIOptions.Placeholder = field.UIOptions.Placeholder.Translate(ctx)
		configField.UIOptions.Label = field.UIOptions.Label.Translate(ctx)
		configField.UIOptions.Text = field.UIOptions.Text.Translate(ctx)
		if field.UIOptions.Action != nil {
			configField.UIOptions.Action = &ConfigFieldUIOptionAction{
				Url: field.UIOptions.Action.Url,
			}
		}

		for _, option := range field.Options {
			configField.Options = append(configField.Options, ConfigFieldOption{
				Label: option.Label.Translate(ctx),
				Value: option.Value,
			})
		}
		g.ConfigFields = append(g.ConfigFields, configField)
	}
}

type ConfigField struct {
	Name        string               `json:"name"`
	Type        string               `json:"type"`
	Title       string               `json:"title"`
	Description string               `json:"description"`
	Required    bool                 `json:"required"`
	Value       any                  `json:"value"`
	UIOptions   ConfigFieldUIOptions `json:"ui_options"`
	Options     []ConfigFieldOption  `json:"options,omitempty"`
}

type ConfigFieldUIOptions struct {
	Placeholder string                     `json:"placeholder,omitempty"`
	Rows        string                     `json:"rows,omitempty"`
	InputType   string                     `json:"input_type,omitempty"`
	Label       string                     `json:"label,omitempty"`
	Action      *ConfigFieldUIOptionAction `json:"action,omitempty"`
	Text        string                     `json:"text,omitempty"`
}

type ConfigFieldOption struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

type ConfigFieldUIOptionAction struct {
	Url string `json:"url"`
}

type UpdatePluginConfigReq struct {
	PluginSlugName string         `validate:"required,gt=1,lte=100" json:"plugin_slug_name"`
	ConfigFields   map[string]any `json:"config_fields"`
}
