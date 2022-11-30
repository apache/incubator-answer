package templaterender

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestPaginator(t *testing.T) {
	list := Paginator(5, 20, 300)
	spew.Dump(list)
}
