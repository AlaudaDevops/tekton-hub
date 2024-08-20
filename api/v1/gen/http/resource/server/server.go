// Code generated by goa v3.18.2, DO NOT EDIT.
//
// resource HTTP server
//
// Command:
// $ goa gen github.com/tektoncd/hub/api/v1/design

package server

import (
	"bufio"
	"context"
	"io"
	"net/http"

	resource "github.com/tektoncd/hub/api/v1/gen/resource"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
	"goa.design/plugins/v3/cors"
)

// Server lists the resource service endpoint HTTP handlers.
type Server struct {
	Mounts                             []*MountPoint
	Query                              http.Handler
	List                               http.Handler
	VersionsByID                       http.Handler
	ByCatalogKindNameVersion           http.Handler
	ByCatalogKindNameVersionReadme     http.Handler
	ByCatalogKindNameVersionYaml       http.Handler
	ByVersionID                        http.Handler
	ByCatalogKindName                  http.Handler
	ByID                               http.Handler
	GetRawYamlByCatalogKindNameVersion http.Handler
	GetLatestRawYamlByCatalogKindName  http.Handler
	CORS                               http.Handler
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

// New instantiates HTTP handlers for all the resource service endpoints using
// the provided encoder and decoder. The handlers are mounted on the given mux
// using the HTTP verb and path defined in the design. errhandler is called
// whenever a response fails to be encoded. formatter is used to format errors
// returned by the service methods prior to encoding. Both errhandler and
// formatter are optional and can be nil.
func New(
	e *resource.Endpoints,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"Query", "GET", "/v1/query"},
			{"List", "GET", "/v1/resources"},
			{"VersionsByID", "GET", "/v1/resource/{id}/versions"},
			{"ByCatalogKindNameVersion", "GET", "/v1/resource/{catalog}/{kind}/{name}/{version}"},
			{"ByCatalogKindNameVersionReadme", "GET", "/v1/resource/{catalog}/{kind}/{name}/{version}/readme"},
			{"ByCatalogKindNameVersionYaml", "GET", "/v1/resource/{catalog}/{kind}/{name}/{version}/yaml"},
			{"ByVersionID", "GET", "/v1/resource/version/{versionID}"},
			{"ByCatalogKindName", "GET", "/v1/resource/{catalog}/{kind}/{name}"},
			{"ByID", "GET", "/v1/resource/{id}"},
			{"GetRawYamlByCatalogKindNameVersion", "GET", "/v1/resource/{catalog}/{kind}/{name}/{version}/raw"},
			{"GetLatestRawYamlByCatalogKindName", "GET", "/v1/resource/{catalog}/{kind}/{name}/raw"},
			{"CORS", "OPTIONS", "/v1/query"},
			{"CORS", "OPTIONS", "/v1/resources"},
			{"CORS", "OPTIONS", "/v1/resource/{id}/versions"},
			{"CORS", "OPTIONS", "/v1/resource/{catalog}/{kind}/{name}/{version}"},
			{"CORS", "OPTIONS", "/v1/resource/{catalog}/{kind}/{name}/{version}/readme"},
			{"CORS", "OPTIONS", "/v1/resource/{catalog}/{kind}/{name}/{version}/yaml"},
			{"CORS", "OPTIONS", "/v1/resource/version/{versionID}"},
			{"CORS", "OPTIONS", "/v1/resource/{catalog}/{kind}/{name}"},
			{"CORS", "OPTIONS", "/v1/resource/{id}"},
			{"CORS", "OPTIONS", "/v1/resource/{catalog}/{kind}/{name}/{version}/raw"},
			{"CORS", "OPTIONS", "/v1/resource/{catalog}/{kind}/{name}/raw"},
		},
		Query:                              NewQueryHandler(e.Query, mux, decoder, encoder, errhandler, formatter),
		List:                               NewListHandler(e.List, mux, decoder, encoder, errhandler, formatter),
		VersionsByID:                       NewVersionsByIDHandler(e.VersionsByID, mux, decoder, encoder, errhandler, formatter),
		ByCatalogKindNameVersion:           NewByCatalogKindNameVersionHandler(e.ByCatalogKindNameVersion, mux, decoder, encoder, errhandler, formatter),
		ByCatalogKindNameVersionReadme:     NewByCatalogKindNameVersionReadmeHandler(e.ByCatalogKindNameVersionReadme, mux, decoder, encoder, errhandler, formatter),
		ByCatalogKindNameVersionYaml:       NewByCatalogKindNameVersionYamlHandler(e.ByCatalogKindNameVersionYaml, mux, decoder, encoder, errhandler, formatter),
		ByVersionID:                        NewByVersionIDHandler(e.ByVersionID, mux, decoder, encoder, errhandler, formatter),
		ByCatalogKindName:                  NewByCatalogKindNameHandler(e.ByCatalogKindName, mux, decoder, encoder, errhandler, formatter),
		ByID:                               NewByIDHandler(e.ByID, mux, decoder, encoder, errhandler, formatter),
		GetRawYamlByCatalogKindNameVersion: NewGetRawYamlByCatalogKindNameVersionHandler(e.GetRawYamlByCatalogKindNameVersion, mux, decoder, encoder, errhandler, formatter),
		GetLatestRawYamlByCatalogKindName:  NewGetLatestRawYamlByCatalogKindNameHandler(e.GetLatestRawYamlByCatalogKindName, mux, decoder, encoder, errhandler, formatter),
		CORS:                               NewCORSHandler(),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "resource" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.Query = m(s.Query)
	s.List = m(s.List)
	s.VersionsByID = m(s.VersionsByID)
	s.ByCatalogKindNameVersion = m(s.ByCatalogKindNameVersion)
	s.ByCatalogKindNameVersionReadme = m(s.ByCatalogKindNameVersionReadme)
	s.ByCatalogKindNameVersionYaml = m(s.ByCatalogKindNameVersionYaml)
	s.ByVersionID = m(s.ByVersionID)
	s.ByCatalogKindName = m(s.ByCatalogKindName)
	s.ByID = m(s.ByID)
	s.GetRawYamlByCatalogKindNameVersion = m(s.GetRawYamlByCatalogKindNameVersion)
	s.GetLatestRawYamlByCatalogKindName = m(s.GetLatestRawYamlByCatalogKindName)
	s.CORS = m(s.CORS)
}

