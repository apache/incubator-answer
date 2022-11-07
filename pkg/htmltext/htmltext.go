package htmltext

import (
	"github.com/grokify/html-strip-tags-go"
	"regexp"
	"strings"
)

// ClearText clear HTML, get the clear text
func ClearText(html string) (text string) {
	var (
		re        *regexp.Regexp
		codeReg   = `(?ism)<(pre)>.*<\/pre>`
		codeRepl  = "{code...}"
		linkReg   = `(?ism)<a.*?[^<]>.*?<\/a>`
		linkRepl  = "[link]"
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
	text = ClearText(html)
	runeText := []rune(text)
	text = string(runeText[0:limit])
	return
}
