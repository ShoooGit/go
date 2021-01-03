// Code generated by goa v3.2.6, DO NOT EDIT.
//
// Users service
//
// Command:
// $ goa gen sample1/design

package users

import (
	"context"
	usersviews "sample1/gen/users/views"
)

// users serves user account relative information.
type Service interface {
	// List all stored users
	ListUser(context.Context) (res Goa2SampleUserCollection, err error)
	// Show user by ID
	GetUser(context.Context, *GetUserPayload) (res *Goa2SampleUser, err error)
	// Add new user and return its ID.
	CreateUser(context.Context, *CreateUserPayload) (res string, err error)
	// Update user item.
	UpdateUser(context.Context, *UpdateUserPayload) (res *Goa2SampleUser, err error)
	// Delete user by id.
	DeleteUser(context.Context, *DeleteUserPayload) (err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "Users"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [5]string{"list user", "get user", "create user", "update user", "delete user"}

// Goa2SampleUserCollection is the result type of the Users service list user
// method.
type Goa2SampleUserCollection []*Goa2SampleUser

// GetUserPayload is the payload type of the Users service get user method.
type GetUserPayload struct {
	ID string
}

// Goa2SampleUser is the result type of the Users service get user method.
type Goa2SampleUser struct {
	// User id
	ID string
	// Name of user
	Name string
	// Email of user
	Email string
}

// CreateUserPayload is the payload type of the Users service create user
// method.
type CreateUserPayload struct {
	// User id
	ID string
	// Name of user
	Name string
	// Email of user
	Email string
}

// UpdateUserPayload is the payload type of the Users service update user
// method.
type UpdateUserPayload struct {
	// User id
	ID string
	// Name of user
	Name *string
	// Email of user
	Email *string
}

// DeleteUserPayload is the payload type of the Users service delete user
// method.
type DeleteUserPayload struct {
	ID string
}

// NewGoa2SampleUserCollection initializes result type Goa2SampleUserCollection
// from viewed result type Goa2SampleUserCollection.
func NewGoa2SampleUserCollection(vres usersviews.Goa2SampleUserCollection) Goa2SampleUserCollection {
	return newGoa2SampleUserCollection(vres.Projected)
}

// NewViewedGoa2SampleUserCollection initializes viewed result type
// Goa2SampleUserCollection from result type Goa2SampleUserCollection using the
// given view.
func NewViewedGoa2SampleUserCollection(res Goa2SampleUserCollection, view string) usersviews.Goa2SampleUserCollection {
	p := newGoa2SampleUserCollectionView(res)
	return usersviews.Goa2SampleUserCollection{Projected: p, View: "default"}
}

// NewGoa2SampleUser initializes result type Goa2SampleUser from viewed result
// type Goa2SampleUser.
func NewGoa2SampleUser(vres *usersviews.Goa2SampleUser) *Goa2SampleUser {
	return newGoa2SampleUser(vres.Projected)
}

// NewViewedGoa2SampleUser initializes viewed result type Goa2SampleUser from
// result type Goa2SampleUser using the given view.
func NewViewedGoa2SampleUser(res *Goa2SampleUser, view string) *usersviews.Goa2SampleUser {
	p := newGoa2SampleUserView(res)
	return &usersviews.Goa2SampleUser{Projected: p, View: "default"}
}

// newGoa2SampleUserCollection converts projected type Goa2SampleUserCollection
// to service type Goa2SampleUserCollection.
func newGoa2SampleUserCollection(vres usersviews.Goa2SampleUserCollectionView) Goa2SampleUserCollection {
	res := make(Goa2SampleUserCollection, len(vres))
	for i, n := range vres {
		res[i] = newGoa2SampleUser(n)
	}
	return res
}

// newGoa2SampleUserCollectionView projects result type
// Goa2SampleUserCollection to projected type Goa2SampleUserCollectionView
// using the "default" view.
func newGoa2SampleUserCollectionView(res Goa2SampleUserCollection) usersviews.Goa2SampleUserCollectionView {
	vres := make(usersviews.Goa2SampleUserCollectionView, len(res))
	for i, n := range res {
		vres[i] = newGoa2SampleUserView(n)
	}
	return vres
}

// newGoa2SampleUser converts projected type Goa2SampleUser to service type
// Goa2SampleUser.
func newGoa2SampleUser(vres *usersviews.Goa2SampleUserView) *Goa2SampleUser {
	res := &Goa2SampleUser{}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	if vres.Name != nil {
		res.Name = *vres.Name
	}
	if vres.Email != nil {
		res.Email = *vres.Email
	}
	return res
}

// newGoa2SampleUserView projects result type Goa2SampleUser to projected type
// Goa2SampleUserView using the "default" view.
func newGoa2SampleUserView(res *Goa2SampleUser) *usersviews.Goa2SampleUserView {
	vres := &usersviews.Goa2SampleUserView{
		ID:    &res.ID,
		Name:  &res.Name,
		Email: &res.Email,
	}
	return vres
}
