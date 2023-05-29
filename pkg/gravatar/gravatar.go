package gravatar

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/url"
)

// GetAvatarURL get avatar url from gravatar by email
func GetAvatarURL(baseURL, email string) string {
	h := md5.New()
	h.Write([]byte(email))
	return baseURL + hex.EncodeToString(h.Sum(nil))
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
	query.Set("s", fmt.Sprintf("%d", sizePixel))
	originalURL.RawQuery = query.Encode()
	return originalURL.String()
}
