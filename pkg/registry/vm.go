package registry

import (
	"encoding/json"
	"errors"

	spinregistry "github.com/erikh/spin/gen/spin_registry"
	"go.etcd.io/bbolt"
)

// VMCreate creates a vm and returns its id, which is auto-generated.
func (db *DB) VMCreate(vm *spinregistry.UpdatedVM) (uint64, error) {
	content, err := json.Marshal(vm)
	if err != nil {
		return 0, err
	}

	var id uint64

	err = db.db.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(vmBucket))

		var err error
		id, err = bucket.NextSequence()
		if err != nil {
			return err
		}

		return bucket.Put(makeKey(id), content)
	})

	return id, err
}

// VMGet retrieves the vm at the id specified.
func (db *DB) VMGet(id uint64) (*spinregistry.UpdatedVM, error) {
	var vm spinregistry.UpdatedVM

	return &vm, db.db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(vmBucket))

		value := bucket.Get(makeKey(id))
		if value == nil {
			return errors.New("doesn't exist")
		}

		return json.Unmarshal(value, &vm)
	})
}

// VMDelete removes the vm at the id specified.
func (db *DB) VMDelete(id uint64) error {
	return db.db.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(vmBucket))
		key := makeKey(id)
		if bucket.Get(key) == nil {
			return errors.New("doesn't exist")
		}

		return bucket.Delete(makeKey(id))
	})
}

// VMUpdate replaces the id with the new vm definition.
func (db *DB) VMUpdate(id uint64, vm *spinregistry.UpdatedVM) error {
	content, err := json.Marshal(vm)
	if err != nil {
		return err
	}

	return db.db.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(vmBucket))
		key := makeKey(id)
		if bucket.Get(key) == nil {
			return errors.New("doesn't exist")
		}

		return bucket.Put(makeKey(id), content)
	})
}

// VMList returns a list of all stored vm ids.
func (db *DB) VMList() ([]uint64, error) {
	var ids []uint64

	return ids, db.db.View(func(tx *bbolt.Tx) error {
		cursor := tx.Bucket([]byte(vmBucket)).Cursor()

		for key, _ := cursor.First(); key != nil; key, _ = cursor.Next() {
			ids = append(ids, decodeKey(key))
		}

		return nil
	})
}
