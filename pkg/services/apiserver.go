package services

import (
	"log"
	"net/http"
	"os"

	"github.com/erikh/spin"
	brokerclient "github.com/erikh/spin/clients/broker"
	registryclient "github.com/erikh/spin/clients/registry"
	spinapiserversvr "github.com/erikh/spin/gen/http/spin_apiserver/server"
	spinapiserver "github.com/erikh/spin/gen/spin_apiserver"
	goahttp "goa.design/goa/v3/http"
	httpmdlwr "goa.design/goa/v3/http/middleware"
	"goa.design/goa/v3/middleware"
)

// APIServer bootstraps a handler suitable for serving a spin-apiserver
func APIServer(showLog bool, brokerConfig brokerclient.Config, registryConfig registryclient.Config) (http.Handler, error) {
	var logger *log.Logger
	if showLog {
		logger = log.New(os.Stderr, "[spin-apiserver] ", log.Ltime)
	}

	spinAPIEndpoints := spinapiserver.NewEndpoints(
		spin.NewSpinApiserver(
			logger,
			brokerclient.New(brokerConfig),
			registryclient.New(registryConfig),
		))

	dec := goahttp.RequestDecoder
	enc := goahttp.ResponseEncoder
	mux := goahttp.NewMuxer()

	spinAPIServer := spinapiserversvr.New(spinAPIEndpoints, mux, dec, enc, errorHandler(logger), nil)
	spinapiserversvr.Mount(mux, spinAPIServer)
	var handler http.Handler = mux
	if showLog {
		handler = httpmdlwr.Log(middleware.NewLogger(logger))(handler)
	}
	handler = httpmdlwr.RequestID()(handler)

	return handler, nil
}
