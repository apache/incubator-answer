package sample

import "github.com/answerdev/answer/internal/plugin"

type Sample struct {
}

func (s *Sample) Info() plugin.Info {
	return plugin.Info{
		Name:        "sample",
		Description: "sample plugin",
		Version:     "0.0.1",
	}
}

func init() {
	plugin.Register(&Sample{})
}
