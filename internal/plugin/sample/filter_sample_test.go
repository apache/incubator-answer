package sample

import (
	"fmt"
	"testing"

	"github.com/answerdev/answer/internal/plugin"
	"github.com/stretchr/testify/assert"
)

func TestFilterSample_FilterText(t *testing.T) {
	// try to call filter plugin for filter text that are not allowed
	err := plugin.CallFilter(func(fn plugin.Filter) error {
		return fn.FilterText("bloody and violent words")
	})
	assert.Error(t, err)
	fmt.Println(err)
}
