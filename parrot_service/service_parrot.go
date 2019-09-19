package parrot

import (
	"context"
	"log"
	serviceparrot "parrot_service/gen/service_parrot"
)

// service-parrot service example implementation.
// The example methods log the requests and return zero values.
type serviceParrotsrvc struct {
	logger *log.Logger
}

// NewServiceParrot returns the service-parrot service implementation.
func NewServiceParrot(logger *log.Logger) serviceparrot.Service {
	return &serviceParrotsrvc{logger}
}

// Postparrot implements postparrot.
func (s *serviceParrotsrvc) Postparrot(ctx context.Context, p *serviceparrot.Parrotpayload) (err error) {
	s.logger.Print("serviceParrot.postparrot")
	return
}

// Listaparrot implements listaparrot.
func (s *serviceParrotsrvc) Listaparrot(ctx context.Context, p *serviceparrot.ListaparrotPayload) (err error) {
	s.logger.Print("serviceParrot.listaparrot")
	return
}

// Listallparrots implements listallparrots.
func (s *serviceParrotsrvc) Listallparrots(ctx context.Context) (err error) {
	s.logger.Print("serviceParrot.listallparrots")
	return
}
