package converter

import (
	"github.com/gomarkdown/markdown"
)

// Markdown2HTML convert markdown to html
func Markdown2HTML(md string) string {
	html := markdown.ToHTML([]byte(md), nil, nil)
	return string(html)
}
