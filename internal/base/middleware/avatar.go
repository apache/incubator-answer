package middleware

import (
	"fmt"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/answerdev/answer/internal/service/service_config"
	"github.com/answerdev/answer/internal/service/uploader"
	"github.com/answerdev/answer/pkg/converter"
	"github.com/gin-gonic/gin"
	"github.com/segmentfault/pacman/log"
)

type AvatarMiddleware struct {
	serviceConfig   *service_config.ServiceConfig
	uploaderService uploader.UploaderService
}

// NewAvatarMiddleware new auth user middleware
func NewAvatarMiddleware(serviceConfig *service_config.ServiceConfig,
	uploaderService uploader.UploaderService,
) *AvatarMiddleware {
	return &AvatarMiddleware{
		serviceConfig:   serviceConfig,
		uploaderService: uploaderService,
	}
}

func (am *AvatarMiddleware) AvatarThumb() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uri := ctx.Request.RequestURI
		if strings.Contains(uri, "/uploads/avatar/") {
			size := converter.StringToInt(ctx.Query("s"))
			uriWithoutQuery, _ := url.Parse(uri)
			filename := filepath.Base(uriWithoutQuery.Path)
			filePath := fmt.Sprintf("%s/avatar/%s", am.serviceConfig.UploadPath, filename)
			var err error
			if size != 0 {
				filePath, err = am.uploaderService.AvatarThumbFile(ctx, filename, size)
				if err != nil {
					log.Error(err)
					ctx.Abort()
				}
			}
			avatarFile, err := os.ReadFile(filePath)
			if err != nil {
				log.Error(err)
				ctx.Abort()
				return
			}
			ctx.Header("content-type", fmt.Sprintf("image/%s", strings.TrimLeft(path.Ext(filePath), ".")))
			_, err = ctx.Writer.Write(avatarFile)
			if err != nil {
				log.Error(err)
			}
			ctx.Abort()
			return

		} else {
			urlInfo, err := url.Parse(uri)
			if err != nil {
				ctx.Next()
				return
			}
			ext := strings.TrimPrefix(filepath.Ext(urlInfo.Path), ".")
			ctx.Header("content-type", fmt.Sprintf("image/%s", ext))
		}
		ctx.Next()
	}
}
