package smetrics

import (
	"context"
	"fmt"
	"log/slog"

	"go.opentelemetry.io/contrib/bridges/otelslog"
	otelmetric "go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/noop"
	"go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/trace"
	oteltrace "go.opentelemetry.io/otel/trace"
)

// NewTelemetry creates a new telemetry instance.
func NewTelemetry(ctx context.Context, name, version string, opt ...Option) (*Telemetry, error) {
	res := newResource(name, version)

	t := &Telemetry{
		lp:     nil,
		mp:     nil,
		tp:     nil,
		log:    nil,
		meter:  nil,
		tracer: nil,
	}

	for _, o := range opt {
		o(svr)
	}

	lp, err := NewLoggerHTTPProvider(ctx, res)
	if err != nil {
		return nil, fmt.Errorf("failed to create logger: %w", err)
	}

	logger := otelslog.NewLogger(name, otelslog.WithLoggerProvider(lp))

	mp, err := NewMeterHTTPProvider(ctx, res)
	if err != nil {
		return nil, fmt.Errorf("failed to create meter: %w", err)
	}
	meter := mp.Meter(name)

	tp, err := NewTracerHTTPProvider(ctx, res)
	if err != nil {
		return nil, fmt.Errorf("failed to create tracer: %w", err)
	}
	tracer := tp.Tracer(name)

	return &Telemetry{
		lp:     lp,
		mp:     mp,
		tp:     tp,
		log:    logger,
		meter:  meter,
		tracer: tracer,
	}, nil
}

// Log logs a message
func (t *Telemetry) Log(ctx context.Context, level slog.Level, msg string, attrs ...slog.Attr) {
	t.log.LogAttrs(ctx, level, msg, attrs...)
}

// NewMeterInt64Histogram creates a new int64 histogram metric.
func (t *Telemetry) NewMeterInt64Histogram(metric Metric) (Int64Histogram, error) { //nolint:ireturn
	histogram, err := t.meter.Int64Histogram(
		metric.Name,
		otelmetric.WithDescription(metric.Description),
		otelmetric.WithUnit(metric.Unit),
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create histogram: %w", err)
	}

	return histogram, nil
}

// NewMeterInt64UpDownCounter creates a new int64 up down counter metric.
func (t *Telemetry) NewMeterInt64UpDownCounter(metric Metric) (Int64UpDownCounter, error) { //nolint:ireturn
	counter, err := t.meter.Int64UpDownCounter(
		metric.Name,
		otelmetric.WithDescription(metric.Description),
		otelmetric.WithUnit(metric.Unit),
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create counter: %w", err)
	}

	return counter, nil
}

// Trace starts a new span with the given name. The span must be ended by calling End.
func (t *Telemetry) Trace(ctx context.Context, name string) (context.Context, oteltrace.Span) { //nolint:ireturn
	//nolint: spancheck
	return t.tracer.Start(ctx, name)
}

// Shutdown shuts down the logger, meter, and tracer.
func (t *Telemetry) Shutdown(ctx context.Context) {
	t.lp.Shutdown(ctx)
	t.mp.Shutdown(ctx)
	t.tp.Shutdown(ctx)
}

var _ TelemetryProvider = (*Telemetry)(nil)
