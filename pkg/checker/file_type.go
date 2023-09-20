package checker

import (
	"image"
	_ "image/gif" // use init to support decode jpeg,jpg,png,gif
	_ "image/jpeg"
	_ "image/png"
	"io"
	"strings"
)

// IsSupportedImageFile currently answers support image type is `image/jpeg,image/jpg,image/png, image/gif`
func IsSupportedImageFile(file io.Reader, ext string) bool {
	ext = strings.TrimPrefix(ext, ".")
	var err error
	switch strings.ToUpper(ext) {
	case "JPG", "JPEG", "PNG", "GIF": // only allow for `image/jpeg,image/jpg,image/png, image/gif`
		_, _, err = image.Decode(file)
	case "ICO":
		// TODO: There is currently no good Golang library to parse whether the image is in ico format.
		return true
	default:
		return false
	}
	return err == nil
}
