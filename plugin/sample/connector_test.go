package sample

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/answerdev/answer/plugin"
	"github.com/gin-gonic/gin"
)

func TestConnector_Info(t *testing.T) {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Request, _ = http.NewRequest("", "", nil)
	c.Request.Header.Set("Accept-Language", "en_US")
	_ = plugin.CallBase(func(base plugin.Base) error {
		info := base.Info()
		fmt.Println(info.Name.Translate(c))
		fmt.Println(info.SlugName)
		fmt.Println(info.Description.Translate(c))
		fmt.Println(info.Author)
		fmt.Println(info.Version)
		return nil
	})
}
