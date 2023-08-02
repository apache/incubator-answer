package migrations

import (
	"context"
	"xorm.io/xorm"
)

func addUserLanguage(ctx context.Context, x *xorm.Engine) error {
	type User struct {
		ID       string `xorm:"not null pk autoincr BIGINT(20) id"`
		Username string `xorm:"not null default '' VARCHAR(50) UNIQUE username"`
		Language string `xorm:"not null default '' VARCHAR(100) language"`
	}
	return x.Context(ctx).Sync(new(User))
}
