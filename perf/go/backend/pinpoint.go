package backend

import (
	"go.skia.org/infra/perf/go/backend/shared"
	pinpoint_service "go.skia.org/infra/pinpoint/go/service"
	pb "go.skia.org/infra/pinpoint/proto/v1"
	"golang.org/x/time/rate"
	"google.golang.org/grpc"
)

// pinpointService implements backend.BackendService, provides a wrapper struct
// for the pinpoint service implementation.
type pinpointService struct {
	server *pinpoint_service.Server
}

// NewPinpointService returns a new instance of the pinpoint service.
func NewPinpointService(t pinpoint_service.TemporalProvider, l *rate.Limiter) *pinpointService {
	return &pinpointService{
		server: pinpoint_service.New(t, l),
	}
}

// GetAuthorizationPolicy returns the authorization policy for the service.
func (service *pinpointService) GetAuthorizationPolicy() shared.AuthorizationPolicy {
	// TODO(ashwinpv) Once validation is done, update this to only allow the FE service account role.
	return shared.AuthorizationPolicy{
		AllowUnauthenticated: true,
	}
}

// RegisterGrpc registers the grpc service with the server instance.
func (service *pinpointService) RegisterGrpc(grpcServer *grpc.Server) {
	pb.RegisterPinpointServer(grpcServer, service.server)
}

// GetServiceDescriptor returns the service descriptor for the service.
func (service *pinpointService) GetServiceDescriptor() grpc.ServiceDesc {
	return pb.Pinpoint_ServiceDesc
}
