package db

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DB struct {
	db *gorm.DB
}

var migrateTypes = []interface{}{
	&Package{},
	&Command{},
	&QueueItem{},
}

func New(dbpath string) (*DB, error) {
	db, err := gorm.Open(sqlite.Open(dbpath), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	for _, m := range migrateTypes {
		db.AutoMigrate(m)
	}

	return &DB{db: db}, nil
}

type StringArray []string

func (sa StringArray) Scan(value interface{}) error {
	content, ok := value.([]byte)
	if !ok {
		return errors.New("not an array")
	}

	return json.Unmarshal(content, &sa)
}

func (sa StringArray) Value() (driver.Value, error) {
	return json.Marshal(sa)
}

type Package struct {
	gorm.Model

	UUID     string
	Enqueued bool
}

type Command struct {
	gorm.Model

	PackageID uint
	Package   Package
	UUID      string

	AddCommand

	Completed bool
}

type AddCommand struct {
	Resource   string
	Action     string
	Parameters StringArray `gorm:"type:text"`
}

func (ac *AddCommand) Validate() error {
	return nil
}

type QueueItem struct {
	gorm.Model
	UUID     string
	Resource string
}

func (db *DB) NewPackage() (string, error) {
	uuid := uuid.New().String()
	return uuid, db.db.Create(&Package{UUID: uuid}).Error
}

func (db *DB) AddToPackage(pkgUUID string, ac *AddCommand) (string, error) {
	if err := ac.Validate(); err != nil {
		return "", err
	}

	var p Package

	if err := db.db.Where("uuid = ?", pkgUUID).First(&p).Error; err != nil {
		return "", err
	}

	uuid := uuid.New().String()

	command := &Command{
		Package:    p,
		UUID:       uuid,
		AddCommand: *ac,
	}

	return uuid, db.db.Create(command).Error
}

func (db *DB) EnqueuePackage(pkgUUID string) ([]string, error) {
	tx := db.db.Begin()
	defer tx.Rollback()

	pkg := &Package{}

	if err := tx.Where("uuid = ?", pkgUUID).First(pkg).Error; err != nil {
		return nil, err
	}

	if pkg.Enqueued {
		return nil, errors.New("already queued")
	}

	commands := []Command{}
	if err := tx.Where("package_id = ?", pkg.ID).Order("created_at").Find(&commands).Error; err != nil {
		return nil, err
	}

	uuids := []string{}

	for _, command := range commands {
		uuids = append(uuids, command.UUID)

		qi := &QueueItem{
			UUID:     command.UUID,
			Resource: command.Resource,
		}

		if err := tx.Create(qi).Error; err != nil {
			return nil, err
		}
	}

	pkg.Enqueued = true
	if err := tx.Save(pkg).Error; err != nil {
		return nil, err
	}

	return uuids, tx.Commit().Error
}
