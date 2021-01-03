package registry

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	spinregistry "code.hollensbe.org/erikh/spin/gen/spin_registry"
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

func TestDBCRUD(t *testing.T) {
	db := makeDB(t)

	vm := &spinregistry.VM{
		Name: "foo",
	}

	id, err := db.VMCreate(vm)
	if err != nil {
		t.Fatal(err)
	}

	if id == 0 {
		t.Fatal("id was equal to 0")
	}

	vm2, err := db.VMGet(id)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(vm, vm2) {
		t.Fatal("vms were not equal")
	}

	vm.Cpus = 1 // set something differently to test update

	if err := db.VMUpdate(id, vm); err != nil {
		t.Fatal(err)
	}

	vm2, err = db.VMGet(id)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(vm, vm2) {
		t.Fatal("vms were not equal")
	}

	ids, err := db.VMList()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(ids, []uint64{id}) {
		t.Fatal("id list was not equal")
	}

	if err := db.VMDelete(id); err != nil {
		t.Fatal(err)
	}

	if _, err := db.VMGet(id); err == nil {
		t.Fatal("vm is still available after delete")
	}

	ids, err = db.VMList()
	if err != nil {
		t.Fatal(err)
	}

	if len(ids) != 0 {
		t.Fatal("vms still available after all have been deleted")
	}

	ids = []uint64{}

	for i := 0; i < 10; i++ {
		id, err := db.VMCreate(vm)
		if err != nil {
			t.Fatal(err)
		}

		ids = append(ids, id)
	}

	ids2, err := db.VMList()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(ids, ids2) {
		t.Fatal("lists don't match")
	}
}

func TestDBCRUDTable(t *testing.T) {
	table := map[string]struct {
		pass bool
		call func(db *DB) error
	}{
		"get non-existent": {
			call: func(db *DB) error {
				_, err := db.VMGet(1)
				return err
			},
		},
		"delete non-existent": {
			call: func(db *DB) error {
				return db.VMDelete(1)
			},
		},
		"update non-existent": {
			call: func(db *DB) error {
				return db.VMUpdate(1, &spinregistry.VM{})
			},
		},
	}

	for name, item := range table {
		db := makeDB(t)
		err := item.call(db)
		if err != nil && item.pass {
			t.Fatalf("[%v]: %v", name, err)
		} else if !item.pass && err == nil {
			t.Fatalf("[%v]: Was expected to fail and did not", name)
		}
	}
}
