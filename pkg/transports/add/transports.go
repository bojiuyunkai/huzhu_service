package add
import (
	"context"

	"huzhu_service/pb"
	endpoints "huzhu_service/pkg/endpoints/add"
	service "huzhu_service/pkg/svc/add"
	"github.com/go-kit/kit/log"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"huzhu_service/pkg/def"
	"github.com/go-kit/kit/endpoint"
	"google.golang.org/grpc"

)


func NewGrpcHandler(logger log.Logger,options []grpctransport.ServerOption)map[string]*grpctransport.Server{
	addsvc:=service.NewAddSvc(logger)
	handlers:= make(map[string]*grpctransport.Server);
	handlers["sum"]= grpctransport.NewServer(
			endpoints.NewSumEndpoint(addsvc,logger),
			decodeGRPCSumRequest,
			encodeGRPCSumResponse,
			options...,
		);
	handlers["concat"]= grpctransport.NewServer(
			endpoints.NewConcatEndpoint(addsvc,logger),
			decodeGRPCConcatRequest,
			encodeGRPCConcatResponse,
			options...,
			)

    return handlers;
}

func NewGRPCClient(conn *grpc.ClientConn,  logger log.Logger)service.AddsvcService{

	// global client middlewares
	options := []grpctransport.ClientOption{
		//zipkinClient,
	}
	// The Sum endpoint is the same thing, with slightly different
	// middlewares to demonstrate how to specialize per-endpoint.
	var sumEndpoint endpoint.Endpoint
	{
		sumEndpoint = grpctransport.NewClient(
			conn,
			"pb.Addsvc",
			"Sum",
			encodeGRPCSumRequest,
			decodeGRPCSumResponse,
			pb.SumReply{},
			options...,
		).Endpoint()
	}

	// The Concat endpoint is the same thing, with slightly different
	// middlewares to demonstrate how to specialize per-endpoint.
	var concatEndpoint endpoint.Endpoint
	{
		concatEndpoint = grpctransport.NewClient(
			conn,
			"pb.Addsvc",
			"Concat",
			encodeGRPCConcatRequest,
			decodeGRPCConcatResponse,
			pb.ConcatReply{},
			options...,
		).Endpoint()
	}

	return endpoints.Endpoints{
		SumEndpoint:    sumEndpoint,
		ConcatEndpoint: concatEndpoint,
	}

}


// decodeGRPCSumRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain request. Primarily useful in a server.
func decodeGRPCSumRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.SumRequest)
	return endpoints.SumRequest{A: req.A, B: req.B}, nil
}

// encodeGRPCSumResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain response to a gRPC reply. Primarily useful in a server.
func encodeGRPCSumResponse(_ context.Context, grpcReply interface{}) (res interface{}, err error) {
	reply := grpcReply.(endpoints.SumResponse)
	return &pb.SumReply{Rs: reply.Rs}, def.GrpcEncodeError(reply.Err)
}

// decodeGRPCConcatRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain request. Primarily useful in a server.
func decodeGRPCConcatRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.ConcatRequest)
	return endpoints.ConcatRequest{A: req.A, B: req.B}, nil
}

// encodeGRPCConcatResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain response to a gRPC reply. Primarily useful in a server.
func encodeGRPCConcatResponse(_ context.Context, grpcReply interface{}) (res interface{}, err error) {
	reply := grpcReply.(endpoints.ConcatResponse)
	return &pb.ConcatReply{Rs: reply.Rs}, def.GrpcEncodeError(reply.Err)
}

// encodeGRPCSumRequest is a transport/grpc.EncodeRequestFunc that converts a
// user-domain Sum request to a gRPC Sum request. Primarily useful in a client.
func encodeGRPCSumRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoints.SumRequest)
	
	return &pb.SumRequest{A: req.A, B: req.B}, nil
}

// decodeGRPCSumResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC Sum reply to a user-domain Sum response. Primarily useful in a client.
func decodeGRPCSumResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.SumReply)
	return endpoints.SumResponse{Rs: reply.Rs}, nil
}

// encodeGRPCConcatRequest is a transport/grpc.EncodeRequestFunc that converts a
// user-domain Concat request to a gRPC Concat request. Primarily useful in a client.
func encodeGRPCConcatRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoints.ConcatRequest)
	return &pb.ConcatRequest{A: req.A, B: req.B}, nil
}

// decodeGRPCConcatResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC Concat reply to a user-domain Concat response. Primarily useful in a client.
func decodeGRPCConcatResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.ConcatReply)
	return endpoints.ConcatResponse{Rs: reply.Rs}, nil
}