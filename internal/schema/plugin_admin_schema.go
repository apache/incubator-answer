package schema

type UpdatePluginStatusReq struct {
	PluginSlugName string `validate:"required,gt=1,lte=100" json:"plugin_slug_name"`
	Enabled        bool   `json:"enabled"`
}
