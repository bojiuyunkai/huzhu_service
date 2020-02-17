package transports
import (
	"context"
	"huzhu_service/pb"

)
func (s *GrpcServer) Echo(ctx context.Context, req *pb.EchoRequest) (rep *pb.EchoReply, err error) {
	_, rp, err := s.echo.ServeGRPC(ctx, req)
	if err != nil {
		return nil, grpcEncodeError(err)
	}
	rep = rp.(*pb.EchoReply)
	return rep, nil
}

func (s *GrpcServer) SayHello(ctx context.Context, req *pb.SayHelloRequest) (rep *pb.SayHelloReply, err error) {
	_, rp, err := s.sayhello.ServeGRPC(ctx, req)
	if err != nil {
		return nil, grpcEncodeError(err)
	}
	rep = rp.(*pb.SayHelloReply)
	return rep, nil
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