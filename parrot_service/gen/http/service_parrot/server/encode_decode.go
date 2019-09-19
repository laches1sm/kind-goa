// Code generated by goa v3.0.6, DO NOT EDIT.
//
// service-parrot HTTP server encoders and decoders
//
// Command:
// $ goa gen parrot_service/design

package server

import (
	"context"
	"io"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodePostparrotResponse returns an encoder for responses returned by the
// service-parrot postparrot endpoint.
func EncodePostparrotResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusCreated)
		return nil
	}
}

// DecodePostparrotRequest returns a decoder for requests sent to the
// service-parrot postparrot endpoint.
func DecodePostparrotRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		payload := NewPostparrotParrotpayload()

		return payload, nil
	}
}

// EncodeListaparrotResponse returns an encoder for responses returned by the
// service-parrot listaparrot endpoint.
func EncodeListaparrotResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusOK)
		return nil
	}
}

// DecodeListaparrotRequest returns a decoder for requests sent to the
// service-parrot listaparrot endpoint.
func DecodeListaparrotRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body ListaparrotRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateListaparrotRequestBody(&body)
		if err != nil {
			return nil, err
		}
		payload := NewListaparrotPayload(&body)

		return payload, nil
	}
}

// EncodeListallparrotsResponse returns an encoder for responses returned by
// the service-parrot listallparrots endpoint.
func EncodeListallparrotsResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusOK)
		return nil
	}
}
