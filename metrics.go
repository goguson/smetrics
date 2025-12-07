package smetrics

import (
	"context"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	otelmetric "go.opentelemetry.io/otel/metric"
	oteltrace "go.opentelemetry.io/otel/trace"
)

// Metric represents a metric that can be collected by the server.
type Metric struct {
	Name        string
	Unit        string
	Description string
}
type Int64Histogram interface {
	Record(ctx context.Context, incr int64, options ...otelmetric.RecordOption)
}
type Int64UpDownCounter interface {
	Add(ctx context.Context, incr int64, options ...otelmetric.AddOption)
}

type Tracer interface {
	Start(ctx context.Context, spanName string, opts ...oteltrace.SpanStartOption) (context.Context, Span)
}

type Span interface {
	// End completes the Span. The Span is considered complete and ready to be
	// delivered through the rest of the telemetry pipeline after this method
	// is called. Therefore, updates to the Span are not allowed after this
	// method has been called.
	End(options ...oteltrace.SpanEndOption)

	// AddEvent adds an event with the provided name and options.
	AddEvent(name string, options ...oteltrace.EventOption)

	// AddLink adds a link.
	// Adding links at span creation using WithLinks is preferred to calling AddLink
	// later, for contexts that are available during span creation, because head
	// sampling decisions can only consider information present during span creation.
	AddLink(link oteltrace.Link)

	// IsRecording returns the recording state of the Span. It will return
	// true if the Span is active and events can be recorded.
	IsRecording() bool

	// RecordError will record err as an exception span event for this span. An
	// additional call to SetStatus is required if the Status of the Span should
	// be set to Error, as this method does not change the Span status. If this
	// span is not being recorded or err is nil then this method does nothing.
	RecordError(err error, options ...oteltrace.EventOption)

	// SpanContext returns the SpanContext of the Span. The returned SpanContext
	// is usable even after the End method has been called for the Span.
	SpanContext() oteltrace.SpanContext

	// SetStatus sets the status of the Span in the form of a code and a
	// description, provided the status hasn't already been set to a higher
	// value before (OK > Error > Unset). The description is only included in a
	// status when the code is for an error.
	SetStatus(code codes.Code, description string)

	// SetName sets the Span name.
	SetName(name string)

	// SetAttributes sets kv as attributes of the Span. If a key from kv
	// already exists for an attribute of the Span it will be overwritten with
	// the value contained in kv.
	SetAttributes(kv ...attribute.KeyValue)

	// TracerProvider returns a TracerProvider that can be used to generate
	// additional Spans on the same telemetry pipeline as the current Span.
	TracerProvider() TracerProvider
}

type TracerProvider interface {
	Tracer(name string, options ...oteltrace.TracerOption) Tracer
}
