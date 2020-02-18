package message

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	service "huzhu_service/pkg/svc/message"
	//pb "huzhu_service/pb"
	"huzhu_service/pkg/endpoints"
	// "fmt"
)
// Endpoints collects all of the endpoints that compose the addsvc service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
type Endpoints struct {
	EchoEndpoint    endpoint.Endpoint `json:""`
	SayHelloEndpoint endpoint.Endpoint `json:""`
}


// New return a new instance of the endpoint that wraps the provided service.
func New(svc service.MsgsvcService, logger log.Logger ) (ep Endpoints) {
	
	var echoEndpoint endpoint.Endpoint = NewEchoEndpoint(svc,logger);
	ep.EchoEndpoint = echoEndpoint

	var sayhelloEndpoint endpoint.Endpoint = NewSayHelloEndpoint(svc,logger);
	ep.SayHelloEndpoint = sayhelloEndpoint;
	return ep
}
func NewEchoEndpoint(svc service.MsgsvcService, logger log.Logger) endpoint.Endpoint{
	var echoEndpoint endpoint.Endpoint
	method := "echo"
	echoEndpoint = MakeEchoEndpoint(svc)
	
	echoEndpoint = endpoints.LoggingMiddleware(log.With(logger, "method", method))(echoEndpoint)
	return echoEndpoint;
}
func NewSayHelloEndpoint(svc service.MsgsvcService, logger log.Logger) endpoint.Endpoint{
	var sayHelloEndpoint endpoint.Endpoint
	method := "sayhello"
	sayHelloEndpoint = MakeSayHelloEndpoint(svc)
	
	sayHelloEndpoint = endpoints.LoggingMiddleware(log.With(logger, "method", method))(sayHelloEndpoint)
	
	return sayHelloEndpoint;
}

// MakeSumEndpoint returns an endpoint that invokes Sum on the service.
// Primarily useful in a server.
func MakeEchoEndpoint(svc service.MsgsvcService) (ep endpoint.Endpoint) {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(EchoRequest)
		if err := req.validate(); err != nil {
			return EchoResponse{}, err
		}
		
		rs, err := svc.Echo(ctx, req.Word)
		return EchoResponse{Rs: rs}, err
	}
}

// MakeConcatEndpoint returns an endpoint that invokes Concat on the service.
// Primarily useful in a server.
func MakeSayHelloEndpoint(svc service.MsgsvcService) (ep endpoint.Endpoint) {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SayHelloRequest)
		
		if err := req.validate(); err != nil {
			return SayHelloResponse{}, err
		}
		rs, err := svc.SayHello(ctx, req.Saidword, req.Want)
		return SayHelloResponse{Rs: rs}, err
	}
}


