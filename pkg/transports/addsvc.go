package transports
import (
	"context"
	

	
	
	
	

	"huzhu_service/pb"
	"huzhu_service/pkg/endpoints"

)
//主要转发
func (s *GrpcServer) Sum(ctx context.Context, req *pb.SumRequest) (rep *pb.SumReply, err error) {
	_, rp, err := s.sum.ServeGRPC(ctx, req)
	if err != nil {
		return nil, grpcEncodeError(err)
	}
	rep = rp.(*pb.SumReply)
	return rep, nil
}
//主要转发
func (s *GrpcServer) Concat(ctx context.Context, req *pb.ConcatRequest) (rep *pb.ConcatReply, err error) {
	_, rp, err := s.concat.ServeGRPC(ctx, req)
	if err != nil {
		return nil, grpcEncodeError(err)
	}
	rep = rp.(*pb.ConcatReply)
	return rep, nil
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
	return &pb.SumReply{Rs: reply.Rs}, grpcEncodeError(reply.Err)
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
	return &pb.ConcatReply{Rs: reply.Rs}, grpcEncodeError(reply.Err)
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