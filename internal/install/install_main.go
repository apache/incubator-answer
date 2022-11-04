package install

import (
	"os"

	"github.com/answerdev/answer/internal/base/translator"
	"github.com/answerdev/answer/internal/cli"
)

var (
	port     = os.Getenv("INSTALL_PORT")
	confPath = ""
)

func Run(configPath string) {
	confPath = configPath
	// initialize translator for return internationalization error when installing.
	_, err := translator.NewTranslator(&translator.I18n{BundleDir: cli.I18nPath})
	if err != nil {
		panic(err)
	}

	installServer := NewInstallHTTPServer()
	if len(port) == 0 {
		port = "80"
	}
	if err = installServer.Run(":" + port); err != nil {
		panic(err)
	}
}
