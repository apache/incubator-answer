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
			sizeStr := ctx.Query("s")
			size := converter.StringToInt(sizeStr)
			uUrl, err := url.Parse(u)
			if err != nil {
				ctx.Next()
				return
			}
			_, urlfileName := filepath.Split(uUrl.Path)
			uploadPath := am.serviceConfig.UploadPath
			filePath := fmt.Sprintf("%s/avatar/%s", uploadPath, urlfileName)
			var avatarfile []byte
			if size == 0 {
				avatarfile, err = os.ReadFile(filePath)
			} else {
				avatarfile, err = am.uploaderService.AvatarThumbFile(ctx, uploadPath, urlfileName, size)
			}
			if err != nil {
				ctx.Next()
				return
			}
			ext := strings.ToLower(path.Ext(filePath)[1:])
			ctx.Header("content-type", fmt.Sprintf("image/%s", ext))
			_, err = ctx.Writer.WriteString(string(avatarfile))
			if err != nil {
				log.Error(err)
			}
			ctx.Abort()
			return

		} else {
			uUrl, err := url.Parse(u)
			if err != nil {
				ctx.Next()
				return
			}
			_, urlfileName := filepath.Split(uUrl.Path)
			uploadPath := am.serviceConfig.UploadPath
			filePath := fmt.Sprintf("%s/%s", uploadPath, urlfileName)
			ext := strings.ToLower(path.Ext(filePath)[1:])
			ctx.Header("content-type", fmt.Sprintf("image/%s", ext))
		}
		ctx.Next()
	}
}
