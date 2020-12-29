package registry

import (
	"encoding/binary"
	"encoding/json"
	"errors"

	"go.etcd.io/bbolt"
)

type DBConfig struct {
	Filename string
}

type DB struct {
	db *bbolt.DB
}

const vmBucket = "vms"

var makeBuckets = []string{vmBucket}

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

func makeKey(id uint64) []byte {
	byt := make([]byte, 8)
	binary.BigEndian.PutUint64(byt, id)
	return byt
}

func decodeKey(key []byte) uint64 {
	return binary.BigEndian.Uint64(key)
}

func (db *DB) Create(vm *VM) (uint64, error) {
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

func (db *DB) Get(id uint64) (*VM, error) {
	var vm VM

	return &vm, db.db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(vmBucket))

		value := bucket.Get(makeKey(id))
		if value == nil {
			return errors.New("doesn't exist")
		}

		return json.Unmarshal(value, &vm)
	})
}

func (db *DB) Delete(id uint64) error {
	return db.db.Update(func(tx *bbolt.Tx) error {
		tx.Bucket([]byte(vmBucket)).Delete(makeKey(id))
		return nil
	})
}

func (db *DB) Update(id uint64, vm *VM) error {
	content, err := json.Marshal(vm)
	if err != nil {
		return err
	}

	return db.db.Update(func(tx *bbolt.Tx) error {
		return tx.Bucket([]byte(vmBucket)).Put(makeKey(id), content)
	})
}

func (db *DB) List() ([]uint64, error) {
	var ids []uint64

	return ids, db.db.View(func(tx *bbolt.Tx) error {
		cursor := tx.Bucket([]byte(vmBucket)).Cursor()

		for key, _ := cursor.First(); key != nil; key, _ = cursor.Next() {
			ids = append(ids, decodeKey(key))
		}

		return nil
	})
}
