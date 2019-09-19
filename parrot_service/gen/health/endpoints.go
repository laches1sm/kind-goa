// Code generated by goa v3.0.6, DO NOT EDIT.
//
// Health endpoints
//
// Command:
// $ goa gen parrot_service/design

package health

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "Health" service endpoints.
type Endpoints struct {
	Check goa.Endpoint
}

// NewEndpoints wraps the methods of the "Health" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		Check: NewCheckEndpoint(s),
	}
}

// Use applies the given middleware to all the "Health" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Check = m(e.Check)
}

// NewCheckEndpoint returns an endpoint function that calls the method "check"
// of service "Health".
func NewCheckEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return nil, s.Check(ctx)
	}
}