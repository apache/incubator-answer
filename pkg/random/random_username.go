package random

import (
	"encoding/hex"
	"math/rand"
)

func UsernameSuffix() string {
	bytes := make([]byte, 2)
	_, _ = rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

func Username() string {
	bytes := make([]byte, 6)
	_, _ = rand.Read(bytes)
	return hex.EncodeToString(bytes)
}
