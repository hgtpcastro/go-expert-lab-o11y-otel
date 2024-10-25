package app

import (
	"github.com/hgtpcastro/go-expert-lab-o11y-otel/internal/pkg/configs/environment"
	"github.com/hgtpcastro/go-expert-lab-o11y-otel/internal/pkg/fxapp"
	"github.com/hgtpcastro/go-expert-lab-o11y-otel/internal/pkg/logger"
	"github.com/hgtpcastro/go-expert-lab-o11y-otel/internal/services/zipcodeservice/internal/shared/configurations/zipcode"
	"go.uber.org/fx"
)

type ZipCodeApplication struct {
	*zipcode.ZipCodeServiceConfigurator
}

func NewZipCodeApplication(
	providers []interface{},
	decorates []interface{},
	options []fx.Option,
	logger logger.Logger,
	environment environment.Environment,
) *ZipCodeApplication {
	app := fxapp.NewApplication(providers, decorates, options, logger, environment)
	return &ZipCodeApplication{
		ZipCodeServiceConfigurator: zipcode.NewZipCodeServiceConfigurator(app),
	}
}
