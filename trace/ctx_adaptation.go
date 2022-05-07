package traceutil

import (
	"context"
	"github.com/gin-gonic/gin"
)

// popCtxByOriginCtx 取出目标 ctx 并生成适配器
func (t trace) popCtxByOriginCtx(ctx context.Context) (targetCtx context.Context) {
	switch ctx := ctx.(type) {
	case *gin.Context:
		// 从 gin.Context 中获取 http Context
		targetCtx = ctx.Request.Context()
	default:
		targetCtx = ctx
	}
	return targetCtx
}
