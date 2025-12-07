package noop

import (
	"context"
	"smetrics"

	oteltrace "go.opentelemetry.io/otel/trace"
)

var _ smetrics.Tracer = (*Tracer)(nil)

type Tracer struct{}

func (n Tracer) Start(ctx context.Context, spanName string, opts ...oteltrace.SpanStartOption) (context.Context, smetrics.Span) {
	return ctx, Span{}
}
