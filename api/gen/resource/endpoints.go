// Code generated by goa v3.4.0, DO NOT EDIT.
//
// resource endpoints
//
// Command:
// $ goa gen github.com/tektoncd/hub/api/design

package resource

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "resource" service endpoints.
type Endpoints struct {
	List                     goa.Endpoint
	VersionsByID             goa.Endpoint
	ByCatalogKindNameVersion goa.Endpoint
}

// NewEndpoints wraps the methods of the "resource" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		List:                     NewListEndpoint(s),
		VersionsByID:             NewVersionsByIDEndpoint(s),
		ByCatalogKindNameVersion: NewByCatalogKindNameVersionEndpoint(s),
	}
}

// Use applies the given middleware to all the "resource" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.List = m(e.List)
	e.VersionsByID = m(e.VersionsByID)
	e.ByCatalogKindNameVersion = m(e.ByCatalogKindNameVersion)
}

// NewListEndpoint returns an endpoint function that calls the method "List" of
// service "resource".
func NewListEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		res, err := s.List(ctx)
		if err != nil {
			return nil, err
		}
		vres := NewViewedResources(res, "default")
		return vres, nil
	}
}

// NewVersionsByIDEndpoint returns an endpoint function that calls the method
// "VersionsByID" of service "resource".
func NewVersionsByIDEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*VersionsByIDPayload)
		return s.VersionsByID(ctx, p)
	}
}

// NewByCatalogKindNameVersionEndpoint returns an endpoint function that calls
// the method "ByCatalogKindNameVersion" of service "resource".
func NewByCatalogKindNameVersionEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*ByCatalogKindNameVersionPayload)
		return s.ByCatalogKindNameVersion(ctx, p)
	}
}
