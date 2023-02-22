package sample

import "github.com/answerdev/answer/plugin"

type Connector struct {
}

func init() {
	plugin.Register(&Connector{})
}

func (c *Connector) Info() plugin.Info {
	return plugin.Info{
		Name:     plugin.MakeTranslator("plugin.connector.name"),
		SlugName: "connector",
		//Description: plugin.MakeTranslator("plugin.connector.description"),
		Author:  "answerdev",
		Version: "0.0.1",
	}
}
