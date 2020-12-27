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

func Broker(dbpath string) (http.Handler, error) {
	logger := log.New(os.Stderr, "[spin] ", log.Ltime)

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
	handler = httpmdlwr.Log(middleware.NewLogger(logger))(handler)
	handler = httpmdlwr.RequestID()(handler)

	return handler, nil
}