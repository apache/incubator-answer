package plugin

type Embed interface {
	Base
}

var (
	// CallReviewer is a function that calls all registered parsers
	CallEmbed,
	registerEmbed = MakePlugin[Embed](false)
)
