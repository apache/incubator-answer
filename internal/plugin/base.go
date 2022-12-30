package plugin

// Info presents the plugin information
type Info struct {
	Name        string
	Description string
	Author      string
	Version     string
	Disabled    bool
}

// Base is the base plugin
type Base interface {
	// Info returns the plugin information
	Info() Info
}

var (
	// CallBase is a function that calls all registered base plugins
	CallBase,
	registerBase = MakePlugin[Base]()
)
