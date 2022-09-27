package entity

// UserCacheInfo 用户缓存信息
type UserCacheInfo struct {
	UserID      string `json:"user_id"`
	UserStatus  int    `json:"user_status"`
	EmailStatus int    `json:"email_status"`
}
