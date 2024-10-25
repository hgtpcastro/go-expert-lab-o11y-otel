package logrous

import (
	"github.com/hgtpcastro/go-expert-lab-o11y-otel/internal/pkg/logger"
	"github.com/hgtpcastro/go-expert-lab-o11y-otel/internal/pkg/logger/configs"
	"go.uber.org/fx"
)

// Module provided to fxlog
// https://uber-go.github.io/fx/modules.html
var Module = fx.Module("logrous_fx",
	// - order is not important in provide
	// - provide can have parameter and will resolve if registered
	// - execute its func only if it requested
	fx.Provide(
		fx.Annotate(
			NewLogrusLogger,
			fx.As(new(logger.Logger)),
		),
		configs.ProvideLogConfigs,
	))

var ModuleFunc = func(l logger.Logger) fx.Option {
	return fx.Module("logrous_fx",

		fx.Provide(configs.ProvideLogConfigs),
		fx.Supply(fx.Annotate(l, fx.As(new(logger.Logger)))),
	)
}
