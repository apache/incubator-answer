package encryption

import (
	"crypto/md5"
	"encoding/hex"
)

// MD5 return md5 hash
func MD5(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}
