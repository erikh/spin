package registry

import (
	"errors"

	"go.etcd.io/bbolt"
)

type DBConfig struct {
	Filename string
}

type DB struct {
	db *bbolt.DB
}

var makeBuckets = []string{"vms"}

func NewDB(c DBConfig) (*DB, error) {
	db, err := bbolt.Open(c.Filename, 0600, &bbolt.Options{FreelistType: bbolt.FreelistMapType})
	if err != nil {
		return nil, err
	}

	return &DB{db: db}, db.Update(func(tx *bbolt.Tx) error {
		for _, bucket := range makeBuckets {
			if _, err := tx.CreateBucketIfNotExists([]byte(bucket)); err != nil {
				return err
			}
		}

		return nil
	})
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
