package translator

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/google/wire"
	myTran "github.com/segmentfault/pacman/contrib/i18n"
	"github.com/segmentfault/pacman/i18n"
	"sigs.k8s.io/yaml"
)

// ProviderSet is providers.
var ProviderSet = wire.NewSet(NewTranslator)
var GlobalTrans i18n.Translator

// LangOption language option
type LangOption struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

// LanguageOptions language
var LanguageOptions []*LangOption

// NewTranslator new a translator
func NewTranslator(c *I18n) (tr i18n.Translator, err error) {
	GlobalTrans, err = myTran.NewTranslator(c.BundleDir)
	if err != nil {
		return nil, err
	}

	i18nFile, err := os.ReadFile(filepath.Join(c.BundleDir, "i18n.yaml"))
	if err != nil {
		return nil, fmt.Errorf("read i18n file failed: %s", err)
	}

	s := struct {
		LangOption []*LangOption `json:"language_options"`
	}{}
	err = yaml.Unmarshal(i18nFile, &s)
	if err != nil {
		return nil, fmt.Errorf("i18n file parsing failed: %s", err)
	}
	LanguageOptions = s.LangOption
	return GlobalTrans, err
}

// CheckLanguageIsValid check user input language is valid
func CheckLanguageIsValid(lang string) bool {
	for _, option := range LanguageOptions {
		if option.Value == lang {
			return true
		}
	}
	return false
}
