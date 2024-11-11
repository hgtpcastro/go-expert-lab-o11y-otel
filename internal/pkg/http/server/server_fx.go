package server

import (
	"context"
	"errors"
	"net/http"

	"github.com/hgtpcastro/go-expert-lab-o11y-otel/internal/pkg/http/server/contracts"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Module(
	"httpserverfx",
	fx.Provide(NewEchoHttpServer),
	fx.Options(fx.Invoke(registerHooks)),
)

func registerHooks(
	lc fx.Lifecycle,
	echoServer contracts.EchoHttpServer,
	log *zap.Logger,
) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := echoServer.RunHttpServer(); !errors.Is(
					err,
					http.ErrServerClosed,
				) {
					log.Sugar().Fatalf(
						"(EchoHttpServer.RunHttpServer) error in running server: {%v}",
						err,
					)
				}
			}()
			echoServer.Log().Sugar().Infof(
				"%s is listening on Host:{%s} Http PORT: {%s}",
				echoServer.Config().Http.Name,
				echoServer.Config().Http.Host,
				echoServer.Config().Http.Port,
			)

			return nil
		},
		OnStop: func(ctx context.Context) error {
			if err := echoServer.GracefulShutdown(ctx); err != nil {
				echoServer.Log().Sugar().
					Errorf("error shutting down echo server: %v", err)
			} else {
				echoServer.Log().Info("echo server shutdown gracefully")
			}
			return nil
		},
	})

}
