package message
import (
	"context"
	"huzhu_service/pb"
	service "huzhu_service/pkg/svc/message"
	endpoints "huzhu_service/pkg/endpoints/message"
	"github.com/go-kit/kit/log"
	grpctransport "github.com/go-kit/kit/transport/grpc"

)


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
			encodeGRPCEchoResponse,
			options...,
	);
    return handlers;
}

// decodeGRPCSumRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain request. Primarily useful in a server.
func decodeGRPCSayHelloRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.SayHelloRequest)
	return req, nil
}
// decodeGRPCSumRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain request. Primarily useful in a server.
func decodeGRPCEchoRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.EchoRequest)
	return req, nil
}
// encodeGRPCSumResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain response to a gRPC reply. Primarily useful in a server.
func encodeGRPCEchoResponse(_ context.Context, grpcReply interface{}) (res interface{}, err error) {
	
	return res,nil
}