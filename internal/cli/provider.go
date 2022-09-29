package cli

import (
	"bytes"

	"github.com/google/wire"
	"github.com/segmentfault/answer/assets"
	"github.com/segmentfault/answer/internal/base/data"
	"github.com/segmentfault/answer/internal/entity"
)

// ProviderSetCli is providers.
var ProviderSetCli = wire.NewSet(NewCli)

type Cli struct {
	DataSource *data.Data
}

var CommandCli *Cli

func NewCli(dataSource *data.Data) *Cli {
	CommandCli = &Cli{DataSource: dataSource}
	return CommandCli
}

// InitDB init db
func (c *Cli) InitDB() (err error) {
	// check db connection
	err = c.DataSource.DB.Ping()
	if err != nil {
		return err
	}

	exist, err := c.DataSource.DB.IsTableExist(&entity.User{})
	if err != nil {
		return err
	}
	if exist {
		return nil
	}

	// create table if not exist
	s := &bytes.Buffer{}
	s.Write(assets.AnswerSql)
	_, err = c.DataSource.DB.Import(s)
	if err != nil {
		return err
	}
	return nil
}
