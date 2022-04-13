// Code generated by goa v3.6.2, DO NOT EDIT.
//
// admin endpoints
//
// Command:
// $ goa gen github.com/tektoncd/hub/api/design

package admin

import (
	"context"

	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// Endpoints wraps the "admin" service endpoints.
type Endpoints struct {
	UpdateAgent   goa.Endpoint
	RefreshConfig goa.Endpoint
}

// NewEndpoints wraps the methods of the "admin" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	// Casting service to Auther interface
	a := s.(Auther)
	return &Endpoints{
		UpdateAgent:   NewUpdateAgentEndpoint(s, a.JWTAuth),
		RefreshConfig: NewRefreshConfigEndpoint(s, a.JWTAuth),
	}
}

// Use applies the given middleware to all the "admin" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.UpdateAgent = m(e.UpdateAgent)
	e.RefreshConfig = m(e.RefreshConfig)
}

// NewUpdateAgentEndpoint returns an endpoint function that calls the method
// "UpdateAgent" of service "admin".
func NewUpdateAgentEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*UpdateAgentPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{"rating:read", "rating:write", "agent:create", "catalog:refresh", "config:refresh", "refresh:token"},
			RequiredScopes: []string{"agent:create"},
		}
		ctx, err = authJWTFn(ctx, p.Token, &sc)
		if err != nil {
			return nil, err
		}
		return s.UpdateAgent(ctx, p)
	}
}

// NewRefreshConfigEndpoint returns an endpoint function that calls the method
// "RefreshConfig" of service "admin".
func NewRefreshConfigEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*RefreshConfigPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{"rating:read", "rating:write", "agent:create", "catalog:refresh", "config:refresh", "refresh:token"},
			RequiredScopes: []string{"config:refresh"},
		}
		ctx, err = authJWTFn(ctx, p.Token, &sc)
		if err != nil {
			return nil, err
		}
		return s.RefreshConfig(ctx, p)
	}
}
