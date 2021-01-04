package registry

import (
	"reflect"
	"testing"

	spinregistry "code.hollensbe.org/erikh/spin/gen/spin_registry"
)

func TestStorageVolumeDBCRUD(t *testing.T) {
	db := makeDB(t)

	if err := db.StorageVolumeDelete("test"); err == nil {
		t.Fatal("no error deleting invalid volume")
	}

	list, err := db.StorageVolumeList()
	if err != nil {
		t.Fatal(err)
	}

	if len(list) > 0 {
		t.Fatal("list yielded invalid items")
	}

	volumeList := []string{"test", "test2", "test3"}
	for _, vol := range volumeList {
		if err := db.StorageVolumeCreate(vol); err != nil {
			t.Fatal(err)
		}
	}

	list, err = db.StorageVolumeList()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(volumeList, list) {
		t.Fatalf("lists were not equal: %v %v", volumeList, list)
	}

	for _, vol := range volumeList {
		if err := db.StorageVolumeDelete(vol); err != nil {
			t.Fatal(err)
		}
	}

	list, err = db.StorageVolumeList()
	if err != nil {
		t.Fatal(err)
	}

	if len(list) > 0 {
		t.Fatal("list yielded invalid items")
	}
}

func TestStorageImageDBCRUD(t *testing.T) {
	db := makeDB(t)

	image := &spinregistry.Storage{
		Image:     "test.raw",
		Volume:    "test",
		ImageSize: makeUintPtr(50),
	}

	if err := db.StorageImageCreate(image); err == nil {
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

	if err := db.StorageVolumeCreate("test"); err != nil {
		t.Fatal(err)
	}

	if _, err := db.StorageImageGet("test", "test.raw"); err == nil {
		t.Fatal("was able to fetch an invalid image for a valid volume")
	}

	if err := db.StorageImageDelete("test", "test.raw"); err == nil {
		t.Fatal("was able to delete an invalid image for a valid volume")
	}

	if err := db.StorageImageCreate(image); err != nil {
		t.Fatal(err)
	}

	getImage, err := db.StorageImageGet("test", "test.raw")
	if err != nil {
		t.Fatal("was able to fetch an invalid image for a valid volume")
	}

	if !reflect.DeepEqual(image, getImage) {
		t.Fatal("storage structs were not equal")
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
			Volume:    "test",
			ImageSize: makeUintPtr(50),
		}

		if err := db.StorageImageCreate(image); err != nil {
			t.Fatal(err)
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
