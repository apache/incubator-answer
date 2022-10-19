package migrations

import (
	"fmt"

	"github.com/segmentfault/answer/internal/base/data"
	"github.com/segmentfault/answer/internal/entity"
)

var (
	tables = []interface{}{
		&entity.Activity{},        // done
		&entity.Answer{},          // index
		&entity.Collection{},      // done
		&entity.CollectionGroup{}, // updated_at
		&entity.Comment{},         // done
		&entity.Config{},          // done
		&entity.Meta{},            // done
		&entity.Notification{},    // done
		&entity.Question{},        // done
		&entity.Report{},          // reported_user_id
		&entity.Revision{},        // done
		&entity.SiteInfo{},        // all
		&entity.Tag{},             // done
		&entity.TagRel{},          // index
		&entity.Uniqid{},          // done
		&entity.User{},            // index
		&Version{},
	}
)

// InitDB init db
func InitDB(dataConf *data.Database) (err error) {
	engine, err := data.NewDB(false, dataConf)
	if err != nil {
		fmt.Println("new database failed: ", err.Error())
		return err
	}

	exist, err := engine.IsTableExist(&Version{})
	if err != nil {
		return fmt.Errorf("check table exists failed: %s", err)
	}
	if exist {
		fmt.Println("[database] already exists")
		return nil
	}

	err = engine.Sync(tables...)
	if err != nil {
		return fmt.Errorf("sync table failed: %s", err)
	}
	return nil
}
