package noop

import (
	"context"

	otelmetric "go.opentelemetry.io/otel/metric"
)

var _ smetrics.Int64UpDownCounter = (*UpDownCounter)(nil)

type UpDownCounter struct {
}

func (n UpDownCounter) Add(ctx context.Context, incr int64, options ...otelmetric.AddOption) {

}
