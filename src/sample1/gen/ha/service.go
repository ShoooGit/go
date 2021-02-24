// Code generated by goa v3.2.6, DO NOT EDIT.
//
// Ha service
//
// Command:
// $ goa gen sample1/design

package ha

import (
	"context"
	haviews "sample1/gen/ha/views"
)

// ha serves.
type Service interface {
	// decied on a theme and card
	DrawCard(context.Context) (res Goa3SampleHaCollection, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "Ha"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [1]string{"draw card"}

// Goa3SampleHaCollection is the result type of the Ha service draw card method.
type Goa3SampleHaCollection []*Goa3SampleHa

// Ha Response
type Goa3SampleHa struct {
	// theme of game
	Theme string
	// card of abc
	Card string
}

// NewGoa3SampleHaCollection initializes result type Goa3SampleHaCollection
// from viewed result type Goa3SampleHaCollection.
func NewGoa3SampleHaCollection(vres haviews.Goa3SampleHaCollection) Goa3SampleHaCollection {
	return newGoa3SampleHaCollection(vres.Projected)
}

// NewViewedGoa3SampleHaCollection initializes viewed result type
// Goa3SampleHaCollection from result type Goa3SampleHaCollection using the
// given view.
func NewViewedGoa3SampleHaCollection(res Goa3SampleHaCollection, view string) haviews.Goa3SampleHaCollection {
	p := newGoa3SampleHaCollectionView(res)
	return haviews.Goa3SampleHaCollection{Projected: p, View: "default"}
}

// newGoa3SampleHaCollection converts projected type Goa3SampleHaCollection to
// service type Goa3SampleHaCollection.
func newGoa3SampleHaCollection(vres haviews.Goa3SampleHaCollectionView) Goa3SampleHaCollection {
	res := make(Goa3SampleHaCollection, len(vres))
	for i, n := range vres {
		res[i] = newGoa3SampleHa(n)
	}
	return res
}

// newGoa3SampleHaCollectionView projects result type Goa3SampleHaCollection to
// projected type Goa3SampleHaCollectionView using the "default" view.
func newGoa3SampleHaCollectionView(res Goa3SampleHaCollection) haviews.Goa3SampleHaCollectionView {
	vres := make(haviews.Goa3SampleHaCollectionView, len(res))
	for i, n := range res {
		vres[i] = newGoa3SampleHaView(n)
	}
	return vres
}

// newGoa3SampleHa converts projected type Goa3SampleHa to service type
// Goa3SampleHa.
func newGoa3SampleHa(vres *haviews.Goa3SampleHaView) *Goa3SampleHa {
	res := &Goa3SampleHa{}
	if vres.Theme != nil {
		res.Theme = *vres.Theme
	}
	if vres.Card != nil {
		res.Card = *vres.Card
	}
	return res
}

// newGoa3SampleHaView projects result type Goa3SampleHa to projected type
// Goa3SampleHaView using the "default" view.
func newGoa3SampleHaView(res *Goa3SampleHa) *haviews.Goa3SampleHaView {
	vres := &haviews.Goa3SampleHaView{
		Theme: &res.Theme,
		Card:  &res.Card,
	}
	return vres
}