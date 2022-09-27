package router

// SwaggerConfig struct describes configure for the Swagger API endpoint
type SwaggerConfig struct {
	Show     bool   `json:"show"`
	Protocol string `json:"protocol"`
	Host     string `json:"host"`
	Address  string `json:"address"`
}
