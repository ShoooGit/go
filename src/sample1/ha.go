package goa3sample

import (
	"context"
	"log"
	ha "sample1/gen/ha"
)

// Ha service example implementation.
// The example methods log the requests and return zero values.
type hasrvc struct {
	logger *log.Logger
}

// NewHa returns the Ha service implementation.
func NewHa(logger *log.Logger) ha.Service {
	return &hasrvc{logger}
}

// decied on a theme and card
func (s *hasrvc) DrawCard(ctx context.Context) (res ha.Goa3SampleHaCollection, err error) {
	s.logger.Print("ha.draw card")
	res = []*ha.Goa3SampleHa{}
	ha := &ha.Goa3SampleHa{Theme: "theme", Card: "A"}
	res = append(res, ha)
	return
}
