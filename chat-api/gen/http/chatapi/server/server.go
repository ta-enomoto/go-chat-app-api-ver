// Code generated by goa v3.4.3, DO NOT EDIT.
//
// chatapi HTTP server
//
// Command:
// $ goa gen chat-api/design

package server

import (
	chatapi "chat-api/gen/chatapi"
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
	"goa.design/plugins/v3/cors"
)

// Server lists the chatapi service endpoint HTTP handlers.
type Server struct {
	Mounts   []*MountPoint
	Getchat  http.Handler
	Postchat http.Handler
	CORS     http.Handler
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

// New instantiates HTTP handlers for all the chatapi service endpoints using
// the provided encoder and decoder. The handlers are mounted on the given mux
// using the HTTP verb and path defined in the design. errhandler is called
// whenever a response fails to be encoded. formatter is used to format errors
// returned by the service methods prior to encoding. Both errhandler and
// formatter are optional and can be nil.
func New(
	e *chatapi.Endpoints,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"Getchat", "GET", "/chatroom/{id}"},
			{"Postchat", "POST", "/chatroom/chat"},
			{"CORS", "OPTIONS", "/chatroom/{id}"},
			{"CORS", "OPTIONS", "/chatroom/chat"},
		},
		Getchat:  NewGetchatHandler(e.Getchat, mux, decoder, encoder, errhandler, formatter),
		Postchat: NewPostchatHandler(e.Postchat, mux, decoder, encoder, errhandler, formatter),
		CORS:     NewCORSHandler(),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "chatapi" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.Getchat = m(s.Getchat)
	s.Postchat = m(s.Postchat)
	s.CORS = m(s.CORS)
}

// Mount configures the mux to serve the chatapi endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountGetchatHandler(mux, h.Getchat)
	MountPostchatHandler(mux, h.Postchat)
	MountCORSHandler(mux, h.CORS)
}

// MountGetchatHandler configures the mux to serve the "chatapi" service
// "getchat" endpoint.
func MountGetchatHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleChatapiOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/chatroom/{id}", f)
}

// NewGetchatHandler creates a HTTP handler which loads the HTTP request and
// calls the "chatapi" service "getchat" endpoint.
func NewGetchatHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeGetchatRequest(mux, decoder)
		encodeResponse = EncodeGetchatResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "getchat")
		ctx = context.WithValue(ctx, goa.ServiceKey, "chatapi")
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

// MountPostchatHandler configures the mux to serve the "chatapi" service
// "postchat" endpoint.
func MountPostchatHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleChatapiOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/chatroom/chat", f)
}

// NewPostchatHandler creates a HTTP handler which loads the HTTP request and
// calls the "chatapi" service "postchat" endpoint.
func NewPostchatHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodePostchatRequest(mux, decoder)
		encodeResponse = EncodePostchatResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "postchat")
		ctx = context.WithValue(ctx, goa.ServiceKey, "chatapi")
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
// service chatapi.
func MountCORSHandler(mux goahttp.Muxer, h http.Handler) {
	h = HandleChatapiOrigin(h)
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("OPTIONS", "/chatroom/{id}", f)
	mux.Handle("OPTIONS", "/chatroom/chat", f)
}

// NewCORSHandler creates a HTTP handler which returns a simple 200 response.
func NewCORSHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})
}

// HandleChatapiOrigin applies the CORS response headers corresponding to the
// origin for the service chatapi.
func HandleChatapiOrigin(h http.Handler) http.Handler {
	origHndlr := h.(http.HandlerFunc)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			origHndlr(w, r)
			return
		}
		if cors.MatchOrigin(origin, "http://172.25.0.2") {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := r.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				w.Header().Set("Access-Control-Allow-Methods", "GET")
				w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, Authorization")
			}
			origHndlr(w, r)
			return
		}
		origHndlr(w, r)
		return
	})
}
