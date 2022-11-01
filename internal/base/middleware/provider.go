package middleware

import (
	"github.com/google/wire"
)

// ProviderSetMiddleware is providers.
var ProviderSetMiddleware = wire.NewSet(
	NewAuthUserMiddleware,
	NewAvatarMiddleware,
)
