package registry

import (
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
