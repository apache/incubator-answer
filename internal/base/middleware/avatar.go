package middleware

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"path/filepath"
	"strings"

	"github.com/answerdev/answer/internal/service/service_config"
	"github.com/answerdev/answer/internal/service/uploader"
	"github.com/answerdev/answer/pkg/converter"
	"github.com/gin-gonic/gin"
)

type AvatarMiddleware struct {
	serviceConfig   *service_config.ServiceConfig
	uploaderService *uploader.UploaderService
}

// NewAvatarMiddleware new auth user middleware
func NewAvatarMiddleware(serviceConfig *service_config.ServiceConfig,
	uploaderService *uploader.UploaderService,
) *AvatarMiddleware {
	return &AvatarMiddleware{
		serviceConfig:   serviceConfig,
		uploaderService: uploaderService,
	}
}

func (am *AvatarMiddleware) AvatarThumb() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		u := ctx.Request.RequestURI
		if strings.Contains(u, "/uploads/avatar/") {
			wstr := ctx.Query("width")
			hstr := ctx.Query("height")
			w := converter.StringToInt(wstr)
			h := converter.StringToInt(hstr)
			uUrl, err := url.Parse(u)
			if err != nil {
				ctx.Next()
				return
			}
			_, urlfileName := filepath.Split(uUrl.Path)
			uploadPath := am.serviceConfig.UploadPath
			filePath := fmt.Sprintf("%s/avatar/%s", uploadPath, urlfileName)
			var avatarfile []byte
			if w == 0 && h == 0 {
				avatarfile, err = ioutil.ReadFile(filePath)
			} else {
				avatarfile, err = am.uploaderService.AvatarThumbFile(ctx, uploadPath, urlfileName, w, h)
			}
			if err != nil {
				ctx.Next()
				return
			}
			ctx.Writer.WriteString(string(avatarfile))
			ctx.Abort()
			return

		}
		ctx.Next()
	}
}
