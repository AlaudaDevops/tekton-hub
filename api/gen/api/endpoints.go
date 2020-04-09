// Code generated by goa v3.1.1, DO NOT EDIT.
//
// api endpoints
//
// Command:
// $ goa gen github.com/tektoncd/hub/api/design

package api

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "api" service endpoints.
type Endpoints struct {
	List goa.Endpoint
}

// NewEndpoints wraps the methods of the "api" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		List: NewListEndpoint(s),
	}
}

// Use applies the given middleware to all the "api" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.List = m(e.List)
}

// NewListEndpoint returns an endpoint function that calls the method "list" of
// service "api".
func NewListEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return nil, s.List(ctx)
	}
}
