package htmltext

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

func TestClearText(t *testing.T) {
	var (
		expected,
		clearedText string
	)

	// test code clear text
	expected = "hello{code...}"
	clearedText = ClearText("<p>hello<pre>var a = \"good\"</pre></p>")
	assert.Equal(t, expected, clearedText)

	// test link clear text
	expected = "hello [example.com]"
	clearedText = ClearText("<p>hello <a href=\"http://example.com/\">example.com</a></p>")
	assert.Equal(t, expected, clearedText)
	clearedText = ClearText("<p>hello<a href=\"https://example.com/\">example.com</a></p>")
	assert.Equal(t, expected, clearedText)

	expected = "hello world"
	clearedText = ClearText("<div> hello</div>\n<div>world</div>")
	assert.Equal(t, expected, clearedText)
}

func TestFetchExcerpt(t *testing.T) {
	var (
		expected,
		text string
	)

	// test english string
	expected = "hello..."
	text = FetchExcerpt("<p>hello world</p>", "...", 5)
	assert.Equal(t, expected, text)

	// test mixed string
	expected = "hello你好..."
	text = FetchExcerpt("<p>hello你好world</p>", "...", 7)
	assert.Equal(t, expected, text)

	// test mixed string with emoticon
	expected = "hello你好😂..."
	text = FetchExcerpt("<p>hello你好😂world</p>", "...", 8)
	assert.Equal(t, expected, text)

	expected = "hello你好"
	text = FetchExcerpt("<p>hello你好</p>", "...", 8)
	assert.Equal(t, expected, text)
}

func TestUrlTitle(t *testing.T) {
	list := []string{
		"hello你好😂...",
		"这是一个，标题，title",
	}
	for _, title := range list {
		formatTitle := UrlTitle(title)
		spew.Dump(formatTitle)
	}
}
