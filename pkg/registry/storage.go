package registry

import "go.etcd.io/bbolt"

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
