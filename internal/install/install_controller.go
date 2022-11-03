package install

import (
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

// 1、校验配置文件 post installation/config-file/check
//2、校验数据库 post installation/db/check
//3、创建配置文件和数据库 post installation/init
//4、配置网站基本信息和超级管理员信息 post installation/base-info

// LangOptions get installation language options
// @Summary get installation language options
// @Description get installation language options
// @Tags Lang
// @Produce json
// @Success 200 {object} handler.RespBody{data=[]*translator.LangOption}
// @Router /installation/language/options [get]
func LangOptions(ctx *gin.Context) {
	handler.HandleResponse(ctx, nil, translator.LanguageOptions)
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
	resp.DbTableExist = cli.CheckDB(allConfig.Data.Database, true)
	handler.HandleResponse(ctx, nil, resp)
}

// CheckDatabase check database if exist when installation
// @Summary check database if exist when installation
// @Description check database if exist when installation
// @Tags installation
// @Accept json
// @Produce json
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
	resp.ConnectionSuccess = cli.CheckDB(dataConf, true)
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
// @Success 200 {object} handler.RespBody{data=install.CheckConfigFileResp{}}
// @Router /installation/init [post]
func InitEnvironment(ctx *gin.Context) {
	req := &CheckDatabaseReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}

	err := cli.InstallConfigFile(confPath)
	if err != nil {
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
		err = errors.BadRequest(reason.ReadConfigFailed)
		handler.HandleResponse(ctx, err, nil)
		return
	}

	if err := migrations.InitDB(c.Data.Database); err != nil {
		log.Error("init database error: ", err.Error())
		handler.HandleResponse(ctx, errors.BadRequest(reason.DatabaseConnectionFailed), schema.ErrTypeAlert)
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
// @Success 200 {object} handler.RespBody{data=install.CheckConfigFileResp{}}
// @Router /installation/base-info [post]
func InitBaseInfo(ctx *gin.Context) {
	req := &InitBaseInfoReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}

	// 修改配置文件
	// 修改管理员和对应信息
	handler.HandleResponse(ctx, nil, nil)
	return
}
