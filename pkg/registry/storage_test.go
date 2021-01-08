package registry

import (
	"path/filepath"
	"reflect"
	"testing"

	spinregistry "github.com/erikh/spin/gen/spin_registry"
)

func TestStorageVolumeDBCRUD(t *testing.T) {
	db := makeDB(t)

	if err := db.StorageVolumeDelete("test"); err == nil {
		t.Fatal("no error deleting invalid volume")
	}

	m, err := db.StorageVolumeList()
	if err != nil {
		t.Fatal(err)
	}

	if len(m) > 0 {
		t.Fatal("list yielded invalid items")
	}

	volumeMap := map[string]string{}
	for _, vol := range []string{"test", "test2", "test3"} {
		volumeMap[vol] = filepath.Join("/tmp", vol)
		if err := db.StorageVolumeCreate(vol, filepath.Join("/tmp", vol)); err != nil {
			t.Fatal(err)
		}
	}

	m, err = db.StorageVolumeList()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(volumeMap, m) {
		t.Fatalf("lists were not equal: %v %v", volumeMap, m)
	}

	for vol := range volumeMap {
		if err := db.StorageVolumeDelete(vol); err != nil {
			t.Fatal(err)
		}
	}

	m, err = db.StorageVolumeList()
	if err != nil {
		t.Fatal(err)
	}

	if len(m) > 0 {
		t.Fatal("list yielded invalid items")
	}
}

func TestStorageImageDBCRUD(t *testing.T) {
	db := makeDB(t)

	volumeName := "test"

	image := &spinregistry.Storage{
		Image:     "test.raw",
		Volume:    &volumeName,
		ImageSize: makeUintPtr(50),
	}

	if _, err := db.StorageImageCreate(image); err == nil {
		t.Fatal("was able to create image with invalid volume")
	}

	if _, err := db.StorageImageGet("test", "test.raw"); err == nil {
		t.Fatal("was able to fetch image with invalid volume")
	}

	if err := db.StorageImageDelete("test", "test.raw"); err == nil {
		t.Fatal("was able to fetch image with invalid volume")
	}

	if _, err := db.StorageImageList("test"); err == nil {
		t.Fatal("was able to list images with invalid volume")
	}

	if err := db.StorageVolumeCreate("test", "/tmp/test"); err != nil {
		t.Fatal(err)
	}

	if _, err := db.StorageImageGet("test", "test.raw"); err == nil {
		t.Fatal("was able to fetch an invalid image for a valid volume")
	}

	if err := db.StorageImageDelete("test", "test.raw"); err == nil {
		t.Fatal("was able to delete an invalid image for a valid volume")
	}

	img, err := db.StorageImageCreate(image)
	if err != nil {
		t.Fatal(err)
	}

	if img.Path != "/tmp/test/test.raw" {
		t.Fatal("image path was invalid")
	}

	getImage, err := db.StorageImageGet("test", "test.raw")
	if err != nil {
		t.Fatal("was able to fetch an invalid image for a valid volume")
	}

	if !reflect.DeepEqual(img, getImage) {
		t.Fatal("storage structs were not equal", getImage, img)
	}

	if err := db.StorageImageDelete("test", "test.raw"); err != nil {
		t.Fatal(err)
	}

	images, err := db.StorageImageList("test")
	if err != nil {
		t.Fatal(err)
	}

	if len(images) != 0 {
		t.Fatal("images found in empty volume")
	}

	imageList := []string{"test1.raw", "test2.raw", "test3.raw"}

	for _, imgName := range imageList {
		image := &spinregistry.Storage{
			Image:     imgName,
			Volume:    &volumeName,
			ImageSize: makeUintPtr(50),
		}

		img, err := db.StorageImageCreate(image)
		if err != nil {
			t.Fatal(err)
		}

		if img.Path != filepath.Join("/tmp/test", imgName) {
			t.Fatal("filepath was not equal to expectation")
		}
	}

	images, err = db.StorageImageList("test")
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(images, imageList) {
		t.Fatalf("image list returned was not equal: %v %v", images, imageList)
	}

	for _, imgName := range imageList {
		if err := db.StorageImageDelete("test", imgName); err != nil {
			t.Fatal(err)
		}
	}

	images, err = db.StorageImageList("test")
	if err != nil {
		t.Fatal(err)
	}

	if len(images) != 0 {
		t.Fatal("images found in empty volume")
	}
}
