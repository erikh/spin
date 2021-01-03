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

// BuildAddVolumeRequest instantiates a HTTP request object with method and
// path set to call the "spin-apiserver" service "add_volume" endpoint
func (c *Client) BuildAddVolumeRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: AddVolumeSpinApiserverPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("spin-apiserver", "add_volume", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeAddVolumeRequest returns an encoder for requests sent to the
// spin-apiserver add_volume server.
func EncodeAddVolumeRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*spinapiserver.AddVolumePayload)
		if !ok {
			return goahttp.ErrInvalidType("spin-apiserver", "add_volume", "*spinapiserver.AddVolumePayload", v)
		}
		body := NewAddVolumeRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("spin-apiserver", "add_volume", err)
		}
		return nil
	}
}

// DecodeAddVolumeResponse returns a decoder for responses returned by the
// spin-apiserver add_volume endpoint. restoreBody controls whether the
// response body should be restored after having been read.
func DecodeAddVolumeResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
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
			return nil, goahttp.ErrInvalidResponse("spin-apiserver", "add_volume", resp.StatusCode, string(body))
		}
	}
}

// BuildRemoveVolumeRequest instantiates a HTTP request object with method and
// path set to call the "spin-apiserver" service "remove_volume" endpoint
func (c *Client) BuildRemoveVolumeRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		volume string
	)
	{
		p, ok := v.(*spinapiserver.RemoveVolumePayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("spin-apiserver", "remove_volume", "*spinapiserver.RemoveVolumePayload", v)
		}
		volume = p.Volume
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: RemoveVolumeSpinApiserverPath(volume)}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("spin-apiserver", "remove_volume", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeRemoveVolumeResponse returns a decoder for responses returned by the
// spin-apiserver remove_volume endpoint. restoreBody controls whether the
// response body should be restored after having been read.
func DecodeRemoveVolumeResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
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
			return nil, goahttp.ErrInvalidResponse("spin-apiserver", "remove_volume", resp.StatusCode, string(body))
		}
	}
}

// BuildInfoVolumeRequest instantiates a HTTP request object with method and
// path set to call the "spin-apiserver" service "info_volume" endpoint
func (c *Client) BuildInfoVolumeRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		volume string
	)
	{
		p, ok := v.(*spinapiserver.InfoVolumePayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("spin-apiserver", "info_volume", "*spinapiserver.InfoVolumePayload", v)
		}
		volume = p.Volume
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: InfoVolumeSpinApiserverPath(volume)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("spin-apiserver", "info_volume", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeInfoVolumeResponse returns a decoder for responses returned by the
// spin-apiserver info_volume endpoint. restoreBody controls whether the
// response body should be restored after having been read.
func DecodeInfoVolumeResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
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
			return nil, goahttp.ErrInvalidResponse("spin-apiserver", "info_volume", resp.StatusCode, string(body))
		}
	}
}

// BuildCreateImageOnVolumeRequest instantiates a HTTP request object with
// method and path set to call the "spin-apiserver" service
// "create_image_on_volume" endpoint
func (c *Client) BuildCreateImageOnVolumeRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: CreateImageOnVolumeSpinApiserverPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("spin-apiserver", "create_image_on_volume", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeCreateImageOnVolumeRequest returns an encoder for requests sent to the
// spin-apiserver create_image_on_volume server.
func EncodeCreateImageOnVolumeRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*spinapiserver.CreateImageOnVolumePayload)
		if !ok {
			return goahttp.ErrInvalidType("spin-apiserver", "create_image_on_volume", "*spinapiserver.CreateImageOnVolumePayload", v)
		}
		body := NewCreateImageOnVolumeRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("spin-apiserver", "create_image_on_volume", err)
		}
		return nil
	}
}

// DecodeCreateImageOnVolumeResponse returns a decoder for responses returned
// by the spin-apiserver create_image_on_volume endpoint. restoreBody controls
// whether the response body should be restored after having been read.
func DecodeCreateImageOnVolumeResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
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
			return nil, goahttp.ErrInvalidResponse("spin-apiserver", "create_image_on_volume", resp.StatusCode, string(body))
		}
	}
}

