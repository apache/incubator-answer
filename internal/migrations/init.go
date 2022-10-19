package migrations

import (
	"fmt"

	"github.com/segmentfault/answer/internal/entity"
	"xorm.io/xorm"
)

var (
	tables = []interface{}{
		&entity.Activity{},
		&entity.Answer{},
		&entity.Collection{},
		&entity.CollectionGroup{},
		&entity.Comment{},
		&entity.Config{},
		&entity.Meta{},
		&entity.Notification{},
		&entity.Question{},
		&entity.Report{},
		&entity.Revision{},
		&entity.SiteInfo{},
		&entity.Tag{},
		&entity.TagRel{},
		&entity.Uniqid{},
		&entity.User{},
	}
)

// InitDB init db
func InitDB(engine *xorm.Engine) (err error) {
	exist, err := engine.IsTableExist(&Version{})
	if err != nil {
		return err
	}
	if exist {
		fmt.Println("[database] already exists")
		return nil
	}

	err = engine.Sync(tables)
	if err != nil {
		return err
	}
	return nil
}
