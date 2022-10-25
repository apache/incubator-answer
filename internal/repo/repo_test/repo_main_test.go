package repo_test

import (
	"database/sql"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/answerdev/answer/internal/base/data"
	"github.com/answerdev/answer/internal/migrations"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/segmentfault/pacman/log"
	"xorm.io/xorm/schemas"
)

var (
	mysqlDBSetting = TestDBSetting{
		Driver:       string(schemas.MYSQL),
		ImageName:    "mariadb",
		ImageVersion: "10.4.7",
		ENV:          []string{"MYSQL_ROOT_PASSWORD=root", "MYSQL_DATABASE=answer", "MYSQL_ROOT_HOST=%"},
		PortID:       "3306/tcp",
		Connection:   "root:root@(localhost:%s)/answer?parseTime=true", // port is not fixed, it will be got by port id
	}
	postgresDBSetting = TestDBSetting{
		Driver:       string(schemas.POSTGRES),
		ImageName:    "postgres",
		ImageVersion: "14",
		ENV:          []string{"POSTGRES_USER=root", "POSTGRES_PASSWORD=root", "POSTGRES_DB=answer", "LISTEN_ADDRESSES='*'"},
		PortID:       "5432/tcp",
		Connection:   "host=localhost port=%s user=root password=root dbname=answer sslmode=disable",
	}
	sqlite3DBSetting = TestDBSetting{
		Driver:     string(schemas.SQLITE),
		Connection: os.TempDir() + "answer-test-data.db",
	}
	dbSettingMapping = map[string]TestDBSetting{
		mysqlDBSetting.Driver:    mysqlDBSetting,
		sqlite3DBSetting.Driver:  sqlite3DBSetting,
		postgresDBSetting.Driver: postgresDBSetting,
	}
	// after all test down will execute tearDown function to clean-up
	tearDown func()
	// dataSource used for repo testing
	dataSource *data.Data
)

func TestMain(t *testing.M) {
	dbSetting, ok := dbSettingMapping[os.Getenv("TEST_DB_DRIVER")]
	if !ok {
		dbSetting = dbSettingMapping[string(schemas.MYSQL)]
	}
	defer func() {
		if tearDown != nil {
			tearDown()
		}
	}()
	if err := initTestDB(dbSetting); err != nil {
		panic(err)
	}
	log.Info("init test database successfully")

	if ret := t.Run(); ret != 0 {
		panic(ret)
	}
}

type TestDBSetting struct {
	Driver       string
	ImageName    string
	ImageVersion string
	ENV          []string
	PortID       string
	Connection   string
}

func initTestDB(dbSetting TestDBSetting) error {
	connection, imageCleanUp, err := initDatabaseImage(dbSetting)
	if err != nil {
		return err
	}
	dbSetting.Connection = connection
	ds, dbCleanUp, err := initDatabase(dbSetting)
	if err != nil {
		return err
	}
	dataSource = ds
	tearDown = func() {
		dbCleanUp()
		log.Info("cleanup test database successfully")
		imageCleanUp()
		log.Info("cleanup test database image successfully")
	}
	return nil
}

func initDatabaseImage(dbSetting TestDBSetting) (connection string, cleanup func(), err error) {
	// sqlite3 don't need to set up image
	if dbSetting.Driver == string(schemas.SQLITE) {
		return dbSetting.Connection, func() {
			log.Info("remove database", dbSetting.Connection)
			_ = os.Remove(dbSetting.Connection)
		}, nil
	}
	pool, err := dockertest.NewPool("")
	pool.MaxWait = time.Minute * 5
	if err != nil {
		return "", nil, fmt.Errorf("could not connect to docker: %s", err)
	}

	//resource, err := pool.Run(dbSetting.ImageName, dbSetting.ImageVersion, dbSetting.ENV)
	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: dbSetting.ImageName,
		Tag:        dbSetting.ImageVersion,
		Env:        dbSetting.ENV,
	}, func(config *docker.HostConfig) {
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})
	if err != nil {
		return "", nil, fmt.Errorf("could not pull resource: %s", err)
	}

	connection = fmt.Sprintf(dbSetting.Connection, resource.GetPort(dbSetting.PortID))
	if err := pool.Retry(func() error {
		db, err := sql.Open(dbSetting.Driver, connection)
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		return "", nil, fmt.Errorf("could not connect to database: %s", err)
	}
	return connection, func() { _ = pool.Purge(resource) }, nil
}

func initDatabase(dbSetting TestDBSetting) (dataSource *data.Data, cleanup func(), err error) {
	dataConf := &data.Database{Driver: dbSetting.Driver, Connection: dbSetting.Connection}
	db, err := data.NewDB(true, dataConf)
	if err != nil {
		return nil, nil, fmt.Errorf("connection to database failed: %s", err)
	}
	err = migrations.InitDB(dataConf)
	if err != nil {
		return nil, nil, fmt.Errorf("migrations init database failed: %s", err)
	}

	dataSource, dbCleanUp, err := data.NewData(db, nil)
	if err != nil {
		return nil, nil, fmt.Errorf("new data failed: %s", err)
	}
	return dataSource, dbCleanUp, nil
}
