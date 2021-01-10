// Code generated by goa v3.2.6, DO NOT EDIT.
//
// spin-registry HTTP client encoders and decoders
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

	spinregistry "github.com/erikh/spin/gen/spin_registry"
	"github.com/erikh/spin/pkg/vm"
	goahttp "goa.design/goa/v3/http"
)

// BuildVMCreateRequest instantiates a HTTP request object with method and path
// set to call the "spin-registry" service "vm_create" endpoint
func (c *Client) BuildVMCreateRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: VMCreateSpinRegistryPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("spin-registry", "vm_create", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeVMCreateRequest returns an encoder for requests sent to the
// spin-registry vm_create server.
func EncodeVMCreateRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*vm.Transient)
		if !ok {
			return goahttp.ErrInvalidType("spin-registry", "vm_create", "*vm.Transient", v)
		}
		body := p
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("spin-registry", "vm_create", err)
		}
		return nil
	}
}

// DecodeVMCreateResponse returns a decoder for responses returned by the
// spin-registry vm_create endpoint. restoreBody controls whether the response
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
				return nil, goahttp.ErrDecodingError("spin-registry", "vm_create", err)
			}
			return body, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("spin-registry", "vm_create", resp.StatusCode, string(body))
		}
	}
}

// BuildVMUpdateRequest instantiates a HTTP request object with method and path
// set to call the "spin-registry" service "vm_update" endpoint
func (c *Client) BuildVMUpdateRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		id uint64
	)
	{
		p, ok := v.(*spinregistry.UpdateVM)
		if !ok {
			return nil, goahttp.ErrInvalidType("spin-registry", "vm_update", "*spinregistry.UpdateVM", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: VMUpdateSpinRegistryPath(id)}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("spin-registry", "vm_update", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeVMUpdateRequest returns an encoder for requests sent to the
// spin-registry vm_update server.
func EncodeVMUpdateRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*spinregistry.UpdateVM)
		if !ok {
			return goahttp.ErrInvalidType("spin-registry", "vm_update", "*spinregistry.UpdateVM", v)
		}
		body := NewVMUpdateRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("spin-registry", "vm_update", err)
		}
		return nil
	}
}

// DecodeVMUpdateResponse returns a decoder for responses returned by the
// spin-registry vm_update endpoint. restoreBody controls whether the response
// body should be restored after having been read.
func DecodeVMUpdateResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
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
			return nil, goahttp.ErrInvalidResponse("spin-registry", "vm_update", resp.StatusCode, string(body))
		}
	}
}

// BuildVMDeleteRequest instantiates a HTTP request object with method and path
// set to call the "spin-registry" service "vm_delete" endpoint
func (c *Client) BuildVMDeleteRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		id uint64
	)
	{
		p, ok := v.(*spinregistry.VMDeletePayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("spin-registry", "vm_delete", "*spinregistry.VMDeletePayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: VMDeleteSpinRegistryPath(id)}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("spin-registry", "vm_delete", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeVMDeleteResponse returns a decoder for responses returned by the
// spin-registry vm_delete endpoint. restoreBody controls whether the response
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
			return nil, goahttp.ErrInvalidResponse("spin-registry", "vm_delete", resp.StatusCode, string(body))
		}
	}
}

// BuildVMGetRequest instantiates a HTTP request object with method and path
// set to call the "spin-registry" service "vm_get" endpoint
func (c *Client) BuildVMGetRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		id uint64
	)
	{
		p, ok := v.(*spinregistry.VMGetPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("spin-registry", "vm_get", "*spinregistry.VMGetPayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: VMGetSpinRegistryPath(id)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("spin-registry", "vm_get", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeVMGetResponse returns a decoder for responses returned by the
// spin-registry vm_get endpoint. restoreBody controls whether the response
// body should be restored after having been read.
func DecodeVMGetResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
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
				body *vm.Transient
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("spin-registry", "vm_get", err)
			}
			return body, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("spin-registry", "vm_get", resp.StatusCode, string(body))
		}
	}
}

