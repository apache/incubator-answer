package translator

import (
	"github.com/google/wire"
	myTran "github.com/segmentfault/pacman/contrib/i18n"
	"github.com/segmentfault/pacman/i18n"
)

// ProviderSet is providers.
var ProviderSet = wire.NewSet(NewTranslator)
var GlobalTrans i18n.Translator

// NewTranslator new a translator
func NewTranslator(c *I18n) (tr i18n.Translator, err error) {
	GlobalTrans, err = myTran.NewTranslator(c.BundleDir)
	return GlobalTrans, err
}
