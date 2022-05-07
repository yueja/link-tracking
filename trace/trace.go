package traceutil

import (
	"context"
	"github.com/sirupsen/logrus"
)

type trace struct{}

type TraceServer interface {
	popCtxByOriginCtx(ctx context.Context) (targetCtx context.Context)
	HandleCtx(ctx context.Context) (fields logrus.Fields)
}

func NewTraceServer() *trace {
	return new(trace)
}
