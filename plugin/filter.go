package plugin

type Filter interface {
	Base
	FilterText(text string) (err error)
}

var (
	// CallFilter is a function that calls all registered parsers
	CallFilter,
	registerFilter = MakePlugin[Filter](false)
)
