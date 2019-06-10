package v1

import "github.com/opentracing/opentracing-go"

// Base ...
type Base struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type TracingBase struct {
	SpanContext opentracing.SpanContext
}
