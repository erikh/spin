package services

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"code.hollensbe.org/erikh/spin"
	spinbrokersvr "code.hollensbe.org/erikh/spin/gen/http/spin_broker/server"
	spinbroker "code.hollensbe.org/erikh/spin/gen/spin_broker"
	goahttp "goa.design/goa/v3/http"
	httpmdlwr "goa.design/goa/v3/http/middleware"
	"goa.design/goa/v3/middleware"
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

	logger := log.New(os.Stderr, "[spin] ", log.Ltime)

	spinBrokerSvc, err := spin.NewSpinBroker(logger, dbpath.Name())
	if err != nil {
		t.Fatal(err)
	}

	spinBrokerEndpoints := spinbroker.NewEndpoints(spinBrokerSvc)
	dec := goahttp.RequestDecoder
	enc := goahttp.ResponseEncoder
	mux := goahttp.NewMuxer()

	spinBrokerServer := spinbrokersvr.New(spinBrokerEndpoints, mux, dec, enc, errorHandler(logger), nil)
	spinbrokersvr.Mount(mux, spinBrokerServer)
	var handler http.Handler = mux
	handler = httpmdlwr.Log(middleware.NewLogger(logger))(handler)
	handler = httpmdlwr.RequestID()(handler)

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

// errorHandler returns a function that writes and logs the given error.
// The function also writes and logs the error unique ID so that it's possible
// to correlate.
func errorHandler(logger *log.Logger) func(context.Context, http.ResponseWriter, error) {
	return func(ctx context.Context, w http.ResponseWriter, err error) {
		id := ctx.Value(middleware.RequestIDKey).(string)
		_, _ = w.Write([]byte("[" + id + "] encoding: " + err.Error()))
		logger.Printf("[%s] ERROR: %s", id, err.Error())
	}
}
