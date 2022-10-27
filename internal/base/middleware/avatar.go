package middleware

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/answerdev/answer/internal/service/service_config"
	"github.com/gin-gonic/gin"
)

type AvatarMiddleware struct {
	serviceConfig *service_config.ServiceConfig
}

// NewAvatarMiddleware new auth user middleware
func NewAvatarMiddleware(serviceConfig *service_config.ServiceConfig) *AvatarMiddleware {
	return &AvatarMiddleware{
		serviceConfig: serviceConfig,
	}
}

func (am *AvatarMiddleware) AvatarThumb() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		url := ctx.Request.RequestURI
		if strings.Contains(url, "/uploads/avatar/") {
			_, fileName := filepath.Split(url)
			filepath := fmt.Sprintf("%s/avatar/%s", am.serviceConfig.UploadPath, fileName)
			f, err := ioutil.ReadFile(filepath)
			if err != nil {
				ctx.Next()
				return
			}
			ctx.Writer.WriteString(string(f))
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
