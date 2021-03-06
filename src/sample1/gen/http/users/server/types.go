// Code generated by goa v3.2.6, DO NOT EDIT.
//
// Users HTTP server types
//
// Command:
// $ goa gen sample1/design

package server

import (
	users "sample1/gen/users"
	usersviews "sample1/gen/users/views"
	"unicode/utf8"

	goa "goa.design/goa/v3/pkg"
)

// CreateUserRequestBody is the type of the "Users" service "create user"
// endpoint HTTP request body.
type CreateUserRequestBody struct {
	// User id
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Name of user
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Email of user
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
}

// UpdateUserRequestBody is the type of the "Users" service "update user"
// endpoint HTTP request body.
type UpdateUserRequestBody struct {
	// Name of user
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Email of user
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
}

// Goa3SampleUserResponseCollection is the type of the "Users" service "list
// user" endpoint HTTP response body.
type Goa3SampleUserResponseCollection []*Goa3SampleUserResponse

// GetUserResponseBody is the type of the "Users" service "get user" endpoint
// HTTP response body.
type GetUserResponseBody struct {
	// User id
	ID string `form:"id" json:"id" xml:"id"`
	// Name of user
	Name string `form:"name" json:"name" xml:"name"`
	// Email of user
	Email string `form:"email" json:"email" xml:"email"`
}

// UpdateUserResponseBody is the type of the "Users" service "update user"
// endpoint HTTP response body.
type UpdateUserResponseBody struct {
	// User id
	ID string `form:"id" json:"id" xml:"id"`
	// Name of user
	Name string `form:"name" json:"name" xml:"name"`
	// Email of user
	Email string `form:"email" json:"email" xml:"email"`
}

// Goa3SampleUserResponse is used to define fields on response body types.
type Goa3SampleUserResponse struct {
	// User id
	ID string `form:"id" json:"id" xml:"id"`
	// Name of user
	Name string `form:"name" json:"name" xml:"name"`
	// Email of user
	Email string `form:"email" json:"email" xml:"email"`
}

// NewGoa3SampleUserResponseCollection builds the HTTP response body from the
// result of the "list user" endpoint of the "Users" service.
func NewGoa3SampleUserResponseCollection(res usersviews.Goa3SampleUserCollectionView) Goa3SampleUserResponseCollection {
	body := make([]*Goa3SampleUserResponse, len(res))
	for i, val := range res {
		body[i] = marshalUsersviewsGoa3SampleUserViewToGoa3SampleUserResponse(val)
	}
	return body
}

// NewGetUserResponseBody builds the HTTP response body from the result of the
// "get user" endpoint of the "Users" service.
func NewGetUserResponseBody(res *usersviews.Goa3SampleUserView) *GetUserResponseBody {
	body := &GetUserResponseBody{
		ID:    *res.ID,
		Name:  *res.Name,
		Email: *res.Email,
	}
	return body
}

// NewUpdateUserResponseBody builds the HTTP response body from the result of
// the "update user" endpoint of the "Users" service.
func NewUpdateUserResponseBody(res *usersviews.Goa3SampleUserView) *UpdateUserResponseBody {
	body := &UpdateUserResponseBody{
		ID:    *res.ID,
		Name:  *res.Name,
		Email: *res.Email,
	}
	return body
}

// NewGetUserPayload builds a Users service get user endpoint payload.
func NewGetUserPayload(id string) *users.GetUserPayload {
	v := &users.GetUserPayload{}
	v.ID = id

	return v
}

// NewCreateUserPayload builds a Users service create user endpoint payload.
func NewCreateUserPayload(body *CreateUserRequestBody) *users.CreateUserPayload {
	v := &users.CreateUserPayload{
		ID:    *body.ID,
		Name:  *body.Name,
		Email: *body.Email,
	}

	return v
}

// NewUpdateUserPayload builds a Users service update user endpoint payload.
func NewUpdateUserPayload(body *UpdateUserRequestBody, id string) *users.UpdateUserPayload {
	v := &users.UpdateUserPayload{
		Name:  body.Name,
		Email: body.Email,
	}
	v.ID = id

	return v
}

// NewDeleteUserPayload builds a Users service delete user endpoint payload.
func NewDeleteUserPayload(id string) *users.DeleteUserPayload {
	v := &users.DeleteUserPayload{}
	v.ID = id

	return v
}

// ValidateCreateUserRequestBody runs the validations defined on Create
// UserRequestBody
func ValidateCreateUserRequestBody(body *CreateUserRequestBody) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Email == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("email", "body"))
	}
	if body.ID != nil {
		if utf8.RuneCountInString(*body.ID) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.id", *body.ID, utf8.RuneCountInString(*body.ID), 1, true))
		}
	}
	if body.ID != nil {
		if utf8.RuneCountInString(*body.ID) > 28 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.id", *body.ID, utf8.RuneCountInString(*body.ID), 28, false))
		}
	}
	return
}
