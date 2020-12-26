package broker

import (
	"encoding/binary"
	"encoding/json"
	"errors"

	"go.etcd.io/bbolt"
)

const (
	BucketQueue    = "queue_items"
	BucketCommands = "commands"
	BucketPackages = "packages"
	BucketStatuses = "statuses"
)

var MakeBuckets = []string{BucketQueue, BucketCommands, BucketPackages, BucketStatuses}

var ErrRecordNotFound = errors.New("Record not found")

type DB struct {
	db *bbolt.DB
}

func New(dbpath string) (*DB, error) {
	db, err := bbolt.Open(dbpath, 0600, &bbolt.Options{FreelistType: bbolt.FreelistMapType})
	if err != nil {
		return nil, err
	}

	return &DB{db: db}, db.Update(func(tx *bbolt.Tx) error {
		for _, bucket := range MakeBuckets {
			if _, err := tx.CreateBucketIfNotExists([]byte(bucket)); err != nil {
				return err
			}
		}

		return nil
	})
}

func (db *DB) Close() error {
	return db.db.Close()
}

func (db *DB) Put(bucket, key string, data interface{}) error {
	return db.db.Update(func(tx *bbolt.Tx) error {
		content, err := json.Marshal(data)
		if err != nil {
			return err
		}

		return tx.Bucket([]byte(bucket)).Put([]byte(key), content)
	})
}

func (db *DB) Get(bucket, key string, value interface{}) error {
	return db.db.View(func(tx *bbolt.Tx) error {
		content := tx.Bucket([]byte(bucket)).Get([]byte(key))
		if content == nil {
			return ErrRecordNotFound
		}

		return json.Unmarshal(content, value)
	})
}

func (db *DB) Insert(bucket string, value interface{}) error {
	return db.db.Update(func(tx *bbolt.Tx) error {
		content, err := json.Marshal(value)
		if err != nil {
			return err
		}

		seq, err := tx.Bucket([]byte(bucket)).NextSequence()
		if err != nil {
			return err
		}

		key := make([]byte, 8)
		binary.BigEndian.PutUint64(key, seq)

		return tx.Bucket([]byte(bucket)).Put(key, content)
	})
}

func (db *DB) Next(bucket string, data interface{}) error {
	tx, err := db.db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	key, value := tx.Bucket([]byte(bucket)).Cursor().First()
	if key == nil {
		return ErrRecordNotFound
	}

	if err := tx.Bucket([]byte(bucket)).Delete(key); err != nil {
		return err
	}

	if err := json.Unmarshal(value, data); err != nil {
		return err
	}

	return tx.Commit()
}
