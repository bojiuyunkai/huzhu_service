package message

import (
	"context"

	//pb "huzhu_service/pb"

)

// Sum implements the service interface, so Endpoints may be used as a service.
// This is primarily useful in the context of a client library.
func (e Endpoints) Echo(ctx context.Context,word string) (rs string, err error) {

	resp, err := e.EchoEndpoint(ctx, EchoRequest{Word:word})
	if err != nil {
		return
	}
	response := resp.(EchoResponse)
	return response.Rs, nil
}



// Concat implements the service interface, so Endpoints may be used as a service.
// This is primarily useful in the context of a client library.
func (e Endpoints) SayHello(ctx context.Context, saidword string, want string) (rs string, err error) {
	resp, err := e.SayHelloEndpoint(ctx, SayHelloRequest{Saidword: saidword, Want: want})
	if err != nil {
		return
	}
	response := resp.(SayHelloResponse)
	return response.Rs, nil
}
