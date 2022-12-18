package translator

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/google/wire"
	myTran "github.com/segmentfault/pacman/contrib/i18n"
	"github.com/segmentfault/pacman/i18n"
	"gopkg.in/yaml.v3"
)

// ProviderSet is providers.
var ProviderSet = wire.NewSet(NewTranslator)
var GlobalTrans i18n.Translator

// LangOption language option
type LangOption struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

// DefaultLangOption default language option. If user config the language is default, the language option is admin choose.
const DefaultLangOption = "Default"

var (
	// LanguageOptions language
	LanguageOptions []*LangOption
)

// NewTranslator new a translator
func NewTranslator(c *I18n) (tr i18n.Translator, err error) {
	entries, err := os.ReadDir(c.BundleDir)
	if err != nil {
		return nil, err
	}

	// read the Bundle resources file from entries
	for _, file := range entries {
		// ignore directory
		if file.IsDir() {
			continue
		}
		// ignore non-YAML file
		if filepath.Ext(file.Name()) != ".yaml" && file.Name() != "i18n.yaml" {
			continue
		}
		buf, err := os.ReadFile(filepath.Join(c.BundleDir, file.Name()))
		if err != nil {
			return nil, fmt.Errorf("read file failed: %s %s", file.Name(), err)
		}

		// parse the backend translation
		originalTr := struct {
			Backend map[string]map[string]interface{} `yaml:"backend"`
			UI      map[string]interface{}            `yaml:"ui"`
		}{}
		if err = yaml.Unmarshal(buf, &originalTr); err != nil {
			return nil, err
		}
		translation := make(map[string]interface{}, 0)
		for k, v := range originalTr.Backend {
			translation[k] = v
		}
		translation["backend"] = originalTr.Backend
		translation["ui"] = originalTr.UI

		content, err := yaml.Marshal(translation)
		if err != nil {
			return nil, fmt.Errorf("marshal translation content failed: %s %s", file.Name(), err)
		}

		// add translator use backend translation
		if err = myTran.AddTranslator(content, file.Name()); err != nil {
			return nil, fmt.Errorf("add translator failed: %s %s", file.Name(), err)
		}
	}
	GlobalTrans = myTran.GlobalTrans

	i18nFile, err := os.ReadFile(filepath.Join(c.BundleDir, "i18n.yaml"))
	if err != nil {
		return nil, fmt.Errorf("read i18n file failed: %s", err)
	}

	s := struct {
		LangOption []*LangOption `yaml:"language_options"`
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
	if lang == DefaultLangOption {
		return true
	}
	for _, option := range LanguageOptions {
		if option.Value == lang {
			return true
		}
	}
	return false
}
