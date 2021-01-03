// Code generated by goa v3.2.6, DO NOT EDIT.
//
// Users HTTP server encoders and decoders
//
// Command:
// $ goa gen sample1/design

package server

import (
	"context"
	"io"
	"net/http"
	usersviews "sample1/gen/users/views"
	"unicode/utf8"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeListUserResponse returns an encoder for responses returned by the
// Users list user endpoint.
func EncodeListUserResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(usersviews.Goa2SampleUserCollection)
		enc := encoder(ctx, w)
		body := NewGoa2SampleUserResponseCollection(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// EncodeGetUserResponse returns an encoder for responses returned by the Users
// get user endpoint.
func EncodeGetUserResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*usersviews.Goa2SampleUser)
		ctx = context.WithValue(ctx, goahttp.ContentTypeKey, "application/json")
		enc := encoder(ctx, w)
		body := NewGetUserResponseBody(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeGetUserRequest returns a decoder for requests sent to the Users get
// user endpoint.
func DecodeGetUserRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			id string

			params = mux.Vars(r)
		)
		id = params["id"]
		payload := NewGetUserPayload(id)

		return payload, nil
	}
}

// EncodeCreateUserResponse returns an encoder for responses returned by the
// Users create user endpoint.
func EncodeCreateUserResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(string)
		enc := encoder(ctx, w)
		body := res
		w.WriteHeader(http.StatusCreated)
		return enc.Encode(body)
	}
}

// DecodeCreateUserRequest returns a decoder for requests sent to the Users
// create user endpoint.
func DecodeCreateUserRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body CreateUserRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateCreateUserRequestBody(&body)
		if err != nil {
			return nil, err
		}
		payload := NewCreateUserPayload(&body)

		return payload, nil
	}
}

// EncodeUpdateUserResponse returns an encoder for responses returned by the
// Users update user endpoint.
func EncodeUpdateUserResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*usersviews.Goa2SampleUser)
		ctx = context.WithValue(ctx, goahttp.ContentTypeKey, "application/json")
		enc := encoder(ctx, w)
		body := NewUpdateUserResponseBody(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeUpdateUserRequest returns a decoder for requests sent to the Users
// update user endpoint.
func DecodeUpdateUserRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body UpdateUserRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}

		var (
			id string

			params = mux.Vars(r)
		)
		id = params["id"]
		if utf8.RuneCountInString(id) < 28 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("id", id, utf8.RuneCountInString(id), 28, true))
		}
		if utf8.RuneCountInString(id) > 28 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("id", id, utf8.RuneCountInString(id), 28, false))
		}
		if err != nil {
			return nil, err
		}
		payload := NewUpdateUserPayload(&body, id)

		return payload, nil
	}
}

// EncodeDeleteUserResponse returns an encoder for responses returned by the
// Users delete user endpoint.
func EncodeDeleteUserResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusNoContent)
		return nil
	}
}

// DecodeDeleteUserRequest returns a decoder for requests sent to the Users
// delete user endpoint.
func DecodeDeleteUserRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			id string

			params = mux.Vars(r)
		)
		id = params["id"]
		payload := NewDeleteUserPayload(id)

		return payload, nil
	}
}

// marshalUsersviewsGoa2SampleUserViewToGoa2SampleUserResponse builds a value
// of type *Goa2SampleUserResponse from a value of type
// *usersviews.Goa2SampleUserView.
func marshalUsersviewsGoa2SampleUserViewToGoa2SampleUserResponse(v *usersviews.Goa2SampleUserView) *Goa2SampleUserResponse {
	res := &Goa2SampleUserResponse{
		ID:    *v.ID,
		Name:  *v.Name,
		Email: *v.Email,
	}

	return res
}
