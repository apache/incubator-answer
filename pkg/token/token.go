package token

import "github.com/google/uuid"

// GenerateToken generate token
func GenerateToken() string {
	uid, _ := uuid.NewUUID()
	return uid.String()
}
