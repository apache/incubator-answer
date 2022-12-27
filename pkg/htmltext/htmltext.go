package htmltext

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/gosimple/slug"
	strip "github.com/grokify/html-strip-tags-go"
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

func UrlTitle(title string) (text string) {
	title = ClearEmoji(title)
	title = slug.Make(title)
	// title = strings.ReplaceAll(title, " ", "-")
	title = url.QueryEscape(title)
	return title
}

func ClearEmoji(s string) string {
	ret := ""
	rs := []rune(s)
	for i := 0; i < len(rs); i++ {
		if len(string(rs[i])) != 4 {
			ret += string(rs[i])
		}
	}
	return ret
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
		return
	}

	text = string(runeText[0:limit]) + trimMarker
	return
}

func GetPicByUrl(Url string) string {
	res, err := http.Get(Url)
	if err != nil {
		return ""
	}
	defer res.Body.Close()
	pix, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return ""
	}
	return string(pix)
}
