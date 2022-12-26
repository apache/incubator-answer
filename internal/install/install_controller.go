package install

import (
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/answerdev/answer/configs"
	"github.com/answerdev/answer/internal/base/conf"
	"github.com/answerdev/answer/internal/base/data"
	"github.com/answerdev/answer/internal/base/handler"
	"github.com/answerdev/answer/internal/base/reason"
	"github.com/answerdev/answer/internal/base/translator"
	"github.com/answerdev/answer/internal/cli"
	"github.com/answerdev/answer/internal/migrations"
	"github.com/answerdev/answer/internal/schema"
	"github.com/gin-gonic/gin"
	"github.com/segmentfault/pacman/errors"
	"github.com/segmentfault/pacman/log"
)

// LangOptions get installation language options
// @Summary get installation language options
// @Description get installation language options
// @Tags Lang
// @Produce json
// @Success 200 {object} handler.RespBody{data=[]translator.LangOption}
// @Router /installation/language/options [get]
func LangOptions(ctx *gin.Context) {
	handler.HandleResponse(ctx, nil, translator.LanguageOptions)
}

// CheckConfigFileAndRedirectToInstallPage if config file not exist try to redirect to install page
// @Summary if config file not exist try to redirect to install page
// @Description if config file not exist try to redirect to install page
// @Tags installation
// @Accept json
// @Produce json
// @Router / [get]
func CheckConfigFileAndRedirectToInstallPage(ctx *gin.Context) {
	if cli.CheckConfigFile(confPath) {
		ctx.Redirect(http.StatusFound, "/50x")
	} else {
		ctx.Redirect(http.StatusFound, "/install")
	}
}

// CheckConfigFile check config file if exist when installation
// @Summary check config file if exist when installation
// @Description check config file if exist when installation
// @Tags installation
// @Accept json
// @Produce json
// @Success 200 {object} handler.RespBody{data=install.CheckConfigFileResp{}}
// @Router /installation/config-file/check [post]
func CheckConfigFile(ctx *gin.Context) {
	resp := &CheckConfigFileResp{}
	resp.ConfigFileExist = cli.CheckConfigFile(confPath)
	if !resp.ConfigFileExist {
		handler.HandleResponse(ctx, nil, resp)
		return
	}
	allConfig, err := conf.ReadConfig(confPath)
	if err != nil {
		log.Error(err)
		err = errors.BadRequest(reason.ReadConfigFailed)
		handler.HandleResponse(ctx, err, nil)
		return
	}
	resp.DBConnectionSuccess = cli.CheckDBConnection(allConfig.Data.Database)
	if resp.DBConnectionSuccess {
		resp.DbTableExist = cli.CheckDBTableExist(allConfig.Data.Database)
	}
	handler.HandleResponse(ctx, nil, resp)
}

// CheckDatabase check database if exist when installation
// @Summary check database if exist when installation
// @Description check database if exist when installation
// @Tags installation
// @Accept json
// @Produce json
// @Param data body install.CheckDatabaseReq  true "CheckDatabaseReq"
// @Success 200 {object} handler.RespBody{data=install.CheckConfigFileResp{}}
// @Router /installation/db/check [post]
func CheckDatabase(ctx *gin.Context) {
	req := &CheckDatabaseReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}

	resp := &CheckDatabaseResp{}
	dataConf := &data.Database{
		Driver:     req.DbType,
		Connection: req.GetConnection(),
	}
	resp.ConnectionSuccess = cli.CheckDBConnection(dataConf)
	if !resp.ConnectionSuccess {
		handler.HandleResponse(ctx, errors.BadRequest(reason.DatabaseConnectionFailed), schema.ErrTypeAlert)
		return
	}
	handler.HandleResponse(ctx, nil, resp)
}

// InitEnvironment init environment
// @Summary init environment
// @Description init environment
// @Tags installation
// @Accept json
// @Produce json
// @Param data body install.CheckDatabaseReq  true "CheckDatabaseReq"
// @Success 200 {object} handler.RespBody{}
// @Router /installation/init [post]
func InitEnvironment(ctx *gin.Context) {
	req := &CheckDatabaseReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}

	// check config file if exist
	if cli.CheckConfigFile(confPath) {
		log.Debug("config file already exists")
		handler.HandleResponse(ctx, nil, nil)
		return
	}

	if err := cli.InstallConfigFile(confPath); err != nil {
		handler.HandleResponse(ctx, errors.BadRequest(reason.InstallConfigFailed), &InitEnvironmentResp{
			Success:            false,
			CreateConfigFailed: true,
			DefaultConfig:      string(configs.Config),
			ErrType:            schema.ErrTypeAlert.ErrType,
		})
		return
	}

	c, err := conf.ReadConfig(confPath)
	if err != nil {
		log.Errorf("read config failed %s", err)
		handler.HandleResponse(ctx, errors.BadRequest(reason.ReadConfigFailed), nil)
		return
	}
	c.Data.Database.Driver = req.DbType
	c.Data.Database.Connection = req.GetConnection()
	c.Data.Cache.FilePath = filepath.Join(cli.CacheDir, cli.DefaultCacheFileName)
	c.I18n.BundleDir = cli.I18nPath
	c.ServiceConfig.UploadPath = cli.UploadFilePath

	if err := conf.RewriteConfig(confPath, c); err != nil {
		log.Errorf("rewrite config failed %s", err)
		handler.HandleResponse(ctx, errors.BadRequest(reason.ReadConfigFailed), nil)
		return
	}
	handler.HandleResponse(ctx, nil, nil)
}

// InitBaseInfo init base info
// @Summary init base info
// @Description init base info
// @Tags installation
// @Accept json
// @Produce json
// @Param data body install.InitBaseInfoReq  true "InitBaseInfoReq"
// @Success 200 {object} handler.RespBody{}
// @Router /installation/base-info [post]
func InitBaseInfo(ctx *gin.Context) {
	req := &InitBaseInfoReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}
	req.FormatSiteUrl()

	c, err := conf.ReadConfig(confPath)
	if err != nil {
		log.Errorf("read config failed %s", err)
		handler.HandleResponse(ctx, errors.BadRequest(reason.ReadConfigFailed), nil)
		return
	}

	if cli.CheckDBTableExist(c.Data.Database) {
		log.Warn("database is already initialized")
		handler.HandleResponse(ctx, nil, nil)
		return
	}

	if err := migrations.InitDB(c.Data.Database); err != nil {
		log.Error("init database error: ", err.Error())
		handler.HandleResponse(ctx, errors.BadRequest(reason.InstallCreateTableFailed), schema.ErrTypeAlert)
		return
	}

	err = migrations.UpdateInstallInfo(c.Data.Database, req.Language, req.SiteName, req.SiteURL, req.ContactEmail,
		req.AdminName, req.AdminPassword, req.AdminEmail)
	if err != nil {
		log.Error(err)
		handler.HandleResponse(ctx, errors.BadRequest(reason.InstallConfigFailed), nil)
		return
	}

	handler.HandleResponse(ctx, nil, nil)
	go func() {
		time.Sleep(1 * time.Second)
		os.Exit(0)
	}()
}
