package trace

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// WithTrace 链路追踪
func WithTrace() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		traceID := ctx.GetHeader("UBER-TRACE-ID")
		if traceID == "" {
			traceID = uuid.New().String()
		}
		spanID := uuid.New().String()
		value := make(map[CtxValueCommonKey]string)
		// 设置新的 ctx value
		cv := GetCtxValue(ctx.Request.Context())
		if cv != nil {
			value = cv.Common
			if value[CtxValueCommonKeyTraceID] == "" {
				value[CtxValueCommonKeyTraceID] = traceID
			}
			value[CtxValueCommonKeySpanID] = spanID
		} else {
			value = NewCtxTraceAndSpan(traceID, spanID)
		}
		ctxValue := NewCtxValue(value)
		cfCtx := SetCtxValue(ctx.Request.Context(), ctxValue)
		ctx.Request = ctx.Request.WithContext(cfCtx)
		ctx.Next()
	}
}

// GetCtxValue 获取ctx value
func GetCtxValue(ctx context.Context) *CtxValue {
	cv := ctx.Value(CtxValueKeyTraceAndSpan)
	if v, ok := cv.(*CtxValue); ok {
		if v.Common == nil {
			v.Common = make(map[CtxValueCommonKey]string)
		}
		return v
	}
	return &CtxValue{Common: make(map[CtxValueCommonKey]string)}
}

// NewCtxTraceAndSpan 生成ctx trace
func NewCtxTraceAndSpan(traceId, spanId string) map[CtxValueCommonKey]string {
	return map[CtxValueCommonKey]string{
		CtxValueCommonKeyTraceID: traceId,
		CtxValueCommonKeySpanID:  spanId,
	}
}

// NewCtxValue 创建cxt value
func NewCtxValue(common map[CtxValueCommonKey]string) *CtxValue {
	if common == nil {
		common = make(map[CtxValueCommonKey]string)
	}
	return &CtxValue{
		Common: common,
	}
}

// SetCtxValue 设置ctx value
func SetCtxValue(ctx context.Context, value *CtxValue) context.Context {
	ctx = context.WithValue(ctx, CtxValueKeyTraceAndSpan, value)
	return ctx
}
