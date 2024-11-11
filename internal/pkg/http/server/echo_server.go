package server

import (
	"context"
	"fmt"
	"strings"

	"github.com/hgtpcastro/go-expert-lab-o11y-otel/internal/pkg/http/server/contracts"
	"github.com/hgtpcastro/go-expert-lab-o11y-otel/internal/pkg/http/server/middlewares/log"
	oteltracing "github.com/hgtpcastro/go-expert-lab-o11y-otel/internal/pkg/http/server/middlewares/otel_tracing"
	"github.com/labstack/echo/v4/middleware"
	"go.opentelemetry.io/otel/propagation"

	"github.com/hgtpcastro/go-expert-lab-o11y-otel/internal/pkg/config"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type echoHttpServer struct {
	echo         *echo.Echo
	config       *config.Config
	log          *zap.Logger
	routeBuilder *contracts.RouteBuilder
}

func NewEchoHttpServer(
	config *config.Config,
	log *zap.Logger,
) contracts.EchoHttpServer {
	e := echo.New()
	e.HideBanner = true

	return &echoHttpServer{
		echo:         e,
		config:       config,
		log:          log,
		routeBuilder: contracts.NewRouteBuilder(e),
	}
}

func (s *echoHttpServer) RunHttpServer() error {
	// https://echo.labstack.com/guide/http_server/
	return s.echo.Start(s.config.Http.Port)
}

func (s *echoHttpServer) GracefulShutdown(ctx context.Context) error {
	err := s.echo.Shutdown(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (s *echoHttpServer) ApplyVersioningFromHeader() {
	s.echo.Pre(apiVersion)
}

func (s *echoHttpServer) RouteBuilder() *contracts.RouteBuilder {
	return s.routeBuilder
}

func (s *echoHttpServer) Log() *zap.Logger {
	return s.log
}

func (s *echoHttpServer) Config() *config.Config {
	return s.config
}

func (s *echoHttpServer) AddMiddlewares(middlewares ...echo.MiddlewareFunc) {
	if len(middlewares) > 0 {
		s.echo.Use(middlewares...)
	}
}

func (s *echoHttpServer) SetupDefaultMiddlewares() {
	skipper := func(c echo.Context) bool {
		return strings.Contains(c.Request().URL.Path, "swagger") ||
			strings.Contains(c.Request().URL.Path, "metrics") ||
			strings.Contains(c.Request().URL.Path, "health") ||
			strings.Contains(c.Request().URL.Path, "favicon.ico")
	}

	// log errors and information
	s.echo.Use(
		log.EchoLogger(
			s.log,
			log.WithSkipper(skipper),
		),
	)

	// propagator := newPropagator()

	s.echo.Use(
		oteltracing.HttpTrace(
			oteltracing.WithSkipper(skipper),
			oteltracing.WithServiceName(s.config.Tracing.ServiceName),
			// oteltracing.WithTracerProvider(),
			// oteltracing.WithPropagators(propagator),
			// oteltracing.WithInstrumentationName("pdi-2024-02"),
		),
	)

	s.echo.Use(middleware.RequestID())
}

// APIVersion Header Based Versioning
func apiVersion(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := c.Request()
		headers := req.Header

		apiVersion := headers.Get("version")

		req.URL.Path = fmt.Sprintf("/%s%s", apiVersion, req.URL.Path)

		return next(c)
	}
}

func newPropagator() propagation.TextMapPropagator {
	return propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	)
}
