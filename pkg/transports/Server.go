package transports
import (
"context"
"huzhu_service/pb"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	
)

//rpc server实现
type GrpcServer struct {
	sum    grpctransport.Handler `json:""`
	concat grpctransport.Handler `json:""`
	echo grpctransport.Handler `json:""`
	sayhello grpctransport.Handler `json:""`
	
}

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
//主要转发
func (s *GrpcServer) Echo(ctx context.Context, req *pb.EchoRequest) (rep *pb.EchoReply, err error) {
	_, rp, err := s.echo.ServeGRPC(ctx, req)
	if err != nil {
		return nil, grpcEncodeError(err)
	}
	rep = rp.(*pb.EchoReply)
	return rep, nil
}
//主要转发
func (s *GrpcServer) SayHello(ctx context.Context, req *pb.SayHelloRequest) (rep *pb.SayHelloReply, err error) {
	_, rp, err := s.sayhello.ServeGRPC(ctx, req)
	if err != nil {
		return nil, grpcEncodeError(err)
	}
	rep = rp.(*pb.SayHelloReply)
	return rep, nil
}
