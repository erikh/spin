package broker

import (
	"encoding/binary"
	"encoding/json"
	"errors"

	"github.com/google/uuid"
	"go.etcd.io/bbolt"
)

const (
	BucketQueue    = "queues"
	BucketPackages = "packages"
	BucketCommands = "commands"
	BucketStatuses = "statuses"
)

var MakeBuckets = []string{BucketQueue, BucketCommands, BucketPackages, BucketStatuses}

var (
	ErrRecordNotFound      = errors.New("Record not found")
	ErrRecordAlreadyExists = errors.New("Record already exists")
)

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

type status struct {
	Status bool
	Reason string
}

type ErrorStatus struct {
	Reason string
}

func (es ErrorStatus) Error() string {
	return es.Reason
}

func (db *DB) CommandStatus(uuid string) error {
	return db.db.View(func(tx *bbolt.Tx) error {
		value := tx.Bucket([]byte(BucketStatuses)).Get([]byte(uuid))
		if value == nil {
			return ErrRecordNotFound
		}

		var s status

		if err := json.Unmarshal(value, &s); err != nil {
			return err
		}

		if !s.Status {
			return ErrorStatus{s.Reason}
		}

		return nil
	})
}

func (db *DB) FinishCommand(uuid string, statusResult bool, reason string) error {
	return db.db.Update(func(tx *bbolt.Tx) error {
		value := tx.Bucket([]byte(BucketCommands)).Get([]byte(uuid))
		if value == nil {
			return ErrRecordNotFound
		}

		value = tx.Bucket([]byte(BucketStatuses)).Get([]byte(uuid))
		if value != nil {
			return ErrRecordAlreadyExists
		}

		s := status{
			Status: statusResult,
			Reason: reason,
		}

		content, err := json.Marshal(s)
		if err != nil {
			return err
		}

		return tx.Bucket([]byte(BucketStatuses)).Put([]byte(uuid), content)
	})
}

type Package struct {
	name []byte
	db   *DB
}

func (db *DB) Package(uuid string) (*Package, error) {
	err := db.db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(BucketPackages)).Bucket([]byte(uuid))
		if bucket == nil {
			return ErrRecordNotFound
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &Package{
		db:   db,
		name: []byte(uuid),
	}, nil
}

func (db *DB) NewPackage() (*Package, error) {
	var u string

	err := db.db.Update(func(tx *bbolt.Tx) error {
	redo:
		u = uuid.New().String()
		_, err := tx.Bucket([]byte(BucketPackages)).CreateBucket([]byte(u))

		if err == bbolt.ErrBucketExists {
			goto redo
		} else if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &Package{
		db:   db,
		name: []byte(u),
	}, nil
}

type Command struct {
	UUID       string
	Resource   string
	Action     string
	Parameters []string
}

func (p *Package) UUID() string {
	return string(p.name)
}

func (p *Package) Add(value *Command) error {
	return p.db.db.Update(func(tx *bbolt.Tx) error {
	redo:
		u := uuid.New().String()
		value.UUID = u
		cmdBucket := tx.Bucket([]byte(BucketCommands))
		if cmdBucket.Get([]byte(u)) != nil {
			goto redo
		}

		content, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if err := cmdBucket.Put([]byte(u), content); err != nil {
			return err
		}

		bucket := tx.Bucket([]byte(BucketPackages)).Bucket(p.name)
		id, err := bucket.NextSequence()
		if err != nil {
			return err
		}

		byt := make([]byte, 8)
		binary.BigEndian.PutUint64(byt, id)

		return bucket.Put(byt, content)
	})
}

func (p *Package) List() ([]*Command, error) {
	var c []*Command

	return c, p.db.db.View(func(tx *bbolt.Tx) error {
		cursor := tx.Bucket([]byte(BucketPackages)).Bucket(p.name).Cursor()

		for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
			var command Command
			if err := json.Unmarshal(v, &command); err != nil {
				return err // XXX skip these items, maybe?
			}

			c = append(c, &command)
		}

		return nil
	})
}

func (p *Package) Enqueue() error {
	commands, err := p.List()
	if err != nil {
		return err
	}

	for _, c := range commands {
		queue, err := p.db.Queue(c.Resource)
		if err != nil {
			return err
		}

		if err := queue.Insert(c); err != nil {
			return err
		}
	}

	return nil
}

func (p *Package) Finished() error {
	commands, err := p.List()
	if err != nil {
		return err
	}

	for _, c := range commands {
		if err := p.db.CommandStatus(c.UUID); err != nil {
			return err
		}
	}

	return nil
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

		queueBucket := tx.Bucket([]byte(BucketQueue)).Bucket(q.name)

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

	queueBucket := tx.Bucket([]byte(BucketQueue)).Bucket(q.name)

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
