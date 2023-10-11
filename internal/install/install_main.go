package install

import (
	"fmt"
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

	// try to install by env
	if installByEnv, err := TryToInstallByEnv(); installByEnv && err != nil {
		fmt.Printf("[auto-install] try to init by env fail: %v\n", err)
	}

	installServer := NewInstallHTTPServer()
	if len(port) == 0 {
		port = "5370"
	}
	fmt.Printf("[SUCCESS] answer installation service will run at: http://localhost:%s/install/ \n", port)
	if err = installServer.Run(":" + port); err != nil {
		panic(err)
	}
}
