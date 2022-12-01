package entity

// UserCacheInfo User Cache Information
type UserCacheInfo struct {
	UserID      string `json:"user_id"`
	UserStatus  int    `json:"user_status"`
	EmailStatus int    `json:"email_status"`
	IsAdmin     bool   `json:"is_admin"`
}
