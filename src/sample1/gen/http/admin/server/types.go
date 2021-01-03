// Code generated by goa v3.2.6, DO NOT EDIT.
//
// Admin HTTP server types
//
// Command:
// $ goa gen sample1/design

package server

import (
	admin "sample1/gen/admin"
	adminviews "sample1/gen/admin/views"
	"unicode/utf8"

	goa "goa.design/goa/v3/pkg"
)

// AdminCreateUserRequestBody is the type of the "Admin" service "admin create
// user" endpoint HTTP request body.
type AdminCreateUserRequestBody struct {
	// User id
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Name of user
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Email of user
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
}

// AdminUpdateUserRequestBody is the type of the "Admin" service "admin update
// user" endpoint HTTP request body.
type AdminUpdateUserRequestBody struct {
	// Name of user
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Email of user
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
}

// UserNumberResponseBody is the type of the "Admin" service "user_number"
// endpoint HTTP response body.
type UserNumberResponseBody struct {
	// グラフデータ
	Data []*DataResponseBody `form:"data" json:"data" xml:"data"`
	// X軸に使用するkey
	X string `form:"x" json:"x" xml:"x"`
	// Y軸に使用するkey
	Y string `form:"y" json:"y" xml:"y"`
	// ドットの大きさに使用するkey
	Size *string `form:"size,omitempty" json:"size,omitempty" xml:"size,omitempty"`
	// ドットの色分けに使用するkey
	Color *string                     `form:"color,omitempty" json:"color,omitempty" xml:"color,omitempty"`
	Guide *StatsGuideTypeResponseBody `form:"guide" json:"guide" xml:"guide"`
}

// Goa3SampleAdminUserResponseCollection is the type of the "Admin" service
// "admin list user" endpoint HTTP response body.
type Goa3SampleAdminUserResponseCollection []*Goa3SampleAdminUserResponse

// AdminGetUserResponseBody is the type of the "Admin" service "admin get user"
// endpoint HTTP response body.
type AdminGetUserResponseBody struct {
	// User id
	ID string `form:"id" json:"id" xml:"id"`
	// Name of user
	Name string `form:"name" json:"name" xml:"name"`
	// Email of user
	Email     string  `form:"email" json:"email" xml:"email"`
	CreatedAt string  `form:"created_at" json:"created_at" xml:"created_at"`
	UpdatedAt string  `form:"updated_at" json:"updated_at" xml:"updated_at"`
	DeletedAt *string `form:"deleted_at,omitempty" json:"deleted_at,omitempty" xml:"deleted_at,omitempty"`
}

// AdminUpdateUserResponseBody is the type of the "Admin" service "admin update
// user" endpoint HTTP response body.
type AdminUpdateUserResponseBody struct {
	// User id
	ID string `form:"id" json:"id" xml:"id"`
	// Name of user
	Name string `form:"name" json:"name" xml:"name"`
	// Email of user
	Email     string  `form:"email" json:"email" xml:"email"`
	CreatedAt string  `form:"created_at" json:"created_at" xml:"created_at"`
	UpdatedAt string  `form:"updated_at" json:"updated_at" xml:"updated_at"`
	DeletedAt *string `form:"deleted_at,omitempty" json:"deleted_at,omitempty" xml:"deleted_at,omitempty"`
}

