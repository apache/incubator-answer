package cli

import (
	"bytes"

	"github.com/segmentfault/answer/assets"
	"github.com/segmentfault/answer/internal/base/data"
	"github.com/segmentfault/answer/internal/entity"
)

// InitDB init db
func InitDB(dataConf *data.Database) (err error) {
	db := data.NewDB(false, dataConf)
	// check db connection
	err = db.Ping()
	if err != nil {
		return err
	}

	exist, err := db.IsTableExist(&entity.User{})
	if err != nil {
		return err
	}
	if exist {
		return nil
	}

	// create table if not exist
	s := &bytes.Buffer{}
	s.Write(assets.AnswerSql)
	_, err = db.Import(s)
	if err != nil {
		return err
	}
	return nil
}
