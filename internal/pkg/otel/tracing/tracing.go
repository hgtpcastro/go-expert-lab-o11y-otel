package tracing

import (
	"context"

	"emperror.dev/errors"
	"github.com/hgtpcastro/go-expert-lab-o11y-otel/internal/pkg/config"
	"github.com/samber/lo"
	"go.opentelemetry.io/contrib/propagators/ot"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/zipkin"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
)

type TracingOpenTelemetry struct {
	//config      *TracingOptions
	config    *config.Config
	appTracer AppTracer
	provider  *tracesdk.TracerProvider
}

func NewOtelTracing(
	//config *TracingOptions,
	config *config.Config,
) (*TracingOpenTelemetry, error) {
	otelTracing := &TracingOpenTelemetry{
		config: config,
	}

	resource, err := otelTracing.newResource()
	if err != nil {
		return nil, errors.WrapIf(err, "failed to create resource")
	}

	appTracer, err := otelTracing.initTracer(resource)
	if err != nil {
		return nil, err
	}

	otelTracing.appTracer = appTracer

	return otelTracing, nil
}

func (o *TracingOpenTelemetry) Shutdown(ctx context.Context) error {
	return o.provider.Shutdown(ctx)
}

func (o *TracingOpenTelemetry) newResource() (*resource.Resource, error) {
	resource, err := resource.New(context.Background(),
		resource.WithFromEnv(),
		resource.WithTelemetrySDK(),
		resource.WithHost(),
		resource.WithOS(),
		resource.WithSchemaURL(semconv.SchemaURL),
		resource.WithAttributes(
			semconv.ServiceName(o.config.Tracing.ServiceName),
			semconv.ServiceVersion(o.config.Tracing.Version),
			attribute.Int64("ID", o.config.Tracing.Id),
			attribute.String("environment", "develop"),
			semconv.TelemetrySDKVersionKey.String("v1.26.0"), // semconv version
			semconv.TelemetrySDKLanguageGo,
		))

	return resource, err
}

func (o *TracingOpenTelemetry) initTracer(
	resource *resource.Resource,
) (AppTracer, error) {
	exporters, err := o.configExporters()
	if err != nil {
		return nil, err
	}

	var sampler tracesdk.Sampler
	if o.config.Tracing.AlwaysOnSampler {
		sampler = tracesdk.AlwaysSample()
	} else {
		sampler = tracesdk.NeverSample()
	}

	batchExporters := lo.Map(
		exporters,
		func(item tracesdk.SpanExporter, index int) tracesdk.TracerProviderOption {
			return tracesdk.WithBatcher(item)
		},
	)

	opts := append(
		batchExporters,
		tracesdk.WithResource(resource),
		tracesdk.WithSampler(sampler),
	)

	provider := tracesdk.NewTracerProvider(opts...)

	otel.SetTracerProvider(provider)
	o.provider = provider

	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(
			ot.OT{}, // should be placed before `TraceContext` for preventing conflict
			propagation.Baggage{},
			propagation.TraceContext{},
		),
	)

	appTracer := NewAppTracer(o.config.Tracing.InstrumentationName)

	return appTracer, nil
}

func (o *TracingOpenTelemetry) configExporters() ([]tracesdk.SpanExporter, error) {
	// ctx := context.Background()
	// traceOpts := []otlptracegrpc.Option{
	// 	otlptracegrpc.WithTimeout(5 * time.Second),
	// 	otlptracegrpc.WithInsecure(),
	// }

	var exporters []tracesdk.SpanExporter

	if o.config.Tracing.ZipkinExporterOptions != nil {
		zipkinExporter, err := zipkin.New(
			o.config.Tracing.ZipkinExporterOptions.Url,
		)
		if err != nil {
			return nil, errors.WrapIf(
				err,
				"failed to create exporter for zipkin",
			)
		}

		exporters = append(exporters, zipkinExporter)
	}

	return exporters, nil
}
