package tracing

import (
	"context"

	"go.opentelemetry.io/otel/trace"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var (
	Module = fx.Module( //nolint:gochecknoglobals
		"oteltracingfx",
		tracingProviders,
		tracingInvokes,
	)

	tracingProviders = fx.Options(fx.Provide( //nolint:gochecknoglobals
		NewOtelTracing,
		fx.Annotate(
			provideTracer,
			fx.ParamTags(`optional:"true"`),
			fx.As(new(AppTracer)),
			fx.As(new(trace.Tracer)),
		),
	))

	tracingInvokes = fx.Options(
		fx.Invoke(registerHooks),
	) //nolint:gochecknoglobals
)

func provideTracer(
	tracingOtel *TracingOpenTelemetry,
) AppTracer {
	return tracingOtel.appTracer
}

// we don't want to register any dependencies here, its func body should execute always even we don't request for that, so we should use `invoke`
func registerHooks(
	lc fx.Lifecycle,
	logger *zap.Logger,
	tracingOtel *TracingOpenTelemetry,
) {
	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			if err := tracingOtel.Shutdown(ctx); err != nil {
				logger.Sugar().Errorf("error in shutting down trace provider: %v", err)
			} else {
				logger.Info("trace provider shutdown gracefully")
			}

			return nil
		},
	})
}
