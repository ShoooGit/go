// Code generated by goa v3.2.6, DO NOT EDIT.
//
// Admin HTTP client types
//
// Command:
// $ goa gen sample1/design

package client

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
	ID string `form:"id" json:"id" xml:"id"`
	// Name of user
	Name string `form:"name" json:"name" xml:"name"`
	// Email of user
	Email string `form:"email" json:"email" xml:"email"`
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
	Data []*DataResponseBody `form:"data,omitempty" json:"data,omitempty" xml:"data,omitempty"`
	// X軸に使用するkey
	X *string `form:"x,omitempty" json:"x,omitempty" xml:"x,omitempty"`
	// Y軸に使用するkey
	Y *string `form:"y,omitempty" json:"y,omitempty" xml:"y,omitempty"`
	// ドットの大きさに使用するkey
	Size *string `form:"size,omitempty" json:"size,omitempty" xml:"size,omitempty"`
	// ドットの色分けに使用するkey
	Color *string                     `form:"color,omitempty" json:"color,omitempty" xml:"color,omitempty"`
	Guide *StatsGuideTypeResponseBody `form:"guide,omitempty" json:"guide,omitempty" xml:"guide,omitempty"`
}

// AdminListUserResponseBody is the type of the "Admin" service "admin list
// user" endpoint HTTP response body.
type AdminListUserResponseBody []*Goa2SampleAdminUserResponse

// AdminGetUserResponseBody is the type of the "Admin" service "admin get user"
// endpoint HTTP response body.
type AdminGetUserResponseBody struct {
	// User id
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Name of user
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Email of user
	Email     *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	CreatedAt *string `form:"created_at,omitempty" json:"created_at,omitempty" xml:"created_at,omitempty"`
	UpdatedAt *string `form:"updated_at,omitempty" json:"updated_at,omitempty" xml:"updated_at,omitempty"`
	DeletedAt *string `form:"deleted_at,omitempty" json:"deleted_at,omitempty" xml:"deleted_at,omitempty"`
}

// AdminUpdateUserResponseBody is the type of the "Admin" service "admin update
// user" endpoint HTTP response body.
type AdminUpdateUserResponseBody struct {
	// User id
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Name of user
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Email of user
	Email     *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	CreatedAt *string `form:"created_at,omitempty" json:"created_at,omitempty" xml:"created_at,omitempty"`
	UpdatedAt *string `form:"updated_at,omitempty" json:"updated_at,omitempty" xml:"updated_at,omitempty"`
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
	Label *string `form:"label,omitempty" json:"label,omitempty" xml:"label,omitempty"`
}

// Goa2SampleAdminUserResponse is used to define fields on response body types.
type Goa2SampleAdminUserResponse struct {
	// User id
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Name of user
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Email of user
	Email     *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	CreatedAt *string `form:"created_at,omitempty" json:"created_at,omitempty" xml:"created_at,omitempty"`
	UpdatedAt *string `form:"updated_at,omitempty" json:"updated_at,omitempty" xml:"updated_at,omitempty"`
	DeletedAt *string `form:"deleted_at,omitempty" json:"deleted_at,omitempty" xml:"deleted_at,omitempty"`
}

// NewAdminCreateUserRequestBody builds the HTTP request body from the payload
// of the "admin create user" endpoint of the "Admin" service.
func NewAdminCreateUserRequestBody(p *admin.CreateUserPayload) *AdminCreateUserRequestBody {
	body := &AdminCreateUserRequestBody{
		ID:    p.ID,
		Name:  p.Name,
		Email: p.Email,
	}
	return body
}

// NewAdminUpdateUserRequestBody builds the HTTP request body from the payload
// of the "admin update user" endpoint of the "Admin" service.
func NewAdminUpdateUserRequestBody(p *admin.UpdateUserPayload) *AdminUpdateUserRequestBody {
	body := &AdminUpdateUserRequestBody{
		Name:  p.Name,
		Email: p.Email,
	}
	return body
}

// NewUserNumberGoa2SampleAdminUserNumberOK builds a "Admin" service
// "user_number" endpoint result from a HTTP "OK" response.
func NewUserNumberGoa2SampleAdminUserNumberOK(body *UserNumberResponseBody) *adminviews.Goa2SampleAdminUserNumberView {
	v := &adminviews.Goa2SampleAdminUserNumberView{
		X:     body.X,
		Y:     body.Y,
		Size:  body.Size,
		Color: body.Color,
	}
	v.Data = make([]*adminviews.DataView, len(body.Data))
	for i, val := range body.Data {
		v.Data[i] = unmarshalDataResponseBodyToAdminviewsDataView(val)
	}
	v.Guide = unmarshalStatsGuideTypeResponseBodyToAdminviewsStatsGuideTypeView(body.Guide)

	return v
}

// NewAdminListUserGoa2SampleAdminUserCollectionOK builds a "Admin" service
// "admin list user" endpoint result from a HTTP "OK" response.
func NewAdminListUserGoa2SampleAdminUserCollectionOK(body AdminListUserResponseBody) adminviews.Goa2SampleAdminUserCollectionView {
	v := make([]*adminviews.Goa2SampleAdminUserView, len(body))
	for i, val := range body {
		v[i] = unmarshalGoa2SampleAdminUserResponseToAdminviewsGoa2SampleAdminUserView(val)
	}
	return v
}

// NewAdminGetUserGoa2SampleAdminUserOK builds a "Admin" service "admin get
// user" endpoint result from a HTTP "OK" response.
func NewAdminGetUserGoa2SampleAdminUserOK(body *AdminGetUserResponseBody) *adminviews.Goa2SampleAdminUserView {
	v := &adminviews.Goa2SampleAdminUserView{
		ID:        body.ID,
		Name:      body.Name,
		Email:     body.Email,
		CreatedAt: body.CreatedAt,
		UpdatedAt: body.UpdatedAt,
		DeletedAt: body.DeletedAt,
	}

	return v
}

// NewAdminUpdateUserGoa2SampleAdminUserOK builds a "Admin" service "admin
// update user" endpoint result from a HTTP "OK" response.
func NewAdminUpdateUserGoa2SampleAdminUserOK(body *AdminUpdateUserResponseBody) *adminviews.Goa2SampleAdminUserView {
	v := &adminviews.Goa2SampleAdminUserView{
		ID:        body.ID,
		Name:      body.Name,
		Email:     body.Email,
		CreatedAt: body.CreatedAt,
		UpdatedAt: body.UpdatedAt,
		DeletedAt: body.DeletedAt,
	}

	return v
}

// ValidateStatsGuideTypeResponseBody runs the validations defined on
// StatsGuideTypeResponseBody
func ValidateStatsGuideTypeResponseBody(body *StatsGuideTypeResponseBody) (err error) {
	if body.X != nil {
		if err2 := ValidateStatsLabelTypeResponseBody(body.X); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if body.Y != nil {
		if err2 := ValidateStatsLabelTypeResponseBody(body.Y); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateStatsLabelTypeResponseBody runs the validations defined on
// StatsLabelTypeResponseBody
func ValidateStatsLabelTypeResponseBody(body *StatsLabelTypeResponseBody) (err error) {
	if body.Label == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("label", "body"))
	}
	return
}

// ValidateGoa2SampleAdminUserResponse runs the validations defined on
// Goa2-SampleAdminUserResponse
func ValidateGoa2SampleAdminUserResponse(body *Goa2SampleAdminUserResponse) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Email == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("email", "body"))
	}
	if body.CreatedAt == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("created_at", "body"))
	}
	if body.UpdatedAt == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("updated_at", "body"))
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
