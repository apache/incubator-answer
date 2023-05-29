package constant

import "time"

const (
	UserStatusChangedCacheKey          = "answer:user:status:"
	UserStatusChangedCacheTime         = 7 * 24 * time.Hour
	UserTokenCacheKey                  = "answer:user:token:"
	UserTokenCacheTime                 = 7 * 24 * time.Hour
	AdminTokenCacheKey                 = "answer:admin:token:"
	AdminTokenCacheTime                = 7 * 24 * time.Hour
	UserTokenMappingCacheKey           = "answer:user-token:mapping:"
	SiteInfoCacheKey                   = "answer:site-info:"
	SiteInfoCacheTime                  = 1 * time.Hour
	ConfigID2KEYCacheKeyPrefix         = "answer:config:id:"
	ConfigKEY2ContentCacheKeyPrefix    = "answer:config:key:"
	ConnectorUserExternalInfoCacheKey  = "answer:connector:"
	ConnectorUserExternalInfoCacheTime = 10 * time.Minute
)
