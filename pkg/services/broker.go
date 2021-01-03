package services

import (
	"log"
	"net/http"
	"os"

	"code.hollensbe.org/erikh/spin"
	spinbrokersvr "code.hollensbe.org/erikh/spin/gen/http/spin_broker/server"
	spinbroker "code.hollensbe.org/erikh/spin/gen/spin_broker"
	goahttp "goa.design/goa/v3/http"
	httpmdlwr "goa.design/goa/v3/http/middleware"
	"goa.design/goa/v3/middleware"
)

// Broker returns a http.Handler representing a spin-broker.
func Broker(dbpath string, showLog bool) (http.Handler, error) {
	var logger *log.Logger

	if showLog {
		logger = log.New(os.Stderr, "[spin-broker] ", log.Ltime)
	}

	spinBrokerSvc, err := spin.NewSpinBroker(logger, dbpath)
	if err != nil {
		return nil, err
	}

	spinBrokerEndpoints := spinbroker.NewEndpoints(spinBrokerSvc)
	dec := goahttp.RequestDecoder
	enc := goahttp.ResponseEncoder
	mux := goahttp.NewMuxer()

	spinBrokerServer := spinbrokersvr.New(spinBrokerEndpoints, mux, dec, enc, errorHandler(logger), nil)
	spinbrokersvr.Mount(mux, spinBrokerServer)
	var handler http.Handler = mux
	if showLog {
		handler = httpmdlwr.Log(middleware.NewLogger(logger))(handler)
	}
	handler = httpmdlwr.RequestID()(handler)

	return handler, nil
}