// DataResponseBody is used to define fields on response body types.
type DataResponseBody struct {
	Key   *string     `form:"key,omitempty" json:"key,omitempty" xml:"key,omitempty"`
	Value interface{} `form:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
}

// StatsGuideTypeResponseBody is used to define fields on response body types.
type StatsGuideTypeResponseBody struct {
	X *StatsLabelTypeResponseBody `form:"x,omitempty" json:"x,omitempty" xml:"x,omitempty"`
	Y *StatsLabelTypeResponseBody `form:"y,omitempty" json:"y,omitempty" xml:"y,omitempty"`
}

// StatsLabelTypeResponseBody is used to define fields on response body types.
type StatsLabelTypeResponseBody struct {
	Label string `form:"label" json:"label" xml:"label"`
}

// Goa3SampleAdminUserResponse is used to define fields on response body types.
type Goa3SampleAdminUserResponse struct {
	// User id
	ID string `form:"id" json:"id" xml:"id"`
	// Name of user
	Name string `form:"name" json:"name" xml:"name"`
	// Email of user
	Email     string  `form:"email" json:"email" xml:"email"`
	CreatedAt string  `form:"created_at" json:"created_at" xml:"created_at"`
	UpdatedAt string  `form:"updated_at" json:"updated_at" xml:"updated_at"`
	DeletedAt *string `form:"deleted_at,omitempty" json:"deleted_at,omitempty" xml:"deleted_at,omitempty"`
}

// NewUserNumberResponseBody builds the HTTP response body from the result of
// the "user_number" endpoint of the "Admin" service.
func NewUserNumberResponseBody(res *adminviews.Goa3SampleAdminUserNumberView) *UserNumberResponseBody {
	body := &UserNumberResponseBody{
		X:     *res.X,
		Y:     *res.Y,
		Size:  res.Size,
		Color: res.Color,
	}
	if res.Data != nil {
		body.Data = make([]*DataResponseBody, len(res.Data))
		for i, val := range res.Data {
			body.Data[i] = marshalAdminviewsDataViewToDataResponseBody(val)
		}
	}
	if res.Guide != nil {
		body.Guide = marshalAdminviewsStatsGuideTypeViewToStatsGuideTypeResponseBody(res.Guide)
	}
	return body
}

// NewGoa3SampleAdminUserResponseCollection builds the HTTP response body from
// the result of the "admin list user" endpoint of the "Admin" service.
func NewGoa3SampleAdminUserResponseCollection(res adminviews.Goa3SampleAdminUserCollectionView) Goa3SampleAdminUserResponseCollection {
	body := make([]*Goa3SampleAdminUserResponse, len(res))
	for i, val := range res {
		body[i] = marshalAdminviewsGoa3SampleAdminUserViewToGoa3SampleAdminUserResponse(val)
	}
	return body
}

// NewAdminGetUserResponseBody builds the HTTP response body from the result of
// the "admin get user" endpoint of the "Admin" service.
func NewAdminGetUserResponseBody(res *adminviews.Goa3SampleAdminUserView) *AdminGetUserResponseBody {
	body := &AdminGetUserResponseBody{
		ID:        *res.ID,
		Name:      *res.Name,
		Email:     *res.Email,
		CreatedAt: *res.CreatedAt,
		UpdatedAt: *res.UpdatedAt,
		DeletedAt: res.DeletedAt,
	}
	return body
}

// NewAdminUpdateUserResponseBody builds the HTTP response body from the result
// of the "admin update user" endpoint of the "Admin" service.
func NewAdminUpdateUserResponseBody(res *adminviews.Goa3SampleAdminUserView) *AdminUpdateUserResponseBody {
	body := &AdminUpdateUserResponseBody{
		ID:        *res.ID,
		Name:      *res.Name,
		Email:     *res.Email,
		CreatedAt: *res.CreatedAt,
		UpdatedAt: *res.UpdatedAt,
		DeletedAt: res.DeletedAt,
	}
	return body
}

// NewAdminGetUserGetUserPayload builds a Admin service admin get user endpoint
// payload.
func NewAdminGetUserGetUserPayload(id string) *admin.GetUserPayload {
	v := &admin.GetUserPayload{}
	v.ID = id

	return v
}

// NewAdminCreateUserCreateUserPayload builds a Admin service admin create user
// endpoint payload.
func NewAdminCreateUserCreateUserPayload(body *AdminCreateUserRequestBody) *admin.CreateUserPayload {
	v := &admin.CreateUserPayload{
		ID:    *body.ID,
		Name:  *body.Name,
		Email: *body.Email,
	}

	return v
}

// NewAdminUpdateUserUpdateUserPayload builds a Admin service admin update user
// endpoint payload.
func NewAdminUpdateUserUpdateUserPayload(body *AdminUpdateUserRequestBody, id string) *admin.UpdateUserPayload {
	v := &admin.UpdateUserPayload{
		Name:  body.Name,
		Email: body.Email,
	}
	v.ID = id

	return v
}

// NewAdminDeleteUserDeleteUserPayload builds a Admin service admin delete user
// endpoint payload.
func NewAdminDeleteUserDeleteUserPayload(id string) *admin.DeleteUserPayload {
	v := &admin.DeleteUserPayload{}
	v.ID = id

	return v
}

// ValidateAdminCreateUserRequestBody runs the validations defined on Admin
// Create UserRequestBody
func ValidateAdminCreateUserRequestBody(body *AdminCreateUserRequestBody) (err error) {
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
		if utf8.RuneCountInString(*body.ID) < 28 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.id", *body.ID, utf8.RuneCountInString(*body.ID), 28, true))
		}
	}
	if body.ID != nil {
		if utf8.RuneCountInString(*body.ID) > 28 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.id", *body.ID, utf8.RuneCountInString(*body.ID), 28, false))
		}
	}
	return
}
