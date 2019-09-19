package parrot

import (
	"context"
	"log"
	health "parrot_service/gen/health"
)

// Health service example implementation.
// The example methods log the requests and return zero values.
type healthsrvc struct {
	logger *log.Logger
}

// NewHealth returns the Health service implementation.
func NewHealth(logger *log.Logger) health.Service {
	return &healthsrvc{logger}
}

// Health check
func (s *healthsrvc) Check(ctx context.Context) (err error) {
	s.logger.Print("health.check")
	return
}
