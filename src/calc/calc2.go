package calcapi

import (
	calc2 "calc/gen/calc2"
	"context"
	"log"
)

// calc2 service example implementation.
// The example methods log the requests and return zero values.
type calc2srvc struct {
	logger *log.Logger
}

// NewCalc2 returns the calc2 service implementation.
func NewCalc2(logger *log.Logger) calc2.Service {
	return &calc2srvc{logger}
}

// Minus implements minus.
func (s *calc2srvc) Minus(ctx context.Context, p *calc2.MinusPayload) (res int, err error) {
	return p.A - p.B, nil
}
