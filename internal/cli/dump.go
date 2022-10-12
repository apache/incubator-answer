package cli

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/segmentfault/answer/internal/base/data"
	"xorm.io/xorm/schemas"
)

// DumpAllData dump all database data to sql
func DumpAllData(dataConf *data.Database, dumpDataPath string) error {
	db, err := data.NewDB(false, dataConf)
	if err != nil {
		return err
	}
	if err = db.Ping(); err != nil {
		return err
	}

	name := filepath.Join(dumpDataPath, fmt.Sprintf("answer_dump_data_%s.sql", time.Now().Format("2006-01-02")))
	return db.DumpAllToFile(name, schemas.MYSQL)
}
