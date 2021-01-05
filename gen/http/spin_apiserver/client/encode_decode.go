// Code generated by goa v3.2.6, DO NOT EDIT.
//
// spin-apiserver HTTP client encoders and decoders
//
// Command:
// $ goa gen code.hollensbe.org/erikh/spin/design

package client

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"

	spinapiserver "code.hollensbe.org/erikh/spin/gen/spin_apiserver"
	goahttp "goa.design/goa/v3/http"
)

// BuildVMCreateRequest instantiates a HTTP request object with method and path
// set to call the "spin-apiserver" service "vm/create" endpoint
func (c *Client) BuildVMCreateRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: VMCreateSpinApiserverPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("spin-apiserver", "vm/create", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeVMCreateRequest returns an encoder for requests sent to the
// spin-apiserver vm/create server.
func EncodeVMCreateRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*spinapiserver.VM)
		if !ok {
			return goahttp.ErrInvalidType("spin-apiserver", "vm/create", "*spinapiserver.VM", v)
		}
		body := NewVMCreateRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("spin-apiserver", "vm/create", err)
		}
		return nil
	}
}

// DecodeVMCreateResponse returns a decoder for responses returned by the
// spin-apiserver vm/create endpoint. restoreBody controls whether the response
// body should be restored after having been read.
func DecodeVMCreateResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
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
				body uint64
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("spin-apiserver", "vm/create", err)
			}
			return body, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("spin-apiserver", "vm/create", resp.StatusCode, string(body))
		}
	}
}

// BuildVMDeleteRequest instantiates a HTTP request object with method and path
// set to call the "spin-apiserver" service "vm/delete" endpoint
func (c *Client) BuildVMDeleteRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		id uint64
	)
	{
		p, ok := v.(*spinapiserver.VMDeletePayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("spin-apiserver", "vm/delete", "*spinapiserver.VMDeletePayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: VMDeleteSpinApiserverPath(id)}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("spin-apiserver", "vm/delete", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeVMDeleteResponse returns a decoder for responses returned by the
// spin-apiserver vm/delete endpoint. restoreBody controls whether the response
// body should be restored after having been read.
func DecodeVMDeleteResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
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
			return nil, goahttp.ErrInvalidResponse("spin-apiserver", "vm/delete", resp.StatusCode, string(body))
		}
	}
}

// BuildControlStartRequest instantiates a HTTP request object with method and
// path set to call the "spin-apiserver" service "control/start" endpoint
func (c *Client) BuildControlStartRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		id uint64
	)
	{
		p, ok := v.(*spinapiserver.ControlStartPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("spin-apiserver", "control/start", "*spinapiserver.ControlStartPayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ControlStartSpinApiserverPath(id)}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("spin-apiserver", "control/start", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeControlStartResponse returns a decoder for responses returned by the
// spin-apiserver control/start endpoint. restoreBody controls whether the
// response body should be restored after having been read.
func DecodeControlStartResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
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
			return nil, goahttp.ErrInvalidResponse("spin-apiserver", "control/start", resp.StatusCode, string(body))
		}
	}
}

// BuildControlStopRequest instantiates a HTTP request object with method and
// path set to call the "spin-apiserver" service "control/stop" endpoint
func (c *Client) BuildControlStopRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		id uint64
	)
	{
		p, ok := v.(*spinapiserver.ControlStopPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("spin-apiserver", "control/stop", "*spinapiserver.ControlStopPayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ControlStopSpinApiserverPath(id)}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("spin-apiserver", "control/stop", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeControlStopResponse returns a decoder for responses returned by the
// spin-apiserver control/stop endpoint. restoreBody controls whether the
// response body should be restored after having been read.
func DecodeControlStopResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
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
			return nil, goahttp.ErrInvalidResponse("spin-apiserver", "control/stop", resp.StatusCode, string(body))
		}
	}
}

// BuildControlShutdownRequest instantiates a HTTP request object with method
// and path set to call the "spin-apiserver" service "control/shutdown" endpoint
func (c *Client) BuildControlShutdownRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		id uint64
	)
	{
		p, ok := v.(*spinapiserver.ControlShutdownPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("spin-apiserver", "control/shutdown", "*spinapiserver.ControlShutdownPayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ControlShutdownSpinApiserverPath(id)}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("spin-apiserver", "control/shutdown", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeControlShutdownResponse returns a decoder for responses returned by
// the spin-apiserver control/shutdown endpoint. restoreBody controls whether
// the response body should be restored after having been read.
func DecodeControlShutdownResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
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
			return nil, goahttp.ErrInvalidResponse("spin-apiserver", "control/shutdown", resp.StatusCode, string(body))
		}
	}
}

// marshalSpinapiserverStorageToStorageRequestBody builds a value of type
// *StorageRequestBody from a value of type *spinapiserver.Storage.
func marshalSpinapiserverStorageToStorageRequestBody(v *spinapiserver.Storage) *StorageRequestBody {
	res := &StorageRequestBody{
		Volume:    v.Volume,
		Image:     v.Image,
		ImageSize: v.ImageSize,
		Cdrom:     v.Cdrom,
	}

	return res
}

// marshalStorageRequestBodyToSpinapiserverStorage builds a value of type
// *spinapiserver.Storage from a value of type *StorageRequestBody.
func marshalStorageRequestBodyToSpinapiserverStorage(v *StorageRequestBody) *spinapiserver.Storage {
	res := &spinapiserver.Storage{
		Volume:    v.Volume,
		Image:     v.Image,
		ImageSize: v.ImageSize,
		Cdrom:     v.Cdrom,
	}

	return res
}
