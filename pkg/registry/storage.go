package registry

import (
	"encoding/json"
	"errors"

	spinregistry "code.hollensbe.org/erikh/spin/gen/spin_registry"
	"go.etcd.io/bbolt"
)

// StorageVolumeCreate creates a volume
func (db *DB) StorageVolumeCreate(name string) error {
	return db.db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.Bucket([]byte(storageBucket)).CreateBucket([]byte(name))
		return err
	})
}

// StorageVolumeDelete creates a volume
func (db *DB) StorageVolumeDelete(name string) error {
	return db.db.Update(func(tx *bbolt.Tx) error {
		return tx.Bucket([]byte(storageBucket)).DeleteBucket([]byte(name))
	})
}

// StorageVolumeList lists all volumes
func (db *DB) StorageVolumeList() ([]string, error) {
	var list []string
	return list, db.db.View(func(tx *bbolt.Tx) error {
		cur := tx.Bucket([]byte(storageBucket)).Cursor()
		for key, _ := cur.First(); key != nil; key, _ = cur.Next() {
			list = append(list, string(key))
		}

		return nil
	})
}

// StorageImageCreate creates an image within a volume.
func (db *DB) StorageImageCreate(s *spinregistry.Storage) error {
	return db.db.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(storageBucket)).Bucket([]byte(s.Volume))
		if bucket == nil {
			return errors.New("volume does not exist")
		}

		byt := []byte(s.Image)
		if bucket.Get(byt) != nil {
			return errors.New("already exists")
		}

		content, err := json.Marshal(s)
		if err != nil {
			return err
		}

		return bucket.Put(byt, content)
	})
}

// StorageImageDelete will delete an image from the volume.
func (db *DB) StorageImageDelete(volume, image string) error {
	return db.db.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(storageBucket)).Bucket([]byte(volume))
		if bucket == nil {
			return errors.New("volume does not exist")
		}

		obj := bucket.Get([]byte(image))
		if obj == nil {
			return errors.New("image does not exist")
		}

		return bucket.Delete([]byte(image))
	})
}

// StorageImageGet will retrieve an image from the volume as a *spinregistry.Storage.
func (db *DB) StorageImageGet(volume, image string) (*spinregistry.Storage, error) {
	var s spinregistry.Storage

	return &s, db.db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(storageBucket)).Bucket([]byte(volume))
		if bucket == nil {
			return errors.New("volume does not exist")
		}

		obj := bucket.Get([]byte(image))
		if obj == nil {
			return errors.New("image does not exist")
		}

		return json.Unmarshal(obj, &s)
	})
}

// StorageImageList will retrieve a list of images for a given volume name.
func (db *DB) StorageImageList(volume string) ([]string, error) {
	var list []string
	return list, db.db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(storageBucket)).Bucket([]byte(volume))
		if bucket == nil {
			return errors.New("volume does not exist")
		}

		cur := bucket.Cursor()
		for key, _ := cur.First(); key != nil; key, _ = cur.Next() {
			list = append(list, string(key))
		}

		return nil
	})
}
