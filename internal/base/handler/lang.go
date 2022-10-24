package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/segmentfault/answer/internal/base/constant"
	"github.com/segmentfault/pacman/i18n"
	"sort"
	"strconv"
	"strings"
)

type LangQ struct {
	Lang string
	Q    float64
}

func parseAcceptLanguage(acptLang string) []LangQ {
	var lqs []LangQ

	langQStrs := strings.Split(acptLang, ",")
	for _, langQStr := range langQStrs {
		trimedLangQStr := strings.Trim(langQStr, " ")

		langQ := strings.Split(trimedLangQStr, ";")
		if len(langQ) == 1 {
			lq := LangQ{langQ[0], 1}
			lqs = append(lqs, lq)
		} else {
			qp := strings.Split(langQ[1], "=")
			q, err := strconv.ParseFloat(qp[1], 64)
			if err != nil {
				panic(err)
			}
			lq := LangQ{langQ[0], q}
			lqs = append(lqs, lq)
		}
	}
	sort.Slice(lqs, func(i, j int) bool {
		return lqs[i].Q > lqs[j].Q
	})
	return lqs
}

func getLangCode(code string) string {
	if strings.Contains(code, "-") {
		return strings.Split(code, "-")[0]
	}
	if strings.Contains(code, "_") {
		return strings.Split(code, "_")[0]
	}
	return strings.ToLower(code)
}

func getLangFromHeader(header string) i18n.Language {
	acceptLanguages := parseAcceptLanguage(header)
	zhCNCode := getLangCode(string(i18n.LanguageChinese))
	enUSCode := getLangCode(string(i18n.LanguageEnglish))
	for _, lq := range acceptLanguages {
		langCode := getLangCode(lq.Lang)
		if langCode == zhCNCode {
			return i18n.LanguageChinese
		}
		if langCode == enUSCode {
			return i18n.LanguageEnglish
		}
	}
	return i18n.DefaultLang
}

// GetLang get language from header
func GetLang(ctx *gin.Context) i18n.Language {
	acceptLanguageHeader := ctx.GetHeader(constant.AcceptLanguageFlag)
	return getLangFromHeader(acceptLanguageHeader)
}
