package plugin

import "github.com/gin-gonic/gin"

// GinContext is a wrapper of gin.Context
// We export it to make it easy to use in plugins
type GinContext = gin.Context

// Register registers a plugin
func Register(p Base) {
	registerBase(p)

	switch pType := p.(type) {
	case Connector:
		registerConnector(pType)
	case Parser:
		registerParser(pType)
	case Filter:
		registerFilter(pType)
	}
}

// Dump returns all registered plugins infos
func Dump() []Info {
	var infos []Info

	CallBase(func(p Base) error {
		infos = append(infos, p.Info())
		return nil
	})

	return infos
}

type Stack[T Base] struct {
	plugins []T
}

type RegisterFn[T Base] func(p T)
type Caller[T Base] func(p T) error
type CallFn[T Base] func(fn Caller[T]) error

// MakePlugin creates a plugin caller and register stack manager
// It returns a register function and a caller function
// The register function is used to register a plugin, it will be called in the plugin's init function
// The caller function is used to call all registered plugins
func MakePlugin[T Base]() (CallFn[T], RegisterFn[T]) {
	stack := Stack[T]{}

	call := func(fn Caller[T]) error {
		for _, p := range stack.plugins {
			// If the plugin is disabled, skip it
			if p.Info().Disabled {
				continue
			}

			if err := fn(p); err != nil {
				return err
			}
		}
		return nil
	}

	register := func(p T) {
		stack.plugins = append(stack.plugins, p)
	}

	return call, register
}
