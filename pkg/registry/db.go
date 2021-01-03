package registry

import (
	"encoding/binary"
	"encoding/json"
	"errors"

	spinregistry "code.hollensbe.org/erikh/spin/gen/spin_registry"
	"go.etcd.io/bbolt"
)

// DBConfig is the configuration of the database.
type DBConfig struct {
	Filename string
}

// DB is the handle into the registry database.
type DB struct {
	db *bbolt.DB
}

const vmBucket = "vms"

var makeBuckets = []string{vmBucket}

// NewDB creates a new *DB based on a DBConfig.
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

// Create creates a vm and returns its id, which is auto-generated.
func (db *DB) Create(vm *spinregistry.VM) (uint64, error) {
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

// Get retrieves the vm at the id specified.
func (db *DB) Get(id uint64) (*spinregistry.VM, error) {
	var vm spinregistry.VM

	return &vm, db.db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(vmBucket))

		value := bucket.Get(makeKey(id))
		if value == nil {
			return errors.New("doesn't exist")
		}

		return json.Unmarshal(value, &vm)
	})
}

// Delete removes the vm at the id specified.
func (db *DB) Delete(id uint64) error {
	return db.db.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(vmBucket))
		key := makeKey(id)
		if bucket.Get(key) == nil {
			return errors.New("doesn't exist")
		}

		return bucket.Delete(makeKey(id))
	})
}

// Update replaces the id with the new vm definition.
func (db *DB) Update(id uint64, vm *spinregistry.VM) error {
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

// List returns a list of all stored vm ids.
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
