package router

import "github.com/google/wire"

// ProviderSetRouter is providers.
var ProviderSetRouter = wire.NewSet(
	NewAnswerAPIRouter,
	NewSwaggerRouter,
	NewStaticRouter,
	NewUIRouter,
	NewTemplateRouter,
	NewPluginAPIRouter,
)
