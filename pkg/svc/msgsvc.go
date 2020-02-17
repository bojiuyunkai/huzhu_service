package svc

import (
	"context"

	"github.com/go-kit/kit/log"
)

// Service describes a service that adds things together
// Implement yor service methods methods.
// e.x: Foo(ctx context.Context, s string)(rs string, err error)
type MsgsvcService interface {
	Echo(ctx context.Context, word string) (rs string, err error)
	SayHello(ctx context.Context, saidword string, want string) (rs string, err error)
}

// Middleware describes a service (as opposed to endpoint) middleware.
type MsgMiddleware func(MsgsvcService) MsgsvcService



// the concrete implementation of service interface
type stubMsgService struct {
	logger log.Logger `json:"logger"`
}

// New return a new instance of the service.
// If you want to add service middleware this is the place to put them.
func NewMsgSvc(logger log.Logger) (s MsgsvcService) {
	var service MsgsvcService
	{
		service = &stubMsgService{logger: logger}
		service = MloggingMiddleware(logger)(service)
	}
	return service
}

// Implement the business logic of Sum
func (this *stubMsgService) Echo(ctx context.Context,word string) (rs string, err error) {
	return "echo "+word,err;
}

// Implement the business logic of Concat
func (this *stubMsgService) SayHello(ctx context.Context, saidword string,want string) (rs string, err error) {
	return saidword+want,err;
}
