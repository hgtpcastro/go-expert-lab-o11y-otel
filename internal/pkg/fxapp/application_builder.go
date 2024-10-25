package fxapp

import (
	"github.com/hgtpcastro/go-expert-lab-o11y-otel/internal/pkg/configs/environment"
	"github.com/hgtpcastro/go-expert-lab-o11y-otel/internal/pkg/fxapp/contracts"
	"github.com/hgtpcastro/go-expert-lab-o11y-otel/internal/pkg/logger"
	loggerConfigs "github.com/hgtpcastro/go-expert-lab-o11y-otel/internal/pkg/logger/configs"
	"github.com/hgtpcastro/go-expert-lab-o11y-otel/internal/pkg/logger/logrous"
	"github.com/hgtpcastro/go-expert-lab-o11y-otel/internal/pkg/logger/models"
	"github.com/hgtpcastro/go-expert-lab-o11y-otel/internal/pkg/logger/zap"
	"go.uber.org/fx"
)

type applicationBuilder struct {
	provides    []interface{}
	decorates   []interface{}
	options     []fx.Option
	logger      logger.Logger
	environment environment.Environment
}

func NewApplicationBuilder(environments ...environment.Environment) contracts.ApplicationBuilder {
	env := environment.ConfigAppEnv(environments...)

	var logger logger.Logger
	logoption, err := loggerConfigs.ProvideLogConfigs(env)
	if err != nil || logoption == nil {
		logger = zap.NewZapLogger(logoption, env)
	} else if logoption.LogType == models.Logrus {
		logger = logrous.NewLogrusLogger(logoption, env)
	} else {
		logger = zap.NewZapLogger(logoption, env)
	}

	return &applicationBuilder{logger: logger, environment: env}
}

func (a *applicationBuilder) ProvideModule(module fx.Option) {
	a.options = append(a.options, module)
}

func (a *applicationBuilder) Provide(constructors ...interface{}) {
	a.provides = append(a.provides, constructors...)
}

func (a *applicationBuilder) Decorate(constructors ...interface{}) {
	a.decorates = append(a.decorates, constructors...)
}

func (a *applicationBuilder) Build() contracts.Application {
	app := NewApplication(a.provides, a.decorates, a.options, a.logger, a.environment)

	return app
}

func (a *applicationBuilder) GetProvides() []interface{} {
	return a.provides
}

func (a *applicationBuilder) GetDecorates() []interface{} {
	return a.decorates
}

func (a *applicationBuilder) Options() []fx.Option {
	return a.options
}

func (a *applicationBuilder) Logger() logger.Logger {
	return a.logger
}

func (a *applicationBuilder) Environment() environment.Environment {
	return a.environment
}
