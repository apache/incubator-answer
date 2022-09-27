package controller_backyard

import "github.com/google/wire"

// ProviderSetController is controller providers.
var ProviderSetController = wire.NewSet(
	NewReportController,
	NewUserBackyardController,
	NewThemeController,
	NewSiteInfoController,
)
