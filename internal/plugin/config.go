package plugin

type ConfigType int

const (
	ConfigTypeInput ConfigType = iota
	ConfigTypeTextarea
	ConfigTypeSelect
	ConfigTypeCheckbox
)

type ConfigField struct {
	Name        string
	Description string
	Required    bool
	Type        ConfigType
	Items       []ConfigFieldItem
}

type ConfigFieldItem struct {
	Name        string
	Label       string
	Value       string
	PlaceHolder string
	Selected    bool
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
