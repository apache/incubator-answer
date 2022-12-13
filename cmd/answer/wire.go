//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/answerdev/answer/internal/base/conf"
	"github.com/answerdev/answer/internal/base/cron"
	"github.com/answerdev/answer/internal/base/data"
	"github.com/answerdev/answer/internal/base/middleware"
	"github.com/answerdev/answer/internal/base/server"
	"github.com/answerdev/answer/internal/base/translator"
	"github.com/answerdev/answer/internal/controller"
	"github.com/answerdev/answer/internal/controller/template_render"
	"github.com/answerdev/answer/internal/controller_backyard"
	"github.com/answerdev/answer/internal/repo"
	"github.com/answerdev/answer/internal/router"
	"github.com/answerdev/answer/internal/service"
	"github.com/answerdev/answer/internal/service/service_config"
	"github.com/google/wire"
	"github.com/segmentfault/pacman"
	"github.com/segmentfault/pacman/log"
)

// initApplication init application.
func initApplication(
	debug bool,
	serverConf *conf.Server,
	dbConf *data.Database,
	cacheConf *data.CacheConf,
	i18nConf *translator.I18n,
	swaggerConf *router.SwaggerConfig,
	serviceConf *service_config.ServiceConfig,
	logConf log.Logger) (*pacman.Application, func(), error) {
	panic(wire.Build(
		server.ProviderSetServer,
		router.ProviderSetRouter,
		controller.ProviderSetController,
		controller_backyard.ProviderSetController,
		templaterender.ProviderSetTemplateRenderController,
		service.ProviderSetService,
		cron.ProviderSetService,
		repo.ProviderSetRepo,
		translator.ProviderSet,
		middleware.ProviderSetMiddleware,
		newApplication,
	))
}
