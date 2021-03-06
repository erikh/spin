package services

import (
	"log"
	"net/http"
	"os"

	"github.com/erikh/spin"
	spinregistrysvr "github.com/erikh/spin/gen/http/spin_registry/server"
	spinregistry "github.com/erikh/spin/gen/spin_registry"
	"github.com/erikh/spin/pkg/registry"
	goahttp "goa.design/goa/v3/http"
	httpmdlwr "goa.design/goa/v3/http/middleware"
	"goa.design/goa/v3/middleware"
)

// Registry bootstraps a handler suitable for serving a spin-registry
func Registry(dbpath string, showLog bool) (http.Handler, error) {
	var logger *log.Logger
	if showLog {
		logger = log.New(os.Stderr, "[spin-registry] ", log.Ltime)
	}

	db, err := registry.NewDB(registry.DBConfig{Filename: dbpath})
	if err != nil {
		return nil, err
	}

	spinRegistryEndpoints := spinregistry.NewEndpoints(spin.NewSpinRegistry(logger, db))
	dec := goahttp.RequestDecoder
	enc := goahttp.ResponseEncoder
	mux := goahttp.NewMuxer()

	spinRegistryServer := spinregistrysvr.New(spinRegistryEndpoints, mux, dec, enc, errorHandler(logger), nil)
	spinregistrysvr.Mount(mux, spinRegistryServer)
	var handler http.Handler = mux
	if showLog {
		handler = httpmdlwr.Log(middleware.NewLogger(logger))(handler)
	}
	handler = httpmdlwr.RequestID()(handler)

	return handler, nil
}
