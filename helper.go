package api

import (
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
)

// GetTraceInfo
func GetTraceInfo(ctx opentracing.SpanContext) string {
	ret := ""
	if jaegerSpanCtx, ok := ctx.(jaeger.SpanContext); ok {
		ret = jaegerSpanCtx.String()
	}
	return ret
}
