// Code generated by goa v3.2.6, DO NOT EDIT.
//
// Viron HTTP server
//
// Command:
// $ goa gen sample1/design

package server

import (
	"context"
	"net/http"
	"regexp"
	viron "sample1/gen/viron"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
	"goa.design/plugins/v3/cors"
)

// Server lists the Viron service endpoint HTTP handlers.
type Server struct {
	Mounts            []*MountPoint
	Authtype          http.Handler
	VironMenuEndpoint http.Handler
	CORS              http.Handler
}

// ErrorNamer is an interface implemented by generated error structs that
// exposes the name of the error as defined in the design.
type ErrorNamer interface {
	ErrorName() string
}

// MountPoint holds information about the mounted endpoints.
type MountPoint struct {
	// Method is the name of the service method served by the mounted HTTP handler.
	Method string
	// Verb is the HTTP method used to match requests to the mounted handler.
	Verb string
	// Pattern is the HTTP request path pattern used to match requests to the
	// mounted handler.
	Pattern string
}

// New instantiates HTTP handlers for all the Viron service endpoints using the
// provided encoder and decoder. The handlers are mounted on the given mux
// using the HTTP verb and path defined in the design. errhandler is called
// whenever a response fails to be encoded. formatter is used to format errors
// returned by the service methods prior to encoding. Both errhandler and
// formatter are optional and can be nil.
func New(
	e *viron.Endpoints,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"Authtype", "GET", "/viron_authtype"},
			{"VironMenuEndpoint", "GET", "/viron"},
			{"CORS", "OPTIONS", "/viron_authtype"},
			{"CORS", "OPTIONS", "/viron"},
			{"CORS", "OPTIONS", "/v1/swagger.json"},
			{"gen/http/openapi.json", "GET", "/v1/swagger.json"},
		},
		Authtype:          NewAuthtypeHandler(e.Authtype, mux, decoder, encoder, errhandler, formatter),
		VironMenuEndpoint: NewVironMenuEndpointHandler(e.VironMenuEndpoint, mux, decoder, encoder, errhandler, formatter),
		CORS:              NewCORSHandler(),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "Viron" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.Authtype = m(s.Authtype)
	s.VironMenuEndpoint = m(s.VironMenuEndpoint)
	s.CORS = m(s.CORS)
}

// Mount configures the mux to serve the Viron endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountAuthtypeHandler(mux, h.Authtype)
	MountVironMenuEndpointHandler(mux, h.VironMenuEndpoint)
	MountCORSHandler(mux, h.CORS)
	MountGenHTTPOpenapiJSON(mux, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "gen/http/openapi.json")
	}))
}

// MountAuthtypeHandler configures the mux to serve the "Viron" service
// "authtype" endpoint.
func MountAuthtypeHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := handleVironOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/viron_authtype", f)
}

// NewAuthtypeHandler creates a HTTP handler which loads the HTTP request and
// calls the "Viron" service "authtype" endpoint.
func NewAuthtypeHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		encodeResponse = EncodeAuthtypeResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "authtype")
		ctx = context.WithValue(ctx, goa.ServiceKey, "Viron")
		var err error
		res, err := endpoint(ctx, nil)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountVironMenuEndpointHandler configures the mux to serve the "Viron"
// service "viron_menu" endpoint.
func MountVironMenuEndpointHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := handleVironOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/viron", f)
}

// NewVironMenuEndpointHandler creates a HTTP handler which loads the HTTP
// request and calls the "Viron" service "viron_menu" endpoint.
func NewVironMenuEndpointHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		encodeResponse = EncodeVironMenuEndpointResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "viron_menu")
		ctx = context.WithValue(ctx, goa.ServiceKey, "Viron")
		var err error
		res, err := endpoint(ctx, nil)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountGenHTTPOpenapiJSON configures the mux to serve GET request made to
// "/v1/swagger.json".
func MountGenHTTPOpenapiJSON(mux goahttp.Muxer, h http.Handler) {
	mux.Handle("GET", "/v1/swagger.json", handleVironOrigin(h).ServeHTTP)
}

// MountCORSHandler configures the mux to serve the CORS endpoints for the
// service Viron.
func MountCORSHandler(mux goahttp.Muxer, h http.Handler) {
	h = handleVironOrigin(h)
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("OPTIONS", "/viron_authtype", f)
	mux.Handle("OPTIONS", "/viron", f)
	mux.Handle("OPTIONS", "/v1/swagger.json", f)
}

// NewCORSHandler creates a HTTP handler which returns a simple 200 response.
func NewCORSHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})
}

// handleVironOrigin applies the CORS response headers corresponding to the
// origin for the service Viron.
func handleVironOrigin(h http.Handler) http.Handler {
	spec0 := regexp.MustCompile(".*localhost.*")
	origHndlr := h.(http.HandlerFunc)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			origHndlr(w, r)
			return
		}
		if cors.MatchOriginRegexp(origin, spec0) {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin")
			w.Header().Set("Access-Control-Max-Age", "600")
			w.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := r.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
				w.Header().Set("Access-Control-Allow-Headers", "content-type")
			}
			origHndlr(w, r)
			return
		}
		origHndlr(w, r)
		return
	})
}
