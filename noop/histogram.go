package noop

import (
	"context"

	"github.com/goguson/smetrics"

	otelmetric "go.opentelemetry.io/otel/metric"
)

var _ smetrics.Int64Histogram = (*Histogram)(nil)

type Histogram struct {
}

func (n Histogram) Record(ctx context.Context, incr int64, options ...otelmetric.RecordOption) {

}
