package contracts

import (
	"context"

	"github.com/hgtpcastro/go-expert-lab-o11y-otel/internal/pkg/config"
	"go.uber.org/zap"
)

type EchoHttpServer interface {
	RunHttpServer() error
	GracefulShutdown(ctx context.Context) error
	ApplyVersioningFromHeader()
	RouteBuilder() *RouteBuilder
	Log() *zap.Logger
	Config() *config.Config
	SetupDefaultMiddlewares()
}
