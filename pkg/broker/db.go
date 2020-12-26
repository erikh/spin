package broker

import (
	"encoding/binary"
	"encoding/json"
	"errors"

	"go.etcd.io/bbolt"
)

const (
	BucketQueue    = "queues"
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

type Queue struct {
	name []byte
	db   *DB
}

func (db *DB) Queue(name string) (*Queue, error) {
	err := db.db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.Bucket([]byte(BucketQueue)).CreateBucketIfNotExists([]byte(name))
		return err
	})

	if err != nil {
		return nil, err
	}

	return &Queue{
		db:   db,
		name: []byte(name),
	}, nil
}

func (q *Queue) Insert(value interface{}) error {
	return q.db.db.Update(func(tx *bbolt.Tx) error {
		content, err := json.Marshal(value)
		if err != nil {
			return err
		}

		queueBucket := tx.Bucket([]byte(BucketQueue)).Bucket([]byte(q.name))

		seq, err := queueBucket.NextSequence()
		if err != nil {
			return err
		}

		key := make([]byte, 8)
		binary.BigEndian.PutUint64(key, seq)

		return queueBucket.Put(key, content)
	})
}

func (q *Queue) Next(data interface{}) error {
	tx, err := q.db.db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	queueBucket := tx.Bucket([]byte(BucketQueue)).Bucket([]byte(q.name))
	key, value := queueBucket.Cursor().First()
	if key == nil {
		return ErrRecordNotFound
	}

	if err := queueBucket.Delete(key); err != nil {
		return err
	}

	if err := json.Unmarshal(value, data); err != nil {
		return err
	}

	return tx.Commit()
}
