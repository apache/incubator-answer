package checker

import (
	"golang.org/x/image/webp"
	"image"
	_ "image/gif" // use init to support decode jpeg,jpg,png,gif
	_ "image/jpeg"
	_ "image/png"
	"io"
	"strings"
)

// IsSupportedImageFile currently answers support image type is
// `image/jpeg, image/jpg, image/png, image/gif, image/webp`
func IsSupportedImageFile(file io.Reader, ext string) bool {
	ext = strings.ToLower(strings.TrimPrefix(ext, "."))
	var err error
	switch ext {
	case "jpg", "jpeg", "png", "gif": // only allow for `image/jpeg,image/jpg,image/png, image/gif`
		_, _, err = image.Decode(file)
	case "ico":
		// TODO: There is currently no good Golang library to parse whether the image is in ico format.
		return true
	case "webp":
		_, err = webp.Decode(file)
	default:
		return false
	}
	return err == nil
}
