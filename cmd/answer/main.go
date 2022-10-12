package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/segmentfault/answer/internal/base/conf"
	"github.com/segmentfault/pacman"
	"github.com/segmentfault/pacman/contrib/conf/viper"
	"github.com/segmentfault/pacman/contrib/log/zap"
	"github.com/segmentfault/pacman/contrib/server/http"
	"github.com/segmentfault/pacman/log"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the project
	Name = "answer"
	// Version is the version of the project
	Version = "unknown"
	// log level
	logLevel = os.Getenv("LOG_LEVEL")
	// log path
	logPath = os.Getenv("LOG_PATH")
)

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	Execute()
	return
}

func runApp() {
	log.SetLogger(zap.NewLogger(
		log.ParseLevel(logLevel), zap.WithName("answer"), zap.WithPath(logPath), zap.WithCallerFullPath()))

	c, err := readConfig()
	if err != nil {
		panic(err)
	}
	app, cleanup, err := initApplication(
		c.Debug, c.Server, c.Data.Database, c.Data.Cache, c.I18n, c.Swaggerui, c.ServiceConfig, log.GetLogger())
	if err != nil {
		panic(err)
	}
	defer cleanup()
	if err := app.Run(); err != nil {
		panic(err)
	}
}

func readConfig() (c *conf.AllConfig, err error) {
	c = &conf.AllConfig{}
	config, err := viper.NewWithPath(confFlag)
	if err != nil {
		return nil, err
	}
	if err = config.Parse(&c); err != nil {
		return nil, err
	}
	return c, nil
}

func newApplication(serverConf *conf.Server, server *gin.Engine) *pacman.Application {
	return pacman.NewApp(
		pacman.WithName(Name),
		pacman.WithVersion(Version),
		pacman.WithServer(http.NewServer(server, serverConf.HTTP.Addr)),
	)
}
