package uploader

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"path"
	"path/filepath"

	"github.com/answerdev/answer/internal/base/reason"
	"github.com/answerdev/answer/internal/service/service_config"
	"github.com/answerdev/answer/pkg/dir"
	"github.com/answerdev/answer/pkg/uid"
	"github.com/davecgh/go-spew/spew"
	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
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

func (us *UploaderService) AvatarThumbFile(ctx *gin.Context, uploadPath, fileName string, w, h int) (
	avatarfile []byte, err error) {

	filePath := fmt.Sprintf("%s/avatar/%s", uploadPath, fileName)
	avatarfile, err = ioutil.ReadFile(filePath)
	spew.Dump("ioutil.ReadFile", err)
	if err != nil {
		return avatarfile, err
	}
	reader := bytes.NewReader(avatarfile)
	img, err := imaging.Decode(reader)
	spew.Dump("imaging.Decode", err)
	if err != nil {
		return avatarfile, err
	}
	new_image := imaging.Resize(img, w, h, imaging.Linear)
	var buf bytes.Buffer
	err = imaging.Encode(&buf, new_image, imaging.JPEG)
	spew.Dump(buf)
	if err != nil {
		return avatarfile, err
	}
	return buf.Bytes(), nil

	// formatImg := imaging.Resize(img, 1024, 0, imaging.Linear)
	// var buf bytes.Buffer
	// err = imaging.Encode(&buf, formatImg, imaging.JPEG)
	// if err != nil {
	// 	return "", err
	// }
	// reader := bytes.NewReader(buf.Bytes())
	// newFilename := fmt.Sprintf("%s%s", uid.IDStr12(), fileExt)
	// avatarFilePath := path.Join(avatarSubPath, newFilename)
	// filePath := path.Join(us.serviceConfig.UploadPath, avatarFilePath)
	// out, err := os.Create(filePath)
	// if err != nil {
	// 	return "", err
	// }
	// defer out.Close()
	// _, err = io.Copy(out, reader)
	// if err != nil {
	// 	return "", err
	// }
	// url = fmt.Sprintf("%s/uploads/%s", us.serviceConfig.WebHost, avatarFilePath)
	// return url, nil
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
