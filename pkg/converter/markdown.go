package converter

import (
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
)

// Markdown2HTML convert markdown to html
func Markdown2HTML(md string) string {
	extensions := parser.HardLineBreak | parser.CommonExtensions
	p := parser.NewWithExtensions(extensions)
	html := markdown.ToHTML([]byte(md), p, nil)
	return string(html)
}
