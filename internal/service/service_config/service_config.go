package service_config

type ServiceConfig struct {
	UploadPath string `json:"upload_path" mapstructure:"upload_path" yaml:"upload_path"`
}
