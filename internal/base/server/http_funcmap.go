package server

import (
	"html/template"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/answerdev/answer/internal/base/translator"
	"github.com/answerdev/answer/internal/controller"
	"github.com/answerdev/answer/internal/schema"
	"github.com/answerdev/answer/pkg/converter"
	"github.com/answerdev/answer/pkg/day"
	"github.com/answerdev/answer/pkg/htmltext"
	"github.com/segmentfault/pacman/i18n"
)

var funcMap = template.FuncMap{
	"replaceHTMLTag": func(src string, tags ...string) string {
		p := `(?U)<(\d+)>.+</(\d+)>`

		re := regexp.MustCompile(p)
		ms := re.FindAllStringSubmatch(src, -1)
		for _, mi := range ms {
			if mi[1] == mi[2] {
				i, err := strconv.Atoi(mi[1])
				if err != nil || len(tags) < i {
					break
				}

				src = strings.ReplaceAll(src, mi[0], tags[i-1])
			}
		}

		return src
	},
	"join": func(sep string, elems ...string) string {
		return strings.Join(elems, sep)
	},
	"templateHTML": func(data string) template.HTML {
		return template.HTML(data)
	},
	"formatLinkNofollow": func(data string) template.HTML {
		return template.HTML(FormatLinkNofollow(data))
	},
	"translator": func(la i18n.Language, data string, params ...interface{}) string {
		trans := translator.GlobalTrans.Tr(la, data)

		if len(params) > 0 && len(params)%2 == 0 {
			for i := 0; i < len(params); i += 2 {
				k := converter.InterfaceToString(params[i])
				v := converter.InterfaceToString(params[i+1])
				trans = strings.ReplaceAll(trans, "{{ "+k+" }}", v)
				trans = strings.ReplaceAll(trans, "{{"+k+"}}", v)
			}
		}

		return trans
	},
	"timeFormatISO": func(tz string, timestamp int64) string {
		_, _ = time.LoadLocation(tz)
		return time.Unix(timestamp, 0).Format("2006-01-02T15:04:05.000Z")
	},
	"translatorTimeFormatLongDate": func(la i18n.Language, tz string, timestamp int64) string {
		trans := translator.GlobalTrans.Tr(la, "ui.dates.long_date_with_time")
		return day.Format(timestamp, trans, tz)
	},
	"translatorTimeFormat": func(la i18n.Language, tz string, timestamp int64) string {
		var (
			now           = time.Now().Unix()
			between int64 = 0
			trans   string
		)
		_, _ = time.LoadLocation(tz)
		if now > timestamp {
			between = now - timestamp
		}

		if between <= 1 {
			return translator.GlobalTrans.Tr(la, "ui.dates.now")
		}

		if between > 1 && between < 60 {
			trans = translator.GlobalTrans.Tr(la, "ui.dates.x_seconds_ago")
			return strings.ReplaceAll(trans, "{{count}}", converter.IntToString(between))
		}

		if between >= 60 && between < 3600 {
			min := math.Floor(float64(between / 60))
			trans = translator.GlobalTrans.Tr(la, "ui.dates.x_minutes_ago")
			return strings.ReplaceAll(trans, "{{count}}", strconv.FormatFloat(min, 'f', 0, 64))
		}

		if between >= 3600 && between < 3600*24 {
			h := math.Floor(float64(between / 3600))
			trans = translator.GlobalTrans.Tr(la, "ui.dates.x_hours_ago")
			return strings.ReplaceAll(trans, "{{count}}", strconv.FormatFloat(h, 'f', 0, 64))
		}

		if between >= 3600*24 &&
			between < 3600*24*366 &&
			time.Unix(timestamp, 0).Format("2006") == time.Unix(now, 0).Format("2006") {
			trans = translator.GlobalTrans.Tr(la, "ui.dates.long_date")
			return day.Format(timestamp, trans, tz)
		}

		trans = translator.GlobalTrans.Tr(la, "ui.dates.long_date_with_year")
		return day.Format(timestamp, trans, tz)
	},
	"wrapComments": func(comments []*schema.GetCommentResp, la i18n.Language, tz string) map[string]interface{} {
		return map[string]interface{}{
			"comments": comments,
			"language": la,
			"timezone": tz,
		}
	},
	"urlTitle": func(title string) string {
		return htmltext.UrlTitle(title)
	},
}

func FormatLinkNofollow(html string) string {
	var hrefRegexp = regexp.MustCompile("(?m)<a.*?[^<]>.*?</a>")
	match := hrefRegexp.FindAllString(html, -1)
	if match != nil {
		for _, v := range match {
			hasNofollow := strings.Contains(v, "rel=\"nofollow\"")
			hasSiteUrl := strings.Contains(v, controller.SiteUrl)
			if !hasSiteUrl {
				if !hasNofollow {
					nofollowUrl := strings.Replace(v, "<a", "<a rel=\"nofollow\"", 1)
					html = strings.Replace(html, v, nofollowUrl, 1)
				}
			}
		}
	}
	return html
}
