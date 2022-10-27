package middleware

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/answerdev/answer/internal/service/service_config"
	"github.com/answerdev/answer/internal/service/uploader"
	"github.com/answerdev/answer/pkg/converter"
	"github.com/davecgh/go-spew/spew"
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
		//?width=100&height=100
		wstr := ctx.Query("width")
		hstr := ctx.Query("height")
		w := converter.StringToInt(wstr)
		h := converter.StringToInt(hstr)
		url := ctx.Request.RequestURI
		if strings.Contains(url, "/uploads/avatar/") {
			_, fileName := filepath.Split(url)
			uploadPath := am.serviceConfig.UploadPath
			filePath := fmt.Sprintf("%s/avatar/%s", uploadPath, fileName)
			if w == 0 && h == 0 {
				avatarfile, err := ioutil.ReadFile(filePath)
				if err != nil {
					ctx.Next()
					return
				}
				ctx.Writer.WriteString(string(avatarfile))
				ctx.Abort()
				return
			} else {
				spew.Dump(w, h, fileName)
				avatarfile, err := am.uploaderService.AvatarThumbFile(ctx, uploadPath, fileName, w, h)
				if err != nil {
					ctx.Next()
					return
				}
				ctx.Writer.WriteString(string(avatarfile))
				ctx.Abort()
				return
			}

		}
		ctx.Next()
	}
}
