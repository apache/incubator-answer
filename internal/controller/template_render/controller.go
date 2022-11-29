package templaterender

import "github.com/google/wire"

// ProviderSetTemplateRenderController is template render controller providers.
var ProviderSetTemplateRenderController = wire.NewSet(
	NewQuestionController,
)
