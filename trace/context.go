package traceutil

import (
	"context"
	"github.com/sirupsen/logrus"
)

func (t trace) HandleCtx(ctx context.Context) (fields logrus.Fields) {
	fields = make(logrus.Fields, 0)
	t.addTraceIDField(ctx, fields) // traceId
	t.addSpanIDField(ctx, fields)  // spanId
	return
}

// addTraceIDField 获得 ctx 中的 traceId, 没有开启默认取 trace_id
func (t trace) addTraceIDField(ctx context.Context, fields logrus.Fields) {
	value := t.getTraceIDByCtx(ctx)
	fields["trace_id"] = value
	return
}

// addSpanIDField 获得 ctx 中的 spanId
func (t trace) addSpanIDField(ctx context.Context, fields logrus.Fields) {
	value := t.getSpanIDByCtx(ctx)
	fields["span_id"] = value
	return
}
