package answercmd

import (
	"fmt"
	"os"
	"time"

	"github.com/answerdev/answer/internal/base/conf"
	"github.com/answerdev/answer/internal/base/constant"
	"github.com/answerdev/answer/internal/base/cron"
	"github.com/answerdev/answer/internal/cli"
	"github.com/answerdev/answer/internal/schema"
	"github.com/gin-gonic/gin"
	"github.com/segmentfault/pacman"
	"github.com/segmentfault/pacman/contrib/log/zap"
	"github.com/segmentfault/pacman/contrib/server/http"
	"github.com/segmentfault/pacman/log"
)

// go build -ldflags "-X github.com/answerdev/answer/cmd.Version=x.y.z"
var (
	// Name is the name of the project
	Name = "answer"
	// Version is the version of the project
	Version = "0.0.0"
	// Revision is the git short commit revision number
	Revision = "-"
	// Time is the build time of the project
	Time = "-"
	// log level
	logLevel = os.Getenv("LOG_LEVEL")
	// log path
	logPath = os.Getenv("LOG_PATH")
)

// Main
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func Main() {
	log.SetLogger(zap.NewLogger(
		log.ParseLevel(logLevel), zap.WithName("answer"), zap.WithPath(logPath), zap.WithCallerFullPath()))
	Execute()
}

func runApp() {
	c, err := conf.ReadConfig(cli.GetConfigFilePath())
	if err != nil {
		panic(err)
	}
	app, cleanup, err := initApplication(
		c.Debug, c.Server, c.Data.Database, c.Data.Cache, c.I18n, c.Swaggerui, c.ServiceConfig, log.GetLogger())
	if err != nil {
		panic(err)
	}
	constant.Version = Version
	constant.Revision = Revision
	schema.AppStartTime = time.Now()
	fmt.Println("answer Version:", constant.Version, " Revision:", constant.Revision)

	defer cleanup()
	if err := app.Run(); err != nil {
		panic(err)
	}
}

func newApplication(serverConf *conf.Server, server *gin.Engine, manager *cron.ScheduledTaskManager) *pacman.Application {
	manager.Run()
	return pacman.NewApp(
		pacman.WithName(Name),
		pacman.WithVersion(Version),
		pacman.WithServer(http.NewServer(server, serverConf.HTTP.Addr)),
	)
}
