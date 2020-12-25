package db

import (
	"context"
	"fmt"
	"reflect"
	"runtime"
	"testing"
	"time"

	"code.hollensbe.org/erikh/spin/pkg/testutil"
	"github.com/erikh/duct"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func makeDB(t *testing.T) (*DB, ConnConfig) {
	composer := duct.New(duct.Manifest{
		{
			Name: "spin-postgres",
			Env: []string{
				"POSTGRES_USER=spin",
				"POSTGRES_PASSWORD=spin",
				"POSTGRES_DB=spin",
			},
			Image:    "postgres:latest",
			BootWait: 2 * time.Second,
			PortForwards: map[int]int{
				5432: 5432,
			},
		},
	}, duct.WithNewNetwork("spin-db-test"))

	go composer.HandleSignals(true)
	if err := composer.Launch(context.Background()); err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		composer.Teardown(context.Background())
	})

	config := ConnConfig{
		User:     "spin",
		Password: "spin",
		Database: "spin",
		Host:     "localhost",
		Port:     5432,
	}
	db, err := New(config)

	if err != nil {
		t.Fatal(err)
	}

	return db, config
}

func TestDBInit(t *testing.T) {
	makeDB(t)
}

type SATest struct {
	gorm.Model
	Array StringArray
}

func TestStringArray(t *testing.T) {
	db, _ := makeDB(t)

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
	db, _ := makeDB(t)
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
	db, _ := makeDB(t)

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

func makePackages(t *testing.T, db *DB) ([]string, map[string]struct{}) {
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

	return packages, uuidMap
}

func TestEnqueueNextSingleResourceSerial(t *testing.T) {
	db, _ := makeDB(t)
	packages, uuidMap := makePackages(t, db)

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

		if _, err := db.EnqueuePackage(pkg); err == nil {
			t.Fatal("double-enqueue of package")
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

func TestEnqueueNextSingleResourceParallel(t *testing.T) {
	db, config := makeDB(t)
	packages, uuidMap := makePackages(t, db)

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

		if _, err := db.EnqueuePackage(pkg); err == nil {
			t.Fatal("double-enqueue of package")
		}
	}

	errChan := make(chan error, runtime.NumCPU())
	resultChan := make(chan *Command, runtime.NumCPU())
	doneChan := make(chan struct{}, runtime.NumCPU())

	for i := 0; i < runtime.NumCPU(); i++ {
		go func() {
			defer func() { doneChan <- struct{}{} }()
			db, err := New(config)
			if err != nil {
				errChan <- err
				return
			}

			for {
				command, err := db.NextQueueItem("resource")
				if err != nil {
					if err != gorm.ErrRecordNotFound {
						errChan <- err
					}
					return
				}

				resultChan <- command
			}
		}()
	}

	doneCount := 0

	for doneCount != runtime.NumCPU() {
		select {
		case err := <-errChan:
			t.Fatal(err)
		case result := <-resultChan:
			fmt.Println(result.UUID)
			if _, ok := uuidMap[result.UUID]; !ok {
				t.Fatalf("could not find queue item %q: %v", result.UUID, uuidMap)
			}

			delete(uuidMap, result.UUID)
		case <-doneChan:
			doneCount++
		}
	}
}
