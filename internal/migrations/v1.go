package migrations

import (
	"xorm.io/xorm"
)

func addUserLanguage(x *xorm.Engine) error {
	type User struct {
		ID       string `xorm:"not null pk autoincr BIGINT(20) id"`
		Username string `xorm:"not null default '' VARCHAR(50) UNIQUE username"`
		Language string `xorm:"not null default '' VARCHAR(100) language"`
	}
	return x.Sync(new(User))
}
