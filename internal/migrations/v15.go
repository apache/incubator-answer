package migrations

import (
	"context"
	"xorm.io/xorm"
)

func addNoticeConfig(ctx context.Context, x *xorm.Engine) error {
	type User struct {
		NoticeConfig string `xorm:"not null TEXT notice_config"`
	}
	return x.Context(ctx).Sync(new(User))
}
