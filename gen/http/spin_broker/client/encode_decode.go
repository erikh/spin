// Code generated by goa v3.2.6, DO NOT EDIT.
//
// spin-broker HTTP client encoders and decoders
//
// Command:
// $ goa gen github.com/erikh/spin/design

package client

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"

	spinbroker "github.com/erikh/spin/gen/spin_broker"
	goahttp "goa.design/goa/v3/http"
)

// BuildNewRequest instantiates a HTTP request object with method and path set
// to call the "spin-broker" service "new" endpoint
func (c *Client) BuildNewRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: NewSpinBrokerPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("spin-broker", "new", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeNewResponse returns a decoder for responses returned by the
// spin-broker new endpoint. restoreBody controls whether the response body
// should be restored after having been read.
func DecodeNewResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body string
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("spin-broker", "new", err)
			}
			return body, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("spin-broker", "new", resp.StatusCode, string(body))
		}
	}
}

// BuildAddRequest instantiates a HTTP request object with method and path set
// to call the "spin-broker" service "add" endpoint
func (c *Client) BuildAddRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		id string
	)
	{
		p, ok := v.(*spinbroker.AddPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("spin-broker", "add", "*spinbroker.AddPayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: AddSpinBrokerPath(id)}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("spin-broker", "add", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeAddRequest returns an encoder for requests sent to the spin-broker add
// server.
func EncodeAddRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*spinbroker.AddPayload)
		if !ok {
			return goahttp.ErrInvalidType("spin-broker", "add", "*spinbroker.AddPayload", v)
		}
		body := NewAddRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("spin-broker", "add", err)
		}
		return nil
	}
}

// DecodeAddResponse returns a decoder for responses returned by the
// spin-broker add endpoint. restoreBody controls whether the response body
// should be restored after having been read.
func DecodeAddResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body string
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("spin-broker", "add", err)
			}
			return body, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("spin-broker", "add", resp.StatusCode, string(body))
		}
	}
}

// BuildEnqueueRequest instantiates a HTTP request object with method and path
// set to call the "spin-broker" service "enqueue" endpoint
func (c *Client) BuildEnqueueRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		id string
	)
	{
		p, ok := v.(*spinbroker.EnqueuePayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("spin-broker", "enqueue", "*spinbroker.EnqueuePayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: EnqueueSpinBrokerPath(id)}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("spin-broker", "enqueue", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeEnqueueResponse returns a decoder for responses returned by the
// spin-broker enqueue endpoint. restoreBody controls whether the response body
// should be restored after having been read.
func DecodeEnqueueResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body []string
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("spin-broker", "enqueue", err)
			}
			return body, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("spin-broker", "enqueue", resp.StatusCode, string(body))
		}
	}
}

// BuildStatusRequest instantiates a HTTP request object with method and path
// set to call the "spin-broker" service "status" endpoint
func (c *Client) BuildStatusRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		id string
	)
	{
		p, ok := v.(*spinbroker.StatusPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("spin-broker", "status", "*spinbroker.StatusPayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: StatusSpinBrokerPath(id)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("spin-broker", "status", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeStatusResponse returns a decoder for responses returned by the
// spin-broker status endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeStatusResponse may return the following errors:
//	- "record_not_found" (type *goa.ServiceError): http.StatusBadRequest
//	- error: internal error
func DecodeStatusResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body StatusResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("spin-broker", "status", err)
			}
			err = ValidateStatusResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("spin-broker", "status", err)
			}
			res := NewStatusResultOK(&body)
			return res, nil
		case http.StatusBadRequest:
			var (
				body StatusRecordNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("spin-broker", "status", err)
			}
			err = ValidateStatusRecordNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("spin-broker", "status", err)
			}
			return nil, NewStatusRecordNotFound(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("spin-broker", "status", resp.StatusCode, string(body))
		}
	}
}

// BuildNextRequest instantiates a HTTP request object with method and path set
// to call the "spin-broker" service "next" endpoint
func (c *Client) BuildNextRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		resource string
	)
	{
		p, ok := v.(*spinbroker.NextPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("spin-broker", "next", "*spinbroker.NextPayload", v)
		}
		resource = p.Resource
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: NextSpinBrokerPath(resource)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("spin-broker", "next", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeNextResponse returns a decoder for responses returned by the
// spin-broker next endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeNextResponse may return the following errors:
//	- "record_not_found" (type *goa.ServiceError): http.StatusBadRequest
//	- error: internal error
func DecodeNextResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body NextResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("spin-broker", "next", err)
			}
			err = ValidateNextResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("spin-broker", "next", err)
			}
			res := NewNextResultOK(&body)
			return res, nil
		case http.StatusBadRequest:
			var (
				body NextRecordNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("spin-broker", "next", err)
			}
			err = ValidateNextRecordNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("spin-broker", "next", err)
			}
			return nil, NewNextRecordNotFound(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("spin-broker", "next", resp.StatusCode, string(body))
		}
	}
}

// BuildCompleteRequest instantiates a HTTP request object with method and path
// set to call the "spin-broker" service "complete" endpoint
func (c *Client) BuildCompleteRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: CompleteSpinBrokerPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("spin-broker", "complete", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeCompleteRequest returns an encoder for requests sent to the
// spin-broker complete server.
func EncodeCompleteRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*spinbroker.CompletePayload)
		if !ok {
			return goahttp.ErrInvalidType("spin-broker", "complete", "*spinbroker.CompletePayload", v)
		}
		body := NewCompleteRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("spin-broker", "complete", err)
		}
		return nil
	}
}

// DecodeCompleteResponse returns a decoder for responses returned by the
// spin-broker complete endpoint. restoreBody controls whether the response
// body should be restored after having been read.
func DecodeCompleteResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			return nil, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("spin-broker", "complete", resp.StatusCode, string(body))
		}
	}
}
