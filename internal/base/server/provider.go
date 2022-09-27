package server

import "github.com/google/wire"

// ProviderSetServer is providers.
var ProviderSetServer = wire.NewSet(NewHTTPServer)
