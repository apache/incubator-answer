package router

// SwaggerConfig struct describes configure for the Swagger API endpoint
type SwaggerConfig struct {
	Show     bool   `json:"show" mapstructure:"show" yaml:"show"`
	Protocol string `json:"protocol" mapstructure:"protocol" yaml:"protocol"`
	Host     string `json:"host" mapstructure:"host" yaml:"host"`
	Address  string `json:"address" mapstructure:"address" yaml:"address"`
}
