package plugin

type ConfigType string
type InputType string

const (
	ConfigTypeInput    ConfigType = "input"
	ConfigTypeTextarea ConfigType = "textarea"
	ConfigTypeCheckbox ConfigType = "checkbox"
	ConfigTypeRadio    ConfigType = "radio"
	ConfigTypeSelect   ConfigType = "select"
	ConfigTypeUpload   ConfigType = "upload"
	ConfigTypeTimezone ConfigType = "timezone"
	ConfigTypeSwitch   ConfigType = "switch"
)

const (
	InputTypeText     InputType = "text"
	InputTypeColor    InputType = "color"
	InputTypeDate     InputType = "date"
	InputTypeDatetime InputType = "datetime-local"
	InputTypeEmail    InputType = "email"
	InputTypeMonth    InputType = "month"
	InputTypeNumber   InputType = "number"
	InputTypePassword InputType = "password"
	InputTypeRange    InputType = "range"
	InputTypeSearch   InputType = "search"
	InputTypeTel      InputType = "tel"
	InputTypeTime     InputType = "time"
	InputTypeUrl      InputType = "url"
	InputTypeWeek     InputType = "week"
)

type ConfigField struct {
	Name        string               `json:"name"`
	Type        ConfigType           `json:"type"`
	Title       Translator           `json:"title"`
	Description Translator           `json:"description"`
	Required    bool                 `json:"required"`
	Value       string               `json:"value"`
	UIOptions   ConfigFieldUIOptions `json:"ui_options"`
	Options     []ConfigFieldOption  `json:"options,omitempty"`
}

type ConfigFieldUIOptions struct {
	Placeholder Translator `json:"placeholder,omitempty"`
	Rows        string     `json:"rows,omitempty"`
	InputType   InputType  `json:"input_type,omitempty"`
}

type ConfigFieldOption struct {
	Label Translator `json:"label"`
	Value string     `json:"value"`
}

type Config interface {
	Base

	// ConfigFields returns the list of config fields
	ConfigFields() []ConfigField

	// ConfigReceiver receives the config data, it calls when the config is saved or initialized.
	// We recommend to unmarshal the data to a struct, and then use the struct to do something.
	// The config is encoded in JSON format.
	// It depends on the definition of ConfigFields.
	ConfigReceiver(config []byte) error
}

var (
	// CallConfig is a function that calls all registered config plugins
	CallConfig,
	registerConfig = MakePlugin[Config](true)
)
