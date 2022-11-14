package htmltext

import (
	"github.com/grokify/html-strip-tags-go"
	"regexp"
	"strings"
)

// ClearText clear HTML, get the clear text
func ClearText(html string) (text string) {
	if len(html) == 0 {
		text = html
		return
	}

	var (
		re        *regexp.Regexp
		codeReg   = `(?ism)<(pre)>.*<\/pre>`
		codeRepl  = "{code...}"
		linkReg   = `(?ism)<a.*?[^<]>(.*)?<\/a>`
		linkRepl  = " [$1] "
		spaceReg  = ` +`
		spaceRepl = " "
	)
	re = regexp.MustCompile(codeReg)
	html = re.ReplaceAllString(html, codeRepl)

	re = regexp.MustCompile(linkReg)
	html = re.ReplaceAllString(html, linkRepl)

	text = strings.NewReplacer(
		"\n", " ",
		"\r", " ",
		"\t", " ",
	).Replace(strip.StripTags(html))

	// replace multiple spaces to one space
	re = regexp.MustCompile(spaceReg)
	text = strings.TrimSpace(re.ReplaceAllString(text, spaceRepl))
	return
}

// FetchExcerpt return the excerpt from the HTML string
func FetchExcerpt(html, trimMarker string, limit int) (text string) {
	if len(html) == 0 {
		text = html
		return
	}

	text = ClearText(html)
	runeText := []rune(text)
	if len(runeText) <= limit {
		text = string(runeText)
	} else {
		text = string(runeText[0:limit])
	}

	text += trimMarker
	return
}
