package uploader

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
	"path/filepath"

	"github.com/answerdev/answer/internal/base/reason"
	"github.com/answerdev/answer/internal/service/service_config"
	"github.com/answerdev/answer/pkg/dir"
	"github.com/answerdev/answer/pkg/uid"
	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
	"github.com/segmentfault/pacman/errors"
)

const (
	avatarSubPath      = "avatar"
	avatarThumbSubPath = "avatar_thumb"
	postSubPath        = "post"
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

var FormatExts = map[string]imaging.Format{
	".jpg":  imaging.JPEG,
	".jpeg": imaging.JPEG,
	".png":  imaging.PNG,
	".gif":  imaging.GIF,
	".tif":  imaging.TIFF,
	".tiff": imaging.TIFF,
	".bmp":  imaging.BMP,
}

func (us *UploaderService) AvatarThumbFile(ctx *gin.Context, uploadPath, fileName string, w, h int) (
	avatarfile []byte, err error) {
	thumbFileName := fmt.Sprintf("%d_%d@%s", w, h, fileName)
	thumbfilePath := fmt.Sprintf("%s/%s/%s", uploadPath, avatarThumbSubPath, thumbFileName)
	avatarfile, err = ioutil.ReadFile(thumbfilePath)
	if err == nil {
		return avatarfile, nil
	}
	filePath := fmt.Sprintf("%s/avatar/%s", uploadPath, fileName)
	avatarfile, err = ioutil.ReadFile(filePath)
	if err != nil {
		return avatarfile, errors.InternalServer(reason.UnknownError).WithError(err).WithStack()
	}
	reader := bytes.NewReader(avatarfile)
	img, err := imaging.Decode(reader)
	if err != nil {
		return avatarfile, errors.InternalServer(reason.UnknownError).WithError(err).WithStack()
	}
	new_image := imaging.Fill(img, w, h, imaging.Center, imaging.Linear)
	var buf bytes.Buffer
	fileSuffix := path.Ext(fileName)

	_, ok := FormatExts[fileSuffix]

	if !ok {
		return avatarfile, fmt.Errorf("img extension not exist")
	}
	err = imaging.Encode(&buf, new_image, formatExts[fileSuffix])

	if err != nil {
		return avatarfile, errors.InternalServer(reason.UnknownError).WithError(err).WithStack()
	}
	thumbReader := bytes.NewReader(buf.Bytes())
	dir.CreateDirIfNotExist(path.Join(us.serviceConfig.UploadPath, avatarThumbSubPath))
	avatarFilePath := path.Join(avatarThumbSubPath, thumbFileName)
	savefilePath := path.Join(us.serviceConfig.UploadPath, avatarFilePath)
	out, err := os.Create(savefilePath)
	if err != nil {
		return avatarfile, errors.InternalServer(reason.UnknownError).WithError(err).WithStack()
	}
	defer out.Close()
	_, err = io.Copy(out, thumbReader)
	if err != nil {
		return avatarfile, errors.InternalServer(reason.UnknownError).WithError(err).WithStack()
	}
	return buf.Bytes(), nil
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
