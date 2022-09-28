package main

import (
	"flag"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/segmentfault/answer/internal/base/conf"
	"github.com/segmentfault/answer/internal/cli"
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
	Version string
	// confFlag is the config flag.
	confFlag string
	// log level
	logLevel = os.Getenv("LOG_LEVEL")
	// log path
	logPath = os.Getenv("LOG_PATH")
)

func init() {
	flag.StringVar(&confFlag, "c", "../../configs/config.yaml", "config path, eg: -c config.yaml")
}

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	flag.Parse()

	args := flag.Args()

	if len(args) < 1 {
		cli.Usage()
		os.Exit(0)
		return
	}

	if args[0] == "init" {
		cli.InitConfig()
		return
	}
	if len(args) >= 3 {
		if args[0] == "run" && args[1] == "-c" {
			confFlag = args[2]
		}
	}

	log.SetLogger(zap.NewLogger(
		log.ParseLevel(logLevel), zap.WithName(Name), zap.WithPath(logPath), zap.WithCallerFullPath()))

	// init config
	c := &conf.AllConfig{}
	config, err := viper.NewWithPath(confFlag)
	if err != nil {
		panic(err)
	}
	if err = config.Parse(&c); err != nil {
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

func newApplication(serverConf *conf.Server, server *gin.Engine) *pacman.Application {
	return pacman.NewApp(
		pacman.WithName(Name),
		pacman.WithVersion(Version),
		pacman.WithServer(http.NewServer(server, serverConf.HTTP.Addr)),
	)
}
