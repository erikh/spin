package services

import (
	"io/ioutil"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
)

func SetupTestBroker(t *testing.T) string {
	dbpath, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatal(err)
	}
	os.Remove(dbpath.Name())
	t.Cleanup(func() {
		os.Remove(dbpath.Name())
	})

	handler, err := Broker(dbpath.Name(), false)
	if err != nil {
		t.Fatal(err)
	}

	server := httptest.NewServer(handler)
	t.Cleanup(func() {
		server.Close()
	})

	u, err := url.Parse(server.URL)
	if err != nil {
		t.Fatal(err)
	}

	return u.Host
}
