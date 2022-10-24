package uploader

import (
	"fmt"
	"mime/multipart"
	"path"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/segmentfault/answer/internal/base/reason"
	"github.com/segmentfault/answer/internal/service/service_config"
	"github.com/segmentfault/answer/pkg/dir"
	"github.com/segmentfault/answer/pkg/uid"
	"github.com/segmentfault/pacman/errors"
)

const (
	avatarSubPath = "avatar"
	postSubPath   = "post"
)

// UploaderService user service
type UploaderService struct {
	serviceConfig *service_config.ServiceConfig
}

// NewUploaderService new upload service
func NewUploaderService(serviceConfig *service_config.ServiceConfig) *UploaderService {
	err := dir.CreateDirIfNotExist(filepath.Join(serviceConfig.UploadPath, avatarSubPath))
	if err != nil {
		panic(err)
	}
	err = dir.CreateDirIfNotExist(filepath.Join(serviceConfig.UploadPath, postSubPath))
	if err != nil {
		panic(err)
	}
	return &UploaderService{
		serviceConfig: serviceConfig,
	}
}

func (us *UploaderService) UploadAvatarFile(ctx *gin.Context, file *multipart.FileHeader, fileExt string) (
	url string, err error) {
	newFilename := fmt.Sprintf("%s%s", uid.IDStr12(), fileExt)
	avatarFilePath := path.Join(avatarSubPath, newFilename)
	return us.uploadFile(ctx, file, avatarFilePath)
}

func (us *UploaderService) UploadPostFile(ctx *gin.Context, file *multipart.FileHeader, fileExt string) (
	url string, err error) {
	newFilename := fmt.Sprintf("%s%s", uid.IDStr12(), fileExt)
	avatarFilePath := path.Join(postSubPath, newFilename)
	return us.uploadFile(ctx, file, avatarFilePath)
}

func (us *UploaderService) uploadFile(ctx *gin.Context, file *multipart.FileHeader, fileSubPath string) (
	url string, err error) {
	filePath := path.Join(us.serviceConfig.UploadPath, fileSubPath)
	if err := ctx.SaveUploadedFile(file, filePath); err != nil {
		return "", errors.InternalServer(reason.UnknownError).WithError(err).WithStack()
	}
	url = fmt.Sprintf("%s/uploads/%s", us.serviceConfig.WebHost, fileSubPath)
	return url, nil
}
