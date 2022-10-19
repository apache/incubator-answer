package migrations

import (
	"fmt"

	"github.com/segmentfault/pacman/log"
	"xorm.io/xorm"
)

const minDBVersion = 0 // answer 1.0.0

// Migration describes on migration from lower version to high version
type Migration interface {
	Description() string
	Migrate(*xorm.Engine) error
}

type migration struct {
	description string
	migrate     func(*xorm.Engine) error
}

// Description returns the migration's description
func (m *migration) Description() string {
	return m.description
}

// Migrate executes the migration
func (m *migration) Migrate(x *xorm.Engine) error {
	return m.migrate(x)
}

// NewMigration creates a new migration
func NewMigration(desc string, fn func(*xorm.Engine) error) Migration {
	return &migration{description: desc, migrate: fn}
}

// Version version
type Version struct {
	ID            int   `xorm:"not null pk autoincr comment('id') INT(11) id"`
	VersionNumber int64 `xorm:"not null default 0 comment('version_number') INT(11) version_number"`
}

// TableName config table name
func (Version) TableName() string {
	return "version"
}

// Use noopMigration when there is a migration that has been no-oped
var noopMigration = func(_ *xorm.Engine) error { return nil }

var migrations = []Migration{
	// 0->1
	NewMigration("this is first version, no operation", noopMigration),
}

// GetCurrentDBVersion returns the current db version
func GetCurrentDBVersion(engine *xorm.Engine) (int64, error) {
	if err := engine.Sync(new(Version)); err != nil {
		return -1, fmt.Errorf("sync version failed: %v", err)
	}

	currentVersion := &Version{ID: 1}
	has, err := engine.Get(currentVersion)
	if err != nil {
		return -1, fmt.Errorf("get first version failed: %v", err)
	}
	if !has {
		_, err := engine.InsertOne(&Version{ID: 1, VersionNumber: 0})
		if err != nil {
			return -1, fmt.Errorf("insert first version failed: %v", err)
		}
		return 0, nil
	}
	return currentVersion.VersionNumber, nil
}

// ExpectedVersion returns the expected db version
func ExpectedVersion() int64 {
	return int64(minDBVersion + len(migrations))
}

// Migrate database to current version
func Migrate(engine *xorm.Engine) error {
	currentDBVersion, err := GetCurrentDBVersion(engine)
	if err != nil {
		return err
	}
	expectedVersion := ExpectedVersion()

	for currentDBVersion < expectedVersion {
		log.Infof("[migrate] current db version is %d, try to migrate version %d, latest version is %d",
			currentDBVersion, currentDBVersion+1, expectedVersion)
		migrationFunc := migrations[currentDBVersion]
		log.Infof("[migrate] try to migrate db version %d, description: %s", currentDBVersion+1, migrationFunc.Description())
		if err := migrationFunc.Migrate(engine); err != nil {
			log.Errorf("[migrate] migrate to db version %d failed: ", currentDBVersion+1, err.Error())
			return err
		}
		log.Infof("[migrate] migrate to db version %d success", currentDBVersion+1)
		if _, err := engine.Update(&Version{ID: 1, VersionNumber: currentDBVersion + 1}); err != nil {
			log.Errorf("[migrate] migrate to db version %d, update failed: %s", currentDBVersion+1, err.Error())
			return err
		}
		currentDBVersion++
	}
	return nil
}
