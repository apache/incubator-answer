package plugin

type Parser interface {
	Base
	Parse(text string) (string, error)
}

var (
	// CallParser is a function that calls all registered parsers
	CallParser,
	registerParser = MakePlugin[Parser](false)
)
