package trace

// CtxValueCommonKey ctx common value key
type CtxValueCommonKey string

// 常量定义
const (
	CtxValueCommonKeySpanID  CtxValueCommonKey = "span_id"
	CtxValueCommonKeyTraceID CtxValueCommonKey = "trace_id"
)

// CtxValueKey ctx value key
type CtxValueKey string

// CtxValueKeyTraceAndSpan 常量定义
const (
	CtxValueKeyTraceAndSpan CtxValueKey = "CtxValueTraceAndSpan"
)

// CtxValue 上下文内容
type CtxValue struct {
	Common map[CtxValueCommonKey]string
}
