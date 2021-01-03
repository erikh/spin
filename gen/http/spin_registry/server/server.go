// Code generated by goa v3.2.6, DO NOT EDIT.
//
// spin-registry HTTP server
//
// Command:
// $ goa gen code.hollensbe.org/erikh/spin/design

package server

import (
	"context"
	"net/http"

	spinregistry "code.hollensbe.org/erikh/spin/gen/spin_registry"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Server lists the spin-registry service endpoint HTTP handlers.
type Server struct {
	Mounts   []*MountPoint
	VMCreate http.Handler
	VMUpdate http.Handler
	VMDelete http.Handler
	VMGet    http.Handler
	VMList   http.Handler
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

// New instantiates HTTP handlers for all the spin-registry service endpoints
// using the provided encoder and decoder. The handlers are mounted on the
// given mux using the HTTP verb and path defined in the design. errhandler is
// called whenever a response fails to be encoded. formatter is used to format
// errors returned by the service methods prior to encoding. Both errhandler
// and formatter are optional and can be nil.
func New(
	e *spinregistry.Endpoints,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"VMCreate", "POST", "/vm/create"},
			{"VMUpdate", "POST", "/vm/update/{id}"},
			{"VMDelete", "POST", "/vm/delete/{id}"},
			{"VMGet", "GET", "/vm/get/{id}"},
			{"VMList", "GET", "/vm/list"},
		},
		VMCreate: NewVMCreateHandler(e.VMCreate, mux, decoder, encoder, errhandler, formatter),
		VMUpdate: NewVMUpdateHandler(e.VMUpdate, mux, decoder, encoder, errhandler, formatter),
		VMDelete: NewVMDeleteHandler(e.VMDelete, mux, decoder, encoder, errhandler, formatter),
		VMGet:    NewVMGetHandler(e.VMGet, mux, decoder, encoder, errhandler, formatter),
		VMList:   NewVMListHandler(e.VMList, mux, decoder, encoder, errhandler, formatter),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "spin-registry" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.VMCreate = m(s.VMCreate)
	s.VMUpdate = m(s.VMUpdate)
	s.VMDelete = m(s.VMDelete)
	s.VMGet = m(s.VMGet)
	s.VMList = m(s.VMList)
}

// Mount configures the mux to serve the spin-registry endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountVMCreateHandler(mux, h.VMCreate)
	MountVMUpdateHandler(mux, h.VMUpdate)
	MountVMDeleteHandler(mux, h.VMDelete)
	MountVMGetHandler(mux, h.VMGet)
	MountVMListHandler(mux, h.VMList)
}

// MountVMCreateHandler configures the mux to serve the "spin-registry" service
// "vm/create" endpoint.
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
// calls the "spin-registry" service "vm/create" endpoint.
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
		ctx = context.WithValue(ctx, goa.MethodKey, "vm/create")
		ctx = context.WithValue(ctx, goa.ServiceKey, "spin-registry")
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

// MountVMUpdateHandler configures the mux to serve the "spin-registry" service
// "vm/update" endpoint.
func MountVMUpdateHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/vm/update/{id}", f)
}

// NewVMUpdateHandler creates a HTTP handler which loads the HTTP request and
// calls the "spin-registry" service "vm/update" endpoint.
func NewVMUpdateHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeVMUpdateRequest(mux, decoder)
		encodeResponse = EncodeVMUpdateResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "vm/update")
		ctx = context.WithValue(ctx, goa.ServiceKey, "spin-registry")
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

// MountVMDeleteHandler configures the mux to serve the "spin-registry" service
// "vm/delete" endpoint.
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
// calls the "spin-registry" service "vm/delete" endpoint.
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
		ctx = context.WithValue(ctx, goa.MethodKey, "vm/delete")
		ctx = context.WithValue(ctx, goa.ServiceKey, "spin-registry")
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

// MountVMGetHandler configures the mux to serve the "spin-registry" service
// "vm/get" endpoint.
func MountVMGetHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/vm/get/{id}", f)
}

// NewVMGetHandler creates a HTTP handler which loads the HTTP request and
// calls the "spin-registry" service "vm/get" endpoint.
func NewVMGetHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeVMGetRequest(mux, decoder)
		encodeResponse = EncodeVMGetResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "vm/get")
		ctx = context.WithValue(ctx, goa.ServiceKey, "spin-registry")
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

// MountVMListHandler configures the mux to serve the "spin-registry" service
// "vm/list" endpoint.
func MountVMListHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/vm/list", f)
}

// NewVMListHandler creates a HTTP handler which loads the HTTP request and
// calls the "spin-registry" service "vm/list" endpoint.
func NewVMListHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		encodeResponse = EncodeVMListResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "vm/list")
		ctx = context.WithValue(ctx, goa.ServiceKey, "spin-registry")
		var err error
		res, err := endpoint(ctx, nil)
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
