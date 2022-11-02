package migrations

import (
	"xorm.io/xorm"
)

func addUserLanguage(x *xorm.Engine) error {
	type User struct {
		Language string `xorm:"not null default '' VARCHAR(100) language"`
	}
	return x.Sync(new(User))
}
