package add

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
type SumRequest struct {
	A int64 `json:"a"`
	B int64 `json:"b"`
}

func (r SumRequest) validate() error {
	if r.A == 0 && r.B == 0 {
		return def.ErrTwoZeroes
	}
	if (r.B > 0 && r.A > (intMax-r.B)) || (r.B < 0 && r.A < (intMin-r.B)) {
		return def.ErrIntOverflow
	}
	return nil
}

// ConcatRequest collects the request parameters for the Concat method.
type ConcatRequest struct {
	A string `json:"a"`
	B string `json:"b"`
}

func (r ConcatRequest) validate() error {
	if len(r.A)+len(r.B) > maxLen {
		return def.ErrMaxSizeExceeded
	}
	return nil
}
