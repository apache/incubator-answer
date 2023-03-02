package plugin

// Info presents the plugin information
type Info struct {
	Name        Translator
	SlugName    string
	Description Translator
	Author      string
	Version     string
	Link        string
}

// Base is the base plugin
type Base interface {
	// Info returns the plugin information
	Info() Info
}

var (
	// CallBase is a function that calls all registered base plugins
	CallBase,
	registerBase = MakePlugin[Base](true)
)
