package contracts

import (
	"github.com/hgtpcastro/go-expert-lab-o11y-otel/internal/pkg/configs/environment"
	"github.com/hgtpcastro/go-expert-lab-o11y-otel/internal/pkg/logger"
	"go.uber.org/fx"
)

type ApplicationBuilder interface {
	// ProvideModule register modules directly instead and modules should not
	// register with `provided` function
	ProvideModule(module fx.Option)

	// Provide register functions constructors as dependency resolver
	Provide(constructors ...interface{})
	Decorate(constructors ...interface{})
	Build() Application

	GetProvides() []interface{}
	GetDecorates() []interface{}
	Options() []fx.Option
	Logger() logger.Logger
	Environment() environment.Environment
}