// MethodNames returns the methods served.
func (s *Server) MethodNames() []string { return resource.MethodNames[:] }

// Mount configures the mux to serve the resource endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountQueryHandler(mux, h.Query)
	MountListHandler(mux, h.List)
	MountVersionsByIDHandler(mux, h.VersionsByID)
	MountByCatalogKindNameVersionHandler(mux, h.ByCatalogKindNameVersion)
	MountByCatalogKindNameVersionReadmeHandler(mux, h.ByCatalogKindNameVersionReadme)
	MountByCatalogKindNameVersionYamlHandler(mux, h.ByCatalogKindNameVersionYaml)
	MountByVersionIDHandler(mux, h.ByVersionID)
	MountByCatalogKindNameHandler(mux, h.ByCatalogKindName)
	MountByIDHandler(mux, h.ByID)
	MountGetRawYamlByCatalogKindNameVersionHandler(mux, h.GetRawYamlByCatalogKindNameVersion)
	MountGetLatestRawYamlByCatalogKindNameHandler(mux, h.GetLatestRawYamlByCatalogKindName)
	MountCORSHandler(mux, h.CORS)
}

// Mount configures the mux to serve the resource endpoints.
func (s *Server) Mount(mux goahttp.Muxer) {
	Mount(mux, s)
}

// MountQueryHandler configures the mux to serve the "resource" service "Query"
// endpoint.
func MountQueryHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleResourceOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/v1/query", f)
}

