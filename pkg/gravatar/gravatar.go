package gravatar

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/url"
)

const (
	defaultURLPrefix = "https://www.gravatar.com/avatar/"
)

// GetAvatarURL get avatar url from gravatar by email
func GetAvatarURL(email string) string {
	h := md5.New()
	h.Write([]byte(email))
	return defaultURLPrefix + hex.EncodeToString(h.Sum(nil))
}

// Resize resize avatar by pixel
func Resize(originalAvatarURL string, sizePixel int) (resizedAvatarURL string) {
	if len(originalAvatarURL) == 0 {
		return
	}
	originalURL, err := url.Parse(originalAvatarURL)
	if err != nil {
		return originalAvatarURL
	}
	query := originalURL.Query()
	query.Set("p", fmt.Sprintf("%d", sizePixel))
	originalURL.RawQuery = query.Encode()
	return originalURL.String()
}
