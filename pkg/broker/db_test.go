package broker

import (
	"fmt"
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

func TestNext(t *testing.T) {
	db := makeDB(t)

	var values []string

	queue, err := db.Queue("std")
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < 10000; i++ {
		value := testutil.RandomString(30, 5)
		if err := queue.Insert(value); err != nil {
			t.Fatal(err)
		}

		values = append(values, value)
	}

	for _, value := range values {
		var nextValue string
		if err := queue.Next(&nextValue); err != nil {
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

	queue, err := db.Queue("std")
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < 10000; i++ {
		value := testutil.RandomString(30, 5)
		if err := queue.Insert(value); err != nil {
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
				if err := queue.Next(&nextValue); err != nil {
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

	if err := queue.Next(&nextValue); err != ErrRecordNotFound {
		t.Fatalf("invalid error occurred after draining queue: %v", err)
	}
}

func TestNextParallelMultiQueue(t *testing.T) {
	db := makeDB(t)

	values := map[string]struct{}{}
	queues := []*Queue{}
	concurrency := runtime.NumCPU() * 2

	for i := 0; i < concurrency; i++ {
		queue, err := db.Queue(fmt.Sprintf("%d", i))
		if err != nil {
			t.Fatal(err)
		}

		queues = append(queues, queue)
	}

	for i := 0; i < 100000; i++ {
		value := testutil.RandomString(30, 5)
		if err := queues[i%concurrency].Insert(value); err != nil {
			t.Fatal(err)
		}

		values[value] = struct{}{}
	}

	errChan := make(chan error, 1)
	valueChan := make(chan string, concurrency)

	for i := 0; i < concurrency; i++ {
		queue := queues[i]
		go func(queue *Queue) {
			for {
				var nextValue string
				if err := queue.Next(&nextValue); err != nil {
					if err != ErrRecordNotFound {
						errChan <- err
					}

					return
				}

				valueChan <- nextValue
			}
		}(queue)
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
}

func TestPackage(t *testing.T) {
	db := makeDB(t)

	packages := []*Package{}
	commands := []*Command{}
	resources := []string{}
	resource_commands := map[string][]*Command{}

	for i := 0; i < 100; i++ {
		resource := testutil.RandomString(30, 5)
		resources = append(resources, resource)
	}

	for i := 0; i < 100; i++ {
		pkg, err := db.NewPackage()
		if err != nil {
			t.Fatal(err)
		}

		for i := 0; i < 100; i++ {
			c := &Command{
				Resource:   resources[i],
				Action:     testutil.RandomString(30, 5),
				Parameters: []string{testutil.RandomString(30, 5)},
			}

			err := pkg.Add(c)
			if err != nil {
				t.Fatal(err)
			}

			commands = append(commands, c)
			resource_commands[resources[i]] = append(resource_commands[resources[i]], c)
		}

		packages = append(packages, pkg)
	}

	for _, pkg := range packages {
		list, err := pkg.List()
		if err != nil {
			t.Fatal(err)
		}

		if len(list) == 0 {
			t.Fatal("list yielded no results")
		}

		for _, c := range list {
			command := commands[0]
			if len(commands) > 0 {
				commands = commands[1:]
			}

			if !reflect.DeepEqual(command, c) {
				t.Fatal("commands did not match")
			}
		}

		if err := pkg.Enqueue(); err != nil {
			t.Fatal(err)
		}
	}

	for _, resource := range resources {
		queue, err := db.Queue(resource)
		if err != nil {
			t.Fatal(err)
		}

		var c Command

		for _, command := range resource_commands[resource] {
			if err := queue.Next(&c); err != nil {
				t.Fatal(err)
			}

			if !reflect.DeepEqual(*command, c) {
				t.Fatalf("values did not match: %v %v", *command, c)
			}

			if err := db.FinishCommand(c.UUID, true, ""); err != nil {
				t.Fatal(err)
			}
		}
	}

	for _, pkg := range packages {
		if err := pkg.Finished(); err != nil {
			t.Fatal(err)
		}
	}
}