// NewQueryHandler creates a HTTP handler which loads the HTTP request and
// calls the "resource" service "Query" endpoint.
func NewQueryHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeQueryRequest(mux, decoder)
		encodeResponse = EncodeQueryResponse(encoder)
		encodeError    = EncodeQueryError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "Query")
		ctx = context.WithValue(ctx, goa.ServiceKey, "resource")
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

// MountListHandler configures the mux to serve the "resource" service "List"
// endpoint.
func MountListHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleResourceOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/v1/resources", f)
}

// NewListHandler creates a HTTP handler which loads the HTTP request and calls
// the "resource" service "List" endpoint.
func NewListHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeListRequest(mux, decoder)
		encodeResponse = EncodeListResponse(encoder)
		encodeError    = EncodeListError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "List")
		ctx = context.WithValue(ctx, goa.ServiceKey, "resource")
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

// MountVersionsByIDHandler configures the mux to serve the "resource" service
// "VersionsByID" endpoint.
func MountVersionsByIDHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleResourceOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/v1/resource/{id}/versions", f)
}

// NewVersionsByIDHandler creates a HTTP handler which loads the HTTP request
// and calls the "resource" service "VersionsByID" endpoint.
func NewVersionsByIDHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeVersionsByIDRequest(mux, decoder)
		encodeResponse = EncodeVersionsByIDResponse(encoder)
		encodeError    = EncodeVersionsByIDError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "VersionsByID")
		ctx = context.WithValue(ctx, goa.ServiceKey, "resource")
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

// MountByCatalogKindNameVersionHandler configures the mux to serve the
// "resource" service "ByCatalogKindNameVersion" endpoint.
func MountByCatalogKindNameVersionHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleResourceOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/v1/resource/{catalog}/{kind}/{name}/{version}", f)
}

// NewByCatalogKindNameVersionHandler creates a HTTP handler which loads the
// HTTP request and calls the "resource" service "ByCatalogKindNameVersion"
// endpoint.
func NewByCatalogKindNameVersionHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeByCatalogKindNameVersionRequest(mux, decoder)
		encodeResponse = EncodeByCatalogKindNameVersionResponse(encoder)
		encodeError    = EncodeByCatalogKindNameVersionError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "ByCatalogKindNameVersion")
		ctx = context.WithValue(ctx, goa.ServiceKey, "resource")
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

// MountByCatalogKindNameVersionReadmeHandler configures the mux to serve the
// "resource" service "ByCatalogKindNameVersionReadme" endpoint.
func MountByCatalogKindNameVersionReadmeHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleResourceOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/v1/resource/{catalog}/{kind}/{name}/{version}/readme", f)
}

// NewByCatalogKindNameVersionReadmeHandler creates a HTTP handler which loads
// the HTTP request and calls the "resource" service
// "ByCatalogKindNameVersionReadme" endpoint.
func NewByCatalogKindNameVersionReadmeHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeByCatalogKindNameVersionReadmeRequest(mux, decoder)
		encodeResponse = EncodeByCatalogKindNameVersionReadmeResponse(encoder)
		encodeError    = EncodeByCatalogKindNameVersionReadmeError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "ByCatalogKindNameVersionReadme")
		ctx = context.WithValue(ctx, goa.ServiceKey, "resource")
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

// MountByCatalogKindNameVersionYamlHandler configures the mux to serve the
// "resource" service "ByCatalogKindNameVersionYaml" endpoint.
func MountByCatalogKindNameVersionYamlHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleResourceOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/v1/resource/{catalog}/{kind}/{name}/{version}/yaml", f)
}

// NewByCatalogKindNameVersionYamlHandler creates a HTTP handler which loads
// the HTTP request and calls the "resource" service
// "ByCatalogKindNameVersionYaml" endpoint.
func NewByCatalogKindNameVersionYamlHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeByCatalogKindNameVersionYamlRequest(mux, decoder)
		encodeResponse = EncodeByCatalogKindNameVersionYamlResponse(encoder)
		encodeError    = EncodeByCatalogKindNameVersionYamlError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "ByCatalogKindNameVersionYaml")
		ctx = context.WithValue(ctx, goa.ServiceKey, "resource")
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

