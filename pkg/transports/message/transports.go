package message
import (
	"context"
	"huzhu_service/pb"
	service "huzhu_service/pkg/svc/message"
	endpoints "huzhu_service/pkg/endpoints/message"
	"github.com/go-kit/kit/log"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
	"github.com/go-kit/kit/endpoint"
	"huzhu_service/pkg/def"
	

)

func NewGRPCClient(conn *grpc.ClientConn,  logger log.Logger)service.MsgsvcService{

	// global client middlewares
	options := []grpctransport.ClientOption{
		//zipkinClient,
	}
	// The Sum endpoint is the same thing, with slightly different
	// middlewares to demonstrate how to specialize per-endpoint.
	var echoEndpoint endpoint.Endpoint
	{
		echoEndpoint = grpctransport.NewClient(
			conn,
			"pb.Msgsvc",
			"Echo",
			encodeGRPCEchoRequest,
			decodeGRPCEchoResponse,
			pb.EchoReply{},
			options...,
		).Endpoint()
	}

	// The Concat endpoint is the same thing, with slightly different
	// middlewares to demonstrate how to specialize per-endpoint.
	var sayHelloEndpoint endpoint.Endpoint
	{
		sayHelloEndpoint = grpctransport.NewClient(
			conn,
			"pb.Msgsvc",
			"SayHello",
			encodeGRPCSayHelloRequest,
			decodeGRPCSayHelloResponse,
			pb.SayHelloReply{},
			options...,
		).Endpoint()
	}

	return endpoints.Endpoints{
		EchoEndpoint:    echoEndpoint,
		SayHelloEndpoint: sayHelloEndpoint,
	}

}
func NewGrpcHandler(logger log.Logger,options []grpctransport.ServerOption)map[string]*grpctransport.Server{
	svc:=service.NewMsgSvc(logger)
	handlers:= make(map[string]*grpctransport.Server);
	handlers["echo"]= grpctransport.NewServer(
			endpoints.NewEchoEndpoint(svc,logger),
			decodeGRPCEchoRequest,
			encodeGRPCEchoResponse,
			options...,
		);
	handlers["sayhello"]= grpctransport.NewServer(
			endpoints.NewSayHelloEndpoint(svc,logger),
			decodeGRPCSayHelloRequest,
			encodeGRPCSayHelloResponse,
			options...,
	);
    return handlers;
}

// encodeGRPCConcatResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain response to a gRPC reply. Primarily useful in a server.
func encodeGRPCSayHelloResponse(_ context.Context, grpcReply interface{}) (res interface{}, err error) {
	reply := grpcReply.(endpoints.SayHelloResponse)
	return &pb.SayHelloReply{Rs: reply.Rs}, def.GrpcEncodeError(reply.Err)
}

// encodeGRPCSumRequest is a transport/grpc.EncodeRequestFunc that converts a
// user-domain Sum request to a gRPC Sum request. Primarily useful in a client.
func encodeGRPCSayHelloRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoints.SayHelloRequest)
	return &pb.SayHelloRequest{Saidword: req.Saidword, Want: req.Want}, nil
	
	
}

// decodeGRPCSumResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC Sum reply to a user-domain Sum response. Primarily useful in a client.
func decodeGRPCSayHelloResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.SayHelloReply)
	return endpoints.SayHelloResponse{Rs: reply.Rs}, nil
}

// encodeGRPCSumRequest is a transport/grpc.EncodeRequestFunc that converts a
// user-domain Sum request to a gRPC Sum request. Primarily useful in a client.
func encodeGRPCEchoRequest(_ context.Context, request interface{}) (interface{}, error) {

	
	r:=request.(endpoints.EchoRequest);
	return &pb.EchoRequest{Word: r.Word}, nil
	
	
}

// decodeGRPCSumResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC Sum reply to a user-domain Sum response. Primarily useful in a client.
func decodeGRPCEchoResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.EchoReply)
	return endpoints.EchoResponse{Rs: reply.Rs}, nil
}



// decodeGRPCSumRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain request. Primarily useful in a server.
func decodeGRPCSayHelloRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.SayHelloRequest)
	return endpoints.SayHelloRequest{Saidword: req.Saidword, Want: req.Want}, nil
}
// decodeGRPCSumRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain request. Primarily useful in a server.
func decodeGRPCEchoRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.EchoRequest)
	return endpoints.EchoRequest{Word: req.Word}, nil
}
// encodeGRPCSumResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain response to a gRPC reply. Primarily useful in a server.
func encodeGRPCEchoResponse(_ context.Context, grpcReply interface{}) (res interface{}, err error) {

	reply := grpcReply.(endpoints.EchoResponse)
	return &pb.EchoReply{Rs: reply.Rs}, def.GrpcEncodeError(reply.Err)
}