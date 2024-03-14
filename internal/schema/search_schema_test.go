package schema

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReplaceSearchContent(t *testing.T) {
	content := "user:aaa [tag] ssssfdfdf-as#fsadf"
	replacedContent, patterns := ReplaceSearchContent(content)
	ret := strings.Join(append(patterns, replacedContent), " ")

	assert.Equal(t, "user:aaa [tag] ssssfdfdf as fsadf", ret)

	content = "user:aaa-sss [tag1] ssssfdfdf-as#fsadf [tag2] score:3"
	replacedContent, patterns = ReplaceSearchContent(content)
	ret = strings.Join(append(patterns, replacedContent), " ")

	assert.Equal(t, "user:aaa-sss score:3 [tag1] [tag2] ssssfdfdf as fsadf", ret)
}
