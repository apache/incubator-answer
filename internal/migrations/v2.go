package migrations

import (
	"xorm.io/xorm"
)

func addTagRecommendedAndReserved(x *xorm.Engine) error {
	type Tag struct {
		Recommend bool `xorm:"not null default false BOOL recommend"`
		Reserved  bool `xorm:"not null default false BOOL reserved"`
	}
	return x.Sync(new(Tag))
}
