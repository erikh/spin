package registry

import (
	"io/ioutil"
	"os"
	"testing"
)

func makeDB(t *testing.T) *DB {
	f, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() { os.Remove(f.Name()) })

	if err := os.Remove(f.Name()); err != nil {
		t.Fatal(err)
	}

	db, err := NewDB(DBConfig{Filename: f.Name()})
	if err != nil {
		t.Fatal(err)
	}

	return db
}

func TestDBInit(t *testing.T) {
	makeDB(t)
}