// BuildDeleteImageOnVolumeRequest instantiates a HTTP request object with
// method and path set to call the "spin-apiserver" service
// "delete_image_on_volume" endpoint
func (c *Client) BuildDeleteImageOnVolumeRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: DeleteImageOnVolumeSpinApiserverPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("spin-apiserver", "delete_image_on_volume", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeDeleteImageOnVolumeRequest returns an encoder for requests sent to the
// spin-apiserver delete_image_on_volume server.
func EncodeDeleteImageOnVolumeRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*spinapiserver.DeleteImageOnVolumePayload)
		if !ok {
			return goahttp.ErrInvalidType("spin-apiserver", "delete_image_on_volume", "*spinapiserver.DeleteImageOnVolumePayload", v)
		}
		body := NewDeleteImageOnVolumeRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("spin-apiserver", "delete_image_on_volume", err)
		}
		return nil
	}
}

// DecodeDeleteImageOnVolumeResponse returns a decoder for responses returned
// by the spin-apiserver delete_image_on_volume endpoint. restoreBody controls
// whether the response body should be restored after having been read.
func DecodeDeleteImageOnVolumeResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
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
			return nil, goahttp.ErrInvalidResponse("spin-apiserver", "delete_image_on_volume", resp.StatusCode, string(body))
		}
	}
}

// BuildResizeImageOnVolumeRequest instantiates a HTTP request object with
// method and path set to call the "spin-apiserver" service
// "resize_image_on_volume" endpoint
func (c *Client) BuildResizeImageOnVolumeRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ResizeImageOnVolumeSpinApiserverPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("spin-apiserver", "resize_image_on_volume", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeResizeImageOnVolumeRequest returns an encoder for requests sent to the
// spin-apiserver resize_image_on_volume server.
func EncodeResizeImageOnVolumeRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*spinapiserver.ResizeImageOnVolumePayload)
		if !ok {
			return goahttp.ErrInvalidType("spin-apiserver", "resize_image_on_volume", "*spinapiserver.ResizeImageOnVolumePayload", v)
		}
		body := NewResizeImageOnVolumeRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("spin-apiserver", "resize_image_on_volume", err)
		}
		return nil
	}
}

// DecodeResizeImageOnVolumeResponse returns a decoder for responses returned
// by the spin-apiserver resize_image_on_volume endpoint. restoreBody controls
// whether the response body should be restored after having been read.
func DecodeResizeImageOnVolumeResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
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
			return nil, goahttp.ErrInvalidResponse("spin-apiserver", "resize_image_on_volume", resp.StatusCode, string(body))
		}
	}
}

// BuildInfoImageOnVolumeRequest instantiates a HTTP request object with method
// and path set to call the "spin-apiserver" service "info_image_on_volume"
// endpoint
func (c *Client) BuildInfoImageOnVolumeRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		volume    string
		imageName string
	)
	{
		p, ok := v.(*spinapiserver.InfoImageOnVolumePayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("spin-apiserver", "info_image_on_volume", "*spinapiserver.InfoImageOnVolumePayload", v)
		}
		volume = p.Volume
		imageName = p.ImageName
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: InfoImageOnVolumeSpinApiserverPath(volume, imageName)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("spin-apiserver", "info_image_on_volume", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeInfoImageOnVolumeResponse returns a decoder for responses returned by
// the spin-apiserver info_image_on_volume endpoint. restoreBody controls
// whether the response body should be restored after having been read.
func DecodeInfoImageOnVolumeResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
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
			return nil, goahttp.ErrInvalidResponse("spin-apiserver", "info_image_on_volume", resp.StatusCode, string(body))
		}
	}
}

// BuildMoveImageRequest instantiates a HTTP request object with method and
// path set to call the "spin-apiserver" service "move_image" endpoint
func (c *Client) BuildMoveImageRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: MoveImageSpinApiserverPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("spin-apiserver", "move_image", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeMoveImageRequest returns an encoder for requests sent to the
// spin-apiserver move_image server.
func EncodeMoveImageRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*spinapiserver.MoveImagePayload)
		if !ok {
			return goahttp.ErrInvalidType("spin-apiserver", "move_image", "*spinapiserver.MoveImagePayload", v)
		}
		body := NewMoveImageRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("spin-apiserver", "move_image", err)
		}
		return nil
	}
}

// DecodeMoveImageResponse returns a decoder for responses returned by the
// spin-apiserver move_image endpoint. restoreBody controls whether the
// response body should be restored after having been read.
func DecodeMoveImageResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
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
			return nil, goahttp.ErrInvalidResponse("spin-apiserver", "move_image", resp.StatusCode, string(body))
		}
	}
}