// MountByVersionIDHandler configures the mux to serve the "resource" service
// "ByVersionId" endpoint.
func MountByVersionIDHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleResourceOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/v1/resource/version/{versionID}", f)
}

// NewByVersionIDHandler creates a HTTP handler which loads the HTTP request
// and calls the "resource" service "ByVersionId" endpoint.
func NewByVersionIDHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeByVersionIDRequest(mux, decoder)
		encodeResponse = EncodeByVersionIDResponse(encoder)
		encodeError    = EncodeByVersionIDError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "ByVersionId")
		ctx = context.WithValue(ctx, goa.ServiceKey, "resource")
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

// MountByCatalogKindNameHandler configures the mux to serve the "resource"
// service "ByCatalogKindName" endpoint.
func MountByCatalogKindNameHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleResourceOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/v1/resource/{catalog}/{kind}/{name}", f)
}

// NewByCatalogKindNameHandler creates a HTTP handler which loads the HTTP
// request and calls the "resource" service "ByCatalogKindName" endpoint.
func NewByCatalogKindNameHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeByCatalogKindNameRequest(mux, decoder)
		encodeResponse = EncodeByCatalogKindNameResponse(encoder)
		encodeError    = EncodeByCatalogKindNameError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "ByCatalogKindName")
		ctx = context.WithValue(ctx, goa.ServiceKey, "resource")
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

// MountByIDHandler configures the mux to serve the "resource" service "ById"
// endpoint.
func MountByIDHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleResourceOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/v1/resource/{id}", f)
}

// NewByIDHandler creates a HTTP handler which loads the HTTP request and calls
// the "resource" service "ById" endpoint.
func NewByIDHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeByIDRequest(mux, decoder)
		encodeResponse = EncodeByIDResponse(encoder)
		encodeError    = EncodeByIDError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "ById")
		ctx = context.WithValue(ctx, goa.ServiceKey, "resource")
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

// MountGetRawYamlByCatalogKindNameVersionHandler configures the mux to serve
// the "resource" service "GetRawYamlByCatalogKindNameVersion" endpoint.
func MountGetRawYamlByCatalogKindNameVersionHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleResourceOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/v1/resource/{catalog}/{kind}/{name}/{version}/raw", f)
}

// NewGetRawYamlByCatalogKindNameVersionHandler creates a HTTP handler which
// loads the HTTP request and calls the "resource" service
// "GetRawYamlByCatalogKindNameVersion" endpoint.
func NewGetRawYamlByCatalogKindNameVersionHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeGetRawYamlByCatalogKindNameVersionRequest(mux, decoder)
		encodeResponse = EncodeGetRawYamlByCatalogKindNameVersionResponse(encoder)
		encodeError    = EncodeGetRawYamlByCatalogKindNameVersionError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "GetRawYamlByCatalogKindNameVersion")
		ctx = context.WithValue(ctx, goa.ServiceKey, "resource")
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
		o := res.(*resource.GetRawYamlByCatalogKindNameVersionResponseData)
		defer o.Body.Close()
		if wt, ok := o.Body.(io.WriterTo); ok {
			n, err := wt.WriteTo(w)
			if err != nil {
				if n == 0 {
					if err := encodeError(ctx, w, err); err != nil {
						errhandler(ctx, w, err)
					}
				} else {
					if f, ok := w.(http.Flusher); ok {
						f.Flush()
					}
					panic(http.ErrAbortHandler) // too late to write an error
				}
			}
			return
		}
		// handle immediate read error like a returned error
		buf := bufio.NewReader(o.Body)
		if _, err := buf.Peek(1); err != nil && err != io.EOF {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
			return
		}
		if _, err := io.Copy(w, buf); err != nil {
			if f, ok := w.(http.Flusher); ok {
				f.Flush()
			}
			panic(http.ErrAbortHandler) // too late to write an error
		}
	})
}

