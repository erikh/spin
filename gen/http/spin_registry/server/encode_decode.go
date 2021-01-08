// Code generated by goa v3.2.6, DO NOT EDIT.
//
// spin-registry HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/erikh/spin/design

package server

import (
	"context"
	"io"
	"net/http"
	"strconv"

	spinregistry "github.com/erikh/spin/gen/spin_registry"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeVMCreateResponse returns an encoder for responses returned by the
// spin-registry vm_create endpoint.
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
// spin-registry vm_create endpoint.
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
		payload := NewVMCreateUpdatedVM(&body)

		return payload, nil
	}
}

// EncodeVMUpdateResponse returns an encoder for responses returned by the
// spin-registry vm_update endpoint.
func EncodeVMUpdateResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusOK)
		return nil
	}
}

// DecodeVMUpdateRequest returns a decoder for requests sent to the
// spin-registry vm_update endpoint.
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
		payload := NewVMUpdateUpdateVM(&body, id)

		return payload, nil
	}
}

// EncodeVMDeleteResponse returns an encoder for responses returned by the
// spin-registry vm_delete endpoint.
func EncodeVMDeleteResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusOK)
		return nil
	}
}

// DecodeVMDeleteRequest returns a decoder for requests sent to the
// spin-registry vm_delete endpoint.
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

// EncodeVMGetResponse returns an encoder for responses returned by the
// spin-registry vm_get endpoint.
func EncodeVMGetResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*spinregistry.UpdatedVM)
		enc := encoder(ctx, w)
		body := NewVMGetResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeVMGetRequest returns a decoder for requests sent to the spin-registry
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

// EncodeVMListResponse returns an encoder for responses returned by the
// spin-registry vm_list endpoint.
func EncodeVMListResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.([]uint64)
		enc := encoder(ctx, w)
		body := res
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// EncodeStorageVolumesListResponse returns an encoder for responses returned
// by the spin-registry storage_volumes_list endpoint.
func EncodeStorageVolumesListResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(map[string]string)
		enc := encoder(ctx, w)
		body := res
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// EncodeStorageVolumesCreateResponse returns an encoder for responses returned
// by the spin-registry storage_volumes_create endpoint.
func EncodeStorageVolumesCreateResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusOK)
		return nil
	}
}

// DecodeStorageVolumesCreateRequest returns a decoder for requests sent to the
// spin-registry storage_volumes_create endpoint.
func DecodeStorageVolumesCreateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body StorageVolumesCreateRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateStorageVolumesCreateRequestBody(&body)
		if err != nil {
			return nil, err
		}
		payload := NewStorageVolumesCreatePayload(&body)

		return payload, nil
	}
}

// EncodeStorageVolumesDeleteResponse returns an encoder for responses returned
// by the spin-registry storage_volumes_delete endpoint.
func EncodeStorageVolumesDeleteResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusOK)
		return nil
	}
}

// DecodeStorageVolumesDeleteRequest returns a decoder for requests sent to the
// spin-registry storage_volumes_delete endpoint.
func DecodeStorageVolumesDeleteRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body StorageVolumesDeleteRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateStorageVolumesDeleteRequestBody(&body)
		if err != nil {
			return nil, err
		}
		payload := NewStorageVolumesDeletePayload(&body)

		return payload, nil
	}
}

// EncodeStorageImagesListResponse returns an encoder for responses returned by
// the spin-registry storage_images_list endpoint.
func EncodeStorageImagesListResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.([]string)
		enc := encoder(ctx, w)
		body := res
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeStorageImagesListRequest returns a decoder for requests sent to the
// spin-registry storage_images_list endpoint.
func DecodeStorageImagesListRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body StorageImagesListRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateStorageImagesListRequestBody(&body)
		if err != nil {
			return nil, err
		}
		payload := NewStorageImagesListPayload(&body)

		return payload, nil
	}
}

// EncodeStorageImagesCreateResponse returns an encoder for responses returned
// by the spin-registry storage_images_create endpoint.
func EncodeStorageImagesCreateResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*spinregistry.Image)
		enc := encoder(ctx, w)
		body := NewStorageImagesCreateResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeStorageImagesCreateRequest returns a decoder for requests sent to the
// spin-registry storage_images_create endpoint.
func DecodeStorageImagesCreateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body StorageImagesCreateRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateStorageImagesCreateRequestBody(&body)
		if err != nil {
			return nil, err
		}
		payload := NewStorageImagesCreateStorage(&body)

		return payload, nil
	}
}

// EncodeStorageImagesDeleteResponse returns an encoder for responses returned
// by the spin-registry storage_images_delete endpoint.
func EncodeStorageImagesDeleteResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusOK)
		return nil
	}
}

// DecodeStorageImagesDeleteRequest returns a decoder for requests sent to the
// spin-registry storage_images_delete endpoint.
func DecodeStorageImagesDeleteRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body StorageImagesDeleteRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateStorageImagesDeleteRequestBody(&body)
		if err != nil {
			return nil, err
		}
		payload := NewStorageImagesDeletePayload(&body)

		return payload, nil
	}
}

// EncodeStorageImagesGetResponse returns an encoder for responses returned by
// the spin-registry storage_images_get endpoint.
func EncodeStorageImagesGetResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*spinregistry.Image)
		enc := encoder(ctx, w)
		body := NewStorageImagesGetResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeStorageImagesGetRequest returns a decoder for requests sent to the
// spin-registry storage_images_get endpoint.
func DecodeStorageImagesGetRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body StorageImagesGetRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateStorageImagesGetRequestBody(&body)
		if err != nil {
			return nil, err
		}
		payload := NewStorageImagesGetPayload(&body)

		return payload, nil
	}
}

// unmarshalImageRequestBodyToSpinregistryImage builds a value of type
// *spinregistry.Image from a value of type *ImageRequestBody.
func unmarshalImageRequestBodyToSpinregistryImage(v *ImageRequestBody) *spinregistry.Image {
	res := &spinregistry.Image{
		Path:   *v.Path,
		Cdrom:  *v.Cdrom,
		Volume: v.Volume,
	}

	return res
}

// unmarshalUpdatedVMRequestBodyToSpinregistryUpdatedVM builds a value of type
// *spinregistry.UpdatedVM from a value of type *UpdatedVMRequestBody.
func unmarshalUpdatedVMRequestBodyToSpinregistryUpdatedVM(v *UpdatedVMRequestBody) *spinregistry.UpdatedVM {
	res := &spinregistry.UpdatedVM{
		Name:   *v.Name,
		Cpus:   *v.Cpus,
		Memory: *v.Memory,
	}
	res.Images = make([]*spinregistry.Image, len(v.Images))
	for i, val := range v.Images {
		res.Images[i] = unmarshalImageRequestBodyToSpinregistryImage(val)
	}

	return res
}

// marshalSpinregistryImageToImageResponseBody builds a value of type
// *ImageResponseBody from a value of type *spinregistry.Image.
func marshalSpinregistryImageToImageResponseBody(v *spinregistry.Image) *ImageResponseBody {
	res := &ImageResponseBody{
		Path:   v.Path,
		Cdrom:  v.Cdrom,
		Volume: v.Volume,
	}

	return res
}
