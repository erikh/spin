package broker

import (
	"encoding/binary"
	"encoding/json"
	"errors"

	"code.hollensbe.org/erikh/spin/pkg/agent/dispatcher"
	"github.com/google/uuid"
	"go.etcd.io/bbolt"
)

const (
	// BucketQueue is the outer bucket for other queue-oriented buckets, keyed by resource.
	BucketQueue = "queues"
	// BucketPackages is the outer bucket for all package queues.
	BucketPackages = "packages"
	// BucketCommands is the bucket for all commands.
	BucketCommands = "commands"
	// BucketStatuses is the bucket for all statuses.
	BucketStatuses = "statuses"
)

var makeBuckets = []string{BucketQueue, BucketCommands, BucketPackages, BucketStatuses}

var (
	// ErrRecordNotFound is returned when records are not found. This in
	// particular can happen in scenarios when requesting the status of an object
	// and it is unavailable.
	ErrRecordNotFound = errors.New("Record not found")

	// ErrRecordAlreadyExists is returned when trying to set a status that has
	// already been set.
	ErrRecordAlreadyExists = errors.New("Record already exists")
)

// DB is our DB handle. It contains a *bbolt.DB handle for internal use.
type DB struct {
	db *bbolt.DB
}

// New creates a new database handle. The file is created if it does not
// already exist with 0600 perms, and some tuning is applied. The outermost
// buckets are also created at this time if they do not exist.
func New(dbpath string) (*DB, error) {
	db, err := bbolt.Open(dbpath, 0600, &bbolt.Options{FreelistType: bbolt.FreelistMapType})
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

// Close the database.
func (db *DB) Close() error {
	return db.db.Close()
}

type status struct {
	Status bool
	Reason string
}

// ErrorStatus is for status reports that are in an errored state. This
// includes the reason as well. The type mostly exists to distinguish it from
// other golang errors.
type ErrorStatus struct {
	Reason string
}

func (es ErrorStatus) Error() string {
	return es.Reason
}

// CommandStatus returns the status of the Command UUID as an error if faulty,
// otherwise nil for success.
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

// FinishCommand reports a status for the Command UUID. If the result is true,
// reason is ignored.
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

// Package is a collection of queue items that are in a waiting or enqueued
// state as a group.
type Package struct {
	name []byte
	db   *DB
}

// Package returns a Package from the provided UUID. If none is found,
// ErrRecordNotFound is returned.
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

// NewPackage creates a new Package and returns it, or error if there was any
// trouble. It will retry on collisions.
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

// Command is a unit of instruction; it contains a UUID, the unique identifier
// of the commmand, a Resource, the type of command to execute, an action, the
// name of the command to execute, and parameters, a collection of items that
// relate to the action for the purposes of execution.
//
// Commands are typically fed to Packages, then the Package is Enqueued, Next()
// calls are made to yield the commands for the resource, the command is
// processed, FinishCommand is called to finish the command, then statuses are
// polled and eventually yielded.
type Command dispatcher.Command

// UUID returns the Package UUID.
func (p *Package) UUID() string {
	return string(p.name)
}

// Add adds a Command to the Package. It will be assigned a UUID which is
// written back to the pointered value. This may be retried. Any error writing
// to the database or marshaling the values will return an error.
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

// List lists the commands for the Package.
func (p *Package) List() ([]Command, error) {
	var c []Command

	return c, p.db.db.View(func(tx *bbolt.Tx) error {
		cursor := tx.Bucket([]byte(BucketPackages)).Bucket(p.name).Cursor()

		for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
			var command Command
			if err := json.Unmarshal(v, &command); err != nil {
				return err // XXX skip these items, maybe?
			}

			c = append(c, command)
		}

		return nil
	})
}

// Enqueue enqueues the package in the relevant resource queues.
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

// Finished returns non-nil if:
// - The package is unfinished (ErrRecordNotFound)
// - The package had errors (ErrorStatus)
//
// Otherwise, it returns nil, which means everything ran fine.
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

// Queue is the notion of a resource-based queue in the broker. These are
// Command structs separated by resource and organized into order.
type Queue struct {
	name []byte
	db   *DB
}

// Queue retrieves the Queue object for the named resource.
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

// Insert inserts an item into the Queue.
func (q *Queue) Insert(value Command) error {
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

// Next returns the next item in the queue, then removes it from the queue.
func (q *Queue) Next() (Command, error) {
	var data Command

	tx, err := q.db.db.Begin(true)
	if err != nil {
		return data, err
	}
	defer tx.Rollback()

	queueBucket := tx.Bucket([]byte(BucketQueue)).Bucket(q.name)

	key, value := queueBucket.Cursor().First()
	if key == nil {
		return data, ErrRecordNotFound
	}

	if err := json.Unmarshal(value, &data); err != nil {
		return data, err
	}

	for _, dep := range data.Dependencies {
		if err := q.db.CommandStatus(dep); err != nil {
			return data, err
		}
	}

	if err := queueBucket.Delete(key); err != nil {
		return data, err
	}

	return data, tx.Commit()
}
