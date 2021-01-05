// Code generated by goa v3.2.6, DO NOT EDIT.
//
// spin-apiserver HTTP server encoders and decoders
//
// Command:
// $ goa gen code.hollensbe.org/erikh/spin/design

package server

import (
	"context"
	"io"
	"net/http"
	"strconv"

	spinapiserver "code.hollensbe.org/erikh/spin/gen/spin_apiserver"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeVMCreateResponse returns an encoder for responses returned by the
// spin-apiserver vm/create endpoint.
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
// spin-apiserver vm/create endpoint.
func DecodeVMCreateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body VMCreateRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateVMCreateRequestBody(&body)
		if err != nil {
			return nil, err
		}
		payload := NewVMCreateVM(&body)

		return payload, nil
	}
}

// EncodeVMDeleteResponse returns an encoder for responses returned by the
// spin-apiserver vm/delete endpoint.
func EncodeVMDeleteResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusOK)
		return nil
	}
}

// DecodeVMDeleteRequest returns a decoder for requests sent to the
// spin-apiserver vm/delete endpoint.
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

// EncodeControlStartResponse returns an encoder for responses returned by the
// spin-apiserver control/start endpoint.
func EncodeControlStartResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusOK)
		return nil
	}
}

// DecodeControlStartRequest returns a decoder for requests sent to the
// spin-apiserver control/start endpoint.
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
// spin-apiserver control/stop endpoint.
func EncodeControlStopResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusOK)
		return nil
	}
}

// DecodeControlStopRequest returns a decoder for requests sent to the
// spin-apiserver control/stop endpoint.
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
// the spin-apiserver control/shutdown endpoint.
func EncodeControlShutdownResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusOK)
		return nil
	}
}

// DecodeControlShutdownRequest returns a decoder for requests sent to the
// spin-apiserver control/shutdown endpoint.
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

// unmarshalStorageRequestBodyToSpinapiserverStorage builds a value of type
// *spinapiserver.Storage from a value of type *StorageRequestBody.
func unmarshalStorageRequestBodyToSpinapiserverStorage(v *StorageRequestBody) *spinapiserver.Storage {
	res := &spinapiserver.Storage{
		Volume:    *v.Volume,
		Image:     *v.Image,
		ImageSize: v.ImageSize,
		Cdrom:     v.Cdrom,
	}

	return res
}
