package message

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	service "huzhu_service/pkg/svc/message"
	pb "huzhu_service/pb"
	"huzhu_service/pkg/endpoints"
)

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
		req := request.(*pb.EchoRequest)
		rs, err := svc.Echo(ctx, req.Word)
		if err!=nil {
			return &pb.EchoReply{
				Code: 500,
				Err:  err.Error(),
			}, nil
		}
		return &pb.EchoReply{
				Code: 200,
				Err:  "",
				Rs:rs,
			}, nil
	}
}

// MakeConcatEndpoint returns an endpoint that invokes Concat on the service.
// Primarily useful in a server.
func MakeSayHelloEndpoint(svc service.MsgsvcService) (ep endpoint.Endpoint) {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pb.SayHelloRequest)
		
		rs, err := svc.SayHello(ctx, req.Saidword,req.Want)
		if err!=nil {
			return &pb.SayHelloReply{
				Code: 500,
				Err:  err.Error(),
			}, nil
		}
		return &pb.SayHelloReply{
				Code: 200,
				Err:  "",
				Rs:rs,
			}, nil
	}
}


