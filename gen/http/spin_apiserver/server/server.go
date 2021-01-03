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
	Mounts              []*MountPoint
	AddVolume           http.Handler
	RemoveVolume        http.Handler
	InfoVolume          http.Handler
	CreateImageOnVolume http.Handler
	DeleteImageOnVolume http.Handler
	ResizeImageOnVolume http.Handler
	InfoImageOnVolume   http.Handler
	MoveImage           http.Handler
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
			{"AddVolume", "POST", "/storage/volume/add"},
			{"RemoveVolume", "POST", "/storage/volume/remove/{volume}"},
			{"InfoVolume", "GET", "/storage/volume/info/{volume}"},
			{"CreateImageOnVolume", "POST", "/storage/volume/image/create"},
			{"DeleteImageOnVolume", "POST", "/storage/volume/image/delete"},
			{"ResizeImageOnVolume", "POST", "/storage/volume/image/resize"},
			{"InfoImageOnVolume", "GET", "/storage/volume/image/info/{volume}/{image_name}"},
			{"MoveImage", "POST", "/storage/volume/image/move"},
		},
		AddVolume:           NewAddVolumeHandler(e.AddVolume, mux, decoder, encoder, errhandler, formatter),
		RemoveVolume:        NewRemoveVolumeHandler(e.RemoveVolume, mux, decoder, encoder, errhandler, formatter),
		InfoVolume:          NewInfoVolumeHandler(e.InfoVolume, mux, decoder, encoder, errhandler, formatter),
		CreateImageOnVolume: NewCreateImageOnVolumeHandler(e.CreateImageOnVolume, mux, decoder, encoder, errhandler, formatter),
		DeleteImageOnVolume: NewDeleteImageOnVolumeHandler(e.DeleteImageOnVolume, mux, decoder, encoder, errhandler, formatter),
		ResizeImageOnVolume: NewResizeImageOnVolumeHandler(e.ResizeImageOnVolume, mux, decoder, encoder, errhandler, formatter),
		InfoImageOnVolume:   NewInfoImageOnVolumeHandler(e.InfoImageOnVolume, mux, decoder, encoder, errhandler, formatter),
		MoveImage:           NewMoveImageHandler(e.MoveImage, mux, decoder, encoder, errhandler, formatter),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "spin-apiserver" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.AddVolume = m(s.AddVolume)
	s.RemoveVolume = m(s.RemoveVolume)
	s.InfoVolume = m(s.InfoVolume)
	s.CreateImageOnVolume = m(s.CreateImageOnVolume)
	s.DeleteImageOnVolume = m(s.DeleteImageOnVolume)
	s.ResizeImageOnVolume = m(s.ResizeImageOnVolume)
	s.InfoImageOnVolume = m(s.InfoImageOnVolume)
	s.MoveImage = m(s.MoveImage)
}

// Mount configures the mux to serve the spin-apiserver endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountAddVolumeHandler(mux, h.AddVolume)
	MountRemoveVolumeHandler(mux, h.RemoveVolume)
	MountInfoVolumeHandler(mux, h.InfoVolume)
	MountCreateImageOnVolumeHandler(mux, h.CreateImageOnVolume)
	MountDeleteImageOnVolumeHandler(mux, h.DeleteImageOnVolume)
	MountResizeImageOnVolumeHandler(mux, h.ResizeImageOnVolume)
	MountInfoImageOnVolumeHandler(mux, h.InfoImageOnVolume)
	MountMoveImageHandler(mux, h.MoveImage)
}

// MountAddVolumeHandler configures the mux to serve the "spin-apiserver"
// service "add_volume" endpoint.
func MountAddVolumeHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/storage/volume/add", f)
}

// NewAddVolumeHandler creates a HTTP handler which loads the HTTP request and
// calls the "spin-apiserver" service "add_volume" endpoint.
func NewAddVolumeHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeAddVolumeRequest(mux, decoder)
		encodeResponse = EncodeAddVolumeResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "add_volume")
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

