package services

import (
	"log"
	"net/http"
	"os"

	"code.hollensbe.org/erikh/spin"
	spinregistrysvr "code.hollensbe.org/erikh/spin/gen/http/spin_registry/server"
	spinregistry "code.hollensbe.org/erikh/spin/gen/spin_registry"
	"code.hollensbe.org/erikh/spin/pkg/registry"
	goahttp "goa.design/goa/v3/http"
	httpmdlwr "goa.design/goa/v3/http/middleware"
	"goa.design/goa/v3/middleware"
)

// Registry bootstraps a handler suitable for serving a spin-registry
func Registry(dbpath string) (http.Handler, error) {
	logger := log.New(os.Stderr, "[spin] ", log.Ltime)

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
	handler = httpmdlwr.Log(middleware.NewLogger(logger))(handler)
	handler = httpmdlwr.RequestID()(handler)

	return handler, nil
}
