package sample

import (
	"fmt"
	"strings"

	"github.com/answerdev/answer/internal/plugin"
)

type FilterSample struct {
}

func (s *FilterSample) Info() plugin.Info {
	return plugin.Info{
		Name:        "filter sample",
		Description: "filter sample plugin",
		Version:     "0.0.1",
	}
}

func (s *FilterSample) FilterText(data string) (err error) {
	if strings.Contains(data, "violent") {
		return fmt.Errorf("bloody and violent words cannot appear in this website")
	}
	return nil
}

func init() {
	plugin.Register(&FilterSample{})
}
