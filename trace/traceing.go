package traceutil

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	trace2 "github.com/yueja/link-tracking/middleware/trace"
	"log"
	"reflect"
)

// GetTraceIDByCtx 从 Ctx 中获取 TraceID
func (t trace) getTraceIDByCtx(ctx context.Context) string {
	ctx = t.popCtxByOriginCtx(ctx)
	sp := opentracing.SpanFromContext(ctx)
	if sp == nil {
		if trace2.GetCtxValue(ctx) != nil {
			return trace2.GetCtxValue(ctx).Common[trace2.CtxValueCommonKeyTraceID]
		}
		log.Println("noSpanId getTraceIDByCtx spanFromContext")
		return ""
	}
	spCtx := sp.Context()
	switch spCtx := spCtx.(type) {
	case jaeger.SpanContext:
		return spCtx.TraceID().String()
	default:
		log.Println("NoSpanId", reflect.TypeOf(spCtx).String())
		return ""
	}
}

// GetSpanIDByCtx 从 Ctx 中获取 SpanID
func (t trace) getSpanIDByCtx(ctx context.Context) string {
	ctx = t.popCtxByOriginCtx(ctx)
	sp := opentracing.SpanFromContext(ctx)
	if sp == nil {
		if trace2.GetCtxValue(ctx) != nil {
			return trace2.GetCtxValue(ctx).Common[trace2.CtxValueCommonKeySpanID]
		}
		log.Println("noSpanId getSpanIDByCtx spanFromContext")
		return ""
	}
	spCtx := sp.Context()
	switch spCtx := spCtx.(type) {
	case jaeger.SpanContext:
		return spCtx.SpanID().String()
	default:
		log.Println("NoSpanId", reflect.TypeOf(spCtx).String())
		return ""
	}
}
