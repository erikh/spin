package registry

import (
	"encoding/json"
	"errors"
	"fmt"
	"path/filepath"

	spinregistry "github.com/erikh/spin/gen/spin_registry"
	"go.etcd.io/bbolt"
)

// StorageVolumeCreate creates a volume
func (db *DB) StorageVolumeCreate(name string, path string) error {
	return db.db.Update(func(tx *bbolt.Tx) error {
		pathBucket := tx.Bucket([]byte(storagePathBucket))

		if pathBucket.Get([]byte(name)) != nil {
			return errors.New("path already in use")
		}

		if err := pathBucket.Put([]byte(name), []byte(path)); err != nil {
			return err
		}

		_, err := tx.Bucket([]byte(storageBucket)).CreateBucket([]byte(name))
		return err
	})
}

// StorageVolumeDelete creates a volume
func (db *DB) StorageVolumeDelete(name string) error {
	return db.db.Update(func(tx *bbolt.Tx) error {
		if err := tx.Bucket([]byte(storagePathBucket)).Delete([]byte(name)); err != nil {
			fmt.Printf("Fell through on storage volume delete with this error pruning the path: %v", err)
		}
		return tx.Bucket([]byte(storageBucket)).DeleteBucket([]byte(name))
	})
}

// StorageVolumeList lists all volumes
func (db *DB) StorageVolumeList() (map[string]string, error) {
	list := map[string]string{}
	return list, db.db.View(func(tx *bbolt.Tx) error {
		cur := tx.Bucket([]byte(storagePathBucket)).Cursor()
		for key, value := cur.First(); key != nil; key, value = cur.Next() {
			list[string(key)] = string(value)
		}

		return nil
	})
}

// StorageImageCreate creates an image within a volume.
func (db *DB) StorageImageCreate(s *spinregistry.Storage) (*spinregistry.Image, error) {
	if s.Cdrom {
		return nil, errors.New("cannot create cdrom images")
	}

	if s.Volume == nil {
		return nil, errors.New("volume cannot be nil")
	}

	image := &spinregistry.Image{}

	return image, db.db.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(storageBucket)).Bucket([]byte(*s.Volume))
		if bucket == nil {
			return errors.New("volume does not exist")
		}

		byt := []byte(s.Image)
		if bucket.Get(byt) != nil {
			return errors.New("image already exists")
		}

		volPath := tx.Bucket([]byte(storagePathBucket)).Get([]byte(*s.Volume))
		if volPath == nil {
			return errors.New("invalid volume")
		}

		image.Path = filepath.Join(string(volPath), s.Image)
		image.Cdrom = false
		image.Volume = s.Volume

		content, err := json.Marshal(image)
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
func (db *DB) StorageImageGet(volume, image string) (*spinregistry.Image, error) {
	var s spinregistry.Image

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
