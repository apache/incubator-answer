package checker

import (
	"image/jpeg"
	"image/png"
	"io"
	"strings"
)

// IsSupportedImageFile currently answers support image type is `image/jpeg,image/jpg,image/png`
func IsSupportedImageFile(file io.Reader, ext string) bool {
	ext = strings.TrimPrefix(ext, ".")
	var err error
	switch strings.ToUpper(ext) {
	case "JPEG":
		_, err = jpeg.Decode(file)
	case "PNG":
		_, err = png.Decode(file)
	case "ICO":
		// TODO: There is currently no good Golang library to parse whether the image is in ico format.
		return true
	case "JPG":
		return true
	default:
		return false
	}
	return err == nil
}
