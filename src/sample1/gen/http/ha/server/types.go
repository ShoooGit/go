// Code generated by goa v3.2.6, DO NOT EDIT.
//
// Ha HTTP server types
//
// Command:
// $ goa gen sample1/design

package server

import (
	haviews "sample1/gen/ha/views"
)

// Goa3SampleUserResponseCollection is the type of the "Ha" service "draw card"
// endpoint HTTP response body.
type Goa3SampleUserResponseCollection []*Goa3SampleUserResponse

// Goa3SampleUserResponse is used to define fields on response body types.
type Goa3SampleUserResponse struct {
	// theme of game
	Theme string `form:"theme" json:"theme" xml:"theme"`
	// card of abc
	Card string `form:"card" json:"card" xml:"card"`
}

// NewGoa3SampleUserResponseCollection builds the HTTP response body from the
// result of the "draw card" endpoint of the "Ha" service.
func NewGoa3SampleUserResponseCollection(res haviews.Goa3SampleUserCollectionView) Goa3SampleUserResponseCollection {
	body := make([]*Goa3SampleUserResponse, len(res))
	for i, val := range res {
		body[i] = marshalHaviewsGoa3SampleUserViewToGoa3SampleUserResponse(val)
	}
	return body
}
