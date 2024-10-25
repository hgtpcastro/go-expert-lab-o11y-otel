package configs

import (
	"github.com/hgtpcastro/go-expert-lab-o11y-otel/internal/pkg/configs/environment"
	"go.uber.org/fx"
)

// Module provided to fxlog
// https://uber-go.github.io/fx/modules.html
var Module = fx.Module(
	"configs_fx",
	fx.Provide(func() environment.Environment {
		return environment.ConfigAppEnv()
	}),
)

var ModuleFunc = func(e environment.Environment) fx.Option {
	return fx.Module(
		"configs_fx",
		fx.Provide(func() environment.Environment {
			return environment.ConfigAppEnv(e)
		}),
	)
}
