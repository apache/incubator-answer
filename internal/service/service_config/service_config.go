package service_config

type ServiceConfig struct {
	SecretKey  string `json:"secret_key" mapstructure:"secret_key"`
	WebHost    string `json:"web_host" mapstructure:"web_host"`
	UploadPath string `json:"upload_path" mapstructure:"upload_path"`
}