// BuildVMListRequest instantiates a HTTP request object with method and path
// set to call the "spin-registry" service "vm_list" endpoint
func (c *Client) BuildVMListRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: VMListSpinRegistryPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("spin-registry", "vm_list", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeVMListResponse returns a decoder for responses returned by the
// spin-registry vm_list endpoint. restoreBody controls whether the response
// body should be restored after having been read.
func DecodeVMListResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
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
				body []uint64
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("spin-registry", "vm_list", err)
			}
			return body, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("spin-registry", "vm_list", resp.StatusCode, string(body))
		}
	}
}

// BuildStorageVolumesListRequest instantiates a HTTP request object with
// method and path set to call the "spin-registry" service
// "storage_volumes_list" endpoint
func (c *Client) BuildStorageVolumesListRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: StorageVolumesListSpinRegistryPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("spin-registry", "storage_volumes_list", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeStorageVolumesListResponse returns a decoder for responses returned by
// the spin-registry storage_volumes_list endpoint. restoreBody controls
// whether the response body should be restored after having been read.
func DecodeStorageVolumesListResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
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
				body map[string]string
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("spin-registry", "storage_volumes_list", err)
			}
			return body, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("spin-registry", "storage_volumes_list", resp.StatusCode, string(body))
		}
	}
}

// BuildStorageVolumesCreateRequest instantiates a HTTP request object with
// method and path set to call the "spin-registry" service
// "storage_volumes_create" endpoint
func (c *Client) BuildStorageVolumesCreateRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: StorageVolumesCreateSpinRegistryPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("spin-registry", "storage_volumes_create", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeStorageVolumesCreateRequest returns an encoder for requests sent to
// the spin-registry storage_volumes_create server.
func EncodeStorageVolumesCreateRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*spinregistry.StorageVolumesCreatePayload)
		if !ok {
			return goahttp.ErrInvalidType("spin-registry", "storage_volumes_create", "*spinregistry.StorageVolumesCreatePayload", v)
		}
		body := NewStorageVolumesCreateRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("spin-registry", "storage_volumes_create", err)
		}
		return nil
	}
}

// DecodeStorageVolumesCreateResponse returns a decoder for responses returned
// by the spin-registry storage_volumes_create endpoint. restoreBody controls
// whether the response body should be restored after having been read.
func DecodeStorageVolumesCreateResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
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
			return nil, goahttp.ErrInvalidResponse("spin-registry", "storage_volumes_create", resp.StatusCode, string(body))
		}
	}
}

// BuildStorageVolumesDeleteRequest instantiates a HTTP request object with
// method and path set to call the "spin-registry" service
// "storage_volumes_delete" endpoint
func (c *Client) BuildStorageVolumesDeleteRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: StorageVolumesDeleteSpinRegistryPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("spin-registry", "storage_volumes_delete", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeStorageVolumesDeleteRequest returns an encoder for requests sent to
// the spin-registry storage_volumes_delete server.
func EncodeStorageVolumesDeleteRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*spinregistry.StorageVolumesDeletePayload)
		if !ok {
			return goahttp.ErrInvalidType("spin-registry", "storage_volumes_delete", "*spinregistry.StorageVolumesDeletePayload", v)
		}
		body := NewStorageVolumesDeleteRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("spin-registry", "storage_volumes_delete", err)
		}
		return nil
	}
}

// DecodeStorageVolumesDeleteResponse returns a decoder for responses returned
// by the spin-registry storage_volumes_delete endpoint. restoreBody controls
// whether the response body should be restored after having been read.
func DecodeStorageVolumesDeleteResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
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
			return nil, goahttp.ErrInvalidResponse("spin-registry", "storage_volumes_delete", resp.StatusCode, string(body))
		}
	}
}

// BuildStorageImagesListRequest instantiates a HTTP request object with method
// and path set to call the "spin-registry" service "storage_images_list"
// endpoint
func (c *Client) BuildStorageImagesListRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: StorageImagesListSpinRegistryPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("spin-registry", "storage_images_list", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeStorageImagesListRequest returns an encoder for requests sent to the
// spin-registry storage_images_list server.
func EncodeStorageImagesListRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*spinregistry.StorageImagesListPayload)
		if !ok {
			return goahttp.ErrInvalidType("spin-registry", "storage_images_list", "*spinregistry.StorageImagesListPayload", v)
		}
		body := NewStorageImagesListRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("spin-registry", "storage_images_list", err)
		}
		return nil
	}
}

