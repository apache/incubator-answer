package demo

import "github.com/answerdev/answer/plugin"

type DemoPlugin struct {
}

func init() {
	plugin.Register(&DemoPlugin{})
}

func (d DemoPlugin) Info() plugin.Info {
	return plugin.Info{
		Name:        plugin.MakeTranslator("i18n.demo.name"),
		SlugName:    "demo_plugin",
		Description: plugin.MakeTranslator("i18n.demo.description"),
		Author:      "answerdev",
		Version:     "0.0.1",
	}
}
