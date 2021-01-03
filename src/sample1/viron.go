package goa3sample

import (
	"context"
	"log"
	viron "sample1/gen/viron"
)

// Viron service example implementation.
// The example methods log the requests and return zero values.
type vironsrvc struct {
	logger *log.Logger
}

// NewViron returns the Viron service implementation.
func NewViron(logger *log.Logger) viron.Service {
	return &vironsrvc{logger}
}

// Add viron_authtype
func (s *vironsrvc) Authtype(ctx context.Context) (res viron.VironAuthtypeCollection, err error) {
	s.logger.Print("viron.authtype")
	return
}

// Add viron_menu
func (s *vironsrvc) VironMenuEndpoint(ctx context.Context) (res *viron.VironMenu, err error) {
	res = &viron.VironMenu{}
	s.logger.Print("viron.viron_menu")
	return
}
