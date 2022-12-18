package cron

import (
	"github.com/google/wire"
)

// ProviderSetService is providers.
var ProviderSetService = wire.NewSet(
	NewScheduledTaskManager,
)
