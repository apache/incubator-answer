package controller_admin

import "github.com/google/wire"

// ProviderSetController is controller providers.
var ProviderSetController = wire.NewSet(
	NewReportController,
	NewUserAdminController,
	NewThemeController,
	NewSiteInfoController,
	NewRoleController,
	NewPluginController,
)
