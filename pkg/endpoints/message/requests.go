package message

import def "huzhu_service/pkg/def"

type Request interface {
	validate() error
}

const (
	intMax = 1<<31 - 1
	intMin = -(intMax + 1)
	maxLen = 10
)

// SumRequest collects the request parameters for the Sum method.
type EchoRequest struct {
	Word string `json:"word"`
}

func (r EchoRequest) validate() error {
	if r.Word == "" {
		return def.ErrTwoZeroes
	}
	
	return nil
}

// ConcatRequest collects the request parameters for the Concat method.
type SayHelloRequest struct {
	Saidword string `json:"saidword"`
	Want string `json:"want"`
}

func (r SayHelloRequest) validate() error {
	
	return nil
}
