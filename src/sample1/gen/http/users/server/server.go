// Code generated by goa v3.2.6, DO NOT EDIT.
//
// Users HTTP server
//
// Command:
// $ goa gen sample1/design

package server

import (
	"context"
	"net/http"
	"regexp"
	users "sample1/gen/users"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
	"goa.design/plugins/v3/cors"
)

// Server lists the Users service endpoint HTTP handlers.
type Server struct {
	Mounts     []*MountPoint
	ListUser   http.Handler
	GetUser    http.Handler
	CreateUser http.Handler
	UpdateUser http.Handler
	DeleteUser http.Handler
	CORS       http.Handler
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

// New instantiates HTTP handlers for all the Users service endpoints using the
// provided encoder and decoder. The handlers are mounted on the given mux
// using the HTTP verb and path defined in the design. errhandler is called
// whenever a response fails to be encoded. formatter is used to format errors
// returned by the service methods prior to encoding. Both errhandler and
// formatter are optional and can be nil.
func New(
	e *users.Endpoints,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"ListUser", "GET", "/api/v1/users"},
			{"GetUser", "GET", "/api/v1/users/{id}"},
			{"CreateUser", "POST", "/api/v1/users"},
			{"UpdateUser", "PUT", "/api/v1/users/{id}"},
			{"DeleteUser", "DELETE", "/api/v1/users/{id}"},
			{"CORS", "OPTIONS", "/api/v1/users"},
			{"CORS", "OPTIONS", "/api/v1/users/{id}"},
		},
		ListUser:   NewListUserHandler(e.ListUser, mux, decoder, encoder, errhandler, formatter),
		GetUser:    NewGetUserHandler(e.GetUser, mux, decoder, encoder, errhandler, formatter),
		CreateUser: NewCreateUserHandler(e.CreateUser, mux, decoder, encoder, errhandler, formatter),
		UpdateUser: NewUpdateUserHandler(e.UpdateUser, mux, decoder, encoder, errhandler, formatter),
		DeleteUser: NewDeleteUserHandler(e.DeleteUser, mux, decoder, encoder, errhandler, formatter),
		CORS:       NewCORSHandler(),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "Users" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.ListUser = m(s.ListUser)
	s.GetUser = m(s.GetUser)
	s.CreateUser = m(s.CreateUser)
	s.UpdateUser = m(s.UpdateUser)
	s.DeleteUser = m(s.DeleteUser)
	s.CORS = m(s.CORS)
}

// Mount configures the mux to serve the Users endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountListUserHandler(mux, h.ListUser)
	MountGetUserHandler(mux, h.GetUser)
	MountCreateUserHandler(mux, h.CreateUser)
	MountUpdateUserHandler(mux, h.UpdateUser)
	MountDeleteUserHandler(mux, h.DeleteUser)
	MountCORSHandler(mux, h.CORS)
}

// MountListUserHandler configures the mux to serve the "Users" service "list
// user" endpoint.
func MountListUserHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := handleUsersOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/api/v1/users", f)
}

// NewListUserHandler creates a HTTP handler which loads the HTTP request and
// calls the "Users" service "list user" endpoint.
func NewListUserHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		encodeResponse = EncodeListUserResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "list user")
		ctx = context.WithValue(ctx, goa.ServiceKey, "Users")
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

// MountGetUserHandler configures the mux to serve the "Users" service "get
// user" endpoint.
func MountGetUserHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := handleUsersOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/api/v1/users/{id}", f)
}

// NewGetUserHandler creates a HTTP handler which loads the HTTP request and
// calls the "Users" service "get user" endpoint.
func NewGetUserHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeGetUserRequest(mux, decoder)
		encodeResponse = EncodeGetUserResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "get user")
		ctx = context.WithValue(ctx, goa.ServiceKey, "Users")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
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

// MountCreateUserHandler configures the mux to serve the "Users" service
// "create user" endpoint.
func MountCreateUserHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := handleUsersOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/api/v1/users", f)
}

// NewCreateUserHandler creates a HTTP handler which loads the HTTP request and
// calls the "Users" service "create user" endpoint.
func NewCreateUserHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeCreateUserRequest(mux, decoder)
		encodeResponse = EncodeCreateUserResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "create user")
		ctx = context.WithValue(ctx, goa.ServiceKey, "Users")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
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

// MountUpdateUserHandler configures the mux to serve the "Users" service
// "update user" endpoint.
func MountUpdateUserHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := handleUsersOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("PUT", "/api/v1/users/{id}", f)
}

// NewUpdateUserHandler creates a HTTP handler which loads the HTTP request and
// calls the "Users" service "update user" endpoint.
func NewUpdateUserHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeUpdateUserRequest(mux, decoder)
		encodeResponse = EncodeUpdateUserResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "update user")
		ctx = context.WithValue(ctx, goa.ServiceKey, "Users")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
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

// MountDeleteUserHandler configures the mux to serve the "Users" service
// "delete user" endpoint.
func MountDeleteUserHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := handleUsersOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("DELETE", "/api/v1/users/{id}", f)
}

// NewDeleteUserHandler creates a HTTP handler which loads the HTTP request and
// calls the "Users" service "delete user" endpoint.
func NewDeleteUserHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeDeleteUserRequest(mux, decoder)
		encodeResponse = EncodeDeleteUserResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "delete user")
		ctx = context.WithValue(ctx, goa.ServiceKey, "Users")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
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

// MountCORSHandler configures the mux to serve the CORS endpoints for the
// service Users.
func MountCORSHandler(mux goahttp.Muxer, h http.Handler) {
	h = handleUsersOrigin(h)
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("OPTIONS", "/api/v1/users", f)
	mux.Handle("OPTIONS", "/api/v1/users/{id}", f)
}

// NewCORSHandler creates a HTTP handler which returns a simple 200 response.
func NewCORSHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})
}

// handleUsersOrigin applies the CORS response headers corresponding to the
// origin for the service Users.
func handleUsersOrigin(h http.Handler) http.Handler {
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
