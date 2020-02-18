package message

import (
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
)

var (
	_ httptransport.Headerer    = (*EchoResponse)(nil)
	_ httptransport.StatusCoder = (*EchoResponse)(nil)
	_ httptransport.Headerer    = (*SayHelloResponse)(nil)
	_ httptransport.StatusCoder = (*SayHelloResponse)(nil)
)

// SumResponse collects the response values for the Sum method.
type EchoResponse struct {
	Rs  string `json:"rs"`
	Err error `json:"err"`
}

func (r EchoResponse) StatusCode() int {
	return http.StatusOK // TBA
}

func (r EchoResponse) Headers() http.Header {
	return http.Header{}
}

// ConcatResponse collects the response values for the Concat method.
type SayHelloResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

func (r SayHelloResponse) StatusCode() int {
	return http.StatusOK // TBA
}

func (r SayHelloResponse) Headers() http.Header {
	return http.Header{}
}
