// Code generated by goa v3.2.6, DO NOT EDIT.
//
// spin-apiserver HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/erikh/spin/design

package server

import (
	"context"
	"io"
	"net/http"
	"strconv"

	"github.com/erikh/spin/pkg/vm"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeVMCreateResponse returns an encoder for responses returned by the
// spin-apiserver vm_create endpoint.
func EncodeVMCreateResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(uint64)
		enc := encoder(ctx, w)
		body := res
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeVMCreateRequest returns a decoder for requests sent to the
// spin-apiserver vm_create endpoint.
func DecodeVMCreateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body *vm.Create
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		payload := body

		return payload, nil
	}
}

// EncodeVMDeleteResponse returns an encoder for responses returned by the
// spin-apiserver vm_delete endpoint.
func EncodeVMDeleteResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusOK)
		return nil
	}
}

// DecodeVMDeleteRequest returns a decoder for requests sent to the
// spin-apiserver vm_delete endpoint.
func DecodeVMDeleteRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			id  uint64
			err error

			params = mux.Vars(r)
		)
		{
			idRaw := params["id"]
			v, err2 := strconv.ParseUint(idRaw, 10, 64)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("id", idRaw, "unsigned integer"))
			}
			id = v
		}
		if err != nil {
			return nil, err
		}
		payload := NewVMDeletePayload(id)

		return payload, nil
	}
}

// EncodeVMListResponse returns an encoder for responses returned by the
// spin-apiserver vm_list endpoint.
func EncodeVMListResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.([]uint64)
		enc := encoder(ctx, w)
		body := res
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// EncodeVMGetResponse returns an encoder for responses returned by the
// spin-apiserver vm_get endpoint.
func EncodeVMGetResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*vm.Transient)
		enc := encoder(ctx, w)
		body := res
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeVMGetRequest returns a decoder for requests sent to the spin-apiserver
// vm_get endpoint.
func DecodeVMGetRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			id  uint64
			err error

			params = mux.Vars(r)
		)
		{
			idRaw := params["id"]
			v, err2 := strconv.ParseUint(idRaw, 10, 64)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("id", idRaw, "unsigned integer"))
			}
			id = v
		}
		if err != nil {
			return nil, err
		}
		payload := NewVMGetPayload(id)

		return payload, nil
	}
}

// EncodeVMUpdateResponse returns an encoder for responses returned by the
// spin-apiserver vm_update endpoint.
func EncodeVMUpdateResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusOK)
		return nil
	}
}

// DecodeVMUpdateRequest returns a decoder for requests sent to the
// spin-apiserver vm_update endpoint.
func DecodeVMUpdateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body VMUpdateRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateVMUpdateRequestBody(&body)
		if err != nil {
			return nil, err
		}

		var (
			id uint64

			params = mux.Vars(r)
		)
		{
			idRaw := params["id"]
			v, err2 := strconv.ParseUint(idRaw, 10, 64)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("id", idRaw, "unsigned integer"))
			}
			id = v
		}
		if err != nil {
			return nil, err
		}
		payload := NewVMUpdatePayload(&body, id)

		return payload, nil
	}
}

// EncodeControlStartResponse returns an encoder for responses returned by the
// spin-apiserver control_start endpoint.
func EncodeControlStartResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusOK)
		return nil
	}
}

// DecodeControlStartRequest returns a decoder for requests sent to the
// spin-apiserver control_start endpoint.
func DecodeControlStartRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			id  uint64
			err error

			params = mux.Vars(r)
		)
		{
			idRaw := params["id"]
			v, err2 := strconv.ParseUint(idRaw, 10, 64)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("id", idRaw, "unsigned integer"))
			}
			id = v
		}
		if err != nil {
			return nil, err
		}
		payload := NewControlStartPayload(id)

		return payload, nil
	}
}

// EncodeControlStopResponse returns an encoder for responses returned by the
// spin-apiserver control_stop endpoint.
func EncodeControlStopResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusOK)
		return nil
	}
}

// DecodeControlStopRequest returns a decoder for requests sent to the
// spin-apiserver control_stop endpoint.
func DecodeControlStopRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			id  uint64
			err error

			params = mux.Vars(r)
		)
		{
			idRaw := params["id"]
			v, err2 := strconv.ParseUint(idRaw, 10, 64)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("id", idRaw, "unsigned integer"))
			}
			id = v
		}
		if err != nil {
			return nil, err
		}
		payload := NewControlStopPayload(id)

		return payload, nil
	}
}

// EncodeControlShutdownResponse returns an encoder for responses returned by
// the spin-apiserver control_shutdown endpoint.
func EncodeControlShutdownResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusOK)
		return nil
	}
}

// DecodeControlShutdownRequest returns a decoder for requests sent to the
// spin-apiserver control_shutdown endpoint.
func DecodeControlShutdownRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			id  uint64
			err error

			params = mux.Vars(r)
		)
		{
			idRaw := params["id"]
			v, err2 := strconv.ParseUint(idRaw, 10, 64)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("id", idRaw, "unsigned integer"))
			}
			id = v
		}
		if err != nil {
			return nil, err
		}
		payload := NewControlShutdownPayload(id)

		return payload, nil
	}
}
