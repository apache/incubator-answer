package service_config

type ServiceConfig struct {
	SecretKey  string `json:"secret_key" mapstructure:"secret_key" yaml:"secret_key"`
	UploadPath string `json:"upload_path" mapstructure:"upload_path" yaml:"upload_path"`
}
