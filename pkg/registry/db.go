package registry

import (
	"errors"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DBConfig struct {
	Filename string
}

type DB struct {
	db *gorm.DB
}

var migrationTables = []interface{}{}

func NewDB(c DBConfig) (*DB, error) {
	db, err := gorm.Open(sqlite.Open(c.Filename), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	for _, table := range migrationTables {
		if err := db.AutoMigrate(table); err != nil {
			return nil, err
		}
	}

	return &DB{db: db}, nil
}

func (db *DB) Get(id uint64) (*VM, error) {
	return nil, errors.New("unimplemented")
}

func (db *DB) Create(vm *VM) (uint64, error) {
	return 0, errors.New("unimplemented")
}

func (db *DB) Delete(id uint64) error {
	return errors.New("unimplemented")
}

func (db *DB) Update(id uint64, vm *VM) error {
	return errors.New("unimplemented")
}

func (db *DB) List() ([]uint64, error) {
	return nil, errors.New("unimplemented")
}
