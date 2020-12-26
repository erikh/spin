// Code generated by goa v3.2.6, DO NOT EDIT.
//
// spin-broker HTTP server encoders and decoders
//
// Command:
// $ goa gen code.hollensbe.org/erikh/spin/design

package server

import (
	"context"
	"io"
	"net/http"

	spinbroker "code.hollensbe.org/erikh/spin/gen/spin_broker"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeNewResponse returns an encoder for responses returned by the
// spin-broker new endpoint.
func EncodeNewResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(string)
		enc := encoder(ctx, w)
		body := res
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// EncodeAddResponse returns an encoder for responses returned by the
// spin-broker add endpoint.
func EncodeAddResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(string)
		enc := encoder(ctx, w)
		body := res
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeAddRequest returns a decoder for requests sent to the spin-broker add
// endpoint.
func DecodeAddRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body AddRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateAddRequestBody(&body)
		if err != nil {
			return nil, err
		}

		var (
			id string

			params = mux.Vars(r)
		)
		id = params["id"]
		payload := NewAddPayload(&body, id)

		return payload, nil
	}
}

// EncodeEnqueueResponse returns an encoder for responses returned by the
// spin-broker enqueue endpoint.
func EncodeEnqueueResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.([]string)
		enc := encoder(ctx, w)
		body := res
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeEnqueueRequest returns a decoder for requests sent to the spin-broker
// enqueue endpoint.
func DecodeEnqueueRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			id string

			params = mux.Vars(r)
		)
		id = params["id"]
		payload := NewEnqueuePayload(id)

		return payload, nil
	}
}

// EncodeStatusResponse returns an encoder for responses returned by the
// spin-broker status endpoint.
func EncodeStatusResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*spinbroker.StatusResult)
		enc := encoder(ctx, w)
		body := NewStatusResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeStatusRequest returns a decoder for requests sent to the spin-broker
// status endpoint.
func DecodeStatusRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			id string

			params = mux.Vars(r)
		)
		id = params["id"]
		payload := NewStatusPayload(id)

		return payload, nil
	}
}

// EncodeNextResponse returns an encoder for responses returned by the
// spin-broker next endpoint.
func EncodeNextResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*spinbroker.NextResult)
		enc := encoder(ctx, w)
		body := NewNextResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeNextRequest returns a decoder for requests sent to the spin-broker
// next endpoint.
func DecodeNextRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			resource string

			params = mux.Vars(r)
		)
		resource = params["resource"]
		payload := NewNextPayload(resource)

		return payload, nil
	}
}

// EncodeCompleteResponse returns an encoder for responses returned by the
// spin-broker complete endpoint.
func EncodeCompleteResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusOK)
		return nil
	}
}

// DecodeCompleteRequest returns a decoder for requests sent to the spin-broker
// complete endpoint.
func DecodeCompleteRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body CompleteRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateCompleteRequestBody(&body)
		if err != nil {
			return nil, err
		}
		payload := NewCompletePayload(&body)

		return payload, nil
	}
}
