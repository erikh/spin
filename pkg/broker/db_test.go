package broker

import (
	"io/ioutil"
	"os"
	"reflect"
	"runtime"
	"testing"

	"code.hollensbe.org/erikh/spin/pkg/testutil"
)

func makeDB(t *testing.T) *DB {
	f, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatal(err)
	}
	os.Remove(f.Name())
	t.Cleanup(func() { os.Remove(f.Name()) })

	db, err := New(f.Name())
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() { db.Close() })
	return db
}

func TestDBInit(t *testing.T) {
	makeDB(t)
}

func TestGetPut(t *testing.T) {
	db := makeDB(t)

	if err := db.Put(BucketCommands, "one", []string{"one", "two", "three"}); err != nil {
		t.Fatal(err)
	}

	var result []string

	if err := db.Get(BucketCommands, "one", &result); err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(result, []string{"one", "two", "three"}) {
		t.Fatal("result did not match insert")
	}
}

func TestNext(t *testing.T) {
	db := makeDB(t)

	var values []string

	for i := 0; i < 10000; i++ {
		value := testutil.RandomString(30, 5)
		if err := db.Insert(BucketCommands, value); err != nil {
			t.Fatal(err)
		}

		values = append(values, value)
	}

	for _, value := range values {
		var nextValue string
		if err := db.Next(BucketCommands, &nextValue); err != nil {
			t.Fatal(err)
		}

		if value != nextValue {
			t.Fatal("values do not match")
		}
	}
}

func TestNextParallel(t *testing.T) {
	db := makeDB(t)

	values := map[string]struct{}{}

	for i := 0; i < 10000; i++ {
		value := testutil.RandomString(30, 5)
		if err := db.Insert(BucketCommands, value); err != nil {
			t.Fatal(err)
		}

		values[value] = struct{}{}
	}

	concurrency := runtime.NumCPU() * 2

	errChan := make(chan error, 1)
	valueChan := make(chan string, concurrency)

	for i := 0; i < concurrency; i++ {
		go func() {
			for {
				var nextValue string
				if err := db.Next(BucketCommands, &nextValue); err != nil {
					if err != ErrRecordNotFound {
						errChan <- err
					}

					return
				}

				valueChan <- nextValue
			}
		}()
	}

	for len(values) > 0 {
		select {
		case err := <-errChan:
			t.Fatal(err)
		case nextValue := <-valueChan:
			if _, ok := values[nextValue]; !ok {
				t.Fatal("value was already returned")
			}

			delete(values, nextValue)
		}
	}

	var nextValue string

	if err := db.Next(BucketCommands, &nextValue); err != ErrRecordNotFound {
		t.Fatalf("invalid error occurred after draining queue: %v", err)
	}
}
