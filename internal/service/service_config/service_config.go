package service_config

type ServiceConfig struct {
	SecretKey  string `json:"secret_key" mapstructure:"secret_key" yaml:"secret_key"`
	WebHost    string `json:"web_host" mapstructure:"web_host" yaml:"web_host"`
	UploadPath string `json:"upload_path" mapstructure:"upload_path" yaml:"upload_path"`
}
