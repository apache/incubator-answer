package day

import (
	"time"
)

var placeholder = map[string]string{
	"YY":   "06",      // 06	year
	"YYYY": "2006",    // 2006	full year
	"M":    "1",       // 1-12	month
	"MM":   "01",      // 01-12	month
	"MMM":  "Jan",     // Jan-Dec month
	"MMMM": "January", // January-December month
	"D":    "2",       // 1-31	date
	"DD":   "02",      // 01-31	date preset 0
	"H":    "15",      // 00-23	hour 24
	"HH":   "15",      // 00-23	hour 24
	"h":    "3",       // 1-12	hour 12
	"hh":   "03",      // 01-12	hour 12
	"m":    "4",       // 0-59	minute
	"mm":   "04",      // 00-59	minute
	"s":    "5",       // 0-59	second
	"ss":   "05",      // 00-59	second
	"A":    "PM",      // AM / PM
	"a":    "pm",      // am / pm
	"[at]": "at",      // at string
}

func Format(unix int64, format, tz string) (formatted string) {
	/*l := len(placeholders) - 1
	for i := l; i >= 0; i-- {
		format = strings.ReplaceAll(format, placeholders[i].old, placeholders[i].new)
	}*/
	toFormat := ""
	for len(format) > 0 {
		to, suffix := nextStdChunk(format)
		toFormat += to
		format = suffix
	}

	_, _ = time.LoadLocation(tz)
	formatted = time.Unix(unix, 0).Format(toFormat)
	return
}

func nextStdChunk(from string) (to, suffix string) {
	if len(from) == 0 {
		to = ""
		suffix = ""
		return
	}

	s := string(from[0])
	old := ""
	switch s {
	case "Y":
		if len(from) >= 4 && from[:4] == "YYYY" {
			old = "YYYY"
		} else if len(from) >= 2 && from[:2] == "YY" {
			old = "YY"
		}
	case "M":
		for i := 4; i > 0; i-- {
			if len(from) >= i {
				switch from[:i] {
				case "MMMM":
					old = "MMMM"
				case "MMM":
					old = "MMM"
				case "MM":
					old = "MM"
				case "M":
					old = "M"
				}
			}
			if old != "" {
				break
			}
		}
	case "D":
		for i := 2; i >= 0; i-- {
			if len(from) >= i {
				switch from[:i] {
				case "DD":
					old = "DD"
				case "D":
					old = "D"
				}
			}
			if old != "" {
				break
			}
		}
	case "H":
		for i := 2; i >= 0; i-- {
			if len(from) >= i {
				switch from[:i] {
				case "HH":
					old = "HH"
				case "H":
					old = "H"
				}
			}
			if old != "" {
				break
			}
		}
	case "h":
		for i := 2; i >= 0; i-- {
			if len(from) >= i {
				switch from[:i] {
				case "hh":
					old = "hh"
				case "h":
					old = "h"
				}
			}
			if old != "" {
				break
			}
		}
	case "m":
		for i := 2; i >= 0; i-- {
			if len(from) >= i {
				switch from[:i] {
				case "mm":
					old = "mm"
				case "m":
					old = "m"
				}
			}
			if old != "" {
				break
			}
		}
	case "s":
		for i := 2; i >= 0; i-- {
			if len(from) >= i {
				switch from[:i] {
				case "ss":
					old = "ss"
				case "s":
					old = "s"
				}
			}
			if old != "" {
				break
			}
		}
	case "A":
		old = "A"
	case "a":
		old = "a"
	case "[":
		if len(from) >= 4 && from[:4] == "[at]" {
			old = "[at]"
		}
	default:
		old = s
	}

	to, ok := placeholder[old]
	if !ok {
		to = old
	}

	suffix = from[len(old):]
	return
}
