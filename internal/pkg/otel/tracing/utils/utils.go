package utils

import (
	"context"
	"net/http"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/semconv/v1.13.0/httpconv"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"go.opentelemetry.io/otel/trace"
)

type traceContextKeyType int

const parentSpanKey traceContextKeyType = iota + 1

func ContextWithParentSpan(
	parent context.Context,
	span trace.Span,
) context.Context {
	return context.WithValue(parent, parentSpanKey, span)
}

// HttpTraceStatusFromSpanWithCode create an error span with specific status code if we have an error and a successful span when error is nil with a specific status
func HttpTraceStatusFromSpanWithCode(
	span trace.Span,
	err error,
	code int,
) error {
	if err != nil {
		// stackTraceError := errorUtils.ErrorsWithStack(err)

		// https://opentelemetry.io/docs/instrumentation/go/manual/#record-errors
		span.SetAttributes(
			attribute.String("exception.message", err.Error()),
			// attribute.String(telemetrytags.Exceptions.Stacktrace, stackTraceError),
		)
		span.RecordError(err)
	}

	if code > 0 {
		// httpconv doesn't exist in semconv v1.21.0, and it moved to `opentelemetry-go-contrib` pkg
		// https://github.com/open-telemetry/opentelemetry-go/pull/4362
		// https://github.com/open-telemetry/opentelemetry-go/issues/4081
		// using ClientStatus instead of ServerStatus for consideration of 4xx status as error
		span.SetStatus(httpconv.ClientStatus(code))
		span.SetAttributes(semconv.HTTPStatusCode(code))
	} else {
		span.SetStatus(codes.Error, "")
		span.SetAttributes(semconv.HTTPStatusCode(http.StatusInternalServerError))
	}

	return err
}
