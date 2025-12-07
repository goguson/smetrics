package noop

import (
	"smetrics"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	oteltrace "go.opentelemetry.io/otel/trace"
)

var _ smetrics.Span = (*Span)(nil)

type Span struct{}

func (n Span) End(options ...oteltrace.SpanEndOption) {

}

func (n Span) AddEvent(name string, options ...oteltrace.EventOption) {

}

func (n Span) AddLink(link oteltrace.Link) {

}

func (n Span) IsRecording() bool {
	return false
}

func (n Span) RecordError(err error, options ...oteltrace.EventOption) {
}

func (n Span) SpanContext() oteltrace.SpanContext {
	return oteltrace.SpanContext{}
}

func (n Span) SetStatus(code codes.Code, description string) {

}

func (n Span) SetName(name string) {

}

func (n Span) SetAttributes(kv ...attribute.KeyValue) {
}

func (n Span) TracerProvider() smetrics.TracerProvider {
	return TracerProvider{}
}
