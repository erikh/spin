// Code generated by goa v3.2.6, DO NOT EDIT.
//
// spin-apiserver HTTP server
//
// Command:
// $ goa gen code.hollensbe.org/erikh/spin/design

package server

import (
	"context"
	"net/http"

	spinapiserver "code.hollensbe.org/erikh/spin/gen/spin_apiserver"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Server lists the spin-apiserver service endpoint HTTP handlers.
type Server struct {
	Mounts          []*MountPoint
	VMCreate        http.Handler
	VMDelete        http.Handler
	ControlStart    http.Handler
	ControlStop     http.Handler
	ControlShutdown http.Handler
}

// ErrorNamer is an interface implemented by generated error structs that
// exposes the name of the error as defined in the design.
type ErrorNamer interface {
	ErrorName() string
}

// MountPoint holds information about the mounted endpoints.
type MountPoint struct {
	// Method is the name of the service method served by the mounted HTTP handler.
	Method string
	// Verb is the HTTP method used to match requests to the mounted handler.
	Verb string
	// Pattern is the HTTP request path pattern used to match requests to the
	// mounted handler.
	Pattern string
}

// New instantiates HTTP handlers for all the spin-apiserver service endpoints
// using the provided encoder and decoder. The handlers are mounted on the
// given mux using the HTTP verb and path defined in the design. errhandler is
// called whenever a response fails to be encoded. formatter is used to format
// errors returned by the service methods prior to encoding. Both errhandler
// and formatter are optional and can be nil.
func New(
	e *spinapiserver.Endpoints,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"VMCreate", "POST", "/vm/create"},
			{"VMDelete", "POST", "/vm/delete/{id}"},
			{"ControlStart", "POST", "/control/start/{id}"},
			{"ControlStop", "POST", "/control/stop/{id}"},
			{"ControlShutdown", "POST", "/control/shutdown/{id}"},
		},
		VMCreate:        NewVMCreateHandler(e.VMCreate, mux, decoder, encoder, errhandler, formatter),
		VMDelete:        NewVMDeleteHandler(e.VMDelete, mux, decoder, encoder, errhandler, formatter),
		ControlStart:    NewControlStartHandler(e.ControlStart, mux, decoder, encoder, errhandler, formatter),
		ControlStop:     NewControlStopHandler(e.ControlStop, mux, decoder, encoder, errhandler, formatter),
		ControlShutdown: NewControlShutdownHandler(e.ControlShutdown, mux, decoder, encoder, errhandler, formatter),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "spin-apiserver" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.VMCreate = m(s.VMCreate)
	s.VMDelete = m(s.VMDelete)
	s.ControlStart = m(s.ControlStart)
	s.ControlStop = m(s.ControlStop)
	s.ControlShutdown = m(s.ControlShutdown)
}

// Mount configures the mux to serve the spin-apiserver endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountVMCreateHandler(mux, h.VMCreate)
	MountVMDeleteHandler(mux, h.VMDelete)
	MountControlStartHandler(mux, h.ControlStart)
	MountControlStopHandler(mux, h.ControlStop)
	MountControlShutdownHandler(mux, h.ControlShutdown)
}

// MountVMCreateHandler configures the mux to serve the "spin-apiserver"
// service "vm_create" endpoint.
func MountVMCreateHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/vm/create", f)
}

// NewVMCreateHandler creates a HTTP handler which loads the HTTP request and
// calls the "spin-apiserver" service "vm_create" endpoint.
func NewVMCreateHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeVMCreateRequest(mux, decoder)
		encodeResponse = EncodeVMCreateResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "vm_create")
		ctx = context.WithValue(ctx, goa.ServiceKey, "spin-apiserver")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountVMDeleteHandler configures the mux to serve the "spin-apiserver"
// service "vm_delete" endpoint.
func MountVMDeleteHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/vm/delete/{id}", f)
}

// NewVMDeleteHandler creates a HTTP handler which loads the HTTP request and
// calls the "spin-apiserver" service "vm_delete" endpoint.
func NewVMDeleteHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeVMDeleteRequest(mux, decoder)
		encodeResponse = EncodeVMDeleteResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "vm_delete")
		ctx = context.WithValue(ctx, goa.ServiceKey, "spin-apiserver")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountControlStartHandler configures the mux to serve the "spin-apiserver"
// service "control_start" endpoint.
func MountControlStartHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/control/start/{id}", f)
}

// NewControlStartHandler creates a HTTP handler which loads the HTTP request
// and calls the "spin-apiserver" service "control_start" endpoint.
func NewControlStartHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeControlStartRequest(mux, decoder)
		encodeResponse = EncodeControlStartResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "control_start")
		ctx = context.WithValue(ctx, goa.ServiceKey, "spin-apiserver")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountControlStopHandler configures the mux to serve the "spin-apiserver"
// service "control_stop" endpoint.
func MountControlStopHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/control/stop/{id}", f)
}

// NewControlStopHandler creates a HTTP handler which loads the HTTP request
// and calls the "spin-apiserver" service "control_stop" endpoint.
func NewControlStopHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeControlStopRequest(mux, decoder)
		encodeResponse = EncodeControlStopResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "control_stop")
		ctx = context.WithValue(ctx, goa.ServiceKey, "spin-apiserver")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountControlShutdownHandler configures the mux to serve the "spin-apiserver"
// service "control_shutdown" endpoint.
func MountControlShutdownHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/control/shutdown/{id}", f)
}

// NewControlShutdownHandler creates a HTTP handler which loads the HTTP
// request and calls the "spin-apiserver" service "control_shutdown" endpoint.
func NewControlShutdownHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeControlShutdownRequest(mux, decoder)
		encodeResponse = EncodeControlShutdownResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "control_shutdown")
		ctx = context.WithValue(ctx, goa.ServiceKey, "spin-apiserver")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}
