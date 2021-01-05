package registry

import (
	"encoding/binary"

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

const (
	vmBucket      = "registry_vms"
	storageBucket = "registry_storage"
)

var makeBuckets = []string{vmBucket, storageBucket}

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

func makeUintPtr(i uint64) *uint64 {
	return &i
}

func makeKey(id uint64) []byte {
	byt := make([]byte, 8)
	binary.BigEndian.PutUint64(byt, id)
	return byt
}

func decodeKey(key []byte) uint64 {
	return binary.BigEndian.Uint64(key)
}
