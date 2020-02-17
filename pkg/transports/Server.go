package transports
import (

	grpctransport "github.com/go-kit/kit/transport/grpc"
	
)

//rpc server实现
type GrpcServer struct {
	sum    grpctransport.Handler `json:""`
	concat grpctransport.Handler `json:""`
	echo grpctransport.Handler `json:""`
	sayhello grpctransport.Handler `json:""`
	
}

