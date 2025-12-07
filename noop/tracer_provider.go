package noop

import (
	"github.com/goguson/smetrics"

	oteltrace "go.opentelemetry.io/otel/trace"
)

var _ smetrics.TracerProvider = (*TracerProvider)(nil)

type TracerProvider struct{}

func (n TracerProvider) Tracer(name string, options ...oteltrace.TracerOption) smetrics.Tracer {
	return Tracer{}
}