// DecodeStorageImagesListResponse returns a decoder for responses returned by
// the spin-registry storage_images_list endpoint. restoreBody controls whether
// the response body should be restored after having been read.
func DecodeStorageImagesListResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
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
				return nil, goahttp.ErrDecodingError("spin-registry", "storage_images_list", err)
			}
			return body, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("spin-registry", "storage_images_list", resp.StatusCode, string(body))
		}
	}
}

// BuildStorageImagesCreateRequest instantiates a HTTP request object with
// method and path set to call the "spin-registry" service
// "storage_images_create" endpoint
func (c *Client) BuildStorageImagesCreateRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: StorageImagesCreateSpinRegistryPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("spin-registry", "storage_images_create", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeStorageImagesCreateRequest returns an encoder for requests sent to the
// spin-registry storage_images_create server.
func EncodeStorageImagesCreateRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*vm.Storage)
		if !ok {
			return goahttp.ErrInvalidType("spin-registry", "storage_images_create", "*vm.Storage", v)
		}
		body := p
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("spin-registry", "storage_images_create", err)
		}
		return nil
	}
}

// DecodeStorageImagesCreateResponse returns a decoder for responses returned
// by the spin-registry storage_images_create endpoint. restoreBody controls
// whether the response body should be restored after having been read.
func DecodeStorageImagesCreateResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
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
				body *vm.Image
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("spin-registry", "storage_images_create", err)
			}
			return body, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("spin-registry", "storage_images_create", resp.StatusCode, string(body))
		}
	}
}

// BuildStorageImagesDeleteRequest instantiates a HTTP request object with
// method and path set to call the "spin-registry" service
// "storage_images_delete" endpoint
func (c *Client) BuildStorageImagesDeleteRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: StorageImagesDeleteSpinRegistryPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("spin-registry", "storage_images_delete", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeStorageImagesDeleteRequest returns an encoder for requests sent to the
// spin-registry storage_images_delete server.
func EncodeStorageImagesDeleteRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*spinregistry.StorageImagesDeletePayload)
		if !ok {
			return goahttp.ErrInvalidType("spin-registry", "storage_images_delete", "*spinregistry.StorageImagesDeletePayload", v)
		}
		body := NewStorageImagesDeleteRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("spin-registry", "storage_images_delete", err)
		}
		return nil
	}
}

// DecodeStorageImagesDeleteResponse returns a decoder for responses returned
// by the spin-registry storage_images_delete endpoint. restoreBody controls
// whether the response body should be restored after having been read.
func DecodeStorageImagesDeleteResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
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
			return nil, goahttp.ErrInvalidResponse("spin-registry", "storage_images_delete", resp.StatusCode, string(body))
		}
	}
}

// BuildStorageImagesGetRequest instantiates a HTTP request object with method
// and path set to call the "spin-registry" service "storage_images_get"
// endpoint
func (c *Client) BuildStorageImagesGetRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: StorageImagesGetSpinRegistryPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("spin-registry", "storage_images_get", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeStorageImagesGetRequest returns an encoder for requests sent to the
// spin-registry storage_images_get server.
func EncodeStorageImagesGetRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*spinregistry.StorageImagesGetPayload)
		if !ok {
			return goahttp.ErrInvalidType("spin-registry", "storage_images_get", "*spinregistry.StorageImagesGetPayload", v)
		}
		body := NewStorageImagesGetRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("spin-registry", "storage_images_get", err)
		}
		return nil
	}
}

// DecodeStorageImagesGetResponse returns a decoder for responses returned by
// the spin-registry storage_images_get endpoint. restoreBody controls whether
// the response body should be restored after having been read.
func DecodeStorageImagesGetResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
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
				body *vm.Image
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("spin-registry", "storage_images_get", err)
			}
			return body, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("spin-registry", "storage_images_get", resp.StatusCode, string(body))
		}
	}
}
