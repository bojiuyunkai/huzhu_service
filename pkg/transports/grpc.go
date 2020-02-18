package transports
import (
	"github.com/go-kit/kit/log"
	
	grpctransport "github.com/go-kit/kit/transport/grpc"
	//"github.com/go-kit/kit/endpoint"
	//"google.golang.org/grpc"
	
	 addtransports "huzhu_service/pkg/transports/add"
	 msgtransports "huzhu_service/pkg/transports/message"
	 
	 //"huzhu_service/pb"
	 //service "huzhu_service/pkg/svc/add"
	

)

// MakeGRPCServer makes a set of endpoints available as a gRPC server.
func MakeGRPCServer(logger log.Logger) (*GrpcServer) { // Zipkin GRPC Server Trace can either be instantiated per gRPC method with a
	// provided operation name or a global tracing service can be instantiated
	// In this example, we demonstrate a global Zipkin tracing service with
	// Go kit gRPC Interceptor

	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
	}
	addHandlers:=addtransports.NewGrpcHandler(logger,options);
	msgHandlers:= msgtransports.NewGrpcHandler(logger,options);
	return &GrpcServer{
		sum: addHandlers["sum"],
		concat: addHandlers["concat"],
		echo: msgHandlers["echo"],
		sayhello:msgHandlers["sayhello"],

	}
}

// NewGRPCClient returns an AddService backed by a gRPC server at the other end
// of the conn. The caller is responsible for constructing the conn, and
// eventually closing the underlying transport. We bake-in certain middlewares,
/*// implementing the client library pattern.
func NewGRPCClient(conn *grpc.ClientConn,  logger log.Logger) service.AddsvcService { // We construct a single ratelimiter middleware, to limit the total outgoing
	// QPS from this client to all methods on the remote instance. We also
	// construct per-endpoint circuitbreaker middlewares to demonstrate how
	// that's done, although they could easily be combined into a single breaker
	// for the entire remote instance, too.
	//limiter := ratelimit.NewErroringLimiter(rate.NewLimiter(rate.Every(time.Second), 100))

	// Zipkin GRPC Client Trace can either be instantiated per gRPC method with a
	// provided operation name or a global tracing client can be instantiated
	// without an operation name and fed to each Go kit client as ClientOption.
	// In the latter case, the operation name will be the endpoint's grpc method
	// path.
	//
	// In this example, we demonstrace a global tracing client.
	//zipkinClient := zipkin.GRPCClientTrace(zipkinTracer)

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
}*/


