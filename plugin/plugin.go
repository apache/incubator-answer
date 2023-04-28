package plugin

import (
	"encoding/json"

	"github.com/answerdev/answer/internal/base/handler"
	"github.com/answerdev/answer/internal/base/translator"
	"github.com/gin-gonic/gin"
)

// GinContext is a wrapper of gin.Context
// We export it to make it easy to use in plugins
type GinContext = gin.Context

// StatusManager is a manager that manages the status of plugins
// Init Plugins:
// json.Unmarshal([]byte(`{"plugin1": true, "plugin2": false}`), &plugin.StatusManager)
// Dump Status:
// json.Marshal(plugin.StatusManager)
var StatusManager = statusManager{
	status: make(map[string]bool),
}

// Register registers a plugin
func Register(p Base) {
	registerBase(p)

	if _, ok := p.(Config); ok {
		registerConfig(p.(Config))
	}

	if _, ok := p.(Connector); ok {
		registerConnector(p.(Connector))
	}

	if _, ok := p.(Parser); ok {
		registerParser(p.(Parser))
	}

	if _, ok := p.(Filter); ok {
		registerFilter(p.(Filter))
	}

	if _, ok := p.(Storage); ok {
		registerStorage(p.(Storage))
	}

	if _, ok := p.(Cache); ok {
		registerCache(p.(Cache))
	}

	if _, ok := p.(UserCenter); ok {
		registerUserCenter(p.(UserCenter))
	}

	if _, ok := p.(Agent); ok {
		registerAgent(p.(Agent))
	}
}

type Stack[T Base] struct {
	plugins []T
}

type RegisterFn[T Base] func(p T)
type Caller[T Base] func(p T) error
type CallFn[T Base] func(fn Caller[T]) error

// MakePlugin creates a plugin caller and register stack manager
// The parameter super presents if the plugin can be disabled.
// It returns a register function and a caller function
// The register function is used to register a plugin, it will be called in the plugin's init function
// The caller function is used to call all registered plugins
func MakePlugin[T Base](super bool) (CallFn[T], RegisterFn[T]) {
	stack := Stack[T]{}

	call := func(fn Caller[T]) error {
		for _, p := range stack.plugins {
			// If the plugin is disabled, skip it
			if !super && !StatusManager.IsEnabled(p.Info().SlugName) {
				continue
			}

			if err := fn(p); err != nil {
				return err
			}
		}
		return nil
	}

	register := func(p T) {
		for _, plugin := range stack.plugins {
			if plugin.Info().SlugName == p.Info().SlugName {
				panic("plugin " + p.Info().SlugName + " is already registered")
			}
		}
		stack.plugins = append(stack.plugins, p)
	}

	return call, register
}

type statusManager struct {
	status map[string]bool
}

func (m *statusManager) Enable(name string, enabled bool) {
	m.status[name] = enabled
}

func (m *statusManager) IsEnabled(name string) bool {
	if status, ok := m.status[name]; ok {
		return status
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface.
func (m *statusManager) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.status)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (m *statusManager) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &m.status)
}

// Translate translates the key to the current language of the context
func Translate(ctx *GinContext, key string) string {
	return translator.Tr(handler.GetLang(ctx), key)
}

// TranslateFn presents a generator of translated string.
// We use it to delegate the translation work outside the plugin.
type TranslateFn func(ctx *GinContext) string

// Translator contains a function that translates the key to the current language of the context
type Translator struct {
	Fn TranslateFn
}

// MakeTranslator generates a translator from the key
func MakeTranslator(key string) Translator {
	t := func(ctx *GinContext) string {
		return Translate(ctx, key)
	}
	return Translator{Fn: t}
}

// Translate translates the key to the current language of the context
func (t Translator) Translate(ctx *GinContext) string {
	if &t == nil || t.Fn == nil {
		return ""
	}
	return t.Fn(ctx)
}