// MountRemoveVolumeHandler configures the mux to serve the "spin-apiserver"
// service "remove_volume" endpoint.
func MountRemoveVolumeHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/storage/volume/remove/{volume}", f)
}

// NewRemoveVolumeHandler creates a HTTP handler which loads the HTTP request
// and calls the "spin-apiserver" service "remove_volume" endpoint.
func NewRemoveVolumeHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeRemoveVolumeRequest(mux, decoder)
		encodeResponse = EncodeRemoveVolumeResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "remove_volume")
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

// MountInfoVolumeHandler configures the mux to serve the "spin-apiserver"
// service "info_volume" endpoint.
func MountInfoVolumeHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/storage/volume/info/{volume}", f)
}

// NewInfoVolumeHandler creates a HTTP handler which loads the HTTP request and
// calls the "spin-apiserver" service "info_volume" endpoint.
func NewInfoVolumeHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeInfoVolumeRequest(mux, decoder)
		encodeResponse = EncodeInfoVolumeResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "info_volume")
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

// MountCreateImageOnVolumeHandler configures the mux to serve the
// "spin-apiserver" service "create_image_on_volume" endpoint.
func MountCreateImageOnVolumeHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/storage/volume/image/create", f)
}

// NewCreateImageOnVolumeHandler creates a HTTP handler which loads the HTTP
// request and calls the "spin-apiserver" service "create_image_on_volume"
// endpoint.
func NewCreateImageOnVolumeHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeCreateImageOnVolumeRequest(mux, decoder)
		encodeResponse = EncodeCreateImageOnVolumeResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "create_image_on_volume")
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

// MountDeleteImageOnVolumeHandler configures the mux to serve the
// "spin-apiserver" service "delete_image_on_volume" endpoint.
func MountDeleteImageOnVolumeHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/storage/volume/image/delete", f)
}

// NewDeleteImageOnVolumeHandler creates a HTTP handler which loads the HTTP
// request and calls the "spin-apiserver" service "delete_image_on_volume"
// endpoint.
func NewDeleteImageOnVolumeHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeDeleteImageOnVolumeRequest(mux, decoder)
		encodeResponse = EncodeDeleteImageOnVolumeResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "delete_image_on_volume")
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

// MountResizeImageOnVolumeHandler configures the mux to serve the
// "spin-apiserver" service "resize_image_on_volume" endpoint.
func MountResizeImageOnVolumeHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/storage/volume/image/resize", f)
}

// NewResizeImageOnVolumeHandler creates a HTTP handler which loads the HTTP
// request and calls the "spin-apiserver" service "resize_image_on_volume"
// endpoint.
func NewResizeImageOnVolumeHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeResizeImageOnVolumeRequest(mux, decoder)
		encodeResponse = EncodeResizeImageOnVolumeResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "resize_image_on_volume")
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

// MountInfoImageOnVolumeHandler configures the mux to serve the
// "spin-apiserver" service "info_image_on_volume" endpoint.
func MountInfoImageOnVolumeHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/storage/volume/image/info/{volume}/{image_name}", f)
}

// NewInfoImageOnVolumeHandler creates a HTTP handler which loads the HTTP
// request and calls the "spin-apiserver" service "info_image_on_volume"
// endpoint.
func NewInfoImageOnVolumeHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeInfoImageOnVolumeRequest(mux, decoder)
		encodeResponse = EncodeInfoImageOnVolumeResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "info_image_on_volume")
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

// MountMoveImageHandler configures the mux to serve the "spin-apiserver"
// service "move_image" endpoint.
func MountMoveImageHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/storage/volume/image/move", f)
}

// NewMoveImageHandler creates a HTTP handler which loads the HTTP request and
// calls the "spin-apiserver" service "move_image" endpoint.
func NewMoveImageHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeMoveImageRequest(mux, decoder)
		encodeResponse = EncodeMoveImageResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "move_image")
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
