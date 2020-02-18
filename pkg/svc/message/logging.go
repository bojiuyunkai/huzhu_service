package message

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type loggingMiddleware struct {
	logger log.Logger    `json:""`
	next   MsgsvcService `json:""`
}

// LoggingMiddleware takes a logger as a dependency
// and returns a ServiceMiddleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next MsgsvcService) MsgsvcService {
		return loggingMiddleware{level.Info(logger), next}
	}
}

func (lm loggingMiddleware) Echo(ctx context.Context, word string) (rs string, err error) {
	defer func(begin time.Time) {
		lm.logger.Log("method", "Echo", "word", word,  "err", err)
	}(time.Now())

	return lm.next.Echo(ctx, word)
}

func (lm loggingMiddleware) SayHello(ctx context.Context, saidword string,want string) (rs string, err error) {
	defer func(begin time.Time) {
		lm.logger.Log("method", "SayHello", "saidword", saidword, "want", want, "err", err)
	}(time.Now())

	return lm.next.SayHello(ctx, saidword,want)
}