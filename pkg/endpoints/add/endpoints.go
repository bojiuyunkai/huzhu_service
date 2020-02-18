package add

import (
	"context"
	

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"huzhu_service/pkg/endpoints"
	
	service "huzhu_service/pkg/svc/add"
)

// Endpoints collects all of the endpoints that compose the addsvc service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
type Endpoints struct {
	SumEndpoint    endpoint.Endpoint `json:""`
	ConcatEndpoint endpoint.Endpoint `json:""`
}

// New return a new instance of the endpoint that wraps the provided service.
func New(svc service.AddsvcService, logger log.Logger ) (ep Endpoints) {
	
	var sumEndpoint endpoint.Endpoint = NewSumEndpoint(svc,logger);
	ep.SumEndpoint = sumEndpoint

	var concatEndpoint endpoint.Endpoint = NewConcatEndpoint(svc,logger);
	ep.ConcatEndpoint = concatEndpoint;
	return ep
}
func NewSumEndpoint(svc service.AddsvcService, logger log.Logger) endpoint.Endpoint{
	var sumEndpoint endpoint.Endpoint
	method := "sum"
	sumEndpoint = MakeSumEndpoint(svc)
	
	sumEndpoint = endpoints.LoggingMiddleware(log.With(logger, "method", method))(sumEndpoint)
	return sumEndpoint;
}
func NewConcatEndpoint(svc service.AddsvcService, logger log.Logger) endpoint.Endpoint{
	var concatEndpoint endpoint.Endpoint
	method := "concat"
	concatEndpoint = MakeConcatEndpoint(svc)
	
	concatEndpoint = endpoints.LoggingMiddleware(log.With(logger, "method", method))(concatEndpoint)
	
	return concatEndpoint;
}

// MakeSumEndpoint returns an endpoint that invokes Sum on the service.
// Primarily useful in a server.
func MakeSumEndpoint(svc service.AddsvcService) (ep endpoint.Endpoint) {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SumRequest)
		if err := req.validate(); err != nil {
			return SumResponse{}, err
		}
		rs, err := svc.Sum(ctx, req.A, req.B)
		return SumResponse{Rs: rs}, err
	}
}

// MakeConcatEndpoint returns an endpoint that invokes Concat on the service.
// Primarily useful in a server.
func MakeConcatEndpoint(svc service.AddsvcService) (ep endpoint.Endpoint) {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ConcatRequest)
		if err := req.validate(); err != nil {
			return ConcatResponse{}, err
		}
		rs, err := svc.Concat(ctx, req.A, req.B)
		return ConcatResponse{Rs: rs}, err
	}
}

// Sum implements the service interface, so Endpoints may be used as a service.
// This is primarily useful in the context of a client library.
func (e Endpoints) Sum(ctx context.Context, a int64, b int64) (rs int64, err error) {
	resp, err := e.SumEndpoint(ctx, SumRequest{A: a, B: b})
	if err != nil {
		return
	}
	response := resp.(SumResponse)
	return response.Rs, nil
}



// Concat implements the service interface, so Endpoints may be used as a service.
// This is primarily useful in the context of a client library.
func (e Endpoints) Concat(ctx context.Context, a string, b string) (rs string, err error) {
	resp, err := e.ConcatEndpoint(ctx, ConcatRequest{A: a, B: b})
	if err != nil {
		return
	}
	response := resp.(ConcatResponse)
	return response.Rs, nil
}