// MountGetLatestRawYamlByCatalogKindNameHandler configures the mux to serve
// the "resource" service "GetLatestRawYamlByCatalogKindName" endpoint.
func MountGetLatestRawYamlByCatalogKindNameHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleResourceOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/v1/resource/{catalog}/{kind}/{name}/raw", f)
}

// NewGetLatestRawYamlByCatalogKindNameHandler creates a HTTP handler which
// loads the HTTP request and calls the "resource" service
// "GetLatestRawYamlByCatalogKindName" endpoint.
func NewGetLatestRawYamlByCatalogKindNameHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeGetLatestRawYamlByCatalogKindNameRequest(mux, decoder)
		encodeResponse = EncodeGetLatestRawYamlByCatalogKindNameResponse(encoder)
		encodeError    = EncodeGetLatestRawYamlByCatalogKindNameError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "GetLatestRawYamlByCatalogKindName")
		ctx = context.WithValue(ctx, goa.ServiceKey, "resource")
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
		o := res.(*resource.GetLatestRawYamlByCatalogKindNameResponseData)
		defer o.Body.Close()
		if wt, ok := o.Body.(io.WriterTo); ok {
			n, err := wt.WriteTo(w)
			if err != nil {
				if n == 0 {
					if err := encodeError(ctx, w, err); err != nil {
						errhandler(ctx, w, err)
					}
				} else {
					if f, ok := w.(http.Flusher); ok {
						f.Flush()
					}
					panic(http.ErrAbortHandler) // too late to write an error
				}
			}
			return
		}
		// handle immediate read error like a returned error
		buf := bufio.NewReader(o.Body)
		if _, err := buf.Peek(1); err != nil && err != io.EOF {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
			return
		}
		if _, err := io.Copy(w, buf); err != nil {
			if f, ok := w.(http.Flusher); ok {
				f.Flush()
			}
			panic(http.ErrAbortHandler) // too late to write an error
		}
	})
}

// MountCORSHandler configures the mux to serve the CORS endpoints for the
// service resource.
func MountCORSHandler(mux goahttp.Muxer, h http.Handler) {
	h = HandleResourceOrigin(h)
	mux.Handle("OPTIONS", "/v1/query", h.ServeHTTP)
	mux.Handle("OPTIONS", "/v1/resources", h.ServeHTTP)
	mux.Handle("OPTIONS", "/v1/resource/{id}/versions", h.ServeHTTP)
	mux.Handle("OPTIONS", "/v1/resource/{catalog}/{kind}/{name}/{version}", h.ServeHTTP)
	mux.Handle("OPTIONS", "/v1/resource/{catalog}/{kind}/{name}/{version}/readme", h.ServeHTTP)
	mux.Handle("OPTIONS", "/v1/resource/{catalog}/{kind}/{name}/{version}/yaml", h.ServeHTTP)
	mux.Handle("OPTIONS", "/v1/resource/version/{versionID}", h.ServeHTTP)
	mux.Handle("OPTIONS", "/v1/resource/{catalog}/{kind}/{name}", h.ServeHTTP)
	mux.Handle("OPTIONS", "/v1/resource/{id}", h.ServeHTTP)
	mux.Handle("OPTIONS", "/v1/resource/{catalog}/{kind}/{name}/{version}/raw", h.ServeHTTP)
	mux.Handle("OPTIONS", "/v1/resource/{catalog}/{kind}/{name}/raw", h.ServeHTTP)
}

// NewCORSHandler creates a HTTP handler which returns a simple 204 response.
func NewCORSHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(204)
	})
}

// HandleResourceOrigin applies the CORS response headers corresponding to the
// origin for the service resource.
func HandleResourceOrigin(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			h.ServeHTTP(w, r)
			return
		}
		if cors.MatchOrigin(origin, "*") {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin")
			if acrm := r.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				w.Header().Set("Access-Control-Allow-Methods", "GET")
				w.WriteHeader(204)
				return
			}
			h.ServeHTTP(w, r)
			return
		}
		h.ServeHTTP(w, r)
		return
	})
}
