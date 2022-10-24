package handler

import (
	"github.com/segmentfault/pacman/i18n"
	"testing"
)

func TestGetLangFromHeader(t *testing.T) {
	type GetLangTest struct {
		get      string
		expected i18n.Language
	}
	tests := []GetLangTest{
		{"en", i18n.LanguageEnglish},
		{"en_US", i18n.LanguageEnglish},
		{"en-US", i18n.LanguageEnglish},
		{"zh", i18n.LanguageChinese},
		{"zh-CN", i18n.LanguageChinese},
		{"zh_CN", i18n.LanguageChinese},
		{"fr", i18n.DefaultLang},
		{"zh,en;q=0.9,en-US;q=0.8,zh-CN;q=0.7,zh-TW;q=0.6,la;q=0.5,ja;q=0.4,id;q=0.3,fr;q=0.2", i18n.LanguageChinese},
	}
	for _, test := range tests {
		got := getLangFromHeader(test.get)
		if got != test.expected {
			t.Errorf("input: %v, got=%v, expected=%v\n", test.get, got, test.expected)
		}
	}
}
