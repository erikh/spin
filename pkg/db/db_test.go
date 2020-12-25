package db

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	"code.hollensbe.org/erikh/spin/pkg/testutil"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func makeDB(t *testing.T) *DB {
	f, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatal(err)
	}
	os.Remove(f.Name())
	t.Cleanup(func() {
		os.Remove(f.Name())
	})

	db, err := New(f.Name())
	if err != nil {
		t.Fatal(err)
	}

	return db
}

func TestDBInit(t *testing.T) {
	makeDB(t)
}

type SATest struct {
	gorm.Model
	Array StringArray
}

func TestStringArray(t *testing.T) {
	db := makeDB(t)

	if err := db.db.AutoMigrate(&SATest{}); err != nil {
		t.Fatal(err)
	}

	test := &SATest{Array: []string{"foo", "bar", "baz"}}
	if err := db.db.Create(test).Error; err != nil {
		t.Fatal(err)
	}

	test2 := &SATest{}

	if err := db.db.Where("id = ?", test.ID).First(test2).Error; err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(test.Array, test2.Array) {
		t.Fatalf("not equal: %v %v", test.Array, test2.Array)
	}
}

func TestNewPackage(t *testing.T) {
	db := makeDB(t)
	for i := 0; i < 1000; i++ {
		u, err := db.NewPackage()
		if err != nil {
			t.Fatal(err)
		}

		if _, err := uuid.Parse(u); err != nil {
			t.Fatal(err)
		}
	}
}

func TestAddToPackage(t *testing.T) {
	db := makeDB(t)

	for i := 0; i < 100; i++ {
		u, err := db.NewPackage()
		if err != nil {
			t.Fatal(err)
		}

		for i := 0; i < 100; i++ {
			u2, err := db.AddToPackage(u, &AddCommand{
				Resource:   testutil.RandomString(30, 5),
				Action:     testutil.RandomString(30, 5),
				Parameters: []string{testutil.RandomString(30, 5)},
			})

			if err != nil {
				t.Fatal(err)
			}

			if _, err := uuid.Parse(u2); err != nil {
				t.Fatal(err)
			}
		}
	}
}

func TestEnqueueNextSingleResource(t *testing.T) {
	db := makeDB(t)

	packages := []string{}
	uuidMap := map[string]struct{}{}

	for i := 0; i < 100; i++ {
		u, err := db.NewPackage()
		if err != nil {
			t.Fatal(err)
		}

		for i := 0; i < 100; i++ {
			u2, err := db.AddToPackage(u, &AddCommand{
				Resource:   "resource",
				Action:     testutil.RandomString(30, 5),
				Parameters: []string{testutil.RandomString(30, 5)},
			})

			if err != nil {
				t.Fatal(err)
			}

			if _, err := uuid.Parse(u2); err != nil {
				t.Fatal(err)
			}

			uuidMap[u2] = struct{}{}
		}

		packages = append(packages, u)
	}

	for _, pkg := range packages {
		uuids, err := db.EnqueuePackage(pkg)
		if err != nil {
			t.Fatal(err)
		}

		for _, uuid := range uuids {
			if _, ok := uuidMap[uuid]; !ok {
				t.Fatal("enqueued bogus data")
			}
		}
	}

	for len(uuidMap) > 0 {
		command, err := db.NextQueueItem("resource")
		if err != nil {
			t.Fatalf("Error: %v, uuidMap: %v", err, uuidMap)
		}

		if _, ok := uuidMap[command.UUID]; !ok {
			t.Fatal("could not find queue item")
		}

		delete(uuidMap, command.UUID)
	}
}
